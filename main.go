package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)

	http.HandleFunc("/data", func(writer http.ResponseWriter, request *http.Request) {

		keys, ok := request.URL.Query()["text"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'text' is missing")
			return
		}

		log.Printf("%s %s \n", time.Now().Format(time.RFC850), keys[0])
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
		return
	}
}
