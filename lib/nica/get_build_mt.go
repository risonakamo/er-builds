// multi thread version of get build2

package nica

import (
	"er-builds/lib/dak_gg"
	"sync"

	"github.com/imroc/req/v3"
)

func GetBuilds2_mt(
	buildIds []int,
	traitSkills dak_gg.TraitSkillMap,
	client *req.Client,

	workers int,
) []NicaBuild2 {
	// submit build ids for workers to get
	var jobsCh chan int=make(chan int)

	// workers submit results to here
	var buildResultCh chan NicaBuild2=make(chan NicaBuild2)

	var workerWg sync.WaitGroup

	for range workers {
		workerWg.Add(1)
		go getBuildWorker(
			jobsCh,
			buildResultCh,
			traitSkills,
			client,
			&workerWg,
		)
	}
}

// build getter worker. gets jobs from jobs ch, gets the build and submits
// into results ch
func getBuildWorker(
	jobsCh <-chan int,
	jobSubmitCh chan<- NicaBuild2,
	traitSkills dak_gg.TraitSkillMap,
	client *req.Client,

	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for job := range jobsCh {
		var build NicaBuild2=GetBuild2(
			job,
			traitSkills,
			client,
		)

		jobSubmitCh<-build
	}
}