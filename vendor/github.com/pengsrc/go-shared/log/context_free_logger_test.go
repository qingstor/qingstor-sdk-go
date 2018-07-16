package log

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pengsrc/go-shared/buffer"
)

func TestNewContextFreeLogger(t *testing.T) {
	buf := buffer.GlobalBytesPool().Get()
	defer buf.Free()

	logger, err := NewLogger(buf, "INFO")
	assert.NoError(t, err)

	l := NewContextFreeLogger(logger)

	l.Debug("DEBUG message")
	l.Info("INFO message")

	assert.NotContains(t, buf.String(), "DEBUG message")
	assert.Contains(t, buf.String(), "INFO message")
	buf.Reset()

	l.Logger.SetCallerFlag(true)

	l.Info("Hello World!")
	assert.Contains(t, buf.String(), "source=log/context_free_logger_test.go")
	t.Log(buf.String())
	buf.Reset()
}
