build:
	protoc -I=. --go_out=. ./pb/p2p.proto
