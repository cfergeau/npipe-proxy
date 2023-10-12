package main

import (
	"context"
	"fmt"
	"net"

	"inet.af/tcpproxy"
)

func startProxy(ipPort string, npipePath string) (tcpproxy.Proxy, error) {
	var proxy tcpproxy.Proxy
	// listen for connections on host tcp port
	/*
		proxy.ListenFunc = func(_, _ string) (net.Listener, error) {
			return net.Listen("tcp", ":8080")
		}
	*/
	proxy.AddRoute(ipPort, &tcpproxy.DialProxy{
		Addr: fmt.Sprintf("npipe:%d", npipePath),
		// when there's a connection to ipPort, connect to the specified named pipe path
		DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
			return net.Dial("unix", npipePath)
		},
	})

	return proxy, proxy.Start()
}

func main() {
	proxy, err := startProxy(":8080", `/tmp/unix.sock`)
	if err != nil {
		panic(err.Error())
	}

	proxy.Wait()
}
