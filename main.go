package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

// Standard Channels
// VERSION_ROLLING_BITS = 16
// NONCE_BITS = 32

type MessageFrame struct {
	ExtensionType uint16
	MsgType       uint8
	// MsgLength     uint24
	Payload []byte
}

type SignatureNoiseMessage struct {
	Version       uint16
	ValidFrom     uint32
	NotValidAfter uint32
	// Signature     ed25519
}

type Certificate struct {
	Version       uint16
	ValidFrom     uint32
	NotValidAfter uint32
	// PublicKey PUBKEY
	// AuthorityPublicKey PUBKEY
	// Signature ed25519
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	// Check for config
	// Check for pool
	// Start server

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleConnection(c)
	}
}
