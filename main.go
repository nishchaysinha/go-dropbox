package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://api.dropboxapi.com/2/team/groups/list"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + "sl.BYLvpogW6WMascdGy5TV2QxND8DZRGSXYx6jXMe7CtIexh73MNw9xGh5jjrhkrCLfjV2dY7TZ8Rvz8zhPZEsHKdv5gyFHvWchalgyaSFM0NsgsGn8splzmm_MZNJGFNI5gSFwtnI2bXO9om5t0hG"

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
