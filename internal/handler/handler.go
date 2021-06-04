package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sheghun/containerized-webapp/internal/service"
)

type reqBody struct {
	Number int `json:number`
}

func FindHighestPrime(w http.ResponseWriter, r *http.Request) {
	var body reqBody
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"error": "invalid json body payload should be {'number': 55}" }`)
		log.Printf("error occured reading request body: %v", err)
		return
	}

	if body.Number == 0 {
		w.WriteHeader(400)
		io.WriteString(w, `{"error": "Number field is required and must be an integer greater than 0"}`)
		return
	}

	primeNumber := service.FindHighestPrime(body.Number)
	resString := `{"number": ` + fmt.Sprintf("%d", primeNumber) + `}`
	fmt.Println(resString)
	io.WriteString(w, resString)
	return
}

func RenderWebPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html, err := ioutil.ReadFile("./web/build/index.html")
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "error occurred rending page")
		return
	}

	w.Write(html)
}
