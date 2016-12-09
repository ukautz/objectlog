package objectlog

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMultiObjectLog(t *testing.T) {
	l1 := NewBufferObjectLog()
	l2 := NewBufferObjectLog()
	lg := NewMultiLogger(l1, l2)
	lg.Debug("From Debug")
	lg.Info("From Info")
	lg.Warn("From Warn")
	lg.Error("From Error")
	lg.Fatal("From Fatal")
	expect := strings.Join([]string{
		"[DBG] From Debug",
		"[INF] From Info",
		"[WRN] From Warn",
		"[ERR] From Error",
		"[FTL] From Fatal",
	}, "\n") + "\n"
	assert.Equal(t, expect, l1.Buffer().String())
	assert.Equal(t, expect, l2.Buffer().String())
}
