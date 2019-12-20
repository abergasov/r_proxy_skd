package main

import (
	"net/http"
	"r_proxy_skd/app"
	"strconv"
)

var conf = app.NewAppConfig()

func main() {
	http.HandleFunc("/listen/command", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{ok: true}"))
	})
	go app.LogMessage("Proxy server started", app.Info, []string{
		"request proxy at " + conf.ServerUrl,
		"delay: " + strconv.Itoa(conf.DelayTimeout),
		"machine id: " + conf.MachineId,
		"ip: " + conf.IPAddress,
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		go app.LogMessage("Proxy start error", app.Critical, []string{err.Error()})
	}
}
