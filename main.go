package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://localhost:7700",
	})

	jsonFile, _ := os.Open("movies.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var movies []map[string]interface{}
	json.Unmarshal(byteValue, &movies)

	_, err := client.Index("movies").AddDocuments(movies)
	if err != nil {
		panic(err)
	}
}
