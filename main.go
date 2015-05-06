package main

import (
  "net/http"
  "os"
)

func main() {
  var port string
  if len(os.Args) > 1 {
    port = os.Args[1]
  } else {
    port = "8000"
  }

  println("Serving files in the current directory on port " + port)
  http.Handle("/", http.FileServer(http.Dir(".")))
  err := http.ListenAndServe(":" + port, nil)
  if err != nil {
    println("ListenAndServe: ", err)
  }
}
