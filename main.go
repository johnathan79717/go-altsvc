package main

import (
  "net/http"
  "io"
)

func handle(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "OK");
}

func main() {
  http.HandleFunc("/", handle);
  go http.ListenAndServe(":8080", nil)
  http.ListenAndServe(":8081", nil)
}
