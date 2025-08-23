package controllers

import (
	"golangR99/constants"
	"golangR99/models"
	"sync"

	"golang.org/x/time/rate"
)

// only this many http requests/ goroutines can run at a time
const limit = 6

// use semaphore to cap # http requests to limit
var httpLimiter = rate.NewLimiter(rate.Limit(limit), 1)

// use semaphore to cap # goroutines to limit
var semaphore = make(chan struct{}, limit)

func GetAllConcurrently(characters, psychubes []string, nodes []models.Node) []string {
	var wg sync.WaitGroup
	resultsChan := make(chan string, len(characters)+len(psychubes))
	// get characters
	wg.Add(1)
	go func() {
		defer wg.Done()
		fetchCharacters(characters, resultsChan)
	}()
	// get psychubes
	wg.Add(1)
	go func() {
		defer wg.Done()
		fetchPsychubes(psychubes, nodes, resultsChan)

	}()
	// close channel when both funcs are done
	go func() {
		wg.Wait()
		close(resultsChan)
	}()
	// collect results
	var results []string
	for result := range resultsChan {
		results = append(results, result)
	}
	// return results
	return results
}

func GetLists() ([]models.Node, []models.Node) {
	characters := getList(constants.CHARACTERS_ENDPOINT)
	psychubes := getList(constants.PSYCHUBES_ENDPOINT)

	return characters, psychubes
}
