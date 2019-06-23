package main

import "net/http"

func main() {
	http.HandleFunc("/hello",
		func(writer http.ResponseWriter, r *http.Request) {
			writer.Write([]byte("helloWorld"))
		})
	http.ListenAndServe(":8082", nil)
}
