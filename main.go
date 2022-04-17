package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/data", func(writer http.ResponseWriter, request *http.Request) {

		keys, ok := request.URL.Query()["text"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'text' is missing")
			return
		}

		fLog, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		logstr := fmt.Sprintf("%s %s \n", time.Now().Format(time.RFC850), keys[0])
		fLog.WriteString(logstr)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
		return
	}
}
