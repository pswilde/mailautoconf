package main

import (
	"fmt"
  "net/http"
  "log"
	"mailautoconf/web/handler"
	"mailautoconf/global"
)

func main() {
	global.NewConfig()
  http.HandleFunc("/", handler.WebHandler)
	fmt.Println("Starting up Web Listener on port 8010")
  log.Fatal(http.ListenAndServe(":8010", nil))
}
