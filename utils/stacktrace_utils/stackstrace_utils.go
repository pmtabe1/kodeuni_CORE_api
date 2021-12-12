package stacktrace_utils

import "github.com/pkg/errors"

type IStacktraceUtils interface {
}

type StacktraceUtils struct {
}

func GenerateStackstraceWithMessageCapture(err error) (errorTrace string) {
	errorTrace = errors.Wrap(err, err.Error()).Error()
	return errorTrace
}

func GenerateStackstrace(err error) (errorTrace string) {
	errorTrace = errors.Wrap(err, err.Error()).Error()
	return errorTrace
}