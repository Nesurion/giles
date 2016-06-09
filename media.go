package main

import imdb "github.com/eefret/gomdb"

type Media struct {
	Imdb     *imdb.MovieResult
	DirName  string
	Archived bool
}
