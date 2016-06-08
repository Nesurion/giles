package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	imdb "github.com/eefret/gomdb"
	"github.com/gin-gonic/gin"
)

const MovieTempPath string = "/Users/nesurion/Development/giles/example/temp/movies"
const MovieArchivPath string = "/Users/nesurion/Development/giles/example/archiv/movies"

const TvTempPath string = "/Users/nesurion/Development/giles/example/temp/tv"
const TvArchivPath string = "/Users/nesurion/Development/giles/example/archiv/tv"

func main() {
	srv := gin.Default()
	srv.LoadHTMLGlob("templates/*")
	srv.Static("/css", "./assets/css")
	srv.Static("/js", "./assets/js")

	srv.GET("/", func(c *gin.Context) {
		// load temp movies
		movies := loadMovies(MovieTempPath, false)
		// add archived movies
		movies = append(movies, loadMovies(MovieArchivPath, true)...)

		// load temp shows
		shows := loadShows(TvTempPath, false)
		// add archived shows
		shows = append(shows, loadShows(TvArchivPath, true)...)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Movies": movies,
			"TV":     shows,
		})
	})

	srv.POST("/api/archiv/movie", func(c *gin.Context) {
		filePath := c.PostForm("path")
		title := c.PostForm("title")
		_, dest := path.Split(filePath)
		err := os.Rename(path.Clean(filePath), path.Join(MovieArchivPath, dest))
		if err != nil {
			fmt.Printf("Failed to move dir: %s", err)
		}
		fmt.Printf("%s: %s\n", title, filePath)
		c.JSON(200, gin.H{
			"status": "archiving",
			"title":  title,
		})
	})

	srv.POST("/api/archiv/show", func(c *gin.Context) {
		filePath := c.PostForm("path")
		title := c.PostForm("title")
		_, dest := path.Split(filePath)
		err := os.Rename(path.Clean(filePath), path.Join(TvArchivPath, dest))
		if err != nil {
			fmt.Printf("Failed to move dir: %s", err)
		}
		fmt.Printf("%s: %s\n", title, filePath)
		c.JSON(200, gin.H{
			"status": "archiving",
			"title":  title,
		})
	})

	srv.POST("/api/delete", func(c *gin.Context) {
		filePath := c.PostForm("path")
		title := c.PostForm("title")
		err := os.RemoveAll(filePath)
		if err != nil {
			fmt.Printf("failed to delete dir: %s", err)
		}
		fmt.Printf("%s: %s\n", title, filePath)
		c.JSON(200, gin.H{
			"status": "deleting",
			"title":  title,
		})
	})

	srv.Run(":8080")
}

func loadMovies(sourcePath string, archiv bool) []*Media {
	var movies []*Media
	files, err := ioutil.ReadDir(sourcePath)
	if err != nil {
		fmt.Printf("Failed to load movies: %s\n", err)
	}
	for _, f := range files {
		// -6 gives the title without the year
		title := f.Name()[:len(f.Name())-6]
		// remove ", The" from the title
		title = strings.Split(title, ",")[0]

		year := f.Name()[len(f.Name())-6:]
		year = year[1:len(year)]
		year = year[:len(year)-1]

		movie, err := imdb.MovieByTitle(title, year)
		if err != nil {
			fmt.Printf("Failed to load movie %s from imdb", f.Name())
		}
		movies = append(movies, &Media{
			Imdb:     movie,
			Path:     path.Join(sourcePath, f.Name()),
			Archived: archiv,
		})
	}
	return movies
}

func loadShows(sourcePath string, archiv bool) []*Media {
	var shows []*Media
	files, err := ioutil.ReadDir(sourcePath)
	if err != nil {
		fmt.Printf("Failed to load shows: %s\n", err)
	}
	for _, f := range files {
		res, err := imdb.Search(f.Name(), "")
		if err != nil {
			fmt.Printf("Failed to load show %s from imdb", f.Name())
		}
		show, err := imdb.MovieByImdbID(res.Search[0].ImdbID)
		if err != nil {
			fmt.Printf("Failed to load show %s from imdb", f.Name())
		}
		shows = append(shows, &Media{
			Imdb:     show,
			Path:     path.Join(sourcePath, f.Name()),
			Archived: archiv,
		})
	}
	return shows
}
