package main

import (
	"./omdb"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func printMovie(m *omdb.Movie) {
	fmt.Printf("The poster url for the movie %q is %q\n", m.Title, m.Poster)
}

func downLoadPoster(m *omdb.Movie) {
	posterURL := m.Poster
	posterFileName := m.Title + "_poster.jpg"
	res, resErr := http.Get(posterURL)
	if resErr != nil {
		log.Fatalf("Error occured when acessing %s: %s", posterURL, resErr)
	}
	file, fileErr := os.Create(posterFileName)
	if fileErr != nil {
		log.Fatalf("Unable to create file %s: %s", posterFileName, fileErr)
	}
	io.Copy(file, res.Body)
	fmt.Println("Downlaod complete!")
}

func main() {
	result, err := omdb.SearchMovies(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	printMovie(result)
	downLoadPoster(result)
}
