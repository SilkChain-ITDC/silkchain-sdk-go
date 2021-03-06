/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package resource

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/context/api/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/txn"

	"github.com/hyperledger/fabric-sdk-go/pkg/fab/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateChaincodeInstallProposal(t *testing.T) {
	c := setupTestClient()
	peer := mocks.MockPeer{MockName: "Peer1", MockURL: "peer1.example.com", MockRoles: []string{}, MockCert: nil, Payload: []byte("A"), Status: 200}

	request := ChaincodeInstallRequest{
		Name:    "examplecc",
		Path:    "github.com/examplecc",
		Version: "1",
		Package: &ChaincodePackage{},
	}

	txid, err := txn.NewHeader(c.clientContext, fab.SystemChannel)
	assert.Nil(t, err, "create transaction ID failed")

	prop, err := CreateChaincodeInstallProposal(txid, request)
	assert.Nil(t, err, "CreateChaincodeInstallProposal failed")

	_, err = txn.SendProposal(c.clientContext, prop, []fab.ProposalProcessor{&peer})
	assert.Nil(t, err, "sending mock proposal failed")
}
