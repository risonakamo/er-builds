// multi thread version of get build2

package nica

import (
	"er-builds/lib/dak_gg"
	"er-builds/lib/oer_api"
	"sync"

	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
)

func GetBuilds2_mt(
	buildIds []int,
	traitSkills dak_gg.TraitSkillMap,
	langfileDict oer_api.OerLangDict,
	client *req.Client,

	workers int,
) []NicaBuild2 {
	// submit build ids for workers to get
	var jobsCh chan int=make(chan int)

	// workers submit results to here
	var buildResultCh chan NicaBuild2=make(chan NicaBuild2)

	// final array of collected results appears here
	var finalResultCh chan []NicaBuild2=make(chan []NicaBuild2)

	var workerWg sync.WaitGroup

	// start main getter workers
	for range workers {
		workerWg.Add(1)
		go getBuildWorker(
			jobsCh,
			buildResultCh,
			traitSkills,
			langfileDict,
			client,
			&workerWg,
		)
	}

	// start result collector
	go resultCollectWorker(buildResultCh,finalResultCh)

	// submit all builds
	for buildIdI := range buildIds {
		jobsCh<-buildIds[buildIdI]
	}

	// all builds submitted. close the jobs ch
	close(jobsCh)

	// wait for all workers to finish
	workerWg.Wait()

	// close results ch to get collector to send out result
	close(buildResultCh)

	// get the final result
	var finalResult []NicaBuild2=<-finalResultCh
	close(finalResultCh)

	return finalResult
}

// build getter worker. gets jobs from jobs ch, gets the build and submits
// into results ch
func getBuildWorker(
	jobsCh <-chan int,
	jobSubmitCh chan<- NicaBuild2,
	traitSkills dak_gg.TraitSkillMap,
	langfileDict oer_api.OerLangDict,
	client *req.Client,

	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for job := range jobsCh {
		var build NicaBuild2=GetBuild2(
			job,
			traitSkills,
			langfileDict,
			client,
		)

		if len(build.ItemInfos)==0 {
			log.Warn().
			Int("build id",job).
			Msg("got an empty build")
		} else {
			jobSubmitCh<-build
		}

	}
}

// recvs builds on result ch and collects into array. upon
// results channel closing, submits array into final results ch
func resultCollectWorker(
	resultsCh <-chan NicaBuild2,
	finalResultCh chan<- []NicaBuild2,
) {
	var collected []NicaBuild2
	for build := range resultsCh {
		collected=append(collected,build)
	}

	finalResultCh<-collected
}