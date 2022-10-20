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

	// Configuration is the configuration for the ClickHouseKeeper cluster
	Configuration ClickHouseKeeperConfiguration `json:"configuration,omitempty"`

	// Cluster
	Cluster ClickHouseKeeperCluster `json:"cluster"`
}

type ClickHouseKeeperCluster struct {
	Name string `json:"name"`

	Nodes uint `json:"nodes,omitempty"`
}

type ClickHouseKeeperConfiguration struct {
	// listen_host
	// Listen specified address.
	ListenHost []string `json:"listen_host,omitempty"`

	// logger
	Logger Logger `json:"logger,omitempty"`

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

	RaftConfiguration RaftConfiguration `json:"raftConfiguration,omitempty"`

	// OpenSSL
	OpenSSL OpenSSL `json:"openSSL,omitempty"`
}

type RaftConfiguration struct {
	// min_session_timeout_ms
	// Default is 10000ms, Min client session timeout.
	MinSessionTimeoutDuration time.Duration `json:"minSessionTimeoutDuration,omitempty"`

	// session_timeout_ms
	// Default is 100000ms, Max client session timeout.
	SessionTimeoutDuration time.Duration `json:"sessionTimeoutDuration,omitempty"`

	// operation_timeout_ms
	// Default is 10000ms, Default client operation timeout.
	OperationTimeoutDuration time.Duration `json:"operationTimeoutDuration,omitempty"`

	// dead_session_check_period_ms
	// Default is 500ms, How often leader will check sessions to consider them dead and remove.
	DeadSessionCheckPeriodDuration time.Duration `json:"dead_session_check_period_ms"`

	// heart_beat_interval_ms
	// Defaults 500ms, Heartbeat interval between quorum nodes.
	HeartBeatIntervalDuration time.Duration `json:"heartBeatIntervalDuration,omitempty"`

	// election_timeout_lower_bound_ms
	// Default is 1000ms, Lower bound of election timer (avoid too often leader elections)
	ElectionTimeoutLowerBoundDuration time.Duration `json:"electionTimeoutLowerBound,omitempty"`

	// election_timeout_upper_bound_ms
	// Default is 2000ms, Upper bound of election timer (avoid too often leader elections).
	ElectionTimeoutUpperBoundDuration time.Duration `json:"electionTimeoutUpperBound,omitempty"`

	// eserved_log_items
	// Default is 100000, How many log items to store (don't remove during compaction).
	EservedLogItems uint64 `json:"eservedLogItems,omitempty"`

	// snapshot_distance
	// Default is 100000, How many log items we have to collect to write new snapshot.
	SnapshotDistance uint64 `json:"snapshotDistance,omitempty"`

	// auto_forwarding
	// Default is true, Allow to forward write requests from followers to leader.
	AutoForwarding bool `json:"autoForwarding,omitempty"`

	// shutdown_timeout
	// Default is 5000ms, How much time we will wait until RAFT shutdown.
	ShutdownTimeout time.Duration `json:"shutdownTimeout,omitempty"`

	// session_shutdown_timeout
	// Default is 10000ms, How much time we will wait until sessions are closed during shutdown.
	SessionShutdownTimeoutDuration time.Duration `json:"sessionShutdownTimeoutDuration,omitempty"`

	// startup_timeout
	// Default is 180000ms, How much time we will wait until RAFT to start. \
	StartupTimeoutDuration time.Duration `json:"startupTimeoutDuration,omitempty"`

	// raft_logs_level
	// Default is information, Log internal RAFT logs into main server log level. Valid values: 'trace', 'debug', 'information', 'warning', 'error', 'fatal', 'none'.
	RaftLogsLevel string `json:"raftLogsLevel,omitempty"`

	// rotate_log_storage_interval
	// Default is 100000, How many records will be stored in one log storage file.
	RotateLogStorageInterval uint64 `json:"rotateLogStorageInterval,omitempty"`

	// snapshots_to_keep
	// Default is 3, How many compressed snapshots to keep on disk.
	SnapshotsToKeep int `json:"snapshotsToKeep,omitempty"`

	// stale_log_gap
	// Default is 10000, When node became stale and should receive snapshots from leader.
	StaleLogGap uint64 `json:"staleLogGap,omitempty"`

	// fresh_log_gap,
	// Default is 200, When node became fresh.
	FreshLogGap uint64 `json:"freshLogGap,omitempty"`

	// max_requests_batch_size,
	// Default is 100, Max size of batch in requests count before it will be sent to RAFT.
	MaxRequestsBatchSize uint64 `json:"maxRequestsBatchSize,omitempty"`

	// quorum_reads
	// Default is false, Execute read requests as writes through whole RAFT consesus with similar speed.
	QuorumReads bool `json:"quorumReads,omitempty"`

	// force_sync
	// Default is  true, Call fsync on each change in RAFT changelog.
	ForceSync bool `json:"forceSync,omitempty"`

	// compress_logs
	// Default is true, Write compressed coordination logs in ZSTD format.
	CompressLogs bool `json:"compressLogs,omitempty"`

	// compress_snapshots_with_zstd_format
	// Default is true, Write compressed snapshots in ZSTD format (instead of custom LZ4).
	CompressSnapshotsWithZstdFormat bool `json:"compressSnapshotsWithZstdFormat,omitempty"`

	// configuration_change_tries_count
	// Default 20, How many times we will try to apply configuration change (add/remove server) to the cluster.
	ConfigurationChangeTriesCount uint64 `json:"configurationChangeTriesCount,omitempty"`
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
