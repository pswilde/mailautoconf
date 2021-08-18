package structs

import "net/http"


type Session struct {
  ID string
  IP string
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
  InMail Service           `yaml:"InMail" json:",omitempty"`
  OutMail Service          `yaml:"OutMail" json:",omitempty"`
  Calendar Service         `yaml:"Calendar" json:",omitempty"`
  AddressBook Service      `yaml:"AddressBook" json:",omitempty"`
  WebMail Service          `yaml:"WebMail" json:",omitempty"`
  OtherServices []Service  `yaml:"OtherServices" json:",omitempty"`

}
type Service struct {
  Name string             `yaml:"Name" json:",omitempty"`
  Enabled bool            `yaml:"Enabled" json:",omitempty"`
  Type string             `yaml:"Type" json:",omitempty"`
  Server string           `yaml:"Server" json:",omitempty"`
  Port int                `yaml:"Port" json:",omitempty"`
  SocketType string       `yaml:"SocketType" json:",omitempty"`
  SPA bool                `yaml:"SPA" json:",omitempty"`
  UsernameIsFQDN bool     `yaml:"UsernameIsFQDN" json:",omitempty"`
  RequireLocalDomain bool `yaml:"RequireLocalDomain" json:",omitempty"`
  NoAuthRequired bool     `yaml:"NoAuthRequired" json:",omitempty"`
  Authentication string   `yaml:"Authentication" json:",omitempty"`
  // For Outgoing Mail
  POPAuth bool            `yaml:"POPAuth" json:",omitempty"`
  SMTPLast bool           `yaml:"SMTPLast" json:",omitempty"`
  // For WebMail (Unused)
  UsernameDivID string    `yaml:"UsernameDivID" json:",omitempty"`
  UsernameDivName string  `yaml:"UsernameDivName" json:",omitempty"`
  PasswordDivName string  `yaml:"PasswordDivName" json:",omitempty"`
  SubmitButtonID string   `yaml:"SubmitButtonID" json:",omitempty"`
  SubmitButtonName string `yaml:"SubmitButtonName" json:",omitempty"`
}
type Response struct {
  Url string                       `json:"url"`
  ContentType string               `json:"content_type"`
  Message string                   `json:"message"`
  Content map[string]interface{}              `json:"content,omitempty"`
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
