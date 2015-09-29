package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	self := `package main%simport (
	"flag"
	"fmt"
	"log"
	"net/http"%s)%sfunc main() {
	self := %s

	port := flag.Int("port", 3000, "Port to run the server on")
	n, t, bt, cpd := %q, %q, %q, %q

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, self, n+n, n, n+n, bt+self+bt, n, t, bt, cpd, n, n)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(cpd, *port), nil))%s}%s`

	port := flag.Int("port", 3000, "Port to run the server on")
	n, t, bt, cpd := "\n", "\t", "`", ":%d"

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, self, n+n, n, n+n, bt+self+bt, n, t, bt, cpd, n, n)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(cpd, *port), nil))
}
