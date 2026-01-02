package elasticsearch7

import (
	"fmt"
	"io"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func GetESInfo() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9207",
		},
		// ...
	}
	es, err := elasticsearch.NewClient(cfg)
	fmt.Println(elasticsearch.Version)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	fmt.Println(res.String())
}
