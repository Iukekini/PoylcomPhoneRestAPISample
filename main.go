package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"gopkg.in/inconshreveable/log15.v2"
)

var log = log15.New()

func main() {

	ipaddress := flag.String("ip", "", "IP Address of Polycom Phone")
	username := flag.String("username", "Polycom", "API Username")
	password := flag.String("password", "", "Password for API")
	action := flag.String("action", "networkinfo", "Method to call")

	flag.Parse()

	if *ipaddress == "" {
		log.Error("No IPAddress Specified for Phone")
		os.Exit(1)
	}
	if *password == "" {
		log.Error("No Password Specified")
		os.Exit(1)
	}

	switch *action {
	case "networkinfo":
		networkResponse := networkInfo(*ipaddress, *username, *password)
		b, err := json.MarshalIndent(networkResponse, "", "  ")
		if err != nil {
			log.Error(err.Error())
			os.Exit(1);
		}
		os.Stdout.Write(b)
		break
	}

}

//NetworkInfo provides details about the phones's network information
func networkInfo(phoneIP string, username string, password string) NetworkInfoResponse {
	var path = "https://" + phoneIP + "/api/v1/mgmt/network/info"
	log.Info("Attempting to connect to phone via:", path, nil)

	var httpResponse = callPhoneAPI(path, username, password, "", "GET")

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}
	log.Info("Successfully called phone API")

	var response NetworkInfoResponse
	json.Unmarshal(body, &response)
	return response

}

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	return nil
}

func callPhoneAPI(url string, username string, password string, body string, method string) *http.Response {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := new(http.Client)

	resp, err := http.NewRequest(method, url, bytes.NewBufferString(body))

	resp.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	resp.Header.Add("Accept", "application/json")

	response, err := client.Do(resp)

	if err != nil {
		log.Error("Error getting response: ", err.Error())
	}
	if response.StatusCode != 200 {
		log.Error(strconv.Itoa(response.StatusCode) + ":" + response.Status)
		os.Exit(1)
	}
	


	return response

}
