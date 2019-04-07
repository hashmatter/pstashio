package main

import (
	"context"
	pb "github.com/hashmatter/pstashio/pb"
	bservice "github.com/ipfs/go-blockservice"
	//	cid "github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	syncds "github.com/ipfs/go-datastore/sync"
	bstore "github.com/ipfs/go-ipfs-blockstore"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	ipld "github.com/ipfs/go-ipld-format"
	dag "github.com/ipfs/go-merkledag"
	"io"
	"log"
	"os"
)

// scheduler parameters
const (
	// size of the blocks in kb
	blockSize = 256
	// number of continguous blocks a peer will be responsible to provide
	blocksPeer = 10
	// obcurity factor is the number of extra (i.e. not requested) resources the
	// peer will request (some) blocks from
	obsFactor = 4
)

type server struct {
	// a server keeps the current network state
	state *netstate
	// a dag service keeps the metadata of resources stored in the network, ie.
	// how the resources are split into blocks and its links
	dag ipld.DAGService
	// a block service stores the raw data alongside with the block CIDs
	blocks bservice.BlockService
	// cids of the resource roots stored in the blockstore
	resourceRoots []ipld.Node
}

func newServerCtx() *server {
	bs := bstore.NewBlockstore(syncds.MutexWrap(ds.NewMapDatastore()))
	bserv := bservice.New(bs, offline.Exchange(bs))
	dserv := dag.NewDAGService(bserv)

	return &server{
		state:  newState(),
		dag:    dserv,
		blocks: bserv,
	}
}

// register peers in the network
func (*server) RegisterPeer(ctx context.Context, req *pb.RegisterPeerRequest) (*pb.RegisterPeerResponse, error) {
	log.Println("Received peer request: ", req)
	return &pb.RegisterPeerResponse{Status: "OK"}, nil
}

func (*server) GetProviders(ctx context.Context, req *pb.GetProvidersRequest) (*pb.GetProvidersResponse, error) {
	log.Printf("GetProviders (%v)\n", req)

	return &pb.GetProvidersResponse{}, nil
}

func (s *server) addResource(path string) (ipld.Node, error) {
	fd, err := os.Open(path)
	if err != nil {
		return dag.NodeWithData(nil), err
	}

	node, err := createDag(fd, s.dag, blockSize)
	if err != nil {
		return dag.NodeWithData(nil), err
	}

	s.resourceRoots = append(s.resourceRoots, node)
	return node, nil
}

func createDag(r io.Reader, dagServ ipld.DAGService, bsize int) (ipld.Node, error) {
	bf := make([]byte, bsize)
	nodes := []*dag.ProtoNode{}

	for {
		_, err := io.ReadFull(r, bf)
		if err == io.EOF {
			break
		}

		if err == io.ErrUnexpectedEOF {
			break
		}

		if err != nil {
			return dag.NodeWithData(nil), err
		}

		// needed because for... does not copy pointer contents
		b := make([]byte, bsize)
		copy(b, bf)
		n := dag.NodeWithData(b)
		nodes = append(nodes, n)
	}

	ctx := context.Background()
	root := dag.NodeWithData(nil)

	for _, n := range nodes {
		root.AddNodeLink(n.Cid().String(), n)
		err := dagServ.Add(ctx, n)
		if err != nil {
			return dag.NodeWithData(nil), err
		}
	}

	err := dagServ.Add(ctx, root)
	if err != nil {
		return dag.NodeWithData(nil), err
	}
	return root, nil
}
