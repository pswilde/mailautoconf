package global
import (
  . "mailautoconf/structs"
  "github.com/vaughan0/go-ini"
  "fmt"
)

// Global variables
var ThisSession Session
var MainConfig Config
const defaultConfigDir string = "default-config/"
const configDir string = "config/"

func NewConfig() Config {
  fmt.Println("Loading Config…")
  cfg := "default-config/config.default.ini"
  conf, err := ini.LoadFile(cfg)
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(conf)
  fmt.Println("Loading Services…")
  srv := "default-config/services.default.ini"
  serv, err2 := ini.LoadFile(srv)
  if err2 != nil {
    fmt.Println(err2.Error())
  }
  fmt.Println(serv)
  newcfg := Config{}
  return newcfg
}
