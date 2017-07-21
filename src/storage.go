package main


import (
  "fmt"
  "github.com/belfinor/Helium/log"
  "io/ioutil"
  "os"
  "time"
)


type Storage struct {
    Input chan []byte
    File  *os.File
}


func InitStorage() *Storage {
  cfg := GetConfig()

  st := &Storage {
    Input: make( chan []byte, cfg.Storage.Buffer ),
  }

  st.openLog()

  log.Info( fmt.Sprintf( "storage current file=%s/%d", cfg.Storage.Path, cfg.Storage.LogId ) )
  log.Info( fmt.Sprintf( "storage buffer size=%d", cfg.Storage.Buffer ) )

  return st
}


func (s *Storage) openLog() {
  cfg := GetConfig()

  file_name := fmt.Sprintf( "%s/%d", cfg.Storage.Path, cfg.Storage.LogId )
  var err error

  if s.File != nil {
    s.File.Close()
  }

  if s.File, err = os.OpenFile( file_name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0664 ) ; err != nil {
    log.Error( err.Error() )
    panic(err)
  }

  ioutil.WriteFile( cfg.Storage.Index, []byte( fmt.Sprintf( "%d", cfg.Storage.LogId ) ), 664 )

  log.Info( "start log " + file_name )
}


func (s *Storage) Rotate() {

  cfg := GetConfig()

  cfg.Storage.LogId++

  s.openLog()

  remove_name := fmt.Sprintf( "%s/%d", cfg.Storage.Path, cfg.Storage.LogId - cfg.Storage.Save )
  os.Remove( remove_name )
}


func (s *Storage) Push( data []byte ) {
  if data == nil ||  len(data) <= 2 {
    return  
  }

  block := make( []byte, len(data) )
  copy( block, data )

  s.Input <- block
}


func (s *Storage) Writer() {
  log.Info( "start storage writer" )

  start  := time.Now().Unix()
  last   := start

  cfg := GetConfig()

  period := cfg.Storage.Period

  for {
    select {
      case data := <- s.Input:
        if _, err := s.File.Write( data ) ; err != nil {
          panic(err)
        }
      case <- time.After( time.Second ):
    }

    last = time.Now().Unix()

    if last - start >= period {
      s.Rotate()
      start = last
    }
  }
}

