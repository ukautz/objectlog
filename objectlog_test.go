package objectlog

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDefaultFormatter(t *testing.T) {
	withARgs := DefaultFormatter(OBJECT_LOG_LEVEL_INFO, "(PREFIX) ", " (SUFFIX)", "The Message with arg \"%s\"", []interface{}{"ARG1"}, map[string]interface{}{
		"foo": "bar",
	})
	assert.Equal(t, `(PREFIX) The Message with arg "ARG1" (SUFFIX) :: {"foo":"bar"}`, withARgs)

	noArgs := DefaultFormatter(OBJECT_LOG_LEVEL_INFO, "(PREFIX) ", " (SUFFIX)", "The Message with arg \"%s\"", []interface{}{"ARG1"}, nil)
	assert.Equal(t, `(PREFIX) The Message with arg "ARG1" (SUFFIX)`, noArgs)
}

func TestObjectLog(t *testing.T) {
	lg := NewBufferObjectLog()
	ol := NewObjectLog(lg)
	ol.LogDebug("Hello %s", "foo1")
	ol.LogInfo("Hello %s", "foo2")
	ol.LogWarn("Hello %s", "foo3")
	ol.LogError("Hello %s", "foo4")
	ol.LogFatal("Hello %s", "foo5")
	assert.Equal(t, strings.Join([]string{
		"[DBG] Hello foo1",
		"[INF] Hello foo2",
		"[WRN] Hello foo3",
		"[ERR] Hello foo4",
		"[FTL] Hello foo5",
	}, "\n")+"\n", lg.Buffer().String())
}

func TestObjectLog_Args(t *testing.T) {
	lg := NewBufferObjectLog()
	ol := NewObjectLog(lg)
	assert.Equal(t, map[string]interface{}{}, ol.LogArgs())
	ol.SetLogArg("foo", "bar")
	ol.SetLogArg("baz", "zoing")
	assert.Equal(t, map[string]interface{}{"baz": "zoing", "foo": "bar"}, ol.LogArgs())
	ol.SetLogArgs(map[string]interface{}{"foo": "bar"})
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, ol.LogArgs())
	ol.LogDebug("Hello %s", "foo1")
	ol.LogInfo("Hello %s", "foo2")
	ol.LogWarn("Hello %s", "foo3")
	ol.LogError("Hello %s", "foo4")
	ol.LogFatal("Hello %s", "foo5")
	assert.Equal(t, strings.Join([]string{
		`[DBG] Hello foo1 :: {"foo":"bar"}`,
		`[INF] Hello foo2 :: {"foo":"bar"}`,
		`[WRN] Hello foo3 :: {"foo":"bar"}`,
		`[ERR] Hello foo4 :: {"foo":"bar"}`,
		`[FTL] Hello foo5 :: {"foo":"bar"}`,
	}, "\n")+"\n", lg.Buffer().String())
}
func TestObjectLog_BothIx(t *testing.T) {
	lg := NewBufferObjectLog()
	ol := NewObjectLog(lg)
	assert.Equal(t, "", ol.LogPrefix())
	assert.Equal(t, "", ol.LogSuffix())
	ol.SetLogPrefix("PRE ")
	ol.SetLogSuffix(" SUF")
	assert.Equal(t, "PRE ", ol.LogPrefix())
	assert.Equal(t, " SUF", ol.LogSuffix())
	ol.LogDebug("Hello %s", "foo1")
	ol.LogInfo("Hello %s", "foo2")
	ol.LogWarn("Hello %s", "foo3")
	ol.LogError("Hello %s", "foo4")
	ol.LogFatal("Hello %s", "foo5")
	assert.Equal(t, strings.Join([]string{
		`[DBG] PRE Hello foo1 SUF`,
		`[INF] PRE Hello foo2 SUF`,
		`[WRN] PRE Hello foo3 SUF`,
		`[ERR] PRE Hello foo4 SUF`,
		`[FTL] PRE Hello foo5 SUF`,
	}, "\n")+"\n", lg.Buffer().String())
}

func TestObjectLog_Clone(t *testing.T) {
	lg := NewBufferObjectLog()
	from := NewObjectLog(lg).
		SetLogPrefix("PREFIX").
		SetLogSuffix("SUFFIX").
		SetLogArg("foo", "bar").
		SetLogArg("baz", "zoing")
	to := from.LogCloneObjectLog()

	// assure all copied
	assert.Equal(t, "PREFIX", to.LogPrefix())
	assert.Equal(t, "SUFFIX", to.LogSuffix())
	assert.Equal(t, map[string]interface{}{
		"foo": "bar",
		"baz": "zoing",
	}, to.LogArgs())

	// assure modification on clone does not affect orig
	to.SetLogPrefix("PREFIX2").SetLogSuffix("SUFFIX2").SetLogArg("bla", "bla")
	assert.Equal(t, "PREFIX2", to.LogPrefix())
	assert.Equal(t, "SUFFIX2", to.LogSuffix())
	assert.Equal(t, map[string]interface{}{
		"foo": "bar",
		"baz": "zoing",
		"bla": "bla",
	}, to.LogArgs())
	assert.Equal(t, "PREFIX", from.LogPrefix())
	assert.Equal(t, "SUFFIX", from.LogSuffix())
	assert.Equal(t, map[string]interface{}{
		"foo": "bar",
		"baz": "zoing",
	}, from.LogArgs())
}