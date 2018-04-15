package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	httpsPort := os.Getenv("HTTPS_PORT")
	secretDir := os.Getenv("SECRET_DIR")
	autocertWhitelist := os.Getenv("HOST_WHITELIST")
	if autocertWhitelist == "" {
		autocertWhitelist = "quine.space"
	}

	// Setup and handle autocert - a.k.a. Let's Encrypt
	m := &autocert.Manager{
		Cache:      autocert.DirCache(secretDir),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(autocertWhitelist),
	}

	// Redirect to https
	httpRedirectHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Connection", "close")
		url := "https://" + req.Host + req.URL.String()
		http.Redirect(w, req, url, http.StatusTemporaryRedirect)
	})

	// Run the http autocert server with https redirect fallback
	go http.ListenAndServe(":"+httpPort, m.HTTPHandler(httpRedirectHandler))

	// TLS Config based on https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/
	tlsConfig := &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		// GetCertificate provided by the autocert manager
		GetCertificate:           m.GetCertificate,
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
	}

	// Setup the TLS server
	srv := &http.Server{
		Addr:         ":" + httpsPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintf(w, quine, backtick+quine+backtick, backtick)
		}),
	}

	log.Println(srv.ListenAndServeTLS("", ""))
}

//
// Everything below this is purely for the quine
//

var (
	quine = `package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	httpsPort := os.Getenv("HTTPS_PORT")
	secretDir := os.Getenv("SECRET_DIR")
	autocertWhitelist := os.Getenv("HOST_WHITELIST")
	if autocertWhitelist == "" {
		autocertWhitelist = "quine.space"
	}

	// Setup and handle autocert - a.k.a. Let's Encrypt
	m := &autocert.Manager{
		Cache:      autocert.DirCache(secretDir),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(autocertWhitelist),
	}

	// Redirect to https
	httpRedirectHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Connection", "close")
		url := "https://" + req.Host + req.URL.String()
		http.Redirect(w, req, url, http.StatusTemporaryRedirect)
	})

	// Run the http autocert server with https redirect fallback
	go http.ListenAndServe(":"+httpPort, m.HTTPHandler(httpRedirectHandler))

	// TLS Config based on https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/
	tlsConfig := &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		// GetCertificate provided by the autocert manager
		GetCertificate:           m.GetCertificate,
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
	}

	// Setup the TLS server
	srv := &http.Server{
		Addr:         ":" + httpsPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintf(w, quine, backtick+quine+backtick, backtick)
		}),
	}

	log.Println(srv.ListenAndServeTLS("", ""))
}

//
// Everything below this is purely for the quine
//

var (
	quine = %s
	backtick = %q
)
`
	backtick = "`"
)
