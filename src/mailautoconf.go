package main

import (
	// "fmt"
  "net/http"
  // "log"
	"mailautoconf/web/handler"
	"mailautoconf/global"
	"mailautoconf/global/logger"
)

func main() {
	global.NewConfig()
  http.HandleFunc("/", handler.WebHandler)
	logger.Log("Starting up Web Listener on port 8010")
  logger.Fatal(http.ListenAndServe(":8010", nil))
}
