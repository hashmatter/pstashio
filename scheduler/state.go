package main

import (
	"errors"
	pb "github.com/hashmatter/pstashio/pb"
)

var noPeerErr = errors.New("Peer not found")
var noResourceErr = errors.New("Resource not found")

// keeps the state of the network in memory; a network state keeps a log of the
// current peers in the network and the blocks they are storing. the state keeps
// a two-dimensional mapping between peers and blocks, which allow for querying
// which peers are storing which blocks given a peer id and a block id
type netstate struct {
	peers  []peer
	blocks blocksMap

	// a resource list keeps the log of which resources are stored currently in
	// the network and its block ids
	resourceList map[string][]string
}

func newState() *netstate {
	return &netstate{}
}

// TODO: peer ID validation
func (s *netstate) addPeer(p peer) error {
	s.peers = append(s.peers, p)
	return nil
}

func (s *netstate) addBlockPeer(pid string, bid string) error {
	// checks if peer exists
	exists := false
	for _, p := range s.peers {
		if p.id == pid {
			exists = true
			break
		}
	}
	if exists == false {
		return noPeerErr
	}

	// TODO: check reps
	bp := blockPeer{blockId: bid, peerId: pid}
	s.blocks.blockPeers[bid] = bp
	s.blocks.blockPeers[pid] = bp
	return nil
}

// returns the list of IDs of the blocks of a given resource
func (s *netstate) listBlocks(rid string) ([]string, error) {
	lst := s.resourceList[rid]
	// resource does not exist
	if lst == nil {
		return []string{}, noResourceErr
	}
	return lst, nil
}

type peer struct {
	id           string
	ip           string
	cachedBlocks *[]pb.Block
}

// (time) efficient block/peer mapping
type blockPeer struct {
	blockId string
	peerId  string
}

type blocksMap struct {
	// indexes relation peer-block using block id
	blockPeers map[string]blockPeer
	// indexes relation peer-block using peer id
	peersBlock map[string]blockPeer
}
