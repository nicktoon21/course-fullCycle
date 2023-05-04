package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()
	log.Println("Request started!")
	defer log.Println("Request finished!")
	select {
	case <-time.After(5*time.Second):
		log.Println("Request successfully!")
		w.Write([]byte("Request successfully!!!"))
	case  <-ctx.Done():
    log.Println("Request cancelled for client!!!")
		http.Error(w, "Request cancelled for client!", http.StatusRequestTimeout)
	}
}