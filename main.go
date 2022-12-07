package main

import (
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://localhost:7700",
	})

	// https://docs.meilisearch.com/learn/core_concepts/primary_key.html#primary-field
	// md5("C=%s," +"G=%s,K=%s,NS=%s,N=%s")

	client.CreateIndex(&meilisearch.IndexConfig{
		Uid:        "k8s",
		PrimaryKey: "reference_number",
	})

	documents := []map[string]interface{}{
		{
			"reference_number": 287947,
			"title":            "Diary of a Wimpy Kid",
			"author":           "Jeff Kinney",
			"genres":           []string{"comedy", "humor"},
			"price":            5.00,
		},
	}
	client.Index("k8s").AddDocuments(documents, "reference_number")
}
