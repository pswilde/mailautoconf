package responses
import (
  "mailautoconf/global"
  "mailautoconf/global/logger"
  . "mailautoconf/global/structs"
  // "text/template"
  "fmt"
  // "path"
  "strings"
  "bytes"
  "regexp"
)
var email string

func MozAutoconfig() string {
  // The below link has config-v1.1.xml information
  // https://wiki.mozilla.org/Thunderbird:Autoconfiguration:ConfigFileFormat

  // parse the querystring
  logger.CheckError(global.ThisSession.Request.ParseForm())

  // build the response
  response := Response{}
  response.Email = global.ThisSession.Request.FormValue("emailaddress")
  email = response.Email
  response.Config = global.MainConfig

  // set content type to XML
  global.ThisSession.ContentType = "application/xml"

  // execute the template
  var result bytes.Buffer
  template := global.Templates["autoconfig.xml"]
  err := template.Execute(&result, response)
  logger.CheckError(err)

  // return our string of xml
  return result.String()
}
func MsAutoDiscoverXML() string {
  // MS Outlook Autodiscover.xml
  //
  // Example POST Request (sent from client) :
  // <?xml version="1.0" \?\>
  // <Autodiscover xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://schemas.microsoft.com/exchange/autodiscover/outlook/requestschema/2006">
  //   <Request>
  //     <EMailAddress>your@email.address</EMailAddress>
  //     <AcceptableResponseSchema>http://schemas.microsoft.com/exchange/autodiscover/outlook/responseschema/2006a</AcceptableResponseSchema>
  //   </Request>
  // </Autodiscover>

  // Parse the form to get the values
  logger.CheckError(global.ThisSession.Request.ParseForm())

  // convert the input to a string so we can extract the email address
  form := fmt.Sprintf("%s",global.ThisSession.Request.Form)

  // fine the EMailAddress section
  find := regexp.MustCompile(`\<EMailAddress\>(.*?)\<\/EMailAddress\>`)
  email = find.FindString(form)

  // replace the tags
  replace := regexp.MustCompile(`\<[\/]?EMailAddress\>`)
  email = replace.ReplaceAllString(email,``)

  logger.Log("Session ",global.ThisSession.ID ," Request for email : ",email)
  // build the reponse
  response := Response{}
  response.Email = email
  response.Config = global.MainConfig

  // execute the template
  template := global.Templates["autodiscover.xml"]
  global.ThisSession.ContentType = "application/xml"
  var result bytes.Buffer
  err := template.Execute(&result, response)
  logger.CheckError(err)

  // return our string of xml
  return result.String()
}
func MsAutoDiscoverJSON() string {
  // MS Outlook Autodiscover.json - undocumented
  //
  // Example Request
  // /autodiscover/autodiscover.json?Email=you@your.domain&Protocol=Autodiscoverv1&RedirectCount=1
  email = global.ThisSession.Request.FormValue("Email")
  protocol := global.ThisSession.Request.FormValue("Protocol")
  global.ThisSession.ContentType = "application/json"
  switch strings.ToLower(protocol) {
  case "autodiscoverv1":
    response := MSAutodiscoverJSONResponse{}
    response.Protocol = "AutodiscoverV1"
    response.Url = fmt.Sprintf("%s/Autodiscover/Autodiscover.xml", global.MainConfig.BaseURL)
    return global.JSONify(response)
  default:
    response := MSAutodiscoverJSONError{}
    response.ErrorCode = "InvalidProtocol";
    response.ErrorMessage = fmt.Sprintf("The given protocol value '%s' is invalid. Supported values are 'AutodiscoverV1'", protocol)
    return global.JSONify(response)
  }
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
  content := global.JSONifyConfig(global.MainConfig)
  return content
}
