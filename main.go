package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

const api_token = "00uIquHxh7sJ5PMR24Gc3H_aszSraxYe2mgGzBuTs6"
const instance_uri = "https://trial-7480774.okta.com/"

func main() {

	//get PROD Access group ID
	client := &http.Client{}
	req, err := http.NewRequest("GET", instance_uri+"/api/v1/groups?q=Audited&limit=10", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Authorization", "SSWS "+api_token)

	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	groupId := gjson.Get(string(responseData), "#.id|0").String()
	fmt.Println("GroupID: " + groupId)

	//get user by email
	req, err = http.NewRequest("GET", instance_uri+"/api/v1/users?q=janosh.oros@indebted.co&limit=1", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Authorization", "SSWS "+api_token)

	response, err = client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	userId := gjson.Get(string(responseData), "#.id|0").String()
	fmt.Println("UserID: " + userId)

	//add user to group
	req, err = http.NewRequest("PUT", instance_uri+"/api/v1/groups/"+groupId+"/users/"+userId, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Authorization", "SSWS "+api_token)

	response, err = client.Do(req)

	//fmt.Println(req)

	if err != nil {
		fmt.Print(err.Error())
	}

	if response.StatusCode == 204 {
		fmt.Println("User successfully added to the group!")
	} else {
		fmt.Println("Error!" + string(response.StatusCode))
	}
}
