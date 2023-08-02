package main 

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
	"log"
  "net/http"
  "time"
)

func basic_get(url string, resp_type string, header string, token string, debug bool) string {
  c := http.Client{Timeout: time.Duration(1) * time.Second}
  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    log.Fatalln(err)
  }

  req.Header.Add("Accept", resp_type)
  req.Header.Set(header, token)
  if debug {
    log.Printf("%v %v", req.Method, req.URL)
  }
  resp, err := c.Do(req)
  if err != nil {
    log.Fatalln(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  sb := string(body)
  return sb
}

func basic_post(url string, resp_type string, data map[string]string) string {
  //Encode the data
  json_data, err := json.Marshal(data)
  if err != nil {
    log.Fatalln(err)
  }

  responseBody := bytes.NewBuffer(json_data)
  //Leverage Go's HTTP Post function to make request
  resp, err := http.Post(url, resp_type, responseBody)
  //Handle Error
  if err != nil {
     log.Fatalf("An Error Occured %v", err)
  }
  defer resp.Body.Close()
  //Read the response body
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
     log.Fatalln(err)
  }
  sb := string(body)
  log.Printf(sb)
  return sb
}

