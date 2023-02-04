package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadDotenv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")

	}
	return os.Getenv(key)
}

func wrapper_emptyBody(token string, method string, endpoint string) {
	var bearer = "Bearer " + token
	base_url := "https://api.dropboxapi.com/2/"
	url := base_url + endpoint
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))

}

func wrapper_fullBody(token string, method string, endpoint string, reqbody []byte) {
	var bearer = "Bearer " + token

	base_url := "https://api.dropboxapi.com/2/"
	url := base_url + endpoint
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(reqbody))
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))

}

func main() {

	postBody, _ := json.Marshal(map[string]string{
		"include_removed": "false",
		"limit":           "100",
	})

	dotenv := loadDotenv("BEARER_TOKEN")

	wrapper_emptyBody(dotenv, "POST", "team/groups/list")
	println("\n")
	wrapper_emptyBody(dotenv, "POST", "team/get_info")
	println("\n")
	wrapper_fullBody(dotenv, "POST", "team/members/list_v2", postBody)

}
