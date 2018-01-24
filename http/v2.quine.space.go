package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, quine, backtick+quine+backtick, backtick)
	})
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

// Here there be quines.
var (
	quine = `package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, quine, backtick+quine+backtick, backtick)
	})
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

// Here there be quines.
var (
	quine = %s
	backtick = %q
)
`
	backtick = "`"
)
