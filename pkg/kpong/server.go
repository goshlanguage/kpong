package kpong

import (
	"fmt"
	"net"

	"github.com/ccding/go-stun/stun"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Listen uses STUN (Session Traversal Utilities for NAT) for UDP Hole Punching to establish peer to peer connections
// See more in RFC 3489 and RFC 5389
// Or see the readme in https://github.com/ccding/go-stun
func Listen() {
	_, openAddress, err := stun.NewClient().Discover()
	if err != nil {
		panic(errors.Wrap(err, "Unable to start new server, "))
	}

	fmt.Println("Host match has begun, listening on: ", openAddress)
	port := fmt.Sprintf(":%v", openAddress.Port())

	conn, err := net.Listen("tcp", port)
	if err != nil {
		panic(errors.Wrap(err, "Failed to start server: "))
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(conn); err != nil {
		panic(errors.Wrap(err, "Failed to start server: "))
	}
}
