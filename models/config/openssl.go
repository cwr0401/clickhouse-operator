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

	"github.com/cwr0401/clickhouse-operator/utils"
)

type OpenSSLVerificationMode string

const (
	MethodRelaxed OpenSSLVerificationMode = "relaxed"
	MethodStrict  OpenSSLVerificationMode = "strict"
	MethodOnce    OpenSSLVerificationMode = "once"
	MethodNone    OpenSSLVerificationMode = "none"
)

type CertificateHandler string

const (
	ConsoleCertificateHandler CertificateHandler = "ConsoleCertificateHandler"
	AcceptCertificateHandler  CertificateHandler = "AcceptCertificateHandler"
	RejectCertificateHandler  CertificateHandler = "RejectCertificateHandler"
)

type KeyPassphraseHandler string

const (
	KeyConsoleHandler KeyPassphraseHandler = "KeyConsoleHandler"
	KeyFileHandler    KeyPassphraseHandler = "KeyFileHandler"
)

type SSLProtocolName string

const (
	ProtoSSLv2  SSLProtocolName = "sslv2"
	ProtoSSLv3  SSLProtocolName = "sslv3"
	ProtoTLSv1  SSLProtocolName = "tlsv1"
	ProtoTLSv11 SSLProtocolName = "tlsv1_1"
	ProtoTLSv12 SSLProtocolName = "tlsv1_2"
	ProtoTLSv13 SSLProtocolName = "tlsv1_3"
)

type OpenSSLConfig struct {

	// The path to the file with the secret key of the PEM certificate.
	// The file may contain a key and certificate at the same time.
	// openssl req -subj "/CN=localhost" -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout /etc/clickhouse-server/server.key -out /etc/clickhouse-server/server.crt
	PrivateKeyFile string `xml:"privateKeyFile,omitempty"`

	// The path to the client/server certificate file in PEM format.
	// You can omit it if `privateKeyFile` contains the certificate.
	CertificateFile string `xml:"certificateFile,omitempty"`

	// The path to the file or directory that contains trusted CA certificates.
	// If this points to a file, it must be in PEM format and can contain several CA certificates.
	// If this points to a directory, it must contain one .pem file per CA certificate.
	// The filenames are looked up by the CA subject name hash value.
	// Details can be found in the man page of [SSL_CTX_load_verify_locations](https://www.openssl.org/docs/man3.0/man3/SSL_CTX_load_verify_locations.html).
	CAConfig string `xml:"caConfig,omitempty"`

	// The method for checking the node’s certificates.
	// Details are in the description of the [Context](https://github.com/ClickHouse-Extras/poco/blob/master/NetSSL_OpenSSL/include/Poco/Net/Context.h) class.
	// Possible values: `none`, `relaxed`, `strict`, `once`.
	// Default is "relaxed".
	// "None" use for self-signed.
	VerificationMode OpenSSLVerificationMode `xml:"verificationMode,omitempty"`

	// The maximum length of the verification chain. Verification will fail if the certificate chain length exceeds the set value.
	// Default is 9.
	VerificationDepth uint `xml:"verificationDepth,omitempty"`

	// Wether built-in CA certificates for OpenSSL will be used.
	// ClickHouse assumes that builtin CA certificates are in the file `/etc/ssl/cert.pem` (resp. the directory `/etc/ssl/certs`)
	// or in file (resp. directory) specified by the environment variable `SSL_CERT_FILE` (resp. `SSL_CERT_DIR`).
	LoadDefaultCAFile bool `xml:"loadDefaultCAFile,omitempty"`

	// Supported OpenSSL encryptions.
	// Default is `ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH`.
	CipherList string `xml:"cipherList,omitempty"`

	// Enables or disables caching sessions.
	// Must be used in combination with `sessionIdContext`.
	// Acceptable values: `true`, `false`.
	// Default is `false`.
	CacheSessions bool `xml:"cacheSessions,omitempty"`

	// A unique set of random characters that the server appends to each generated identifier.
	// The length of the string must not exceed `SSL_MAX_SSL_SESSION_ID_LENGTH`.
	// This parameter is always recommended since it helps avoid problems both if the server caches the session and
	// if the client requested caching.
	// Default value: `${application.name}`
	SessionIdContext string `xml:"sessionIdContext,omitempty"`

	//
	// The maximum number of sessions that the server caches.
	// A value of 0 means unlimited sessions.
	// Default is [1024\*20](https://github.com/ClickHouse/boringssl/blob/master/include/openssl/ssl.h#L1978).
	SessionCacheSize uint64 `xml:"sessionCacheSize,omitempty"`

	// Time for caching the session on the server.
	// Default is [2h](https://github.com/ClickHouse/boringssl/blob/master/include/openssl/ssl.h#L1926)
	SessionTimeout string `xml:"sessionTimeout,omitempty"`

	// If enabled, verify that the certificate CN or SAN matches the peer hostname.
	// Default is false
	ExtendedVerification bool `xml:"extendedVerification,omitempty"`

	// Require a TLSv1 connection. Acceptable values: `true`, `false`.
	// Default is false.
	RequireTLSv1 bool `xml:"requireTLSv1,omitempty"`

	// Require a TLSv1.1 connection. Acceptable values: `true`, `false`.
	// Default is false.
	RequireTLSv1_1 bool `xml:"requireTlSv11,omitempty"`

	// Require a TLSv1.2 connection. Acceptable values: `true`, `false`.
	// Default is false.
	RequireTLSv1_2 bool `xml:"requireTlSv12,omitempty"`

	// Activates OpenSSL FIPS mode. Supported if the library’s OpenSSL version supports FIPS.
	// Default is false.
	Fips bool `xml:"fips,omitempty"`

	// Class (PrivateKeyPassphraseHandler subclass) that requests the passphrase for accessing the private key.
	// For example: `<privateKeyPassphraseHandler>`, `<name>KeyFileHandler</name>`, `<options><password>test</password></options>`, `</privateKeyPassphraseHandler>`.
	// Default is `KeyConsoleHandler`.
	PrivateKeyPassphraseHandler PrivateKeyPassphraseHandler `xml:"privateKeyPassphraseHandler,omitempty"`

	// Class (a subclass of CertificateHandler) for verifying invalid certificates.
	// For example: `<invalidCertificateHandler> <name>ConsoleCertificateHandler</name> </invalidCertificateHandler>` .
	// Default is `ConsoleCertificateHandler`
	InvalidCertificateHandler InvalidCertificateHandler `xml:"invalidCertificateHandler,omitempty"`

	// Protocols that are not allowed to use.
	// Default is "".
	DisableProtocols string `xml:"disableProtocols,omitempty"`

	// Preferred server ciphers on the client.
	// Default is false.
	PreferServerCiphers bool `xml:"preferServerCiphers,omitempty"`

	// dhparams are optional. You can delete the <dhParamsFile> element.
	// To generate dhparams, use the following command:
	// openssl dhparam -out /etc/clickhouse-server/dhparam.pem 4096
	// Only file format with BEGIN DH PARAMETERS is supported.

	DhParamsFile string `xml:"dhParamsFile,omitempty"`
}

type PrivateKeyPassphraseHandler struct {
	Name    KeyPassphraseHandler  `xml:"name,omitempty"`
	Options utils.MapStringString `xml:"options,omitempty"`
}

type InvalidCertificateHandler struct {
	Name CertificateHandler `xml:"name,omitempty"`
}

type OpenSSL struct {
	XMLName xml.Name `xml:"openSSL"`
	// Server OpenSSL settings.
	Server OpenSSLConfig `xml:"server,omitempty"`

	// Client OpenSSL settings.
	Client OpenSSLConfig `xml:"client,omitempty"`
}
