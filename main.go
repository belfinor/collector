package main


import (
  "flag"
  "github.com/belfinor/Helium/daemon"
  "github.com/belfinor/Helium/io/stream/writer"
  "github.com/belfinor/Helium/log"
)


var ST *writer.Stream


func main() {
    conf   := ""
    is_daemon := false

    flag.StringVar( &conf, "c", "collector.json", "config file name" )
    flag.BoolVar( &is_daemon, "d", false, "run as daemon" )        

    flag.Parse()

    cfg, err := LoadConfig( conf )
    if err != nil {
        panic( err )
    }

    if is_daemon {
        daemon.Run( &cfg.Daemon )
    }
   
    log.Init( &cfg.Log )
    
    log.Info( "collector start" )

    ST = writer.New( &cfg.Storage )

    server := &Server{ Host: cfg.Server.Host, Port: cfg.Server.Port }
    server.Start()
}

