package zmqutil

import (
	zmq "github.com/pebbe/zmq4"
	"syscall"
)

// IsRetryError returns true if err indicates an interruption (as
// opposed to a failure) in a zmq_recv* or zmq_send* call; otherwise, it
// returns false.
//
// The errors considered to not represent failure cases are EAGAIN and
// EINTR.
func IsRetryError(err error) bool {
	if err == syscall.EAGAIN {
		return true
	}
	if err == syscall.EINTR {
		return true
	}
	return false
}

// RetryRecvMessageBytes retries a RecvMessageBytes call until
// successful or a non-retryable error code is returned.
func RetryRecvMessageBytes(sock *zmq.Socket, flags zmq.Flag) ([][]byte, error) {
	var err error
	err = syscall.EAGAIN
	var b [][]byte
	for IsRetryError(err) {
		b, err = sock.RecvMessageBytes(flags)
	}
	return b, err
}


// RetryRecv retries a Recv call until successful or a non-retryable
// error code is returned.
func RetryRecv(soc *zmq.Socket, flags zmq.Flag) (string, error) {
	var err error
	err = syscall.EAGAIN
	var data string
	for IsRetryError(err) {
		data, err = soc.Recv(flags)
	}
	return data, err
}

// RetryRecvMessage retries a RecvMessage call until successful or a
// non-retryable error code is returned.
func RetryRecvMessage(soc *zmq.Socket, flags zmq.Flag) ([]string, error) {
	var err error
	err = syscall.EAGAIN
	var data []string
	for IsRetryError(err) {
		data, err = soc.RecvMessage(flags)
	}
	return data, err
}

// RetryRecvBytes retries a RecvBytes call until successful or a
// non-retryable error code is returned.
func RetryRecvBytes(soc *zmq.Socket, flags zmq.Flag) ([]byte, error) {
	var err error
	err = syscall.EAGAIN
	var data []byte
	for IsRetryError(err) {
		data, err = soc.RecvBytes(flags)
	}
	return data, err
}

// RetrySend retries a Send call until successful or a non-retryable
// error code is returned.
func RetrySend(soc *zmq.Socket, data string, flags zmq.Flag) (int, error) {
	var err error
	err = syscall.EAGAIN
	var written int
	for IsRetryError(err) {
		written, err = soc.Send(data, flags)
	}
	return written, err
}

// RetrySendMessage retries a SendMessage call until successful or a
// non-retryable error code is returned.
func RetrySendMessage(soc *zmq.Socket, parts ...interface{}) (int, error) {
	var err error
	err = syscall.EAGAIN
	var written int
	for IsRetryError(err) {
		written, err = soc.SendMessage(parts...)
	}
	return written, err
}

// RetrySendBytes retries a SendBytes call until successful or a
// non-retryable error code is returned.
func RetrySendBytes(soc *zmq.Socket, data []byte, flags zmq.Flag) (int, error) {
	var err error
	err = syscall.EAGAIN
	var written int
	for IsRetryError(err) {
		written, err = soc.SendBytes(data, flags)
	}
	return written, err
}
