package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

/*This var is a pointer towards template.Template that is a
pointer to help process the html.*/
var tpl *template.Template

/*This init function, once it's initialised, makes it so that each html file
in the templates folder is parsed i.e. they all get looked through once and
then stored in the memory ready to go when needed*/
// func init() {
// 	tpl = template.Must(template.ParseGlob("templates/*html"))
// }

var ArtistsID []int
var ArtistsImage []string
var ArtistsName []string
var ArtistsMembers []string
var ArtistsCreationDate int
var ArtistsFirstAlbum string
var ArtistsLocations []string
var ArtistsConcertDates []string
var ArtistsDatesLocations map[string][]string

type TotalInfo struct {
	ArtistID              []int
	ArtistImage           []string
	ArtistName            []string
	ArtistMembers         [][]string
	ArtistCreationDate    []int
	ArtistFirstAlbum      []string
	ArtistLocations       [][]string
	ArtistConcertDates    [][]string
	ArtistsDatesLocations []map[string][]string
}

var Totale TotalInfo

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Dates struct {
	Dates []dates `json:"index"`
}

type dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Locations struct {
	Locations []locations `json:"index"`
}

type locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Relations struct {
	Index []struct {
		ID             int
		DatesLocations map[string][]string
	}
}

// type Relations struct {
// 	Relations []relations
// }

// type relations struct {
// 	ID             int                 `json:"id"`
// 	DatesLocations map[string][]string `json:"datesLocations"`
// }

func main() {

	UnmarshalArtistData()

	fmt.Println(Totale.ArtistImage[0])
	fmt.Println(Totale.ArtistsDatesLocations[0])
	fmt.Println(Totale.ArtistLocations[0])
	fmt.Println(Totale.ArtistConcertDates[0])

	// var x IT
	// for i := 0; i < 52; i++ {
	// 	x[i] = TotalInfo{ArtistsID, TotalInfo.ArtistImage[i], TotalInfo.ArtistName[i], TotalInfo.ArtistMembers[i], TotalInfo.ArtistCreationDate[i], TotalInfo.ArtistFirstAlbum[i], TotalInfo.ArtistLocations[i], TotalInfo.ArtistConcertDates[i], TotalInfo.ArtistsDatesLocations[i]}

	// }

	// fmt.Println(x[0])
}

func UnmarshalArtistData() {

	responseArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic("Couldn't get Artists info from API")
	}
	defer responseArtists.Body.Close()

	responseArtistsData, err := ioutil.ReadAll(responseArtists.Body)
	if err != nil {
		panic("Couldn't read data for Artists!")
	}

	var responseObjectArtists []Artists

	json.Unmarshal(responseArtistsData, &responseObjectArtists)

	for i := 0; i < 52; i++ {
		Totale.ArtistID = append(Totale.ArtistID, responseObjectArtists[i].ID)
		//fmt.Println(Totale.ArtistID)

	}

	for i := 0; i < 52; i++ {
		Totale.ArtistName = append(Totale.ArtistName, responseObjectArtists[i].Name)
	}

	for i := 0; i < len(responseObjectArtists); i++ {
		Totale.ArtistImage = append(Totale.ArtistImage, responseObjectArtists[i].Image)

		for i := 0; i < len(responseObjectArtists); i++ {
			Totale.ArtistCreationDate = append(Totale.ArtistCreationDate, responseObjectArtists[i].CreationDate)

		}

		for i := 0; i < len(responseObjectArtists); i++ {
			Totale.ArtistFirstAlbum = append(Totale.ArtistFirstAlbum, responseObjectArtists[i].FirstAlbum)

		}

		responseRelations, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
		if err != nil {
			panic("Couldn't get the relations data!")
		}

		responseData, err := ioutil.ReadAll(responseRelations.Body)
		if err != nil {
			panic("Couldn't read data for the Relations")
		}

		var responseObjectRelations Relations

		json.Unmarshal(responseData, &responseObjectRelations)

		//var DLMap []map[string][]string

		for _, value := range responseObjectRelations.Index {
			Totale.ArtistsDatesLocations = append(Totale.ArtistsDatesLocations, value.DatesLocations)
		}

		// fmt.Println(DLMap[0])

		// for i := 0; i < 52; i++ {
		// 	for y, v := range responseObjectRelations.Index[i].DatesLocations {
		// 		for _, tmap := range oMap {
		// 			//for key, value := range tmap{
		// 			tmap[y] = v
		// 		}

		// 	}
		// }

		// fmt.Println(len(oMap))
		//fmt.Printf("dl:  %+v\n", responseObjectRelations.Index)

		// for i := 0; i < 52; i++ {
		// 	for y, v := range responseObjectRelations.Index[0].DatesLocations {

		// 		fmt.Println(y, v)

		// 	}
		// 	fmt.Println(len(responseObjectRelations.Index[0].DatesLocations))
		// 	fmt.Println()
		// }

		responseDates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
		if err != nil {
			panic("Couldn't get Dates info from the API!")
		}
		defer responseDates.Body.Close()

		responseDatesData, err := ioutil.ReadAll(responseDates.Body)
		if err != nil {
			panic("Couldn't read data for Dates")
		}

		var responseObjectDates Dates
		json.Unmarshal(responseDatesData, &responseObjectDates)

		for i := 0; i < len(responseObjectDates.Dates); i++ {
			Totale.ArtistConcertDates = append(Totale.ArtistConcertDates, responseObjectDates.Dates[i].Dates)
		}

		responseLocations, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
		if err != nil {
			panic("Couldn't get Location info from API")
		}
		defer responseLocations.Body.Close()

		responseLocationsData, err := ioutil.ReadAll(responseLocations.Body)
		if err != nil {
			panic("Couldn't read data for Locations!")
		}

		var responseObjectLocations Locations
		json.Unmarshal(responseLocationsData, &responseObjectLocations)

		//fmt.Println(responseObjectLocations.Locations[0].Locations)

		for i := 0; i < len(responseObjectLocations.Locations); i++ {
			Totale.ArtistLocations = append(Totale.ArtistLocations, responseObjectLocations.Locations[i].Locations)
		}

	}
}

func Requests() {

	http.HandleFunc("/", index)
	http.HandleFunc("/info", artistInfo)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started on: http://localhost:8080")
}

func index(w http.ResponseWriter, r *http.Request) {

	//-------------Create a struct to hold unmarshalled data-----------
	// var IT TI

	if r.URL.Path != "/" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "index.html", Totale)
	}
}

func artistInfo(w http.ResponseWriter, r *http.Request) {

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		panic("Couldn't get the relations data!")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Couldn't read data for the Artists")
	}

	var responseObject Relations

	json.Unmarshal(responseData, &responseObject)

	if r.URL.Path != "/info" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "info.html", responseObject.Index)
	}

}
