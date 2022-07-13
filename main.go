package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

const api_token="00uIquHxh7sJ5PMR24Gc3H_aszSraxYe2mgGzBuTs6"
const instance_uri="https://trial-7480774.okta.com/"

func main() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", instance_uri + "/api/v1/groups?q=PROD&limit=10", nil))
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Authorization", "SSWS "+api_token)
	client := &http.Client{}
	response, err := client.Do(req)

    if err != nil {
        fmt.Print(err.Error())
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}