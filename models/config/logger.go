package config

type LoggerLevel string

const (
	LoggerLevelNone        LoggerLevel = "none"
	LoggerLevelFatal       LoggerLevel = "fatal"
	LoggerLevelCritical    LoggerLevel = "critical"
	LoggerLevelError       LoggerLevel = "error"
	LoggerLevelWarning     LoggerLevel = "warning"
	LoggerLevelNotice      LoggerLevel = "notice"
	LoggerLevelInformation LoggerLevel = "information"
	LoggerLevelDebug       LoggerLevel = "debug"
	LoggerLevelTrace       LoggerLevel = "trace"
)

type Logger struct {
	// level
	// Possible levels: none (turns off logging)、fatal、critical、error、warning、notice、information、debug、trace.
	Level LoggerLevel `xml:"level"`

	// log
	// Default is "/var/log/clickhouse-keeper/clickhouse-keeper.log"
	Log string `xml:"log,omitempty"`

	// stream_compress
	// Default is false,
	StreamCompress bool `xml:"stream_compress,omitempty"`

	// errorlog
	// Default is "/var/log/clickhouse-keeper/clickhouse-keeper.err.log"
	ErrorLog string `xml:"errorlog,omitempty"`

	// size
	// Default is 100M,
	Size uint64 `xml:"size,omitempty"`

	// compress
	// Default is true,
	Compress bool `xml:"compress,omitempty"`

	// count
	// Default is 1,
	Count int `xml:"count,omitempty"`

	// flush
	// Default is true,
	Flush bool `xml:"flush,omitempty"`

	// rotateOnOpen
	// Default is false,
	RotateOnOpen bool `xml:"rotateOnOpen,omitempty"`

	// default is true
	Console bool `xml:"console"`

	// logger.formatting
	Formatting LoggerFormatting `xml:"formatting,omitempty"`

	// levels
	Levels LoggerLevels `xml:"levels,omitempty"`
}

type LoggerLevels struct {
	ConfigReloader string `xml:"configReloader,omitempty"`

	Logger []LoggerLevelsLogger `xml:"logger,omitempty"`
}

type LoggerLevelsLogger struct {
	Name  string `xml:"name,omitempty"`
	Level string `xml:"level,omitempty"`
}

type LoggerFormatting struct {
	// type
	// Specify log format(for now, JSON only)
	Type string `xml:"type,omitempty"`

	Names JsonKeyNames `xml:"names,omitempty"`
}

type JsonKeyNames struct {
	// DataTime json key name, default is "date_time".
	DateTime string `xml:"date_time,omitempty"`

	// ThreadName json key name, default is "thread_name".
	ThreadName string `xml:"thread_name,omitempty"`

	// ThreadID json key name, default is "thread_id"
	ThreadID string `xml:"thread_id,omitempty"`

	// Level json key name, default is "level"
	Level string `xml:"level,omitempty"`

	// QueryID json key name, default is "query_id".
	QueryID string `xml:"query_id,omitempty"`

	// LoggerName json key name, default is "logger_name".
	LoggerName string `xml:"logger_name,omitempty"`

	// Message json key name, default is "message".
	Message string `xml:"message,omitempty"`

	// SourceFile json key name, default is "source_file".
	SourceFile string `xml:"source_file,omitempty"`

	// SourceLine json key name, default is "source_line".
	SourceLine int `xml:"source_line,omitempty"`
}
