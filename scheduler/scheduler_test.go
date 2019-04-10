package main

import (
	"log"
	"testing"
)

func TestFilterPeersBlocks(t *testing.T) {

	ps := []Peer{Peer{"peer1"}, Peer{"peer3"}}
	state := State{}
	initStateA(&state)

	pbs := filterPeersBlocks(ps, state)

	// expected lenght of peerblocks after filer is 6 (3 blocks from p1 and 3
	// blocks from p3)
	expLen := 6
	if len(pbs) != expLen {
		t.Errorf("set of PeerBlocks has wrong size (%v != %v)", expLen, len(pbs))
		log.Println(pbs)
	}

	for _, pb := range pbs {
		if pb.peerId == "peer2" {
			t.Error("peer2 should not be included in the set after filtering")
			log.Println(pbs)
		}
	}
}

// sets state A
func initStateA(s *State) {
	// peer1 is storing 3 blocks of resource r0
	s.addPeerBlockEntry(Peer{"peer1"}, Block{"bA"}, "r0")
	s.addPeerBlockEntry(Peer{"peer1"}, Block{"bB"}, "r0")
	s.addPeerBlockEntry(Peer{"peer1"}, Block{"bC"}, "r0")

	// peer2 is storing 3 blocks of resource r0 and 2 blocks of r1
	s.addPeerBlockEntry(Peer{"peer2"}, Block{"bA"}, "r0")
	s.addPeerBlockEntry(Peer{"peer2"}, Block{"bB"}, "r0")
	s.addPeerBlockEntry(Peer{"peer2"}, Block{"bC"}, "r0")
	s.addPeerBlockEntry(Peer{"peer2"}, Block{"bD"}, "r1")
	s.addPeerBlockEntry(Peer{"peer2"}, Block{"bE"}, "r1")

	// peer3 is storing 4 blocks of r1
	s.addPeerBlockEntry(Peer{"peer3"}, Block{"bD"}, "r1")
	s.addPeerBlockEntry(Peer{"peer3"}, Block{"bE"}, "r1")
	s.addPeerBlockEntry(Peer{"peer3"}, Block{"bF"}, "r1")
}
