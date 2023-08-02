package main

import (
  "bytes"
  "encoding/json"
  "fmt"
	"log"
  "os"
)

type Artifactory struct {
	Url   string
	User  string
	Password string
  Header string
  ResponseType string
}

var art = Artifactory{}

var storage Storage

func Initialize() {
	log.Println("Initializeing")
  art.Url = os.Getenv("ARTIFACTORY_URL")
  art.User = os.Getenv("ARTIFACTORY_USER")
  art.Password = os.Getenv("ARTIFACTORY_API")
  art.Header = "X-Jfrog-Art-Api"
  art.ResponseType = "application/json"
  log.Println(art)
}

func PrettyJson(data string) string {
  var out bytes.Buffer
  err := json.Indent(&out, []byte(data), "", "  ")
  if err != nil {
    fmt.Println(err)
  }

  return string(out.Bytes())
}

func GetRepositories(debug bool) Storage {
  url := fmt.Sprintf("%s/api/storageinfo", art.Url)
  var data = basic_get(url, art.ResponseType, art.Header, art.Password, debug)
  log.Println(PrettyJson(data))

  err := json.Unmarshal([]byte(data), &storage)
  if err != nil {
    fmt.Println(err)
    panic(err)
  }

  return storage
}
