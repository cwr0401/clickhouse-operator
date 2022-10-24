/*
Copyright 2022 chenweiran.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/cwr0401/clickhouse-operator/utils"
)

func TestEmptyOpenSSL(t *testing.T) {
	var confg = OpenSSL{
		Client: OpenSSLConfig{
			PrivateKeyFile: "/tmp/openssl",
			PrivateKeyPassphraseHandler: PrivateKeyPassphraseHandler{
				Name: KeyFileHandler,
				Options: utils.MapStringString{
					"password": "test",
				},
			},
		},
	}
	output, err := xml.Marshal(confg)
	if err != nil {
		t.Errorf("xml.Marshal failed: %s\n", err)
	}

	fmt.Printf("xml: %s\n", string(output))
}

func TestOpenSSLConfigUnmarshal(t *testing.T) {
	config_string := `
<openSSL>
	<server> <!-- Used for https server AND secure tcp port -->
		<!-- openssl req -subj "/CN=localhost" -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout /etc/clickhouse-server/server.key -out /etc/clickhouse-server/server.crt -->
		<!-- <certificateFile>/etc/clickhouse-server/server.crt</certificateFile>
		<privateKeyFile>/etc/clickhouse-server/server.key</privateKeyFile> -->
		<!-- dhparams are optional. You can delete the <dhParamsFile> element.
			 To generate dhparams, use the following command:
			  openssl dhparam -out /etc/clickhouse-server/dhparam.pem 4096
			 Only file format with BEGIN DH PARAMETERS is supported.
		  -->
		<!-- <dhParamsFile>/etc/clickhouse-server/dhparam.pem</dhParamsFile>-->
		<verificationMode>none</verificationMode>
		<loadDefaultCAFile>true</loadDefaultCAFile>
		<cacheSessions>true</cacheSessions>
		<disableProtocols>sslv2,sslv3</disableProtocols>
		<preferServerCiphers>true</preferServerCiphers>
	</server>

	<client> <!-- Used for connecting to https dictionary source and secured Zookeeper communication -->
		<loadDefaultCAFile>true</loadDefaultCAFile>
		<cacheSessions>true</cacheSessions>
		<disableProtocols>sslv2,sslv3</disableProtocols>
		<preferServerCiphers>true</preferServerCiphers>
		<privateKeyPassphraseHandler>
			<name>KeyFileHandler</name>
			<options>
				<password>test</password>
			</options>
		</privateKeyPassphraseHandler>
		<!-- Use for self-signed: <verificationMode>none</verificationMode> -->
		<invalidCertificateHandler>
			<!-- Use for self-signed: <name>AcceptCertificateHandler</name> -->
			<name>RejectCertificateHandler</name>
		</invalidCertificateHandler>
	</client>
</openSSL>
`
	var config OpenSSL
	err := xml.Unmarshal([]byte(config_string), &config)
	if err == nil {
		fmt.Printf("%v\n", config)
	}

}
