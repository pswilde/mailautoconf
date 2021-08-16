package responses
import (
  "mailautoconf/global"
  . "mailautoconf/structs"
  "html/template"
  // "fmt"
)
func MozAutoconfig() string {
  // The below link has config-v1.1.xml information
  // https://wiki.mozilla.org/Thunderbird:Autoconfiguration:ConfigFileFormat
  tmpl := "templates/autoconfig.html"
  t := template.Must(template.ParseFiles(tmpl))
  data := struct {
        Name string
        Skills []string
    }{
        Name: "John Doe",
        Skills: []string{
            "C++",
            "Java",
            "Python",
        },
    }
  t.Execute(global.ThisSession.ResponseWriter, data)
  return ""
}
func MsAutoDiscoverXML() string {
  return ""
}
func MsAutoDiscoverJSON() string {
  return ""
}
func DefaultResponse() string {
  response := Response{}
  response.Url = global.ThisSession.Path
  global.ThisSession.ContentType = "application/json"
  response.ContentType = global.ThisSession.ContentType
  response.Message = "Success! Things are working! Please request a valid URL i.e. /mail/config-v1.1.xml";
  return global.JSONify(response)
}
func OurConfig() string {
  global.ThisSession.ContentType = "application/json"
  content := global.JSONify(global.MainConfig)
  return content
}
