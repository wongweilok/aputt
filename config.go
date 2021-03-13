package main

import (
	"bufio"
	"os"
)

const (
	fileDir string = "/aputt/"
	fileName string = "config"
)

var (
	configDir, _ = os.UserConfigDir()
	configPath string = configDir + fileDir
)

func checkConfig() bool {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func createConfig() {
	os.Mkdir(configPath, 0644)
}

func writeConfig(data []byte) {
	err := os.WriteFile(configPath + fileName, data, 0644)
	if err != nil {
		panic(err)
	}
}

func readConfig() string {
	file, err := os.Open(configPath + fileName)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(file)
	sc.Scan()

	return sc.Text()
}
