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

	"github.com/cwr0401/clickhouse-operator/models/config"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClickHouseKeeperSpec defines the desired state of ClickHouseKeeper
type ClickHouseKeeperSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ClickHouseKeeper. Edit clickhousekeeper_types.go to remove/update
	//Foo string `json:"foo,omitempty"`

	// Configuration is the configuration for the ClickHouseKeeper cluster.
	Configuration ClickHouseKeeperConfiguration `json:"configuration,omitempty"`

	// Cluster
	Cluster ClickHouseKeeperCluster `json:"cluster"`
}

type ClickHouseKeeperCluster struct {
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:default=1
	// +optional
	QuorumNodes uint `json:"QuorumNodes,omitempty"`

	//
	ServiceName string `json:"serviceName,omitempty"`

	// template
}

type ClickHouseKeeperConfiguration struct {

	// max_thread_pool_size
	// default value 100
	MaxThreadPoolSize uint `json:"maxThreadPoolSize,omitempty"`

	// max_thread_pool_free_size
	// default value 1000
	MaxThreadPoolFreeSize uint `json:"maxThreadPoolFreeSize,omitempty"`

	// thread_pool_queue_size
	// default value 10000
	ThreadPoolQueueSize uint `json:"threadPoolQueueSize,omitempty"`

	// max_connections
	// default value 1024
	MaxConnections uint `json:"maxConnections,omitempty"`

	// listen_host
	// Listen specified address.
	// Default value "::", enable IPv4 and IPv6.
	ListenHost []string `json:"listenHost,omitempty"`

	// logger
	Logger Logger `json:"logger,omitempty"`

	// ClickHouse Keeper configuration
	KepperServer ClickHouseKeeperServerConfig `json:"kepperServer,omitempty"`

	// OpenSSL
	OpenSSL OpenSSL `json:"openSSL,omitempty"`
}

type ClickHouseKeeperServerConfig struct {
	// Port for a client to connect (default for ZooKeeper is 2181).
	TcpPort uint `json:"tcpPort,omitempty"`

	// Secure port for an SSL connection between client and keeper-server. 0 is disabled.
	TcpPortSecure int `json:"tcpPortSecure,omitempty"`

	// storage_path
	StoragePath string `json:"storagePath,omitempty"`

	// Path to coordination logs, just like ZooKeeper it is best to store logs on non-busy nodes.
	// Default is "/var/lib/clickhouse/coordination/logs"
	LogStoragePath string `json:"logStoragePath,omitempty"`

	// Path to coordination snapshots.
	// Default is "/var/lib/clickhouse/coordination/snapshots".
	SnapshotStoragePath string `json:"snapshotStoragePath,omitempty"`

	RaftConfiguration RaftConfiguration `json:"raftConfiguration,omitempty"`
}

type RaftConfiguration struct {
	// Min client session timeout.
	// Default is 10000ms.
	// +kubebuilder:default=10s
	// +optional
	MinSessionTimeoutDuration time.Duration `json:"minSessionTimeoutDuration,omitempty"`

	// Max client session timeout.
	// Default is 100000ms.
	// +kubebuilder:default=100s
	// +optional
	SessionTimeoutDuration time.Duration `json:"sessionTimeoutDuration,omitempty"`

	// Default client operation timeout.
	// Default is 10000ms.
	// +kubebuilder:default=10ms
	// +optional
	OperationTimeoutDuration time.Duration `json:"operationTimeoutDuration,omitempty"`

	// How often leader will check sessions to consider them dead and remove.
	// Default is 500ms.
	// +kubebuilder:default=500ms
	// +optional
	DeadSessionCheckPeriodDuration time.Duration `json:"deadSessionCheckPeriodDuration,omitempty"`

	// Heartbeat interval between quorum nodes.
	// Defaults 500ms.
	// +kubebuilder:default=500ms
	// +optional
	HeartBeatIntervalDuration time.Duration `json:"heartBeatIntervalDuration,omitempty"`

	// Lower bound of election timer (avoid too often leader elections).
	// Default is 1000ms.
	// +kubebuilder:default=1s
	// +optional
	ElectionTimeoutLowerBoundDuration time.Duration `json:"electionTimeoutLowerBoundDuration,omitempty"`

	// Upper bound of election timer (avoid too often leader elections).
	// Default is 2000ms.
	// +kubebuilder:default=2s
	// +optional
	ElectionTimeoutUpperBoundDuration time.Duration `json:"electionTimeoutUpperBoundDuration,omitempty"`

	// How many log items to store (don't remove during compaction).
	// Default is 100000.
	// +kubebuilder:default=100000
	// +optional
	EservedLogItems uint64 `json:"eservedLogItems,omitempty"`

	// How many log items we have to collect to write new snapshot.
	// Default is 100000.
	// +kubebuilder:default=100000
	// +optional
	SnapshotDistance uint64 `json:"snapshotDistance,omitempty"`

	// Allow to forward write requests from followers to leader.
	// Default is true.
	// +kubebuilder:default=true
	// +optional
	AutoForwarding bool `json:"autoForwarding,omitempty"`

	// How much time we will wait until RAFT shutdown.
	// Default is 5000ms.
	// +kubebuilder:default=5s
	// +optional
	ShutdownTimeoutDuration time.Duration `json:"shutdownTimeoutDuration,omitempty"`

	// How much time we will wait until sessions are closed during shutdown.
	// Default is 10000ms.
	// +kubebuilder:default=10s
	// +optional
	SessionShutdownTimeoutDuration time.Duration `json:"sessionShutdownTimeoutDuration,omitempty"`

	// How much time we will wait until RAFT to start.
	// Default is 180000ms.
	// +kubebuilder:default=3m
	// +optional
	StartupTimeoutDuration time.Duration `json:"startupTimeoutDuration,omitempty"`

	// Log internal RAFT logs into main server log level. Valid values: 'trace', 'debug', 'information', 'warning', 'error', 'fatal', 'none'.
	// Default is information.
	// +kubebuilder:validation:Enum=trace;debug;information;warning;error;fatal;none
	// +kubebuilder:default=information
	// +optional
	RaftLogsLevel config.LoggerLevel `json:"raftLogsLevel,omitempty"`

	// How many records will be stored in one log storage file.
	// Default is 100000.
	// +kubebuilder:default=100000
	// +optional
	RotateLogStorageInterval uint64 `json:"rotateLogStorageInterval,omitempty"`

	// How many compressed snapshots to keep on disk.
	// Default is 3.
	// +kubebuilder:default=3
	// +optional
	SnapshotsToKeep uint64 `json:"snapshotsToKeep,omitempty"`

	// When node became stale and should receive snapshots from leader.
	// Default is 10000.
	// +kubebuilder:default=10000
	// +optional
	StaleLogGap uint64 `json:"staleLogGap,omitempty"`

	// When node became fresh.
	// Default is 200.
	// +kubebuilder:default=200
	// +optional
	FreshLogGap uint64 `json:"freshLogGap,omitempty"`

	// Max size of batch in requests count before it will be sent to RAFT.
	// Default is 100.
	// +kubebuilder:default=100
	// +optional
	MaxRequestsBatchSize uint64 `json:"maxRequestsBatchSize,omitempty"`

	// Execute read requests as writes through whole RAFT consesus with similar speed.
	// Default is false.
	// +kubebuilder:default=true
	// +optional
	QuorumReads bool `json:"quorumReads,omitempty"`

	// Call fsync on each change in RAFT changelog.
	// Default is true.
	// +kubebuilder:default=true
	// +optional
	ForceSync bool `json:"forceSync,omitempty"`

	// Write compressed coordination logs in ZSTD format.
	// Default is true.
	// +kubebuilder:default=true
	// +optional
	CompressLogs bool `json:"compressLogs,omitempty"`

	// Write compressed snapshots in ZSTD format (instead of custom LZ4).
	// Default is true.
	// +kubebuilder:default=true
	// +optional
	CompressSnapshotsWithZstdFormat bool `json:"compressSnapshotsWithZstdFormat,omitempty"`

	// How many times we will try to apply configuration change (add/remove server) to the cluster.
	// Default 20.
	// +kubebuilder:default=20
	// +optional
	ConfigurationChangeTriesCount uint64 `json:"configurationChangeTriesCount,omitempty"`
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
