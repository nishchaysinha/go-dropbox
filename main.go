package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")

	}
	return os.Getenv(key)
}

func main() {

	dotenv := goDotEnvVariable("BEARER_TOKEN")

	url := "https://api.dropboxapi.com/2/team/groups/list"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + dotenv

	// Create a new request using http
	req, err := http.NewRequest("POST", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
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
