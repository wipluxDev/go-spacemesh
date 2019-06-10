package discovery

import (
	"bytes"
	"errors"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/spacemeshos/go-spacemesh/p2p/discovery/pb"
	"github.com/spacemeshos/go-spacemesh/p2p/node"
	"github.com/spacemeshos/go-spacemesh/p2p/p2pcrypto"
	"github.com/spacemeshos/go-spacemesh/p2p/server"
)

func (p *protocol) newPingRequestHandler() func(msg server.Message) []byte {
	return func(msg server.Message) []byte {
		p.logger.Info("handle ping request")
		req := &pb.Ping{}
		if err := proto.Unmarshal(msg.Bytes(), req); err != nil {
			p.logger.Error("failed to deserialize ping message err=", err)
			return nil
		}

		go func() {
			if err := p.verifyPinger(msg.Metadata().FromAddress, req); err != nil {
				//todo: black list ?
				p.logger.Error("didn't add peer to addrbook. msg contents were not valid err=", err)
			}
		}()

		//pong
		payload, err := proto.Marshal(&pb.Ping{
			Me:     &pb.NodeInfo{NodeId: p.local.PublicKey().Bytes(), TCPAddress: p.localTcpAddress, UDPAddress: p.localUdpAddress},
			ToAddr: msg.Metadata().FromAddress.String()})
		// TODO: include the resolve To address

		if err != nil {
			p.logger.Error("Error marshaling response message (Ping)")
			return nil
		}

		return payload
	}
}

func extractAddress(from net.Addr, addrString string) (string, error) {
	//extract port
	_, port, err := net.SplitHostPort(addrString)
	if err != nil {
		return "", err
	}

	addr, _, err := net.SplitHostPort(from.String())
	if err != nil {
		addr = from.String()
	}

	ip := net.ParseIP(addr)
	if ip == nil {
		return "", errors.New("canno't parse incoming ip")
	}

	return net.JoinHostPort(addr, port), nil
}


func (p *protocol) verifyPinger(from net.Addr, pi *pb.Ping) error {
	k, err := p2pcrypto.NewPubkeyFromBytes(pi.Me.NodeId)

	if err != nil {
		return err
	}

	tcp, err := extractAddress(from, pi.Me.TCPAddress)
	if err != nil {
		return err
	}
	udp, err := extractAddress(from, pi.Me.UDPAddress)
	if err != nil {
		return nil
	}
	// todo : Validate ToAddr or drop it.
	// todo: decide on best way to know our ext address

	dn := NodeInfoFromNode(node.New(k, tcp), udp)

	if p.needPong(k) {
		p.table.AddAddress(dn, dn)
		if ka := p.table.Find(k); ka == nil {
			// ping the actual source of this node info
			p.table.AddAddress(dn, dn)
		} else {
			if ka.na.udpAddress != udp {
				p.table.AddAddress(dn, dn)
				// todo : what do we do here ?
				// address replaced ?
				// maybe adversary tries to make us remove'
			}
		}

		err = p.Ping(k)

		if err != nil {
			p.table.RemoveAddress(k)
			return err
		}
	}

	return nil
}

// Ping notifies `peer` about our p2p identity.
func (p *protocol) Ping(peer p2pcrypto.PublicKey) error {
	p.logger.Debug("send ping request Peer: %v", peer)

	data := &pb.Ping{Me: &pb.NodeInfo{NodeId: p.local.PublicKey().Bytes(), TCPAddress: p.localTcpAddress, UDPAddress: p.localUdpAddress}, ToAddr: ""}
	payload, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	ch := make(chan []byte)
	foo := func(msg []byte) {
		defer close(ch)
		p.logger.Debug("handle ping response from %v", peer.String())
		data := &pb.Ping{}
		if err := proto.Unmarshal(msg, data); err != nil {
			p.logger.Error("could not unmarshal block data")
			return
		}

		// todo: if we pinged it we already have id so no need to update
		// todo : but what if id or listen address has changed ?

		ch <- data.Me.NodeId
	}

	err = p.msgServer.SendRequest(PINGPONG, payload, peer, foo)

	if err != nil {
		return err
	}

	timeout := time.NewTimer(MessageTimeout) // todo: check whether this is useless because of `requestLifetime`
	select {
	case id := <-ch:
		if id == nil {
			return errors.New("failed sending message")
		}
		if !bytes.Equal(id, peer.Bytes()) {
			return errors.New("got pong with different public key")
		}
	case <-timeout.C:
		return errors.New("ping timeouted")
	}

	return nil
}
