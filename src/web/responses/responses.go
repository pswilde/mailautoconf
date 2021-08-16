package responses
import (
  "mailautoconf/global"
  . "mailautoconf/structs"
  "text/template"
  "fmt"
  "path"
  "strings"
  "bytes"
  "regexp"
)
var email string
var fmap = template.FuncMap{
    "lower": strings.ToLower,
    "parseUsername": parseUsername,
    "onoff": chooseOnOff,
  }
func MozAutoconfig() string {
  // The below link has config-v1.1.xml information
  // https://wiki.mozilla.org/Thunderbird:Autoconfiguration:ConfigFileFormat
  tmpl := "templates/autoconfig.xml"
  response := Response{}
  response.Email = global.ThisSession.Request.FormValue("emailaddress")
  email = response.Email
  response.Config = global.MainConfig

  name := path.Base(tmpl)
  t, err1 := template.New(name).Funcs(fmap).ParseFiles(tmpl)
  if err1 != nil {
    panic (err1)
  }
  global.ThisSession.ContentType = "application/xml"
  var result bytes.Buffer
  err := t.Execute(&result, response)
  if err != nil {
    fmt.Println(err)
  }
  return result.String()
}
func MsAutoDiscoverXML() string {
  // Example POST Request (sent from client) :
  // <?xml version="1.0" \?\>
  // <Autodiscover xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://schemas.microsoft.com/exchange/autodiscover/outlook/requestschema/2006">
  //   <Request>
  //     <EMailAddress>your@email.address</EMailAddress>
  //     <AcceptableResponseSchema>http://schemas.microsoft.com/exchange/autodiscover/outlook/responseschema/2006a</AcceptableResponseSchema>
  //   </Request>
  // </Autodiscover>
  tmpl := "templates/autodiscover.xml"
  email = global.ThisSession.Request.FormValue("EMailAddress")
  response := Response{}
  response.Config = global.MainConfig
  name := path.Base(tmpl)
  t, err1 := template.New(name).Funcs(fmap).ParseFiles(tmpl)
  if err1 != nil {
    panic (err1)
  }
  global.ThisSession.ContentType = "application/xml"
  var result bytes.Buffer
  err := t.Execute(&result, response)
  if err != nil {
    fmt.Println(err)
  }
  return result.String()
}
func MsAutoDiscoverJSON() string {
  // Example Request
  // /autodiscover/autodiscover.json?Email=you@your.domain&Protocol=Autodiscoverv1&RedirectCount=1
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
func parseUsername(svc Service) string {
  if email == "" {
    return "not-provided"
  }
  if svc.UsernameIsFQDN && !svc.RequireLocalDomain{
    return email
  } else if svc.UsernameIsFQDN && svc.RequireLocalDomain {
    re := regexp.MustCompile(`[^@(%40)]+$`)
    domain := re.FindString(email)
    localemail := strings.Replace(email, domain,
                          global.MainConfig.LocalDomain,1)
    return localemail
  } else {
    re := regexp.MustCompile(`^[^@(%40)]+`)
    username := re.FindString(email)
    return username
  }
}
func chooseOnOff(value bool) string {
  if value {
    return "on"
  } else {
    return "off"
  }
}
