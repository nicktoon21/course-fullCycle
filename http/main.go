package main

import (
	"io"
	"net/http"
)


func main() {
	http.Handle("/", blog{})
  http.ListenAndServe(":8080", nil)
}

type blog struct{}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello World!")
}

// func main() {
// 		req, err := http.Get("https://www.google.com")
// 		if err != nil {
// 			panic(err)
// 		}

// 		defer req.Body.Close()

// 		res, err := io.ReadAll(req.Body)
// 		if err != nil {
// 			panic(err)
// 		}

// 		println(string(res))
// }