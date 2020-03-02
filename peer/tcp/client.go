package tcp

import (
	"goNet"
	. "goNet/log"
	"net"
)

type client struct {
	goNet.PeerIdentify
	session *session
}

func (c *client) Start() {
	conn, err := net.Dial("tcp", c.Addr())
	if err != nil {
		Log.Fatalf("#tcp(%v) connect failed %v", c.Type(), err.Error())
		return
	}
	Log.Infof("#tcp(%v) connect(%v) success", c.Type(), conn.RemoteAddr())
	c.session = newSession(conn)
	go c.session.recvLoop()
}

func (c *client) Stop() {
	c.session.Close()
}

func init() {
	identify := goNet.PeerIdentify{}
	identify.SetType(goNet.PEER_CLIENT)
	goNet.RegisterPeer(&server{PeerIdentify: identify})
}