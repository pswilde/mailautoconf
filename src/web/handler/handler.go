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
  ThisSession = Session{}
  ThisSession.ResponseWriter = w
  ThisSession.Request = r
  ThisSession.Path = strings.ToLower(r.URL.Path[1:])

  switch ThisSession.Path {
  case "mail/config-v1.1.xml":
    ThisSession.WebContent = responses.MozAutoconfig()
  case "autodiscover/autodiscover.xml":
    ThisSession.WebContent = responses.MsAutoDiscoverXML()
  case "autodiscover/autodiscover.json":
    ThisSession.WebContent = responses.MsAutoDiscoverJSON()
  default:
    ThisSession.WebContent = responses.DefaultResponse()
  }

  writeWebOutput()
}

func writeWebOutput () {
  fmt.Fprintf(ThisSession.ResponseWriter, ThisSession.WebContent)
}
