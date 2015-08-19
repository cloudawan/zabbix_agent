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
	"github.com/cloudawan/kubernetes_management_utility/configuration"
)

var configurationContent = `
{
	"requestURL": "http://127.0.0.1:8080/api/v1/watch/events",
	"zabbixCommand": "zabbix_sender -z 127.0.0.1 -p 10051 -s node -k trap -o"
}
`

var LocalConfiguration *configuration.Configuration

func init() {
	var err error
	LocalConfiguration, err = configuration.CreateConfiguration("zabbix_agent", configurationContent)
	if err != nil {
		panic(err)
	}
}
