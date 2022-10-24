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

package v1alpha1

import (
	"github.com/cwr0401/clickhouse-operator/models/config"
)

type Logger struct {
	// Logger level.
	// Possible levels: none (turns off logging)、fatal、critical、error、warning、notice、information、debug、trace.
	// +kubebuilder:validation:Enum=none;fatal;critical;error;warning;notice;information;debug;trace
	// +kubebuilder:default=information
	// +optional
	Level config.LoggerLevel `json:"level,omitempty"`

	// Clickhouse Keeper log file path.
	// Default is "/var/log/clickhouse-keeper/clickhouse-keeper.log"
	// +kubebuilder:default=/var/log/clickhouse-keeper/clickhouse-keeper.log
	// +optional
	Log string `json:"log,omitempty"`

	// Default is false.
	// +kubebuilder:default=true
	// +optional
	StreamCompress bool `json:"streamCompress,omitempty"`

	// Clickhouse Keeper error log file path.
	// Default is "/var/log/clickhouse-keeper/clickhouse-keeper.err.log".
	// +kubebuilder:default=/var/log/clickhouse-keeper/clickhouse-keeper.err.log
	// +optional
	ErrorLog string `json:"errorLog,omitempty"`

	// Log rotate file size.
	// Default is 104857600 bytes(100M).
	// +kubebuilder:default=104857600
	// +optional
	Size uint64 `json:"size,omitempty"`

	// Enable log file compress.
	// Default is true.
	// +kubebuilder:default=true
	// +optional
	Compress bool `json:"compress,omitempty"`

	// Log rotate file count.
	// Default is 1.
	// +kubebuilder:default=1
	// +optional
	Count uint `json:"count,omitempty"`

	// Default is true,
	// +kubebuilder:default=true
	// +optional
	Flush bool `json:"flush,omitempty"`

	// Default is false
	// +kubebuilder:default=false
	// +optional
	RotateOnOpen bool `json:"rotateOnOpen,omitempty"`

	// Enable send log to Console.
	// Default is true.
	// +kubebuilder:default=true
	// +optional
	Console bool `json:"console,omitempty"`

	// +optional
	Formatting LoggerFormatting `json:"formatting,omitempty"`

	// +optional
	Levels LoggerLevels `json:"levels,omitempty"`
}

type LoggerLevels struct {
	// +optional
	ConfigReloader string `json:"configReloader,omitempty"`

	// +optional
	Logger []LoggerLevelsLogger `json:"logger,omitempty"`
}

type LoggerLevelsLogger struct {
	Name  string `json:"name,omitempty"`
	Level string `json:"level,omitempty"`
}

type LoggerFormatting struct {
	// Specify log format(for now, JSON only)
	// +kubebuilder:validation:nullable
	// +kubebuilder:validation:Enum=none;json
	// +optional
	Type string `json:"type,omitempty"`

	// +optional
	Names JsonKeyNames `json:"names,omitempty"`
}

type JsonKeyNames struct {
	// DataTime json key name, default is "date_time".
	// +optional
	DateTime string `json:"dateTime,omitempty"`

	// ThreadName json key name, default is "thread_name".
	// +optional
	ThreadName string `json:"threadName,omitempty"`

	// ThreadID json key name, default is "thread_id"
	// +optional
	ThreadID string `json:"threadID,omitempty"`

	// Level json key name, default is "level"
	// +optional
	Level string `json:"level,omitempty"`

	// QueryID json key name, default is "query_id".
	// +optional
	QueryID string `json:"queryID,omitempty"`

	// LoggerName json key name, default is "logger_name".
	// +optional
	LoggerName string `json:"loggerName,omitempty"`

	// Message json key name, default is "message".
	// +optional
	Message string `json:"message,omitempty"`

	// SourceFile json key name, default is "source_file".
	// +optional
	SourceFile string `json:"sourceFile,omitempty"`

	// SourceLine json key name, default is "source_line".
	// +optional
	SourceLine int `json:"sourceLine,omitempty"`
}

type OpenSSLConfig struct {

	// The path to the file with the secret key of the PEM certificate.
	// The file may contain a key and certificate at the same time.
	// +optional
	PrivateKeyFile string `json:"privateKeyFile,omitempty"`

	// The path to the client/server certificate file in PEM format.
	// You can omit it if `privateKeyFile` contains the certificate.
	// +optional
	CertificateFile string `json:"certificateFile,omitempty"`

	// The path to the file or directory that contains trusted CA certificates.
	// If this points to a file, it must be in PEM format and can contain several CA certificates.
	// If this points to a directory, it must contain one .pem file per CA certificate.
	// The filenames are looked up by the CA subject name hash value.
	// Details can be found in the man page of [SSL_CTX_load_verify_locations](https://www.openssl.org/docs/man3.0/man3/SSL_CTX_load_verify_locations.html).
	// +kubebuilder:default=none
	// +optional
	CAConfig string `json:"caConfig,omitempty"`

	// The method for checking the node’s certificates.
	// Details are in the description of the [Context](https://github.com/ClickHouse-Extras/poco/blob/master/NetSSL_OpenSSL/include/Poco/Net/Context.h) class.
	// Possible values: `none`, `relaxed`, `strict`, `once`.
	// Default is "relaxed".
	// +kubebuilder:validation:Enum=none;relaxed;strict;once
	// +kubebuilder;default=relaxed
	// +optional
	VerificationMode string `json:"verificationMode,omitempty"`

	// The maximum length of the verification chain. Verification will fail if the certificate chain length exceeds the set value.
	// Default is 9.
	// +kubebuilder:default=9
	// +optional
	VerificationDepth uint `json:"verificationDepth,omitempty"`

	// Wether built-in CA certificates for OpenSSL will be used.
	// ClickHouse assumes that builtin CA certificates are in the file `/etc/ssl/cert.pem` (resp. the directory `/etc/ssl/certs`)
	// or in file (resp. directory) specified by the environment variable `SSL_CERT_FILE` (resp. `SSL_CERT_DIR`).
	// +kubebuilder:default=true
	// +optional
	LoadDefaultCAFile bool `json:"loadDefaultCAFile,omitempty"`

	// Supported OpenSSL encryptions.
	// Default is `ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH`.
	// +kubebuilder:default=`ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH`
	// +optional
	CipherList string `json:"cipherList,omitempty"`

	// Enables or disables caching sessions.
	// Must be used in combination with `sessionIdContext`.
	// Acceptable values: `true`, `false`.
	// Default is `false`.
	// +kubebuilder:default=false
	// +optional
	CacheSessions bool `json:"cacheSessions,omitempty"`

	// A unique set of random characters that the server appends to each generated identifier.
	// The length of the string must not exceed `SSL_MAX_SSL_SESSION_ID_LENGTH`.
	// This parameter is always recommended since it helps avoid problems both if the server caches the session and
	// if the client requested caching.
	// Default value: `${application.name}`
	// +optional
	SessionIdContext string `json:"sessionIdContext,omitempty"`

	//
	// The maximum number of sessions that the server caches.
	// A value of 0 means unlimited sessions.
	// Default is [1024\*20](https://github.com/ClickHouse/boringssl/blob/master/include/openssl/ssl.h#L1978).
	// +kubebuilder:default=20480
	// +optional
	SessionCacheSize uint64 `json:"sessionCacheSize,omitempty"`

	// Time for caching the session on the server.
	// Default is [2h](https://github.com/ClickHouse/boringssl/blob/master/include/openssl/ssl.h#L1926)
	// +kubebuilder:default="2h"
	// +optional
	SessionTimeout string `json:"sessionTimeoutSecond,omitempty"`

	// If enabled, verify that the certificate CN or SAN matches the peer hostname.
	// Default is false
	// +kubebuilder:default=false
	// +optional
	ExtendedVerification bool `json:"extendedVerification,omitempty"`

	// Require a TLSv1 connection. Acceptable values: `true`, `false`.
	// Default is false.
	// +kubebuilder:default=false
	// +optional
	RequireTLSv1 bool `json:"requireTLSv1,omitempty"`

	// Require a TLSv1.1 connection. Acceptable values: `true`, `false`.
	// Default is false.
	// +kubebuilder:default=false
	// +optional
	RequireTLSv1_1 bool `json:"requireTlSv11,omitempty"`

	// Require a TLSv1.2 connection. Acceptable values: `true`, `false`.
	// Default is false.
	// +kubebuilder:default=false
	// +optional
	RequireTLSv1_2 bool `json:"requireTlSv12,omitempty"`

	// Activates OpenSSL FIPS mode. Supported if the library’s OpenSSL version supports FIPS.
	// Default is false.
	// +kubebuilder:default=false
	// +optional
	Fips bool `json:"fips,omitempty"`

	// Class (PrivateKeyPassphraseHandler subclass) that requests the passphrase for accessing the private key.
	// For example: `<privateKeyPassphraseHandler>`, `<name>KeyFileHandler</name>`, `<options><password>test</password></options>`, `</privateKeyPassphraseHandler>`.
	// Default is `KeyConsoleHandler`.
	// +optional
	PrivateKeyPassphraseHandler PrivateKeyPassphraseHandler `json:"privateKeyPassphraseHandler,omitempty"`

	// Class (a subclass of CertificateHandler) for verifying invalid certificates.
	// Default is `ConsoleCertificateHandler`
	// +optional
	InvalidCertificateHandler InvalidCertificateHandler `json:"invalidCertificateHandler,omitempty"`

	// Protocols that are not allowed to use.
	// Default is "".
	// +kubebuilder:validation:Enum=sslv2;sslv3;tlsv1;tlsv1_1;tlsv1_2;tlsv1_3
	// +optional
	DisableProtocols []config.SSLProtocolName `json:"disableProtocols,omitempty"`

	// Preferred server ciphers on the client.
	// Default is false.
	// +kubebuilder:default=false
	// +optional
	PreferServerCiphers bool `json:"preferServerCiphers,omitempty"`

	// +optional
	DhParamsFile string `json:"dhParamsFile,omitempty"`
}

// Class (a subclass of CertificateHandler) for verifying invalid certificates.
type InvalidCertificateHandler struct {
	// Default is `ConsoleCertificateHandler`
	// +optional
	Name config.CertificateHandler `json:"name,omitempty"`
}

type PrivateKeyPassphraseHandler struct {
	// Name of the PrivateKeyPassphraseHandler subclass.
	// Default is `KeyConsoleHandler`.
	// +optional
	Name config.KeyPassphraseHandler `json:"name,omitempty"`

	// Options for the PrivateKeyPassphraseHandler subclass.
	// +optional
	Options map[string]string `json:"options,omitempty"`
}

type OpenSSL struct {
	// Server OpenSSL settings.
	// +optional
	Server OpenSSLConfig `json:"server,omitempty"`

	// Client OpenSSL settings.
	// +optional
	Client OpenSSLConfig `json:"client,omitempty"`
}
