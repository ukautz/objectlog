package objectlog

import "bytes"

type (

	// BufferObjectLogger is useful for debugging & testing. It writes all logs in a simple format to
	// the a buffer, which can later on be accessed to evaluate the input, for example to test if a
	// certain log message has been written or not.
	BufferObjectLogger struct {
		buf *bytes.Buffer
	}
)

// NewBufferObjectLog creates new *BufferObjectLogger instance
func NewBufferObjectLog() *BufferObjectLogger {
	return &BufferObjectLogger{
		buf: bytes.NewBuffer(nil),
	}
}

// Buffer provides access to the accumulated buffer
func (this *BufferObjectLogger) Buffer() *bytes.Buffer {
	return this.buf
}

// Clear empties the buffer
func (this *BufferObjectLogger) Clear() {
	this.buf = bytes.NewBuffer(nil)
}

// Debug adds the message to the buffer prefixed by "[DBG] " ended with new line
func (this *BufferObjectLogger) Debug(msg string) {
	this.buf.WriteString("[DBG] " + msg + "\n")
}

// Info adds the message to the buffer prefixed by "[INF] " ended with new line
func (this *BufferObjectLogger) Info(msg string) {
	this.buf.WriteString("[INF] " + msg + "\n")
}

// Warn adds the message to the buffer prefixed by "[WRN] " ended with new line
func (this *BufferObjectLogger) Warn(msg string) {
	this.buf.WriteString("[WRN] " + msg + "\n")
}

// Error adds the message to the buffer prefixed by "[ERR] " ended with new line
func (this *BufferObjectLogger) Error(msg string) {
	this.buf.WriteString("[ERR] " + msg + "\n")
}

// Fatal adds the message to the buffer prefixed by "[FTL] " ended with new line. It DOES NOT EXIT (no call to os.exit)
func (this *BufferObjectLogger) Fatal(msg string) {
	this.buf.WriteString("[FTL] " + msg + "\n")
}
