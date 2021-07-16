package layerfetcher

import "github.com/spacemeshos/go-spacemesh/common/types"

// layerHash is the response for a given layer
type layerHash struct {
	simple      types.Hash32 // the simple hash of a the given layer
	accumulated types.Hash32 // the accumulated hash of all layers up to the given layer
}

// layerBlocks is the response for a given layer hash
type layerBlocks struct {
	Blocks          []types.BlockID
	LatestBlocks    []types.BlockID // LatestBlocks are the blocks received in the last 30 seconds from gossip
	VerifyingVector []types.BlockID // VerifyingVector is the input vector for verifying tortoise
}
