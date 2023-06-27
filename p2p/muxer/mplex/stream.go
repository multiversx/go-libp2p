package mplex

import (
	"time"

	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/network"

	mp "github.com/libp2p/go-mplex"
)

// stream implements network.MuxedStream over mplex.Stream.
type stream mp.Stream

var _ network.MuxedStream = &stream{}
var log = logging.Logger("mplex")

func (s *stream) Read(b []byte) (n int, err error) {
	n, err = s.mplex().Read(b)
	if err == mp.ErrStreamReset {
		log.Debugf("stream.Read error: %s", err)
		err = network.ErrReset
	}

	return n, err
}

func (s *stream) Write(b []byte) (n int, err error) {
	n, err = s.mplex().Write(b)
	if err == mp.ErrStreamReset {
		log.Debugf("stream.Write error: %s", err)
		err = network.ErrReset
	}

	return n, err
}

func (s *stream) Close() error {
	return s.mplex().Close()
}

func (s *stream) CloseWrite() error {
	return s.mplex().CloseWrite()
}

func (s *stream) CloseRead() error {
	return s.mplex().CloseRead()
}

func (s *stream) Reset() error {
	return s.mplex().Reset()
}

func (s *stream) SetDeadline(t time.Time) error {
	return s.mplex().SetDeadline(t)
}

func (s *stream) SetReadDeadline(t time.Time) error {
	return s.mplex().SetReadDeadline(t)
}

func (s *stream) SetWriteDeadline(t time.Time) error {
	return s.mplex().SetWriteDeadline(t)
}

func (s *stream) mplex() *mp.Stream {
	return (*mp.Stream)(s)
}
