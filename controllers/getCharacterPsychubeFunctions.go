package controllers

import (
	"context"
	"encoding/json"
	"golangR99/constants"
	"golangR99/models"
	"log"
	"sync"
)

// concurrent functions
func fetchCharacters(characters []string, results chan<- string) {
	var wg sync.WaitGroup
	for _, name := range characters {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			// use semaphore to keep gaining struct{}s until hitting its cap
			semaphore <- struct{}{}
			// goroutines won't be created by for loop if semaphore is full
			defer func() { <-semaphore }()
			// apply httpLimiter to limit # http requests
			httpLimiter.Wait(context.Background())
			// get character as normal
			character := getCharacter(name)
			result := WriteToFile(character)
			results <- result
		}(name)
	}
	wg.Wait()
}

func fetchPsychubes(psychubes []string, nodes []models.Node, results chan<- string) {
	var wg sync.WaitGroup
	for _, name := range psychubes {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			// use semaphore to keep gaining struct{}s until hitting its cap
			semaphore <- struct{}{}
			// goroutines won't be created by for loop if semaphore is full
			defer func() { <-semaphore }()
			psychube := getPsychube(name, nodes)
			result := WriteToFile(psychube)
			results <- result
		}(name)
	}
	wg.Wait()
}

// individual calling functions
func getCharacter(name string) models.CharacterDB {
	uri := constants.GET_CHARACTER_ENDPOINT(name)
	content := getBytesFromEndpoint(uri)
	// cast to character
	var character models.Character
	err := json.Unmarshal(content, &character)
	if err != nil {
		log.Fatal(err)
	}
	// return character
	new_char := character.Result.Data.CurrentUnit.Nodes[0].Convert()
	return new_char
}

func getPsychube(name string, nodes []models.Node) models.PsychubeDB {
	for _, node := range nodes {
		if node.Name == name {
			return node.ConvertToPsychube()
		}
	}
	return models.PsychubeDB{}
}
