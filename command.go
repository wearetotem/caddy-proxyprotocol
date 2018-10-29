package proxyprotocol

// Command indicates the PROXY command being used.
// https://www.haproxy.org/download/1.8/doc/proxy-protocol.txt
type Command byte

const (
	// CommandLocal indicates the connection was established on purpose by the proxy without being relayed.
	CommandLocal Command = '\x00'

	// CommandProxy the connection was established on behalf of another node, and reflects the original connection endpoints.
	CommandProxy Command = '\x01'
)