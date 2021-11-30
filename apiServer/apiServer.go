package main

import (
	"log"
	"net/http"
	"objectstorage/apiServer/heartbeat"
	"objectstorage/apiServer/locate"
	"objectstorage/apiServer/objects"
	"objectstorage/apiServer/versions"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
