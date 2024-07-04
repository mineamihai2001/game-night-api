package logger

import (
	"context"
	"io"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var levelMap = map[string]zerolog.Level{
	"trace": zerolog.TraceLevel,
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"panic": zerolog.PanicLevel,
}

type TraceIdHandler func(ctx context.Context) string

type TracingHook struct {
	GetTraceId TraceIdHandler
}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	if backGroundEnabled, _ := strconv.ParseBool(os.Getenv("LOG_BACKGROUND_ENABLED")); !backGroundEnabled && ctx == context.Background() {
		return
	}

	traceId := h.GetTraceId(ctx)

	e.Str("trace.id", traceId)
}

func Init(traceIdHandler TraceIdHandler, writers ...io.Writer) {
	if consoleEnabled, _ := strconv.ParseBool(os.Getenv("LOG_CONSOLE_ENABLED")); consoleEnabled {
		// print logs to standard output
		writers = append(writers, os.Stdout)

		// add starting log level
		logLevel, ok := levelMap[os.Getenv("LOG_CONSOLE_LEVEL")]
		if !ok {
			logLevel = zerolog.DebugLevel
		}
		zerolog.SetGlobalLevel(logLevel)
	}

	multi := io.MultiWriter(writers...)

	log.Logger = zerolog.New(multi).
		With().
		Str("service.name", os.Getenv("SERVICE_NAME")).
		Str("service.version", os.Getenv("SERVICE_VERSION")).
		Timestamp().
		Logger().
		Hook(TracingHook{
			GetTraceId: traceIdHandler,
		})

}
