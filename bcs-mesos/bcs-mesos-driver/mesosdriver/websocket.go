/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package mesosdriver

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/websocketDialer"
)

const (
	Module        = "BCS-API-Tunnel-Module"
	RegisterToken = "BCS-API-Tunnel-Token"
	Params        = "BCS-API-Tunnel-Params"
	Cluster       = "BCS-API-Tunnel-ClusterId"
	ModuleName    = "mesos-driver"
)

func (m *MesosDriver) buildWebsocketToApi() error {
	if m.config.RegisterUrl == "" {
		return errors.New("register url is empty")
	}
	bcsApiUrl, err := url.Parse(m.config.RegisterUrl)
	if err != nil {
		return err
	}

	if m.config.RegisterToken == "" {
		return errors.New("register token is empty")
	}
	if m.config.Cluster == "" {
		return errors.New("clusterid is empty")
	}

	var serverAddress string
	if m.config.ServCert.IsSSL {
		serverAddress = fmt.Sprintf("https://%s:%d", m.config.Address, m.config.Port)
	} else {
		serverAddress = fmt.Sprintf("http://%s:%d", m.config.Address, m.config.Port)
	}

	params := map[string]interface{}{
		"address": serverAddress,
	}
	bytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	headers := map[string][]string{
		Module:        {ModuleName},
		Cluster:       {m.config.Cluster},
		RegisterToken: {m.config.RegisterToken},
		Params:        {base64.StdEncoding.EncodeToString(bytes)},
	}

	var tlsConfig *tls.Config
	if m.config.InsecureSkipVerify {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}
	} else {
		// use bcs cacert
		pool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(m.config.ClientCert.CAFile)
		if err != nil {
			return err
		}
		if ok := pool.AppendCertsFromPEM(ca); ok != true {
			return fmt.Errorf("append ca cert failed")
		}
		tlsConfig = &tls.Config{RootCAs: pool}
	}

	go func() {
		for {
			wsURL := fmt.Sprintf("wss://%s/bcsapi/v4/usermanager/v1/websocket/connect", bcsApiUrl.Host)
			blog.Infof("Connecting to %s with token %s", wsURL, m.config.RegisterToken)

			websocketDialer.ClientConnect(context.Background(), wsURL, headers, tlsConfig, nil, func(proto, address string) bool {
				switch proto {
				case "tcp":
					return true
				case "unix":
					return address == "/var/run/docker.sock"
				}
				return false
			})
			time.Sleep(5 * time.Second)
		}
	}()

	return nil
}
