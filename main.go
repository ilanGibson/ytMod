package main

import (
	"flag"
	"log"
	"net/http"

	"ytMod/server"
)

func main() {
	port := flag.String("p", ":8090", "server port")
	flag.Parse()

	Srvr := server.NewServer()
	http.HandleFunc("/api/timer/current", Srvr.GetCurrentWatchTime)
	http.HandleFunc("/api/timer/remain", Srvr.GetRemainingWatchTime)
	http.HandleFunc("/api/timer/start", Srvr.StartTimer)
	http.HandleFunc("/api/timer/end", Srvr.EndTimer)
	http.HandleFunc("/api/timer/bonus", Srvr.AddBonusAllotment)

	log.Printf("running server... %v", *port)
	http.ListenAndServe(*port, nil)
}
