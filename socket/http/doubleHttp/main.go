package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	ports := []string{"9000", "8000"}
	for _, port := range ports {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			printRequest(request)
			writer.Write([]byte("<h1>hello world 1</h1>"))
		})
		go func(port string) {
			log.Fatal(http.ListenAndServe(":"+port, mux))
		}(port)
	}
	select {}
}

func printRequest(request *http.Request) {

	fmt.Println(request.Host)
	fmt.Println(request.Method)
	fmt.Println(request.Referer())
	fmt.Println(request.UserAgent())
}
