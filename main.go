package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type GroupieTracker struct {
	ArtistsUrl   string `json:"artists"`
	LocationsUrl string `json:"locations"`
	DatesUrl     string `json:"dates"`
	RelationUrl  string `json:"relation"`
}

func main() {
	var groupieTrackerObject GroupieTracker

	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	json.Unmarshal(responseData, &groupieTrackerObject)
	fmt.Println(groupieTrackerObject.ArtistsUrl)
	fmt.Println(groupieTrackerObject.LocationsUrl)
	fmt.Println(groupieTrackerObject.DatesUrl)
	fmt.Println(groupieTrackerObject.RelationUrl)

}

// https://github.com/Alika03/groupie-tracker
