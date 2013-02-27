package views

import (
	"fmt"
	"log"
	"net/http"
	//"github.com/NickPresta/shorter/utils"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/gorilla/mux"
)

// This will convert a short URL into the full website
func EmbiggenHandler(w http.ResponseWriter, request *http.Request, session *r.Session) {
	vars := mux.Vars(request)
	key := vars["key"]

	//id := utils.Decode(key)

	var result map[string]interface{}
	err := r.Table("url_mapping").Get(key, "id").Run(session).One(&result)
	if err != nil || len(result) == 0 {
		log.Print(err)
		http.NotFound(w, request)
	}

	url := result["url"]

	fmt.Fprintf(w, "%s", url)
}
