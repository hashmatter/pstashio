module scheduler

require (
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/hashmatter/pstashio/pb v0.0.0
	github.com/ipfs/go-blockservice v0.0.3
	github.com/ipfs/go-cid v0.0.1
	github.com/ipfs/go-datastore v0.0.1
	github.com/ipfs/go-ipfs-blockstore v0.0.1
	github.com/ipfs/go-ipfs-exchange-offline v0.0.1
	github.com/ipfs/go-ipld-format v0.0.1
	github.com/ipfs/go-merkledag v0.0.3
	google.golang.org/grpc v1.19.1
)

replace github.com/hashmatter/pstashio/pb => ../pb
