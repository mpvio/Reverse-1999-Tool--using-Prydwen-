package controllers

import (
	"encoding/json"
	"golangR99/constants"
	"golangR99/models"
	"io"
	"log"
	"net/http"
)

func GetCharacter(name string) models.CharacterDB {
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

func GetPsychube(name string, nodes []models.Node) models.PsychubeDB {
	for _, node := range nodes {
		if node.Name == name {
			return node.ConvertToPsychube()
		}
	}
	return models.PsychubeDB{}
}

func GetLists() ([]models.Node, []models.Node) {
	characters := getList(constants.CHARACTERS_ENDPOINT)
	psychubes := getList(constants.PSYCHUBES_ENDPOINT)

	return characters, psychubes
}

func getList(uri string) []models.Node {
	content := getBytesFromEndpoint(uri)
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

func getBytesFromEndpoint(uri string) []byte {
	// call endpoint
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
	// return content
	return content
}
