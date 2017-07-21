package main


import (
  "github.com/belfinor/Helium/daemon"
  "github.com/belfinor/Helium/log"
  "encoding/json"
  "io/ioutil"
  "strconv"
  "strings"
)


type Config struct {
  Daemon  daemon.Config `json:"daemon"`
  Log     log.Config `json:"log"`
  Server  struct {
    Port   int `json:"port"`
    Host   string `json:"host"`
  } `json:"server"`
  Storage struct {
    Path   string `json:"path"`
    Index  string `json:"index"`
    Buffer int    `json:"buffer"`
    Save   int64  `json:"save"`
    Period int64  `json:"period"`
    LogId  int64  `json:"id"`
  } `json:"storage"`
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

  if cfg.Storage.Buffer < 10 {
    cfg.Storage.Buffer = 10
  }

  // load index file
  if data, err = ioutil.ReadFile( cfg.Storage.Index ) ; err != nil {
    cfg.Storage.LogId = 1
    ioutil.WriteFile( cfg.Storage.Index, []byte("1"), 0664 )
  } else {
    str := strings.TrimSpace( string(data) )
    if cfg.Storage.LogId, err = strconv.ParseInt( str, 10, 64 ) ; err != nil {
      return nil, err
    }
    cfg.Storage.LogId++
  }

  conf = &cfg

  return conf, nil
}


func GetConfig() *Config {
  return conf
}

