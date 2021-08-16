package structs

import "net/http"


type Session struct {
  ResponseWriter http.ResponseWriter
  Request *http.Request
  Path string
  WebContent string
  ContentType string
}
type Config struct {
  Version string           `yaml:"Version"`
  BaseURL string           `yaml:"BaseURL"`
  Domains []string         `yaml:"Domains"`
  LocalDomain string       `yaml:"LocalDomain"`
  InMail Service           `yaml:"InMail"`
  OutMail Service          `yaml:"OutMail"`
  Calendar Service         `yaml:"Calendar"`
  AddressBook Service      `yaml:"AddressBook"`
  WebMail Service          `yaml:"WebMail"`
  OtherServices []Service  `yaml:"OtherServices"`

}
type Service struct {
  Name string             `yaml:"Name"`
  Enabled bool            `yaml:"Enabled"`
  Type string             `yaml:"Type"`
  Server string           `yaml:"Server"`
  Port int                `yaml:"Port"`
  SocketType string       `yaml:"SocketType"`
  SPA bool                `yaml:"SPA"`
  UsernameIsFQDN bool     `yaml:"UsernameIsFQDN"`
  RequireLocalDomain bool `yaml:"RequireLocalDomain"`
  NoAuthRequired bool     `yaml:"NoAuthRequired"`
  Authentication string   `yaml:"Authentication"`
  // For Outgoing Mail
  POPAuth bool            `yaml:"POPAuth"`
  SMTPLast bool           `yaml:"SMTPLast"`
  // For WebMail (Unused)
  UsernameDivID string    `yaml:"UsernameDivID"`
  UsernameDivName string  `yaml:"UsernameDivName"`
  PasswordDivName string  `yaml:"PasswordDivName"`
  SubmitButtonID string   `yaml:"SubmitButtonID"`
  SubmitButtonName string `yaml:"SubmitButtonName"`
}
type Response struct {
  Url string                       `json:"url"`
  ContentType string               `json:"content_type"`
  Message string                   `json:"message"`
  Content map[string]interface{}   `json:"content"`
  Config Config                    `json:"_"`
  Email string                     `json:"_"`
}
type MSAutodiscoverJSONResponse struct {
  // More work to do - handling of MS Autodiscover.json requests
  Protocol string
  Url string
}
type MSAutodiscoverJSONError struct{
  ErrorCode string
  ErrorMessage string
}
