package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	resp, err := http.Get("https://gopub.jota-fab.com/robots.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(resp.Status)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Print(string(body))
	resp.Body.Close()

	type response1 struct {
		Page   int
		Frutis []string
	}

	res1 := &response1{
		Page:   1,
		Frutis: []string{"apple", "banana", "pear"},
	}
	res2, _ := json.Marshal(res1)

	cont := strings.NewReader(string(res2))

	res, err := http.Post("https://gopub.jota-fab.com/ping", "application/json", cont)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(resp.Status)

	body, err = io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}
	res.Body.Close()
	time.Sleep(1000 * time.Millisecond)

	log.Print(string(body))

}
