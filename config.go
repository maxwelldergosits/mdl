package main
import (
  "encoding/json"
  "io/ioutil"
  "os/user"
  "os"
)

var default_config string


func init() {
  user, _ := user.Current()
  default_config = user.HomeDir+"/.dlc"
}


type Config struct {
  Default string // Where files with no mapping specified will go
  Extensions map[string]string // Extension[<ext>] = /path/to/*.<ext>
}

func ReadConfig(config []byte) Config {
  var c Config
  json.Unmarshal(config,&c)
  return c
}

func WriteConfig(c Config) []byte {
  out,_ := json.MarshalIndent(c,"","  ")
  return out
}

func WriteConfigToFile(config []byte, fname string){
  ioutil.WriteFile(fname,config,0777)
}

func ReadConfigFromFile(fname string) ([]byte,error) {
  return ioutil.ReadFile(fname)
}

func LoadDefaultConfig() (Config,error) {

  b,e := ReadConfigFromFile(default_config)
  os.Stdout.Write(b)
  if (e != nil) {
    return Config{},e
  }
  c := ReadConfig(b)
  return c,nil
}

func GetFolder(c Config, ext string) string{
  val, ok := c.Extensions[ext]
  if !ok {
    return c.Default
  } else {
    return val
  }

}
