package main


import (
  "bytes"
  "github.com/belfinor/Helium/pack"
)


type Stream struct {
    data []byte
}


func (s *Stream) Write( src []byte ) {
  s.data = bytes.Join( [][]byte{  s.data, src }, []byte{} )
  size := int16(0)
  
  list := s.data

  for len(list) > 2 {
    if pack.Decode( list, &size ) != nil {
      break
    }
    size = size + 2
    if len(list) > int(size) {
      ST.Push( list[:size] )
      list = list[size:]
    } else if len(list) == int(size) {
      ST.Push( list )
      list = []byte{}
    } else {
      break;
    }
  }

  if len(list) > 0 {
    s.data = list
  } else {
    s.data = []byte{}
  }
}

