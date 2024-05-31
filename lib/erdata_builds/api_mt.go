// multithread implementation for builds get function

package erdata_builds

import (
	"sync"

	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
)

type GetRouteDataJob struct {
	Character string
	Weapon string
	Versions []string

	PageStart int
	PageEnd int
}

// multithread version of get route data 2
func GetRouteData2Mt(
    character string,
    weapon string,
    pages int,
    versions []string,

	workers int,
	pagesPerWorker int,
) []ErRoute2 {
	// ch to submit jobs to workers
	var getRouteDataJobsCh chan GetRouteDataJob=make(chan GetRouteDataJob)

	// ch to recv results from workers
	var routeResultsCh chan []ErRoute2=make(chan []ErRoute2)

	// final result ch
	var finalResultCh chan []ErRoute2=make(chan []ErRoute2)

	// early stop ch. result collect submits to this if it is early stopping
	var earlyStopCh chan bool=make(chan bool)

	var workersWg sync.WaitGroup

	// spawn workers
	for i:=0;i<workers;i++ {
		workersWg.Add(1)
		go getRouteDataWorker(
			getRouteDataJobsCh,
			routeResultsCh,
			&workersWg,
		)
	}

	// spawn collection worker
	go resultsCollectWorker(routeResultsCh,finalResultCh,earlyStopCh)

	// create and submit jobs until reached the limit, or a job returned an empty result.
	// since we are submitting jobs starting at page 0, if a single job returns empty, we should
	// immediately stop creating new jobs, as we know all new jobs will always return empty.
	// but, we still need to complete the current jobs, as some of them might be ongoing before
	// the first zero job was submitted.
	var currentPage int=0

	jobSubmit:
	for {
		if currentPage>=pages {
			break
		}

		// try to pull from early stop. if successful, trigger the early stop and stop
		// creating jobs
		select {
		case <-earlyStopCh:
			break jobSubmit

		case getRouteDataJobsCh<-GetRouteDataJob{
			Character: character,
			Weapon: weapon,
			Versions: versions,

			PageStart: currentPage,
			PageEnd: currentPage+pagesPerWorker,
		}:
			currentPage+=pagesPerWorker

		default:
		}
	}

	// all jobs submitted. close the channel to finish workers
	close(getRouteDataJobsCh)

	// wait for all workers to finish
	workersWg.Wait()

	// close route result ch to get collector to send out the final result
	close(routeResultsCh)

	// pull the final result from the collector worker and close final result ch
	var finalResult []ErRoute2=<-finalResultCh
	close(finalResultCh)

	return finalResult
}

// gets route data jobs from job ch. performs retrieval of requested routes.
// submits results to submit ch. closes when jobs ch closes.
// can spawn en-masse. use wg to wait for all to close.
func getRouteDataWorker(
	jobsCh <-chan GetRouteDataJob,
	submitCh chan<- []ErRoute2,

	wg *sync.WaitGroup,
) {
	defer wg.Done()

	var client *req.Client=req.C()
	var job GetRouteDataJob
	for job = range jobsCh {
		var routes []ErRoute2=getRouteDataMultiPage(
			job.Character,
			job.Weapon,
			job.PageStart,
			job.PageEnd,
			true,
			client,
		)

		submitCh<-filterByVersion(routes,job.Versions)
	}
}

// worker to recv results from main route retrieval workers. combines all results into 1 array.
// upon closing of the results ch, will submit a single result to the final result ch. then quits.
// if get a result from the results ch that is empty length, submits to earlystop channel, but doesnt
// stop collecting as there might still be some workers that are still getting non-empty results.
// only sends early stop signal once.
// should only have 1 of these.
func resultsCollectWorker(
	resultsCh <-chan []ErRoute2,
	finalResultCh chan<- []ErRoute2,
	earlyStopCh chan<- bool,
) {
	var collectedResults []ErRoute2

	var workerResult []ErRoute2
	var sentEarlyStop bool=false
	for workerResult = range resultsCh {
		if len(workerResult)==0 && !sentEarlyStop {
			log.Info().Msg("collector worker got empty result, early stopping collection")
			earlyStopCh<-true
			sentEarlyStop=true
		}

		collectedResults=append(collectedResults,workerResult...)
	}

	finalResultCh<-collectedResults
}