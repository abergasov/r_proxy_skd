package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"r_proxy_skd/app"
	"strconv"
	"time"
)

var conf = app.NewAppConfig()
var client = http.Client{
	Timeout: time.Duration(conf.DelayTimeout) * time.Second,
}

/**
Manage controller request
step 1: check local storage
step 2: send request to cloud
step 3: if fail or cloud storage not answer - return localCheck result
*/
func proxyRequest(body []byte, w http.ResponseWriter) {
	//step 1: check local storage
	response := app.CatchMessage(string(body))
	resp, err := json.Marshal(response)
	if err != nil {
		app.LogMessage("Invalid message object", app.Critical, []string{err.Error()})
		return
	}
	for _, message := range response.Messages {
		if message.Operation != "check_access" {
			continue
		}
		if *message.Granted == 1 {
			break
		}
		//step 2: check cloud storage
		answer, err := client.Post(conf.ServerUrl, "application/json", bytes.NewBuffer(body))
		//if server return response
		if err == nil {
			response, _ := ioutil.ReadAll(answer.Body)
			_, _ = w.Write(response)
			return
		}
		//cloud not answering more than delay - return first response from local storage
		_, _ = w.Write(resp)
		return
	}
	_, _ = w.Write(resp)
	return
}

func main() {
	http.HandleFunc("/listen/command", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			app.LogMessage("Error reading body", app.Warning, []string{err.Error()})
			http.Error(w, "wrong request", http.StatusBadRequest)
			return
		}
		proxyRequest(body, w)
	})

	go app.LogMessage("Proxy server starting", app.Info, []string{
		"request proxy at " + conf.ServerUrl,
		"delay seconds: " + strconv.Itoa(conf.DelayTimeout),
		"machine id: " + conf.MachineId,
		"ip: " + conf.IPAddress,
		"listen port: " + conf.ListenPort,
	})
	err := http.ListenAndServe(":"+conf.ListenPort, nil)
	if err != nil {
		app.LogMessage("Proxy failed, error here", app.Critical, []string{err.Error()})
	}
}
