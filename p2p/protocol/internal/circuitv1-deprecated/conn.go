package relay

import (
	"fmt"
	"net"
	"time"

	inet "github.com/libp2p/go-libp2p-net"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
)

type Conn struct {
	stream inet.Stream
	remote pstore.PeerInfo
}

type NetAddr struct {
	Relay  string
	Remote string
}

func (n *NetAddr) Network() string {
	return "libp2p-circuit-relay"
}

func (n *NetAddr) String() string {
	return fmt.Sprintf("relay[%s-%s]", n.Remote, n.Relay)
}

func (c *Conn) Close() error {
	return c.stream.Reset()
}

func (c *Conn) Read(buf []byte) (int, error) {
	return c.stream.Read(buf)
}

func (c *Conn) Write(buf []byte) (int, error) {
	return c.stream.Write(buf)
}

func (c *Conn) SetDeadline(t time.Time) error {
	return c.stream.SetDeadline(t)
}

func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.stream.SetReadDeadline(t)
}

func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.stream.SetWriteDeadline(t)
}

func (c *Conn) RemoteAddr() net.Addr {
	return &NetAddr{
		Relay:  c.stream.Conn().RemotePeer().Pretty(),
		Remote: c.remote.ID.Pretty(),
	}
}

// TODO: is it okay to cast c.Conn().RemotePeer() into a multiaddr? might be "user input"
func (c *Conn) RemoteMultiaddr() ma.Multiaddr {
	proto := ma.ProtocolWithCode(ma.P_P2P).Name
	peerid := c.stream.Conn().RemotePeer().Pretty()
	p2paddr := ma.StringCast(fmt.Sprintf("/%s/%s", proto, peerid))

	circaddr := ma.Cast(ma.CodeToVarint(P_CIRCUIT))
	return p2paddr.Encapsulate(circaddr)
}

func (c *Conn) LocalMultiaddr() ma.Multiaddr {
	return c.stream.Conn().LocalMultiaddr()
}

func (c *Conn) LocalAddr() net.Addr {
	na, err := manet.ToNetAddr(c.stream.Conn().LocalMultiaddr())
	if err != nil {
		log.Error("failed to convert local multiaddr to net addr:", err)
		return nil
	}
	return na
}
