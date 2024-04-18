package config

import (
	"os"
	"runtime"
	"strings"
)

type Connection struct {
	Host         string
	HostName     string
	User         string
	Port         string
	IdentityFile string
}

func GetConfigPath() string {
	var configPath string
	var USER string = os.Getenv("USER")

	switch runtime.GOOS {
	case "darwin":
		configPath = "/Users/" + USER + "/.ssh/config"
	case "linux":
		configPath = "/home/" + USER + "/.ssh/config"
	case "windows":
		configPath = "C:\\Users\\" + USER + "\\.ssh\\config"
	}

	return configPath
}

func Parse(data []byte) []Connection {
	var connections []Connection

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		splitedLine := strings.Split(line, " ")

		key, value := splitedLine[0], splitedLine[1]

		switch key {
		case "Host":
			connections = append(connections, Connection{Host: value})
		case "HostName":
			connections[len(connections)-1].HostName = value
		case "User":
			connections[len(connections)-1].User = value
		case "Port":
			connections[len(connections)-1].Port = value
		case "IdentityFile":
			connections[len(connections)-1].IdentityFile = value
		default:
			continue
		}
	}

	return connections
}

func Get() []Connection {
	configPath := GetConfigPath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}

		os.Create(configPath)
	}

	return Parse(data)
}

func Create() {
	configPath := GetConfigPath()

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		os.Create(configPath)
	}

}
