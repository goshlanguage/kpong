package kpong

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Listen uses STUN (Session Traversal Utilities for NAT) for UDP Hole Punching to establish peer to peer connections
// See more in RFC 3489 and RFC 5389
// Or see the readme in https://github.com/ccding/go-stun
func Listen(port string) {
	fmt.Println("Starting server")
	conn, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to start listener: "))
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(conn); err != nil {
		fmt.Println(errors.Wrap(err, "Failed to start server: "))
	}
}
