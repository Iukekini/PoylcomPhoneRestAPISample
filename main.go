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

//Setup Logging
var log = log15.New()

func main() {

	//Setup the Flags from the commandline. 
	ipaddress := flag.String("ip", "", "IP Address of Polycom Phone")
	username := flag.String("username", "Polycom", "API Username")
	password := flag.String("password", "", "Password for API")
	action := flag.String("action", "networkinfo", "Method to call")

	flag.Parse()

	//Check to see if we got an IP Address of a phone. 
	if *ipaddress == "" {
		log.Error("No IPAddress Specified for Phone")
		os.Exit(1)
	}

	//Make sure that a password was also passed in. 
	if *password == "" {
		log.Error("No Password Specified")
		os.Exit(1)
	}

	
	switch *action {
		case "networkinfo":
			//Get the Networking Information from Phone. 
			networkResponse := networkInfo(*ipaddress, *username, *password)
			//Print Response out. Nicely formatted. 
			printObject(networkResponse)
			break
	}

}

//NetworkInfo provides details about the phones's network information. IPAddress/ Vlan / DHCP / ETC
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


// callPhoneAPI calls a given VVX Phones with the creds and path requested. You can pass a body and select what method(GET/POST) will be used to make the call. 
//It returns the response if valid. If not it exits
func callPhoneAPI(url string, username string, password string, body string, method string) *http.Response {

	//Skip validation of cert. Per Polycom Insturctions. 
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := new(http.Client)

	resp, err := http.NewRequest(method, url, bytes.NewBufferString(body))

	//Add Headers
	resp.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	resp.Header.Add("Accept", "application/json")

	//Make call to phone. 
	response, err := client.Do(resp)


	//Check for an error making the call. Timeouts are caught here. 
	if err != nil {
		log.Error("Error getting response: ", err.Error())
		os.Exit(1)
	}

	//Checks to make sure we got a successful call. If not display the message received from the phone and exit. 
	if response.StatusCode != 200 {
		log.Error(strconv.Itoa(response.StatusCode) + ":" + response.Status)
		os.Exit(1)
	}
	


	return response

}


/// printOject takes a struct and geneates a nicely formated json file. This is used to display the informatino to the user. 
func printObject(obj interface {}){
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1);
	}
	os.Stdout.Write(b)

}

/// basicAuth generates the value for the basic auth header. 
func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
