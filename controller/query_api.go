package controller

import (
	"encoding/json"
	"github.com/blevesearch/bleve"
	"io/ioutil"
	"log"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var query Query
	err = json.Unmarshal(b, &query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	index, _ := bleve.Open(query.Index)
	_query := bleve.NewQueryStringQuery(query.Text)
	searchRequest := bleve.NewSearchRequest(_query)
	searchResult, _ := index.Search(searchRequest)

	log.Printf("searchResult %v\n", searchResult)

	log.Printf("hist %+v\n", searchResult.Hits)

	log.Printf("status %+v\n", *searchResult.Status)

	searchResultBytes, err := json.Marshal(searchResult)

	w.Write(searchResultBytes)
	w.WriteHeader(http.StatusOK)
}
