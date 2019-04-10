package main

import (
//	"log"
)

type peerId string
type blockId string
type resourceId string

type Peer struct {
	id peerId
}

type Block struct {
	id blockId
}

type Resource struct {
	id     resourceId
	blocks []blockId
}

type PeerBlock struct {
	peerId  peerId
	blockId blockId
}

type schedulerParams struct {
	obfuscationP float64
	spreadP      float64
}

type State map[peerId]map[blockId]resourceId

func (s *State) addPeerBlockEntry(p Peer, b Block, rid resourceId) {
	if (*s)[p.id] == nil {
		(*s)[p.id] = make(map[blockId]resourceId)
	}
	(*s)[p.id][b.id] = rid
}

// get providers returns a set of blocks and peers to replicate based on the
// requester peer neighborhood, its cached blocks and the resourceId to request.
// this is entry point for the gist of the replication algorithm used by the
// scheduler
func getProviders(neighborPeers []Peer, cachedBlocks []Block,
	resource []Resource, netState State, params schedulerParams) ([]PeerBlock, error) {

	// get all PeerBlocks from neighbor peers
	neighPbs := filterPeersBlocks(neighborPeers, netState)

	// gets all PeerBlocks from its neighbor peers that are not yet cached locally
	nonCachedNeighPbs := filterBlocks(neighPbs, cachedBlocks)

	// select final PeerBlocks based on configuration params and requested
	// resource
	rpbs := getBlocksResourcePeers(resource, neighborPeers)
	finalPbs := schedule(rpbs, nonCachedNeighPbs, params.obfuscationP, params.spreadP)

	return finalPbs, nil
}

// returns all PeerBlocks from neighbor peers in a given state
func filterPeersBlocks(peers []Peer, state State) []PeerBlock {
	pbs := []PeerBlock{}

	for pid, _ := range state {
		for _, p := range peers {
			if pid == p.id {
				for blkId, _ := range state[pid] {
					pb := PeerBlock{peerId: pid, blockId: blkId}
					pbs = append(pbs, pb)
				}
				break
			}
		}
	}

	return pbs
}

// returns all PeerBlocks which contain blocks that are not part of the Block
// slice
func filterBlocks(pbs []PeerBlock, bs []Block) []PeerBlock {
	return nil
}

// get all PeerBlocks from a given resource and set of peers
func getBlocksResourcePeers(resource []Resource, peers []Peer) []PeerBlock {
	return nil
}

// given a set of resource peer blocks a set of not cached Peer Blocks, schedule
// the replication according to the parameters
func schedule(resourcePbs []PeerBlock, noCachedPbs []PeerBlock, obsP, spreadP float64) []PeerBlock {
	return nil
}
