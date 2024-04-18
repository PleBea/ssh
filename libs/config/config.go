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

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func Create(connection Connection) {
	configPath := GetConfigPath()

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		os.Create(configPath)
	}

	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, 0644)
	handleError(err)

	defer file.Close()

	_, err = file.WriteString("\n")
	handleError(err)

	_, err = file.WriteString("Host " + connection.Host + "\n")
	handleError(err)

	_, err = file.WriteString("  HostName " + connection.HostName + "\n")
	handleError(err)

	_, err = file.WriteString("  User " + connection.User + "\n")
	handleError(err)

	_, err = file.WriteString("  Port " + connection.Port + "\n")
	handleError(err)

	file.Sync()

	return
}
