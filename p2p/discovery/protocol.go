package discovery

import (
	"github.com/spacemeshos/go-spacemesh/log"
	"github.com/spacemeshos/go-spacemesh/p2p/node"
	"github.com/spacemeshos/go-spacemesh/p2p/p2pcrypto"
	"github.com/spacemeshos/go-spacemesh/p2p/server"
	"github.com/spacemeshos/go-spacemesh/p2p/service"
	"sync"
	"time"
)

type protocolRoutingTable interface {
	GetAddress() *KnownAddress
	AddAddresses(n []NodeInfo, src NodeInfo)
	AddAddress(n NodeInfo, src NodeInfo)
	AddressCache() []NodeInfo
	RemoveAddress(key p2pcrypto.PublicKey)
	Find(key p2pcrypto.PublicKey) *KnownAddress
}

type protocol struct {
	local     node.Node
	table     protocolRoutingTable
	logger    log.Log
	msgServer *server.MessageServer

	pingpongLock sync.RWMutex
	lastpong     map[p2pcrypto.PublicKey]time.Time
	lastping     map[p2pcrypto.PublicKey]time.Time

	localTcpAddress string
	localUdpAddress string
}

func (d *protocol) SetLocalAddresses(tcp, udp string) {
	d.localTcpAddress = tcp
	d.localUdpAddress = udp
}

// Name is the name if the protocol.
const Name = "/udp/v2/discovery"

// MessageBufSize is the buf size we give to the messages channel
const MessageBufSize = 100

// MessageTimeout is the timeout we tolerate when waiting for a message reply
const MessageTimeout = time.Second * 1 // TODO: Parametrize

const PingPongExpiration = 24 * time.Hour

// PINGPONG is the ping protocol ID
const PINGPONG = 0

// GET_ADDRESSES is the findnode protocol ID
const GET_ADDRESSES = 1

// NewDiscoveryProtocol is a constructor for a protocol protocol provider.
func NewDiscoveryProtocol(local node.Node, rt protocolRoutingTable, svc server.Service, log log.Log) *protocol {
	s := server.NewMsgServer(svc, Name, MessageTimeout, make(chan service.DirectMessage, MessageBufSize), log)
	d := &protocol{
		local:     local,
		table:     rt,
		msgServer: s,
		logger:    log,
	}

	d.SetLocalAddresses(local.Address(), local.Address())

	d.msgServer.RegisterMsgHandler(PINGPONG, d.newPingRequestHandler())
	d.msgServer.RegisterMsgHandler(GET_ADDRESSES, d.newGetAddressesRequestHandler())
	return d
}
