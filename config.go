package main


import (
  "github.com/belfinor/Helium/daemon"
  "github.com/belfinor/Helium/io/stream/writer"
  "github.com/belfinor/Helium/log"
  "encoding/json"
  "io/ioutil"
)


type Config struct {
  Daemon  daemon.Config `json:"daemon"`
  Log     log.Config `json:"log"`
  Server  struct {
    Port   int `json:"port"`
    Host   string `json:"host"`
  } `json:"server"`
  Storage writer.Config `json:"storage"`
}


var conf *Config


func LoadConfig( filename string ) ( *Config, error ) {
  data, err := ioutil.ReadFile( filename )
  if err != nil {
    return nil, err
  }

  var cfg Config

  if err = json.Unmarshal( data, &cfg ) ; err != nil {
    return nil, err
  }

  cfg.Storage.LoadLogId()

  conf = &cfg

  return conf, nil
}


func GetConfig() *Config {
  return conf
}

