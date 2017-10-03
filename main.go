package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

type movie struct {
	name string
	rating string
}

func getMovies()  []movie{

	movies := make([]movie,0)

	doc, err := goquery.NewDocument("https://afisha.tut.by/film-mogilev")
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	log.Print(doc.Find("a.media").Html())

	return movies
}

func main()  {
	movies := getMovies()
	for _, m := range movies{
		log.Print(m.name+m.rating)
	}
}