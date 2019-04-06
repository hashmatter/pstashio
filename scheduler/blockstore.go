package main

import (
	"context"
	bserv "github.com/ipfs/go-blockservice"
	ds "github.com/ipfs/go-datastore"
	syncds "github.com/ipfs/go-datastore/sync"
	bs "github.com/ipfs/go-ipfs-blockstore"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	ipld "github.com/ipfs/go-ipld-format"
	dag "github.com/ipfs/go-merkledag"
	"io"
	"os"
)

func addResource(path string) (ipld.DAGService, error) {
	blockStore := bs.NewBlockstore(syncds.MutexWrap(ds.NewMapDatastore()))
	blockServ := bserv.New(blockStore, offline.Exchange(blockStore))
	dagServ := dag.NewDAGService(blockServ)

	fd, err := os.Open(path)
	if err != nil {
		return dagServ, err
	}

	_, err = createDag(fd, dagServ, blockSize)
	if err != nil {
		return dagServ, err
	}

	return dagServ, nil
}

func createDag(r io.Reader, dagServ ipld.DAGService, bsize int) (ipld.Node, error) {
	bf := make([]byte, bsize)
	nodes := []*dag.ProtoNode{}

	for {
		_, err := io.ReadFull(r, bf)
		if err == io.EOF {
			break
		}

		if err != nil {
			return dag.NodeWithData(nil), err
		}

		pnode := dag.NodeWithData(bf)
		nodes = append(nodes, pnode)
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
