package splithttp

import (
	"context"
	"io"
	gonet "net"

	"github.com/sqkam/xray-core/transport/internet/browser_dialer"
	"github.com/sqkam/xray-core/transport/internet/websocket"
)

// implements splithttp.DialerClient in terms of browser dialer
// has no fields because everything is global state :O)
type BrowserDialerClient struct{}

func (c *BrowserDialerClient) IsClosed() bool {
	panic("not implemented yet")
}

func (c *BrowserDialerClient) OpenStream(ctx context.Context, url string, body io.Reader, uploadOnly bool) (io.ReadCloser, gonet.Addr, gonet.Addr, error) {
	if body != nil {
		panic("not implemented yet")
	}

	conn, err := browser_dialer.DialGet(url)
	dummyAddr := &gonet.IPAddr{}
	if err != nil {
		return nil, dummyAddr, dummyAddr, err
	}

	return websocket.NewConnection(conn, dummyAddr, nil, 0), conn.RemoteAddr(), conn.LocalAddr(), nil
}

func (c *BrowserDialerClient) PostPacket(ctx context.Context, url string, body io.Reader, contentLength int64) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = browser_dialer.DialPost(url, bytes)
	if err != nil {
		return err
	}

	return nil
}
