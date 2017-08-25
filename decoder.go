package main


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-08-25


import (
  "bytes"
  "github.com/belfinor/Helium/pack"
)


type Decoder struct {
  data []byte
}


func (d *Decoder) Write( data []byte ) []byte {

  var list []byte

  if d.data == nil {
    list = bytes.Join( [][]byte{ []byte{}, data }, nil )
  } else {
    list = bytes.Join( [][]byte{ d.data, data }, nil )
  }

  res := []byte{}

  size := int16(0)

  for len(list) > 2 {
    if pack.Decode( list, &size ) != nil {
      break
    }
    size = size + 2
    if len(list) > int(size) {
      res = bytes.Join( [][]byte{ res, list[:size] }, nil )
      list = list[size:]
    } else if len(list) == int(size) {
      res = bytes.Join( [][]byte{ res, list }, nil )
      list = []byte{}
    } else {
      break;
    }
  }

  if len(list) > 0 {
    d.data = list
  } else {
    d.data = []byte{}
  }

  return res
}

