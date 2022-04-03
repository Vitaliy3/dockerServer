package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/ha", func(writer http.ResponseWriter, request *http.Request) {
		log, err := os.OpenFile("logs.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		str := fmt.Sprintf("%s :: %s \n", "new request", time.Now())
		_, err = log.WriteString(str)
		if err != nil {
			panic(err)
			return
		}

		fmt.Println("new request", time.Now())

	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
		return
	}
}
