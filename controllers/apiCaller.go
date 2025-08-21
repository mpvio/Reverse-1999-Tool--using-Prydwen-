package controllers

import (
	"encoding/json"
	"fmt"
	"golangR99/constants"
	"golangR99/models"
	"io"
	"log"
	"net/http"
)

func GetLists() ([]models.Node, []models.Node) {
	characters := getList(constants.CHARACTERS_ENDPOINT)
	psychubes := getList(constants.PSYCHUBES_ENDPOINT)

	for _, val := range characters {
		fmt.Printf("%s: %s\n", val.Name, val.Slug)
	}
	for _, val := range psychubes {
		if val.Description1.Raw == "" {
			fmt.Println(val.Name, ": No Effect")
		} else {
			var desc models.Raw
			err := json.Unmarshal([]byte(val.Description1.Raw), &desc)
			if err != nil {
				fmt.Println(val.Name, ":", err)
			} else {
				fmt.Println(val.Name, ":", desc.Content)
			}
		}
	}

	return characters, psychubes
}

func getList(uri string) []models.Node {
	// read character info
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	// get response body
	body := resp.Body
	if body != nil {
		defer body.Close()
	}
	// convert to bytes
	content, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	// cast to list
	var items models.ItemList
	jsonErr := json.Unmarshal(content, &items)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// get relevant part
	var info []models.Node
	// if characters.nodes is empty, this is a call to psychubes endpoint
	if len(items.Data.Characters.Nodes) > 0 {
		info = items.Data.Characters.Nodes
	} else {
		info = items.Data.Psychubes.Nodes
	}
	// return nodes
	return info
}
