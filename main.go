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
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type IndexStruct struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	ID  int      `json:"id"`
	Loc []string `json:"locations"`
}

type IndexStructDates struct {
	Index []Dates `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// type Relations struct {
// 	Id             int
// 	DatesLocations map[int]string
// }

func main() {
	var groupieTrackerObject GroupieTracker // Storing the struct into a variable so that it can be used

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

	fmt.Println(artistsObject[0].ID)
	fmt.Println(artistsObject[0].Image)
	fmt.Println(artistsObject[0].Name)
	fmt.Println(artistsObject[0].Members)
	fmt.Println(artistsObject[0].CreationDate)
	fmt.Println(artistsObject[0].FirstAlbum)

	var locationsObject IndexStruct

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

	fmt.Println(locationsObject.Index[0].ID)
	fmt.Println(locationsObject.Index[0].Loc[0])
	fmt.Println(locationsObject.Index[0].Loc[1])

	for _, locationsValue := range locationsObject.Index {
		fmt.Println(locationsValue)
	}

	var datesObject IndexStructDates

	response3, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData3, err := ioutil.ReadAll(response3.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData3))
	json.Unmarshal(responseData3, &datesObject)

	for _, datesValue := range datesObject.Index {
		fmt.Println(datesValue)
	}
}

// https://github.com/Alika03/groupie-tracker
