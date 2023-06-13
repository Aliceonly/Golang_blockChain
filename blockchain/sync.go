package blockchain

import (
	"fmt"
	"log"
		"net"
		"sync"
)

type Peer struct {
	conn          net.Conn
	blockchain    *Blockchain
	incomingMsgs  chan []byte
	outgoingMsgs  chan []byte
}

type Network struct {
	peers  map[string]*Peer
	bc     *Blockchain
	mutex  sync.Mutex
}

func NewNetwork(bc *Blockchain) *Network {
	return &Network{
		peers: make(map[string]*Peer),
		bc:    bc,
		mutex: sync.Mutex{},
	}
}

func (netwk *Network) StartListening(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go netwk.HandleConnection(conn)
	}
}

func (net *Network) HandleConnection(conn net.Conn) {
	peer := &Peer{
		conn:         conn,
		blockchain:   net.bc,
		incomingMsgs: make(chan []byte),
		outgoingMsgs: make(chan []byte),
	}

	go peer.readMessages()
	go peer.writeMessages()


	net.mutex.Lock()
	net.peers[conn.RemoteAddr().String()] = peer
	net.mutex.Unlock()
}

func (peer *Peer) readMessages() {
	// Read incoming messages from the connection
}

func (peer *Peer) writeMessages() {
	// Write outgoing messages to the connection
}
