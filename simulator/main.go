package main

import (
  "fmt"
  route "github.com/franciscoarmando63/imersaoFullCycleDelivery/application/route"
  "github.com/joho/godotenv"
)

func init() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(v...:"error loading .env file")
  }
}

func main() {
  route := route.Route{
    ID: "1",
    ClientID: "1",
  }

  route.LoadPositions()
  stringjson, _ := route.ExportJsonPositions()
  fmt.Println(stringjson[1])
}