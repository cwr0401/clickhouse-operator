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

type Logger struct {
	// level
	// Possible levels: none (turns off logging)、fatal、critical、error、warning、notice、information、debug、trace.
	Level string `json:"level,omitempty"`

	// log
	// Default is "/var/log/clickhouse-keeper/clickhouse-keeper.log"
	Log string `json:"log,omitempty"`

	// stream_compress
	// Default is false,
	StreamCompress bool `json:"streamCompress,omitempty"`

	// errorlog
	// Default is "/var/log/clickhouse-keeper/clickhouse-keeper.err.log"
	ErrorLog string `json:"errorLog,omitempty"`

	// size
	// Default is 100M,
	Size uint64 `json:"size,omitempty"`

	// compress
	// Default is true,
	Compress bool `json:"compress,omitempty"`

	// count
	// Default is 1,
	Count int `json:"count,omitempty"`

	// flush
	// Default is true,
	Flush bool `json:"flush,omitempty"`

	// rotateOnOpen
	// Default is false,
	RotateOnOpen bool `json:"rotateOnOpen,omitempty"`

	Console bool `json:"console,omitempty"`

	// logger.formatting.type
	Formatting LoggerFormatting `json:"formatting,omitempty"`

	// levels
	Levels LoggerLevels `json:"levels,omitempty"`
}

type LoggerLevels struct {
	ConfigReloader string `json:"configReloader,omitempty"`

	Logger LoggerLevelsLogger `json:"logger,omitempty"`
}

type LoggerLevelsLogger struct {
	Name  string `json:"name,omitempty"`
	Level string `json:"level,omitempty"`
}

type LoggerFormatting struct {
	// type
	// Specify log format(for now, JSON only)
	Type string `json:"type,omitempty"`

	Names JsonKeyNames `json:"names,omitempty"`
}

type JsonKeyNames struct {
	// DataTime json key name, default is "date_time".
	DateTime string `json:"dateTime,omitempty"`

	// ThreadName json key name, default is "thread_name".
	ThreadName string `json:"threadName,omitempty"`

	// ThreadID json key name, default is "thread_id"
	ThreadID string `json:"threadID,omitempty"`

	// Level json key name, default is "level"
	Level string `json:"level,omitempty"`

	// QueryID json key name, default is "query_id".
	QueryID string `json:"queryID,omitempty"`

	// LoggerName json key name, default is "logger_name".
	LoggerName string `json:"loggerName,omitempty"`

	// Message json key name, default is "message".
	Message string `json:"message,omitempty"`

	// SourceFile json key name, default is "source_file".
	SourceFile string `json:"sourceFile,omitempty"`

	// SourceLine json key name, default is "source_line".
	SourceLine int `json:"sourceLine,omitempty"`
}
