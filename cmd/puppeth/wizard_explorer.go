// Copyright 2017 The go-rueereum Authors
// This file is part of go-rueereum.
//
// go-rueereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-rueereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-rueereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Rue-Foundation/go-rue/log"
)

// deployExplorer creates a new block explorer based on some user input.
func (w *wizard) deployExplorer() {
	// Do some sanity check before the user wastes time on input
	if w.conf.Genesis == nil {
		log.Error("No genesis block configured")
		return
	}
	if w.conf.ruestats == "" {
		log.Error("No ruestats server configured")
		return
	}
	if w.conf.Genesis.Config.Ruehash == nil {
		log.Error("Only ruehash network supported")
		return
	}
	// Select the server to interact with
	server := w.selectServer()
	if server == "" {
		return
	}
	client := w.servers[server]

	// Retrieve any active node configurations from the server
	infos, err := checkExplorer(client, w.network)
	if err != nil {
		infos = &explorerInfos{
			nodePort: 30303, webPort: 80, webHost: client.server,
		}
	}
	existed := err == nil

	chainspec, err := newParityChainSpec(w.network, w.conf.Genesis, w.conf.bootFull)
	if err != nil {
		log.Error("Failed to create chain spec for explorer", "err", err)
		return
	}
	chain, _ := json.MarshalIndent(chainspec, "", "  ")

	// Figure out which port to listen on
	fmt.Println()
	fmt.Printf("Which port should the explorer listen on? (default = %d)\n", infos.webPort)
	infos.webPort = w.readDefaultInt(infos.webPort)

	// Figure which virtual-host to deploy ruestats on
	if infos.webHost, err = w.ensureVirtualHost(client, infos.webPort, infos.webHost); err != nil {
		log.Error("Failed to decide on explorer host", "err", err)
		return
	}
	// Figure out where the user wants to store the persistent data
	fmt.Println()
	if infos.datadir == "" {
		fmt.Printf("Where should data be stored on the remote machine?\n")
		infos.datadir = w.readString()
	} else {
		fmt.Printf("Where should data be stored on the remote machine? (default = %s)\n", infos.datadir)
		infos.datadir = w.readDefaultString(infos.datadir)
	}
	// Figure out which port to listen on
	fmt.Println()
	fmt.Printf("Which TCP/UDP port should the archive node listen on? (default = %d)\n", infos.nodePort)
	infos.nodePort = w.readDefaultInt(infos.nodePort)

	// Set a proper name to report on the stats page
	fmt.Println()
	if infos.ruestats == "" {
		fmt.Printf("What should the explorer be called on the stats page?\n")
		infos.ruestats = w.readString() + ":" + w.conf.ruestats
	} else {
		fmt.Printf("What should the explorer be called on the stats page? (default = %s)\n", infos.ruestats)
		infos.ruestats = w.readDefaultString(infos.ruestats) + ":" + w.conf.ruestats
	}
	// Try to deploy the explorer on the host
	nocache := false
	if existed {
		fmt.Println()
		fmt.Printf("Should the explorer be built from scratch (y/n)? (default = no)\n")
		nocache = w.readDefaultString("n") != "n"
	}
	if out, err := deployExplorer(client, w.network, chain, infos, nocache); err != nil {
		log.Error("Failed to deploy explorer container", "err", err)
		if len(out) > 0 {
			fmt.Printf("%s\n", out)
		}
		return
	}
	// All ok, run a network scan to pick any changes up
	log.Info("Waiting for node to finish booting")
	time.Sleep(3 * time.Second)

	w.networkStats()
}