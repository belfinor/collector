package main


import (
  "fmt"
  "github.com/belfinor/Helium/log"
  "net"
  "strconv"
  "time"
)


type Server struct {
  Host string
  Port int
}


func (s *Server) Start() {
  ln, err := net.Listen( "tcp", s.Host + ":" + strconv.Itoa(s.Port) )
	
  if err != nil {
    log.Error( "bind port error" )
    <- time.After( time.Second * 2 )
    panic( "bind port error" )
  }

  log.Info( fmt.Sprintf( "start listen on %s:%d", s.Host, s.Port ) )

  for {
        
    conn, err := ln.Accept() 

    if err != nil {
      continue
    }
			
    go s.handler(conn)
  }
}


func (s *Server) handler(conn net.Conn) {
  defer conn.Close()
  defer log.Info( "connection closed" )

  st := ST.Inst()

  log.Info( "income connection" )

   buffer := make( []byte, 4098 )

   for {
     n, err := conn.Read( buffer )
     if err != nil {
       break
     }

     if n > 0 {
       st.Write( buffer[0:n] )
     }
   }
}

