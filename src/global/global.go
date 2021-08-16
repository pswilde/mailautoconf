package global
import (
  . "mailautoconf/structs"
  "fmt"
  "github.com/pelletier/go-toml/v2"
  "io/ioutil"
  "os"

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
  cfgfile := defaultConfigDir + "config.default.toml"
  unmarshalConfig(cfgfile, &cfg)
  fmt.Println(cfg)
  customcfgfile := configDir + "config.toml"
  unmarshalConfig(customcfgfile, &cfg)
  fmt.Println(cfg)
  svcfile := defaultConfigDir + "services.default.toml"
  // cfg.Services = []Service{
  //   Service{
  //     Name : "first",
  //   },
  //   Service{
  //     Name : "second",
  //   },
  //   Service{
  //     Name : "third",
  //   },
  // }
  // data, _ := toml.Marshal(cfg)

  // ioutil.WriteFile(svcfile, data, 0)

  unmarshalServices(svcfile, &cfg)

  customsvcfile := configDir + "services.toml"
  unmarshalServices(customsvcfile, &cfg)
  // fmt.Println(cfg)
  fmt.Println("\r\nOur Config :")
  fmt.Println(cfg)
  return cfg
}

func unmarshalConfig(file string, cfg *Config)  {
  if fileExists(file) {
    content, err := ioutil.ReadFile(file)
    if err != nil {
      fmt.Println("Error reading config :", file, " : ", err)
    }
    err2 := toml.Unmarshal(content, &cfg)
    if err2 != nil {
      fmt.Println("Error unmarshalling config :", file, " : ", err2)
    }
  }
}

func unmarshalServices(file string, cfg *Config)  {
  if fileExists(file) {
    content, err := ioutil.ReadFile(file)
    if err != nil {
      fmt.Println("Error reading services :", file, " : ", err)
    }
    customsvcfile := configDir + "services.toml"
    content2, err2 := ioutil.ReadFile(file)
    if err2 != nil {
      fmt.Println("Error reading services :", customsvcfile, " : ", err2)
    }
    content = []byte(fmt.Sprintf(string(content),string(content2)))
    var x map[string]interface{}
    err3 := toml.Unmarshal(content, &x)
    if err3 != nil {
      fmt.Println("Error unmarshalling services :", file, " : ", err3)
    }
    fmt.Println(x)
  }
}

func fileExists(file string) bool {
  exists := false
  if _, err := os.Stat(file); err == nil {
    exists = true
  } else {
    fmt.Println(err)
    fmt.Printf("File %s does not exist\n", file);
  }
  return exists
}
