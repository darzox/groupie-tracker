package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type City struct {
	Name string
	Lon  float64
	Lat  float64
}

type Artist struct {
	Id             int      `json: id`
	Image          string   `json: image`
	Name           string   `json: name`
	Members        []string `json: members`
	CreationDate   int      `json: creationDate`
	FirstAlbum     string   `json: firstAlbum`
	Relations      string   `json: realations`
	DatesLocations map[string][]string
	Cities         []City
}

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

const (
	artists   = "https://groupietrackers.herokuapp.com/api/artists"
	locations = "https://groupietrackers.herokuapp.com/api/locations"
	dates     = "https://groupietrackers.herokuapp.com/api/dates"
	relation  = "https://groupietrackers.herokuapp.com/api/relation"

	keyMapAPI = "G8AAi1OjQHeUfHNZwHvLYR0iAUTAKKAe"
)

var Artists []Artist // unmarshal json into slice of Artists

func parse() *[]Artist {
	res, err := http.Get(artists)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &Artists); err != nil {
		log.Fatal(err)
	}
	return &Artists
}

func parseLocations(id int) {
	res, err := http.Get(Artists[id-1].Relations)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var temp Relations

	if err = json.Unmarshal(body, &temp); err != nil {
		log.Print(err)
	}
	Artists[id-1].DatesLocations = temp.DatesLocations
	c := &Artists[id-1].Cities
	for city := range Artists[id-1].DatesLocations {
		*c = append(*c, geocode(city))
		// TimeOut for Map API
		if len(Artists[id-1].DatesLocations) > 4 {
			time.Sleep(time.Second / 4)
		}
	}
}

type Result struct {
	Position struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"position"`
}

type Data struct {
	Results []Result `json:"results"`
}

func geocode(city string) (coordinates City) {
	convertedCityName := strings.ReplaceAll(city, "_", "-")
	res, err := http.Get("https://api.tomtom.com/search/2/geocode/" + convertedCityName + ".json?key=" + keyMapAPI)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data Data
	if err = json.Unmarshal(body, &data); err != nil {
		log.Print(err)
	}
	var tempCity City
	tempCity.Lat = data.Results[0].Position.Lat
	tempCity.Lon = data.Results[0].Position.Lon
	c := strings.Split(city, "-")[0]
	prettyCityName := strings.ReplaceAll(c, "_", " ")
	prettyCityName = strings.Title(prettyCityName)
	tempCity.Name = prettyCityName
	return tempCity
}
