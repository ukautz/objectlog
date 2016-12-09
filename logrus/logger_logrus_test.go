package logrus

import (
	"bytes"
	lr "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type (
	testLogrusFormatter struct{}
)

func (this *testLogrusFormatter) Format(e *lr.Entry) ([]byte, error) {
	return []byte(e.Level.String() + ": " + e.Message + "\n"), nil
}

func TestLogrusObjectLog(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	l := lr.New()
	l.Out = buf
	l.Level = lr.DebugLevel
	l.Formatter = &testLogrusFormatter{}
	lg := NewLogrusObjectLogger(l)
	lg.Debug("From Debug")
	lg.Info("From Info")
	lg.Warn("From Warn")
	lg.Error("From Error")
	assert.Equal(t, strings.Join([]string{
		"debug: From Debug",
		"info: From Info",
		"warning: From Warn",
		"error: From Error",
	}, "\n")+"\n", buf.String())
}
