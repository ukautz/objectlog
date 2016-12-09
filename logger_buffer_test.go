package objectlog

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBufferObjectLog(t *testing.T) {
	lg := NewBufferObjectLog()
	lg.Debug("From Debug")
	lg.Info("From Info")
	lg.Warn("From Warn")
	lg.Error("From Error")
	lg.Fatal("From Fatal")
	assert.Equal(t, strings.Join([]string{
		"[DBG] From Debug",
		"[INF] From Info",
		"[WRN] From Warn",
		"[ERR] From Error",
		"[FTL] From Fatal",
	}, "\n")+"\n", lg.Buffer().String())
}
