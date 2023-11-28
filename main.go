package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func viewUser(userID string) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", configData.Api.ViewUser, userID), nil)
	if err != nil {
		log.Printf("client: could not create request: %s\n", err)
		return
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
		return
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	log.Printf("result: %s, %d", resBody, res.StatusCode)
}

func main() {
	log.Println("worker started.")
	log.Println("start read config file...")
	readConfigFile("config.yaml")
	data := readDataFile("info.user.csv")
	for {
		row := pickRandomElementIn2dArray(data)
		interval := pickRandomInterval()
		viewUser(row[0])
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
