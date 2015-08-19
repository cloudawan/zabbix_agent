// Copyright 2015 CloudAwan LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"

	"os/exec"
	"strings"
	"time"
)

const (
	bufferSize       = 409600
	checkingInterval = time.Second * 1
)

func main() {
	requestURL, ok := LocalConfiguration.GetString("requestURL")
	if ok != true {
		log.Panicln(requestURL)
		return
	}

	zabbixCommand, ok := LocalConfiguration.GetString("zabbixCommand")
	if ok != true {
		log.Panicln(zabbixCommand)
		return
	}

	commandSlice := strings.Split(zabbixCommand, " ")

	for {
		err := executeLongPolling(requestURL, commandSlice)
		log.Panicln(err)
		time.Sleep(1 * time.Minute)
	}

}

func sendEvent(commandSlice []string) {
	command := exec.Command(commandSlice[0], commandSlice[1:]...)
	out, err := command.CombinedOutput()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(out))
}

func executeLongPolling(requestURL string, commandSliceTemplate []string) error {
	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Panicln(err)
		return err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panicln(err)
		return err
	}
	byteSlice := make([]byte, bufferSize)
	for {
		n, err := response.Body.Read(byteSlice)
		if err != nil {
			log.Panicln(err)
			return err
		}

		commandSlice := append(commandSliceTemplate, string(byteSlice[:n]))
		sendEvent(commandSlice)

		fmt.Println(string(byteSlice[:n]))
		time.Sleep(checkingInterval)
	}

	return nil
}
