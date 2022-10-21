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
)

type ClickHouseKeeperConfig struct {
	// The root element of the configuration.
	XMLName xml.Name `xml:"clickhouse"`

	// max_thread_pool_size
	// default value 100
	MaxThreadPoolSize uint `xml:"max_thread_pool_size,omitempty"`

	// max_thread_pool_free_size
	// default value 1000
	MaxThreadPoolFreeSize uint `xml:"max_thread_pool_free_size,omitempty"`

	// thread_pool_queue_size
	// default value 10000
	ThreadPoolQueueSize uint `xml:"thread_pool_queue_size,omitempty"`

	// max_connections
	// default value 1024
	MaxConnections uint `xml:"max_connections,omitempty"`

	// listen_host  Multiple Values
	// default value "::", enable IPv4 and IPv6.
	ListenHost []string `xml:"listen_host,omitempty"`

	// logger
	Logger Logger `xml:"logger"`

	// openSSL

	// keeper_server
	KepperServer ClickHouseKeeperServerConfig `xml:"kepper_server"`
}

type ClickHouseKeeperServerConfig struct {
	// Port for a client to connect (default for ZooKeeper is 2181).
	TcpPort uint `xml:"tcp_port"`

	// Secure port for an SSL connection between client and keeper-server. 0 is disabled.
	TcpPortSecure int `xml:"tcp_port_secure,omitempty"`

	// storage_path
	StoragePath string `xml:"storage_path,omitempty"`

	// Path to coordination logs, just like ZooKeeper it is best to store logs on non-busy nodes.
	// Default is "/var/lib/clickhouse/coordination/logs"
	LogStoragePath string `xml:"log_storage_path,omitempty"`

	// Path to coordination snapshots.
	// Default is "/var/lib/clickhouse/coordination/snapshots".
	SnapshotStoragePath string `xml:"snapshot_storage_path,omitempty"`

	RaftConfiguration RaftConfiguration `xml:"raft_configuration,omitempty"`
}

type RaftConfiguration struct {

	// Min timeout for client session (ms) (default: 10000).
	// min_session_timeout_ms
	MinSessionTimeoutMs uint64 `xml:"min_session_timeout_ms,omitempty"`

	// Max timeout for client session (ms) (default: 100000).
	// session_timeout_ms
	SessionTimeoutMs uint64 `xml:"session_timeout_ms,omitempty"`

	// Timeout for a single client operation (ms) (default: 10000).
	// operation_timeout_ms
	OperationTimeoutMs uint64 `xml:"operation_timeout_ms,omitempty"`

	// How often ClickHouse Keeper checks for dead sessions and removes them (ms) (default: 500).
	// dead_session_check_period_ms
	DeadSessionCheckPeriodMs uint64 `xml:"dead_session_check_period_ms"`

	// How often a ClickHouse Keeper leader will send heartbeats to followers (ms) (default: 500).
	// heart_beat_interval_ms
	HeartBeatIntervalMs uint64 `xml:"heart_beat_interval_ms,omitempty"`

	// If the follower does not receive a heartbeat from the leader in this interval, then it can initiate leader election (default: 1000).
	// election_timeout_lower_bound_ms
	ElectionTimeoutLowerBoundMs uint64 `xml:"election_timeout_lower_bound_ms,omitempty"`

	// If the follower does not receive a heartbeat from the leader in this interval, then it must initiate leader election (default: 2000).
	// election_timeout_upper_bound_ms
	ElectionTimeoutUpperBoundMs uint64 `xml:"election_timeout_upper_bound_ms,omitempty"`

	// How many coordination log records to store before compaction (default: 100000).
	// eserved_log_items
	EservedLogItems uint64 `xml:"eserved_log_items,omitempty"`

	// How often ClickHouse Keeper will create new snapshots (in the number of records in logs) (default: 100000).
	// snapshot_distance
	SnapshotDistance uint64 `xml:"snapshot_distance,omitempty"`

	// Allow to forward write requests from followers to the leader (default: true).
	// auto_forwarding
	AutoForwarding bool `json:"auto_forwarding,omitempty"`

	// Wait to finish internal connections and shutdown (ms) (default: 5000).
	// shutdown_timeout
	ShutdownTimeout uint64 `json:"shutdown_timeout,omitempty"`

	// session_shutdown_timeout
	// Default is 10000ms, How much time we will wait until sessions are closed during shutdown.
	SessionShutdownTimeoutDuration uint64 `json:"sessionShutdownTimeoutDuration,omitempty"`

	// startup_timeout
	// Default is 180000ms, How much time we will wait until RAFT to start.
	StartupTimeoutDuration uint64 `json:"startupTimeoutDuration,omitempty"`

	// raft_logs_level
	// Default is information, Log internal RAFT logs into main server log level. Valid values: 'trace', 'debug', 'information', 'warning', 'error', 'fatal', 'none'.
	RaftLogsLevel LoggerLevel `json:"raftLogsLevel,omitempty"`

	// How many log records to store in a single file (default: 100000).
	// rotate_log_storage_interval
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
