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
