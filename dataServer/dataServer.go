package main

import (
	"log"
	"net/http"
	"objectstorage/dataServer/heartbeat"
	"objectstorage/dataServer/locate"
	"objectstorage/dataServer/objects"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
