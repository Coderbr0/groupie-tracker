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

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	// Locations    string	   `json:"locations"`
	// ConcertDates string	   `json:"concertDates"`
	// Relations    string	   `json:"relations"`
}

type Locations struct {
	Index int      `json:"index"`
	Id    int      `json:"id"`
	Loc   []string `json:"locations"`
}

// type Dates struct {
// 	Id    int
// 	Dates []string
// }

// type Relations struct {
// 	Id             int
// 	DatesLocations map[int]string
// }

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

	var artistsObject []Artists // []Artists as we have many artists; links to the Artists struct and stores as a variable

	response1, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData1, err := ioutil.ReadAll(response1.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData1))
	json.Unmarshal(responseData1, &artistsObject)
	fmt.Println(artistsObject[0].Id)
	fmt.Println(artistsObject[0].Image)
	fmt.Println(artistsObject[0].Name)
	fmt.Println(artistsObject[0].Members)
	fmt.Println(artistsObject[0].CreationDate)
	fmt.Println(artistsObject[0].FirstAlbum)
	// fmt.Println(artistsObject[0].Locations)
	// fmt.Println(artistsObject[0].ConcertDates)
	// fmt.Println(artistsObject[0].Relations)

	var locationsObject []Locations

	response2, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData2, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData2))
	json.Unmarshal(responseData2, &locationsObject)

	fmt.Println(locationsObject[0].Index)

	// for _, locations := range locationsObject {
	// 	fmt.Println(locations)
	// 	fmt.Println(locationsObject[0])
	// }
}

// https://github.com/Alika03/groupie-tracker
