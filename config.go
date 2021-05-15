/*
   Copyright (C) 2021 Wong Wei Lok <wongweilok@disroot.org>

   This file is part of aputt.

   aputt is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   aputt is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with aputt.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"bufio"
	"os"
)

const (
	fileDir  string = "/aputt/"
	fileName string = "config"
)

var (
	configDir, _        = os.UserConfigDir()
	configPath   string = configDir + fileDir
)

// Check existence of config directory
func checkConfigDir() bool {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return false
	}
	return true
}

// Create config directory
func createConfigDir() {
	os.Mkdir(configPath, 0755)
}

// Check existence of config file
func checkConfig() bool {
	if _, err := os.Stat(configPath + fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

// Create config file (if not exist) and write into config file
func writeConfig(data string) {
	err := os.WriteFile(configPath+fileName, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

// Read config file
func readConfig() string {
	file, err := os.Open(configPath + fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Scan()

	return sc.Text()
}
