package tcp

import (
	"net"
	"testing"
	"time"
)

func Test_Tcp_Detect(t *testing.T) {
	t.Parallel()

	var (
		addr    = "127.0.0.1:65443"
		timeout = time.Millisecond * 10
		ok      bool
		err     error
	)

	t.Run("error", func(t *testing.T) {
		ok, err = Detect(addr, timeout)

		if err == nil {
			t.Error("err should not be nil")
		}
		if ok {
			t.Error("should not be ok")
		}
	})

	t.Run("ok", func(t *testing.T) {
		var ln net.Listener
		ln, err = net.Listen("tcp", addr)

		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}

		defer func() { _ = ln.Close() }()

		ok, err = Detect(addr, timeout)
		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}

		if !ok {
			t.Error("should be ok")
		}
	})
}
