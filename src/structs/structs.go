package structs
// I don't like the name of this package, consider naming it "core" or
// separating out the structs and core functions

import "net/http"

type Session struct {
  ResponseWriter http.ResponseWriter
  Request *http.Request
  Path string
  WebContent string
}
type Config struct {
  Services []Service
}
type Service struct {

}
