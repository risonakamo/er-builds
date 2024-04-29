// multithread implementation for builds get function

package erdata_builds

import "sync"

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
	go resultsCollectWorker(routeResultsCh,finalResultCh)

	// create and submit jobs until reached the limit
	var currentPage int=0
	for {
		if currentPage>=pages {
			break
		}

		getRouteDataJobsCh<-GetRouteDataJob{
			Character: character,
			Weapon: weapon,
			Versions: versions,

			PageStart: currentPage,
			PageEnd: currentPage+pagesPerWorker,
		}

		currentPage+=pagesPerWorker
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

	var job GetRouteDataJob
	for job = range jobsCh {
		var routes []ErRoute2=getRouteDataMultiPage(
			job.Character,
			job.Weapon,
			job.PageStart,
			job.PageEnd,
			true,
		)

		submitCh<-filterByVersion(routes,job.Versions)
	}
}

// worker to recv results from main route retrieval workers. combines all results into 1 array.
// upon closing of the results ch, will submit a single result to the final result ch. then quits.
// should only have 1 of these.
func resultsCollectWorker(
	resultsCh <-chan []ErRoute2,
	finalResultCh chan<- []ErRoute2,
) {
	var collectedResults []ErRoute2

	var workerResult []ErRoute2
	for workerResult = range resultsCh {
		collectedResults=append(collectedResults,workerResult...)
	}

	finalResultCh<-collectedResults
}