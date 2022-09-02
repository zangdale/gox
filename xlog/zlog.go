// github.com/rs/zerolog
package xlog

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ZLog 默认日志
var ZLog = log.Logger

// NullLog 空日志
var NullLog = zerolog.Nop

// CustomeFormatLog 自定义输出的 zerolog
func CustomeFormatLog(writer ...io.Writer) zerolog.Logger {
	var out io.Writer = Stdout
	if len(writer) > 0 {
		out = io.MultiWriter(writer...)
	}
	output := zerolog.ConsoleWriter{Out: out, TimeFormat: time.RFC3339}

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	return zerolog.New(output).With().Timestamp().Logger()
}
