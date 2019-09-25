package keepalive

import (
	"errors"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/log"
	"github.com/spacemeshos/go-spacemesh/p2p/p2pcrypto"
	"github.com/spacemeshos/go-spacemesh/p2p/server"
	"github.com/spacemeshos/go-spacemesh/p2p/service"
	"github.com/spacemeshos/go-spacemesh/rand"
	"sync"
	"time"
)

const Name = "/v1/keepalive"

const PING_PONG = 0

type Message struct {
	Nonce uint64
}

type Protocol struct {
	server *server.MessageServer
	tick   time.Duration

	mtx           sync.RWMutex
	pendingNonces map[uint64]chan error
}
//
type network interface {
	server.Service
}

func NewHeartbeat(network server.Service, tick time.Duration, timeout time.Duration, log log.Log) *Protocol {
	inboxch := make(chan service.DirectMessage, 1000)
	p := &Protocol{
		server: server.NewMsgServer(network.(server.Service), Name, timeout, network.RegisterDirectProtocolWithChannel(Name, inboxch), log),
		tick:   tick,
		pendingNonces: make(map[uint64]chan error, 108),
	}

	p.server.RegisterMsgHandler(PING_PONG, p.Handler)

	return p
}

func (p *Protocol) Handler(message server.Message) []byte {
	msg := &Message{}
	err := types.BytesToInterface(message.Bytes(), msg)

	if err != nil {
		return nil
	}

	newmsg := &Message{Nonce: msg.Nonce}
	b, err := types.InterfaceToBytes(newmsg)

	if err != nil {
		return nil
	}

	return b
}

func (p *Protocol) Ping(server p2pcrypto.PublicKey) error {
	m := &Message{rand.Uint64()}
	serialized, err := types.InterfaceToBytes(m)
	if err != nil {
		return err
	}

	ch := p.addPending(m.Nonce)
	defer p.removePending(m.Nonce)

	err = p.server.SendRequest(PING_PONG, serialized, server, func(msg []byte) {
		deser := &Message{}
		err := types.BytesToInterface(msg, deser)
		if err != nil {
			ch <- err
			return
		}
		if deser.Nonce == m.Nonce {
			ch <- nil
			return
		}

		ch <- errors.New("mismatching nonce")
	})

	if err != nil {
		return err
	}

	return <-ch
}

func (p *Protocol) addPending(nonce uint64) chan error {
	p.mtx.Lock()
	c := make(chan error, 1)
	p.pendingNonces[nonce] = c
	p.mtx.Unlock()
	return c
}

func (p *Protocol) removePending(nonce uint64) chan error {
	p.mtx.Lock()
	c := p.pendingNonces[nonce]
	delete(p.pendingNonces, nonce)
	p.mtx.Unlock()
	return c
}
