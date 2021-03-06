/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sizefilter

import (
	"github.com/hyperledger/fabric/orderer/common/filter"
	ab "github.com/hyperledger/fabric/protos/common"
	logging "github.com/op/go-logging"
)

var logger = logging.MustGetLogger("orderer/common/sizefilter")

// MaxBytesRule rejects messages larger than maxBytes
func MaxBytesRule(maxBytes uint32) filter.Rule {
	return &maxBytesRule{maxBytes: maxBytes}
}

type maxBytesRule struct {
	maxBytes uint32
}

func (r *maxBytesRule) Apply(message *ab.Envelope) (filter.Action, filter.Committer) {
	if size := messageByteSize(message); size > r.maxBytes {
		logger.Warningf("%d byte message payload exceeds maximum allowed %d bytes", size, r.maxBytes)
		return filter.Reject, nil
	}
	return filter.Forward, nil
}

func messageByteSize(message *ab.Envelope) uint32 {
	return uint32(len(message.Payload) + len(message.Signature))
}
