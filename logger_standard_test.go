package objectlog

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestStandardObjectLog(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	l := log.New(buf, "", 0)
	lg := NewStandardLogger(l)
	lg.Debug("From Debug")
	lg.Info("From Info")
	lg.Warn("From Warn")
	lg.Error("From Error")
	assert.Equal(t, strings.Join([]string{
		"[DEBUG] From Debug",
		"[INFO] From Info",
		"[WARN] From Warn",
		"[ERROR] From Error",
	}, "\n")+"\n", buf.String())
}
