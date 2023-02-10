package logger

import "context"

type Logger interface {
	Field(string, interface{}) Field
	Debug(context.Context, string, ...Field)
	Info(context.Context, string, ...Field)
	Error(context.Context, string, ...Field)
}

type (
	Field struct {
		Key       string
		Interface interface{}
	}
)
