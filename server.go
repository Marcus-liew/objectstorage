package main

import (
	"log"
	"net/http"
	"objectstorage/heartbeat"
	"objectstorage/locate"
	"objectstorage/objects"
	"os"
)

func main() {
	//第二章加入函数
	go heartbeat.StartHeartbeat()
	//第二章加入函数
	go locate.StartLocate()
	//第一章函数
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
