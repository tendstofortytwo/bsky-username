package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var (
	port = flag.Int("port", 3000, "port to listen on")
)

func readDB() (map[string]string, error) {
	var mp map[string]string
	file, err := os.Open("db.json")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(file).Decode(&mp)
	if err != nil {
		return nil, err
	}
	return mp, nil
}

func main() {
	flag.Parse()
	mp, err := readDB()
	if err != nil {
		log.Fatalf("could not read db: %v", err)
	}
	http.HandleFunc("GET /.well-known/atproto-did", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET %s", r.Host)

		newMp, err := readDB()
		if err != nil {
			log.Printf("could not read db (%v), trying to use last cached copy", err)
		} else {
			mp = newMp
		}

		did, ok := mp[r.Host]
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Write([]byte(did))
	}))
	hp := net.JoinHostPort("", strconv.Itoa(*port))
	log.Printf("listening on &q", hp)
	http.ListenAndServe(hp, nil)
}
