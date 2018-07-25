package controller

import (
	"encoding/json"
	"github.com/blevesearch/bleve"
	"io/ioutil"
	"net/http"
)

func insert(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var appeal Appeal
	err = json.Unmarshal(b, &appeal)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/*message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "example",
		From: "marty.schoch@gmail.com",
		Body: "bleve indexing is easy",
	}*/

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(appeal.Index, mapping)
	if err != nil {
		panic(err)
	}
	index.Index(appeal.Id, appeal.Source)

	w.WriteHeader(http.StatusOK)
}
