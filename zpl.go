package main

import (
  "bytes"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "net"
  "os"
)

func main() {
  var (
    address      string
    conn         net.Conn
    label        string
    labelBytes   []byte
    payloadBytes []byte
    respBytes    []byte
    err          error
  )

  flag.StringVar(&address, "address", "localhost:11382", "Address of zebra proxy printer")

  flag.Parse()

  if labelBytes, err = ioutil.ReadAll(os.Stdin); err != nil {
    log.Fatal(err)
  }

  label = string(labelBytes)

  if conn, err = net.Dial("tcp", address); err != nil {
    log.Fatal(err)
  }

  defer conn.Close()

  payloadBytes = []byte(fmt.Sprintf("%s\r\n\r\n", label))
  if _, err = conn.Write(payloadBytes); err != nil {
    log.Fatal(err)
  }

  if respBytes, err = ioutil.ReadAll(conn); err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(bytes.TrimSpace(respBytes)))
}
