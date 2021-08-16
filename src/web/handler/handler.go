package handler
import (
  . "mailautoconf/structs"
  . "mailautoconf/global"
  "mailautoconf/web/responses"
  "strings"
  "net/http"
  "fmt"
)
func WebHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Request For :",r.URL)
  ThisSession = Session{}
  ThisSession.ResponseWriter = w
  ThisSession.Request = r
  ThisSession.Path = strings.ToLower(r.URL.Path[1:])
  if ThisSession.Path == "" {
    ThisSession.Path = "none"
  }
  switch ThisSession.Path {
  case "mail/config-v1.1.xml",
       "mail/autoconfig.xml":
    ThisSession.WebContent = responses.MozAutoconfig()
  case "autodiscover/autodiscover.xml":
    ThisSession.WebContent = responses.MsAutoDiscoverXML()
  case "autodiscover/autodiscover.json":
    ThisSession.WebContent = responses.MsAutoDiscoverJSON()
  case "get/config":
    ThisSession.WebContent = responses.OurConfig()
  default:
    ThisSession.WebContent = responses.DefaultResponse()
  }

  writeWebOutput()
}

func writeWebOutput () {
  ThisSession.ResponseWriter.Header().Set("Content-Type", ThisSession.ContentType)
  fmt.Fprintf(ThisSession.ResponseWriter, ThisSession.WebContent)
}
