package app

import (
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net"
	"os"
	"skd_server/app"
	"strconv"
	"strings"
)

type Config struct {
	ServerUrl    string
	BotToken     string
	Chat         string
	DelayTimeout int
	Version      string
	MachineId    string
	IPAddress    string
}

var conf *Config

func NewAppConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	var delay, _ = strconv.Atoi(getEnv("REQUEST_DELAY", "5"))
	conf = &Config{
		ServerUrl:    getEnv("SERVER_URL", ""),
		BotToken:     getEnv("BOT_TOKEN", ""),
		Chat:         getEnv("CHAT_ID", ""),
		DelayTimeout: delay,
		Version:      "0.1",
	}
	mId, err := machineID()
	if err != nil {
		conf.MachineId = err.Error()
	} else {
		conf.MachineId = mId
	}
	conf.IPAddress = getDeviceIp()
	return conf
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func machineID() (string, error) {
	id, err := ioutil.ReadFile("/var/lib/dbus/machine-id")
	if err != nil {
		id, err = ioutil.ReadFile("/etc/machine-id")
	}
	if err != nil {
		app.LogMessage("Error get machine id", app.Critical, []string{err.Error()})
		return "", err
	}
	return strings.TrimSpace(strings.Trim(string(id), "\n")), nil
}

func getDeviceIp() string {
	addressList, err := net.InterfaceAddrs()
	if err != nil {
		app.LogMessage("Error get ip address", app.Critical, []string{err.Error()})
		return err.Error()
	}

	var ipList []string
	for _, a := range addressList {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() == nil {
				continue
			}
			ipList = append(ipList, ipNet.IP.String())
		}
	}
	return strings.Join(ipList[:], ",")
}
