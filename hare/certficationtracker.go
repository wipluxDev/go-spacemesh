package hare

import "github.com/spacemeshos/go-spacemesh/log"

// certificationtracker tracks certification messages
type certficiationTracker struct {
	certificates  map[string]struct{} //tracks PubKey->Certification
	setsCertified map[string]*Set     //tracks the sets certified by a pubkey
}

func newCertificationTracker(expectedSize int) *certficiationTracker {
	ct := &certficiationTracker{}
	ct.certificates = make(map[string]struct{}, expectedSize)

	return ct
}

// OnCertification tracks the provided Certification message
func (ct *certficiationTracker) OnCertification(msg *Msg, certifiedSet *Set) bool {
	pub := msg.PubKey
	if _, exist := ct.certificates[pub.String()]; exist {
		if ct.setsCertified[pub.String()].Equals(certifiedSet) {
			return true
		}
		log.Panic("Certified sets do not match from the same pubkey")

	}

	ct.certificates[pub.String()] = struct{}{}
	ct.setsCertified[pub.String()] = certifiedSet

	return false
}
