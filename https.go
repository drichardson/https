package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "8080", "HTTP port.")
	securePort := flag.String("secure-port", "8443", "HTTPS port.")
	directory := flag.String("root", ".", "The directory to serve")
	certFile := flag.String("cert", "cert.pem", "Public key certificate.")
	keyFile := flag.String("key", "key.pem", "Private key.")

	flag.Parse()

	if _, err := os.Stat(*certFile); os.IsNotExist(err) {
		log.Fatalln("Certificate file", *certFile, "does not exist. Create one with: openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout key.pem -out cert.pem")
	}

	if _, err := os.Stat(*keyFile); os.IsNotExist(err) {
		log.Fatalln("Key file", *keyFile, "does not exist. Create one with: openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout key.pem -out cert.pem")
	}

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	serveHttp := func() {
		log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	}

	go serveHttp()

	log.Printf("Serving %s on HTTPS port: %s\n", *directory, *securePort)
	log.Fatal(http.ListenAndServeTLS(":"+*securePort, *certFile, *keyFile, nil))
}
