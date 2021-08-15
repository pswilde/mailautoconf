package main

import (
	"fmt"
  "net/http"
  "log"
	"mailautoconf/web/handler"
	. "mailautoconf/global"
)

func main() {
	MainConfig = NewConfig()
  http.HandleFunc("/", handler.WebHandler)
	fmt.Println("Starting up Web Listener on port 8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
