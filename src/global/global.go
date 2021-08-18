package global
import (
  . "mailautoconf/global/structs"
  "mailautoconf/global/logger"
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  // "os"
  "encoding/json"
  "text/template"
  "path"
  "regexp"
  "strings"
  "time"
  // "reflect"
)

// Global variables
var ThisSession Session
var MainConfig Config

// Template declaration
var templates []string = []string{"autodiscover.xml","autoconfig.xml"}
var Templates map[string]*template.Template = make(map[string]*template.Template)

const defaultConfigDir string = "default-config/"
const configDir string = "config/"

func NewSessionID() string{
  timecode := time.Now()
  id := timecode.Format("20060102150405.000")
  id = strings.Replace(id,".","",1)

  return id
}
func NewConfig() Config {
  MainConfig = loadConfig()
  loadXMLTemplates()
  return MainConfig
}
func loadConfig() Config {
  cfg := Config{}
  logger.Log("Loading Default Config…")
  cfgfile := defaultConfigDir + "config.default.yaml"
  unmarshalConfig(cfgfile, &cfg)
  logger.Log("Loading Custom Config…")
  customcfgfile := configDir + "config.yaml"
  unmarshalConfig(customcfgfile, &cfg)
  logger.Log("Loading Default Services…")
  svcfile := defaultConfigDir + "services.default.yaml"
  unmarshalConfig(svcfile, &cfg)
  logger.Log("Loading Custom Services…")
  customsvcfile := configDir + "services.yaml"
  unmarshalConfig(customsvcfile, &cfg)
  removeDisabledItems(&cfg)
  return cfg
}
func loadXMLTemplates(){
  for _, tmpl := range templates {
    tmpl := fmt.Sprintf("templates/%s",tmpl)
    name := path.Base(tmpl)
    var fmap = template.FuncMap{
        "lower": strings.ToLower,
        "parseUsername": parseUsername,
        "onoff": chooseOnOff,
      }
    t, err := template.New(name).Funcs(fmap).ParseFiles(tmpl)
    logger.CheckError(err)
    Templates[name] = t
  }
}
func unmarshalConfig(file string, cfg *Config)  {
  if logger.FileExists(file) {
    content, err := ioutil.ReadFile(file)
    if !logger.ErrorOK(err){
      logger.Log("Error reading config :", file, " : ", fmt.Sprintf("%v",err))
    }
    err2 := yaml.Unmarshal(content, &cfg)
    if !logger.ErrorOK(err2){
      logger.Log("Error unmarshaling config :", file, " : ", fmt.Sprintf("%v",err2))
    }
  }
}

func removeDisabledItems(cfg *Config) {
  // Rework this, not pretty
  if !cfg.InMail.Enabled {
    cfg.InMail = Service{}
  }
  if !cfg.OutMail.Enabled {
    cfg.OutMail = Service{}
  }
  if !cfg.Calendar.Enabled {
    cfg.Calendar = Service{}
  }
  if !cfg.AddressBook.Enabled {
    cfg.AddressBook = Service{}
  }
  if !cfg.WebMail.Enabled {
    cfg.WebMail = Service{}
  }
  new_svcs := []Service{}
  for _,svc := range cfg.OtherServices {
    if svc.Enabled {
      new_svcs = append(new_svcs,svc)
    }
  }
  cfg.OtherServices = new_svcs
}

func JSONifyConfig(content Config) string {
  data, err := json.Marshal(content)
  logger.CheckError(err)
  return string(data)
}
func JSONify(content interface{}) string {
  data, err := json.Marshal(content)
  logger.CheckError(err)
  return string(data)
}
func parseUsername(svc Service, email string) string {
  if email == "" {
    return "not-provided"
  }
  if svc.UsernameIsFQDN && !svc.RequireLocalDomain{
    return email
  } else if svc.UsernameIsFQDN && svc.RequireLocalDomain {
    re := regexp.MustCompile(`[^@(%40)]+$`)
    domain := re.FindString(email)
    localemail := strings.Replace(email, domain,
                          MainConfig.LocalDomain,1)
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
// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetSessionIP() string {
  r := ThisSession.Request
  ip := r.RemoteAddr
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		ip = forwarded
	}
  logger.Log("Session ", ThisSession.ID, " Connect From : ", ip)
	return ip
}
