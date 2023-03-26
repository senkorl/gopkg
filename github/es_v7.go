package github

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func Demo7() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9201",
		},
		// ...
	}
	es, err := elasticsearch.NewClient(cfg)
	log.Println(elasticsearch.Version)
	if err != nil {
		log.Println("lianjie shibai")
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}
