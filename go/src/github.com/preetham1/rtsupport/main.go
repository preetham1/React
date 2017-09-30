package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
)
func main(){

//fmt.Println("Hello")
http.HandleFunc("/",handler)
http.ListenAndServe(":4000", nil)

}

func handler(w http.ResponseWriter , r *http.Request){

fmt.Fprintf(w,"Hello from go")


}
