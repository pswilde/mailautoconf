package structs

import "net/http"


type Session struct {
  ResponseWriter http.ResponseWriter
  Request *http.Request
  Path string
  WebContent string
}
type Config struct {
  BaseURL string
  Domains []string
  LogonDomain string
  Services []interface{}
}
type Service struct {
  Name string
  Enabled bool
  Type string
  Server string
  Port int
  SocketType string
  SPA bool
  UsernameIsFQDN bool
  NoAuthRequired bool
  Authentication string
  // For Outgoing Mail
  POPAuth bool
  SMTPLast bool
  // For WebMail (Unused)
  UsernameDivID string
  UsernameDivName string
  PasswordDivName string
  SubmitButtonID string
  SubmitButtonName string
}
