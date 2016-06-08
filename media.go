package main

import imdb "github.com/eefret/gomdb"

type Media struct {
	Imdb     *imdb.MovieResult
	Path     string
	Archived bool
}
