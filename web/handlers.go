package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ViewData struct {
	Number      string
	Description string
}

var notFound ViewData = ViewData{
	Number:      "404",
	Description: "Not found",
}

var badRequest ViewData = ViewData{
	Number:      "405",
	Description: "Method Not Allowed",
}

var serverError ViewData = ViewData{
	Number:      "500",
	Description: "Internal Server Error",
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		t, _ := template.ParseFiles("./ui/html/errors.html")
		t.Execute(w, notFound)
		return
	} else {
		if r.Method != http.MethodGet {
			w.WriteHeader(405)
			t, _ := template.ParseFiles("./ui/html/errors.html")
			t.Execute(w, badRequest)
			return
		}
		ptArtists := parse()
		ts, err := template.ParseFiles("./ui/html/home.html")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			t, _ := template.ParseFiles("./ui/html/errors.html")
			t.Execute(w, serverError)
			return
		}
		err = ts.Execute(w, ptArtists)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			t, _ := template.ParseFiles("./ui/html/errors.html")
			t.Execute(w, serverError)
			return
		}
	}
}

func artist(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(404)
		t, _ := template.ParseFiles("./ui/html/errors.html")
		t.Execute(w, notFound)
		return
	}
	if idInt < 1 || idInt > 52 {
		w.WriteHeader(404)
		t, _ := template.ParseFiles("./ui/html/errors.html")
		t.Execute(w, notFound)
		return
	}
	parseLocations(idInt)
	if err != nil {
		log.Fatal(err)
	}
	t, _ := template.ParseFiles("./ui/html/artist.html")
	t.Execute(w, Artists[idInt-1])
}
