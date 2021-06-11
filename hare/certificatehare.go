package hare

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/config"
	"github.com/spacemeshos/go-spacemesh/log"
)

// CertificateHare is a struct for testing which simply subscribes to gossip and
// listens for the certification messages, but doesn't participate in the hare
type CertificateHare struct {
	Closer
	log.Log
	config config.Config

	network    NetworkService
	beginLayer chan types.LayerID

	//broker for the messages for various layers?

	outputChan chan TerminationOutput

	validate outputValidationFunc

	nid types.NodeID
}

//NewCertificateHare creates a new CertificateHare struct
func NewCertificateHare(conf config.Config, p2p NetworkService, validate outputValidationFunc, beginLayer chan types.LayerID, logger log.Log,
	nid types.NodeID) *CertificateHare {
	h := new(CertificateHare)
	h.Closer = NewCloser()

	h.Log = logger
	h.config = conf

	h.network = p2p
	h.beginLayer = beginLayer

	h.outputChan = make(chan TerminationOutput, 20)
	h.nid = nid

	h.validate = validate

	return h
}

// OutputLoop checks for the output from the Gossip protocol
func (ch *CertificateHare) OutputLoop(ctx context.Context) {
	// for {
	// 	select {
	// 	case out := <-ch.outputChan:
	// 		//set := out.Set()
	// 		//check validity of the collected output?
	// 		//id := out.ID()

	// 		//unregister the listening process for that layer
	// 	}
	// }
}

// Start starts the output collection loop
func (ch *CertificateHare) Start(ctx context.Context) error {
	ch.WithContext(ctx).With().Info("starting protocol", log.String("protocol", "HARE_CERTIFICATE_PROTOCOL"))

	//start the broker

	//start the output collection loop, that simply listens to the gossip and reports messages
	return nil
}
