package global
import (
  . "mailautoconf/structs"
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "os"
  "encoding/json"

)

// Global variables
var ThisSession Session
var MainConfig Config
const defaultConfigDir string = "default-config/"
const configDir string = "config/"

func NewConfig() Config {
  MainConfig = loadConfig()
  return MainConfig
}
func loadConfig() Config {
  cfg := Config{}
  fmt.Println("Loading Default Config…")
  cfgfile := defaultConfigDir + "config.default.yaml"
  unmarshalConfig(cfgfile, &cfg)
  fmt.Println("Loading Custom Config…")
  customcfgfile := configDir + "config.yaml"
  unmarshalConfig(customcfgfile, &cfg)
  fmt.Println("Loading Default Services…")
  svcfile := defaultConfigDir + "services.default.yaml"
  unmarshalConfig(svcfile, &cfg)
  fmt.Println("Loading Custom Services…")
  customsvcfile := configDir + "services.yaml"
  unmarshalConfig(customsvcfile, &cfg)
  return cfg
}

func unmarshalConfig(file string, cfg *Config)  {
  if FileExists(file) {
    content, err := ioutil.ReadFile(file)
    if err != nil {
      fmt.Println("Error reading config :", file, " : ", err)
    }
    err2 := yaml.Unmarshal(content, &cfg)
    if err2 != nil {
      fmt.Println("Error unmarshalling config :", file, " : ", err2)
    }
  }
}


func FileExists(file string) bool {
  exists := false
  if _, err := os.Stat(file); err == nil {
    exists = true
  } else {
    fmt.Println(err)
    fmt.Printf("File %s does not exist\n", file);
  }
  return exists
}

func JSONify(content interface{}) string {
  data, err := json.Marshal(content)
  if err != nil {
    fmt.Println(err)
  }
  return string(data)
}
