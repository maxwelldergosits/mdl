package main
import (
  "strings"
  "path"
  "os/user"
  "os"
  "path/filepath"
  "errors"
)

func DoFile(url string) error {

  config, err := LoadDefaultConfig()

  if (err != nil) {
    return err
  }

  ext := GetExt(url)

  fileName, err := GetFileName(url)

  if (err != nil) {
    return err
  }

  folder := GetFolder(config,ext)

  location := folder+"/"+fileName
  location = filepath.Clean(location)

  usr, _ := user.Current()
  dir := usr.HomeDir + "/"


  if location[:2] == "~/" {
      location = strings.Replace(location, "~/", dir, 1)
  }


  dir = filepath.Dir(location)
  os.MkdirAll(dir,0777)

  data, err := GetBytes(url)
  if (err != nil) {
    return err
  }

  fi,err := os.OpenFile(location,os.O_CREATE,0)
  defer fi.Close()
  fi.Write(data)
  return err
}


func GetExt(url string) string {
  return path.Ext(url)
}

func GetFileName(url string) (string,error) {

  splitted := strings.Split(url,"/")
  if(len(splitted) >= 0) {
    return splitted[len(splitted)-1],nil
  } else {
    return "", errors.New("Invalid path")
  }
}


