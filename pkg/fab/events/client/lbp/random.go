/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package lbp

import (
	"math/rand"

	"github.com/hyperledger/fabric-sdk-go/pkg/context/api/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/logging"
)

var logger = logging.NewLogger("fabric_sdk_go")

// Random implements a random load-balance policy
type Random struct {
}

// NewRandom returns a new Random load-balance policy
func NewRandom() *Random {
	return &Random{}
}

// Choose randomly chooses a peer from the list of peers
func (lbp *Random) Choose(peers []fab.Peer) (fab.Peer, error) {
	if len(peers) == 0 {
		logger.Warnf("No peers to choose from!")
		return nil, nil
	}

	index := rand.Intn(len(peers))
	logger.Debugf("Choosing peer at index %d", index)
	return peers[index], nil
}
