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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClickHouseKeeperSpec defines the desired state of ClickHouseKeeper
type ClickHouseKeeperSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ClickHouseKeeper. Edit clickhousekeeper_types.go to remove/update
	//Foo string `json:"foo,omitempty"`

	// config is the configuration for the ClickHouseKeeper cluster
	Configuration ClickHouseKeeperConfiguration `json:"configuration,omitempty"`
}

type ClickHouseKeeperConfiguration struct {
	// Port for a client to connect (default for ZooKeeper is 2181).
	TcpPort int `json:"tcpPort,omitempty"`

	// Secure port for an SSL connection between client and keeper-server. 0 is disabled.
	TcpPortSecure int `json:"tcpPortSecure,omitempty"`

	// Path to coordination logs, just like ZooKeeper it is best to store logs on non-busy nodes.
	// Default is "/var/lib/clickhouse/coordination/logs"
	LogStoragePath string `json:"logStoragePath,omitempty"`

	// Path to coordination snapshots.
	// Default is "/var/lib/clickhouse/coordination/snapshots".
	SnapshotStoragePath string `json:"snapshotStoragePath,omitempty"`

	RaftConfiguration

	// OpenSSL
	OpenSSL OpenSSL `json:"openSSL,omitempty"`
}

type RaftConfiguration struct {
	MinSessionTimeoutDuration time.Duration `json:"minSessionTimeoutDuration,omitempty"`
	SessionTimeoutDuration    time.Duration `json:"sessionTimeoutDuration,omitempty"`
	OperationTimeoutDuration  time.Duration `json:"operationTimeoutDuration,omitempty"`
}

type OpenSSL struct {
	// Used for secure tcp port
	Server OpenSSLServer `json:"server,omitempty"`
}

type OpenSSLServer struct {
	CertificateFile     string   `json:"certificateFile,omitempty"`
	PrivateKeyFile      string   `json:"privateKeyFile,omitempty"`
	DhParamsFile        string   `json:"dhParamsFile,omitempty"`
	CAConfig            string   `json:"caConfig,omitempty"`
	LoadDefaultCAFile   bool     `json:"loadDefaultCAFile,omitempty"`
	VerificationMode    string   `json:"verificationMode,omitempty"`
	CacheSessions       bool     `json:"cacheSessions,omitempty"`
	DisableProtocols    []string `json:"disableProtocols,omitempty"`
	PreferServerCiphers bool     `json:"preferServerCiphers,omitempty"`
}

// ClickHouseKeeperStatus defines the observed state of ClickHouseKeeper
type ClickHouseKeeperStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClickHouseKeeper is the Schema for the clickhousekeepers API
type ClickHouseKeeper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClickHouseKeeperSpec   `json:"spec,omitempty"`
	Status ClickHouseKeeperStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClickHouseKeeperList contains a list of ClickHouseKeeper
type ClickHouseKeeperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClickHouseKeeper `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClickHouseKeeper{}, &ClickHouseKeeperList{})
}
