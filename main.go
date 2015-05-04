// Copyright (c) 2014-2015 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"github.com/btcsuite/btcd/wire"
	"github.com/davecgh/go-spew/spew"
	"log"
	"net"
)

//WHAT I WANT TO DO NEXT IS TO USE ANOTHER LIBRARY, AND CONNECT TO LOCALHOST

func main() {

	tcpAddr := &net.TCPAddr{IP: net.ParseIP("bitseed.xf2.org"), Port: 8333}
	netAddress, err := wire.NewNetAddress(tcpAddr, wire.SFNodeNetwork)
	pver := wire.ProtocolVersion
	btcnet := wire.MainNet

	// Ensure the command is expected value.
	wantCmd := "addr"
	msg := wire.NewMsgAddr()
	if cmd := msg.Command(); cmd != wantCmd {
		log.Print("NewMsgAddr: wrong command - got %v want %v",
			cmd, wantCmd)
	}

	err = msg.AddAddress(netAddress)
	if err != nil {
		log.Print("There was an error adding this address")
	}

	var buf bytes.Buffer
	err = wire.WriteMessage(&buf, msg, pver, btcnet)
	if err != nil {
		log.Print("WriteMessage  error %v", err)
	}

	rbuf := bytes.NewReader(buf.Bytes())

	resultedBytes, resultMsg, _, err := wire.ReadMessageN(rbuf, pver, btcnet)

	if err != nil {
		log.Printf("ReadMessage error %v", err)
	} else {
		log.Printf("ReadMessage result msg %v",
			spew.Sdump(resultMsg))
	}

	if resultedBytes > 0 {
		log.Print("all good")
	}

}
