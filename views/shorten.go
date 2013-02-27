package views

import (
	"encoding/json"
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
	"io/ioutil"
	"log"
	"net/http"
)

type POSTRequest struct {
	URL string
}

// This will convert a URL into a short url
func ShortenHandler(w http.ResponseWriter, request *http.Request, session *r.Session) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not parse POST body: %v", err), http.StatusInternalServerError)
	}

	var postRequest POSTRequest
	err = json.Unmarshal(body, &postRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not decode JSON: %v", err), http.StatusPreconditionFailed)
	}

	url := postRequest.URL

	var response r.WriteResponse
	row := r.Map{"url": url}
	err = r.Table("url_mapping").Insert(row).Run(session).One(&response)
	if err != nil {
		log.Print(err)
		http.Error(w, "Could not save URL", http.StatusInternalServerError)
	}

	id := response.GeneratedKeys[0]

	fmt.Fprintf(w, "http://%s/%s", request.Host, id)
}
