
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric/gossip/api"
	"github.com/hyperledger/fabric/gossip/comm"
	"github.com/hyperledger/fabric/gossip/common"
	"github.com/hyperledger/fabric/gossip/discovery"
	"github.com/hyperledger/fabric/gossip/filter"
	"github.com/hyperledger/fabric/gossip/gossip"
	proto "github.com/hyperledger/fabric/protos/gossip"
)

type Gossip struct {
	SelfMembershipInfoStub        func() discovery.NetworkMember
	selfMembershipInfoMutex       sync.RWMutex
	selfMembershipInfoArgsForCall []struct{}
	selfMembershipInfoReturns     struct {
		result1 discovery.NetworkMember
	}
	selfMembershipInfoReturnsOnCall map[int]struct {
		result1 discovery.NetworkMember
	}
	SelfChannelInfoStub        func(common.ChainID) *proto.SignedGossipMessage
	selfChannelInfoMutex       sync.RWMutex
	selfChannelInfoArgsForCall []struct {
		arg1 common.ChainID
	}
	selfChannelInfoReturns struct {
		result1 *proto.SignedGossipMessage
	}
	selfChannelInfoReturnsOnCall map[int]struct {
		result1 *proto.SignedGossipMessage
	}
	SendStub        func(msg *proto.GossipMessage, peers ...*comm.RemotePeer)
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		msg   *proto.GossipMessage
		peers []*comm.RemotePeer
	}
	SendByCriteriaStub        func(*proto.SignedGossipMessage, gossip.SendCriteria) error
	sendByCriteriaMutex       sync.RWMutex
	sendByCriteriaArgsForCall []struct {
		arg1 *proto.SignedGossipMessage
		arg2 gossip.SendCriteria
	}
	sendByCriteriaReturns struct {
		result1 error
	}
	sendByCriteriaReturnsOnCall map[int]struct {
		result1 error
	}
	PeersStub        func() []discovery.NetworkMember
	peersMutex       sync.RWMutex
	peersArgsForCall []struct{}
	peersReturns     struct {
		result1 []discovery.NetworkMember
	}
	peersReturnsOnCall map[int]struct {
		result1 []discovery.NetworkMember
	}
	PeersOfChannelStub        func(common.ChainID) []discovery.NetworkMember
	peersOfChannelMutex       sync.RWMutex
	peersOfChannelArgsForCall []struct {
		arg1 common.ChainID
	}
	peersOfChannelReturns struct {
		result1 []discovery.NetworkMember
	}
	peersOfChannelReturnsOnCall map[int]struct {
		result1 []discovery.NetworkMember
	}
	UpdateMetadataStub        func(metadata []byte)
	updateMetadataMutex       sync.RWMutex
	updateMetadataArgsForCall []struct {
		metadata []byte
	}
	UpdateLedgerHeightStub        func(height uint64, chainID common.ChainID)
	updateLedgerHeightMutex       sync.RWMutex
	updateLedgerHeightArgsForCall []struct {
		height  uint64
		chainID common.ChainID
	}
	UpdateChaincodesStub        func(chaincode []*proto.Chaincode, chainID common.ChainID)
	updateChaincodesMutex       sync.RWMutex
	updateChaincodesArgsForCall []struct {
		chaincode []*proto.Chaincode
		chainID   common.ChainID
	}
	GossipStub        func(msg *proto.GossipMessage)
	gossipMutex       sync.RWMutex
	gossipArgsForCall []struct {
		msg *proto.GossipMessage
	}
	PeerFilterStub        func(channel common.ChainID, messagePredicate api.SubChannelSelectionCriteria) (filter.RoutingFilter, error)
	peerFilterMutex       sync.RWMutex
	peerFilterArgsForCall []struct {
		channel          common.ChainID
		messagePredicate api.SubChannelSelectionCriteria
	}
	peerFilterReturns struct {
		result1 filter.RoutingFilter
		result2 error
	}
	peerFilterReturnsOnCall map[int]struct {
		result1 filter.RoutingFilter
		result2 error
	}
	AcceptStub        func(acceptor common.MessageAcceptor, passThrough bool) (<-chan *proto.GossipMessage, <-chan proto.ReceivedMessage)
	acceptMutex       sync.RWMutex
	acceptArgsForCall []struct {
		acceptor    common.MessageAcceptor
		passThrough bool
	}
	acceptReturns struct {
		result1 <-chan *proto.GossipMessage
		result2 <-chan proto.ReceivedMessage
	}
	acceptReturnsOnCall map[int]struct {
		result1 <-chan *proto.GossipMessage
		result2 <-chan proto.ReceivedMessage
	}
	JoinChanStub        func(joinMsg api.JoinChannelMessage, chainID common.ChainID)
	joinChanMutex       sync.RWMutex
	joinChanArgsForCall []struct {
		joinMsg api.JoinChannelMessage
		chainID common.ChainID
	}
	LeaveChanStub        func(chainID common.ChainID)
	leaveChanMutex       sync.RWMutex
	leaveChanArgsForCall []struct {
		chainID common.ChainID
	}
	SuspectPeersStub        func(s api.PeerSuspector)
	suspectPeersMutex       sync.RWMutex
	suspectPeersArgsForCall []struct {
		s api.PeerSuspector
	}
	IdentityInfoStub        func() api.PeerIdentitySet
	identityInfoMutex       sync.RWMutex
	identityInfoArgsForCall []struct{}
	identityInfoReturns     struct {
		result1 api.PeerIdentitySet
	}
	identityInfoReturnsOnCall map[int]struct {
		result1 api.PeerIdentitySet
	}
	StopStub         func()
	stopMutex        sync.RWMutex
	stopArgsForCall  []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Gossip) SelfMembershipInfo() discovery.NetworkMember {
	fake.selfMembershipInfoMutex.Lock()
	ret, specificReturn := fake.selfMembershipInfoReturnsOnCall[len(fake.selfMembershipInfoArgsForCall)]
	fake.selfMembershipInfoArgsForCall = append(fake.selfMembershipInfoArgsForCall, struct{}{})
	fake.recordInvocation("SelfMembershipInfo", []interface{}{})
	fake.selfMembershipInfoMutex.Unlock()
	if fake.SelfMembershipInfoStub != nil {
		return fake.SelfMembershipInfoStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.selfMembershipInfoReturns.result1
}

func (fake *Gossip) SelfMembershipInfoCallCount() int {
	fake.selfMembershipInfoMutex.RLock()
	defer fake.selfMembershipInfoMutex.RUnlock()
	return len(fake.selfMembershipInfoArgsForCall)
}

func (fake *Gossip) SelfMembershipInfoReturns(result1 discovery.NetworkMember) {
	fake.SelfMembershipInfoStub = nil
	fake.selfMembershipInfoReturns = struct {
		result1 discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) SelfMembershipInfoReturnsOnCall(i int, result1 discovery.NetworkMember) {
	fake.SelfMembershipInfoStub = nil
	if fake.selfMembershipInfoReturnsOnCall == nil {
		fake.selfMembershipInfoReturnsOnCall = make(map[int]struct {
			result1 discovery.NetworkMember
		})
	}
	fake.selfMembershipInfoReturnsOnCall[i] = struct {
		result1 discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) SelfChannelInfo(arg1 common.ChainID) *proto.SignedGossipMessage {
	fake.selfChannelInfoMutex.Lock()
	ret, specificReturn := fake.selfChannelInfoReturnsOnCall[len(fake.selfChannelInfoArgsForCall)]
	fake.selfChannelInfoArgsForCall = append(fake.selfChannelInfoArgsForCall, struct {
		arg1 common.ChainID
	}{arg1})
	fake.recordInvocation("SelfChannelInfo", []interface{}{arg1})
	fake.selfChannelInfoMutex.Unlock()
	if fake.SelfChannelInfoStub != nil {
		return fake.SelfChannelInfoStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.selfChannelInfoReturns.result1
}

func (fake *Gossip) SelfChannelInfoCallCount() int {
	fake.selfChannelInfoMutex.RLock()
	defer fake.selfChannelInfoMutex.RUnlock()
	return len(fake.selfChannelInfoArgsForCall)
}

func (fake *Gossip) SelfChannelInfoArgsForCall(i int) common.ChainID {
	fake.selfChannelInfoMutex.RLock()
	defer fake.selfChannelInfoMutex.RUnlock()
	return fake.selfChannelInfoArgsForCall[i].arg1
}

func (fake *Gossip) SelfChannelInfoReturns(result1 *proto.SignedGossipMessage) {
	fake.SelfChannelInfoStub = nil
	fake.selfChannelInfoReturns = struct {
		result1 *proto.SignedGossipMessage
	}{result1}
}

func (fake *Gossip) SelfChannelInfoReturnsOnCall(i int, result1 *proto.SignedGossipMessage) {
	fake.SelfChannelInfoStub = nil
	if fake.selfChannelInfoReturnsOnCall == nil {
		fake.selfChannelInfoReturnsOnCall = make(map[int]struct {
			result1 *proto.SignedGossipMessage
		})
	}
	fake.selfChannelInfoReturnsOnCall[i] = struct {
		result1 *proto.SignedGossipMessage
	}{result1}
}

func (fake *Gossip) Send(msg *proto.GossipMessage, peers ...*comm.RemotePeer) {
	fake.sendMutex.Lock()
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		msg   *proto.GossipMessage
		peers []*comm.RemotePeer
	}{msg, peers})
	fake.recordInvocation("Send", []interface{}{msg, peers})
	fake.sendMutex.Unlock()
	if fake.SendStub != nil {
		fake.SendStub(msg, peers...)
	}
}

func (fake *Gossip) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *Gossip) SendArgsForCall(i int) (*proto.GossipMessage, []*comm.RemotePeer) {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return fake.sendArgsForCall[i].msg, fake.sendArgsForCall[i].peers
}

func (fake *Gossip) SendByCriteria(arg1 *proto.SignedGossipMessage, arg2 gossip.SendCriteria) error {
	fake.sendByCriteriaMutex.Lock()
	ret, specificReturn := fake.sendByCriteriaReturnsOnCall[len(fake.sendByCriteriaArgsForCall)]
	fake.sendByCriteriaArgsForCall = append(fake.sendByCriteriaArgsForCall, struct {
		arg1 *proto.SignedGossipMessage
		arg2 gossip.SendCriteria
	}{arg1, arg2})
	fake.recordInvocation("SendByCriteria", []interface{}{arg1, arg2})
	fake.sendByCriteriaMutex.Unlock()
	if fake.SendByCriteriaStub != nil {
		return fake.SendByCriteriaStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendByCriteriaReturns.result1
}

func (fake *Gossip) SendByCriteriaCallCount() int {
	fake.sendByCriteriaMutex.RLock()
	defer fake.sendByCriteriaMutex.RUnlock()
	return len(fake.sendByCriteriaArgsForCall)
}

func (fake *Gossip) SendByCriteriaArgsForCall(i int) (*proto.SignedGossipMessage, gossip.SendCriteria) {
	fake.sendByCriteriaMutex.RLock()
	defer fake.sendByCriteriaMutex.RUnlock()
	return fake.sendByCriteriaArgsForCall[i].arg1, fake.sendByCriteriaArgsForCall[i].arg2
}

func (fake *Gossip) SendByCriteriaReturns(result1 error) {
	fake.SendByCriteriaStub = nil
	fake.sendByCriteriaReturns = struct {
		result1 error
	}{result1}
}

func (fake *Gossip) SendByCriteriaReturnsOnCall(i int, result1 error) {
	fake.SendByCriteriaStub = nil
	if fake.sendByCriteriaReturnsOnCall == nil {
		fake.sendByCriteriaReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendByCriteriaReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Gossip) Peers() []discovery.NetworkMember {
	fake.peersMutex.Lock()
	ret, specificReturn := fake.peersReturnsOnCall[len(fake.peersArgsForCall)]
	fake.peersArgsForCall = append(fake.peersArgsForCall, struct{}{})
	fake.recordInvocation("Peers", []interface{}{})
	fake.peersMutex.Unlock()
	if fake.PeersStub != nil {
		return fake.PeersStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.peersReturns.result1
}

func (fake *Gossip) PeersCallCount() int {
	fake.peersMutex.RLock()
	defer fake.peersMutex.RUnlock()
	return len(fake.peersArgsForCall)
}

func (fake *Gossip) PeersReturns(result1 []discovery.NetworkMember) {
	fake.PeersStub = nil
	fake.peersReturns = struct {
		result1 []discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) PeersReturnsOnCall(i int, result1 []discovery.NetworkMember) {
	fake.PeersStub = nil
	if fake.peersReturnsOnCall == nil {
		fake.peersReturnsOnCall = make(map[int]struct {
			result1 []discovery.NetworkMember
		})
	}
	fake.peersReturnsOnCall[i] = struct {
		result1 []discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) PeersOfChannel(arg1 common.ChainID) []discovery.NetworkMember {
	fake.peersOfChannelMutex.Lock()
	ret, specificReturn := fake.peersOfChannelReturnsOnCall[len(fake.peersOfChannelArgsForCall)]
	fake.peersOfChannelArgsForCall = append(fake.peersOfChannelArgsForCall, struct {
		arg1 common.ChainID
	}{arg1})
	fake.recordInvocation("PeersOfChannel", []interface{}{arg1})
	fake.peersOfChannelMutex.Unlock()
	if fake.PeersOfChannelStub != nil {
		return fake.PeersOfChannelStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.peersOfChannelReturns.result1
}

func (fake *Gossip) PeersOfChannelCallCount() int {
	fake.peersOfChannelMutex.RLock()
	defer fake.peersOfChannelMutex.RUnlock()
	return len(fake.peersOfChannelArgsForCall)
}

func (fake *Gossip) PeersOfChannelArgsForCall(i int) common.ChainID {
	fake.peersOfChannelMutex.RLock()
	defer fake.peersOfChannelMutex.RUnlock()
	return fake.peersOfChannelArgsForCall[i].arg1
}

func (fake *Gossip) PeersOfChannelReturns(result1 []discovery.NetworkMember) {
	fake.PeersOfChannelStub = nil
	fake.peersOfChannelReturns = struct {
		result1 []discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) PeersOfChannelReturnsOnCall(i int, result1 []discovery.NetworkMember) {
	fake.PeersOfChannelStub = nil
	if fake.peersOfChannelReturnsOnCall == nil {
		fake.peersOfChannelReturnsOnCall = make(map[int]struct {
			result1 []discovery.NetworkMember
		})
	}
	fake.peersOfChannelReturnsOnCall[i] = struct {
		result1 []discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) UpdateMetadata(metadata []byte) {
	var metadataCopy []byte
	if metadata != nil {
		metadataCopy = make([]byte, len(metadata))
		copy(metadataCopy, metadata)
	}
	fake.updateMetadataMutex.Lock()
	fake.updateMetadataArgsForCall = append(fake.updateMetadataArgsForCall, struct {
		metadata []byte
	}{metadataCopy})
	fake.recordInvocation("UpdateMetadata", []interface{}{metadataCopy})
	fake.updateMetadataMutex.Unlock()
	if fake.UpdateMetadataStub != nil {
		fake.UpdateMetadataStub(metadata)
	}
}

func (fake *Gossip) UpdateMetadataCallCount() int {
	fake.updateMetadataMutex.RLock()
	defer fake.updateMetadataMutex.RUnlock()
	return len(fake.updateMetadataArgsForCall)
}

func (fake *Gossip) UpdateMetadataArgsForCall(i int) []byte {
	fake.updateMetadataMutex.RLock()
	defer fake.updateMetadataMutex.RUnlock()
	return fake.updateMetadataArgsForCall[i].metadata
}

func (fake *Gossip) UpdateLedgerHeight(height uint64, chainID common.ChainID) {
	fake.updateLedgerHeightMutex.Lock()
	fake.updateLedgerHeightArgsForCall = append(fake.updateLedgerHeightArgsForCall, struct {
		height  uint64
		chainID common.ChainID
	}{height, chainID})
	fake.recordInvocation("UpdateLedgerHeight", []interface{}{height, chainID})
	fake.updateLedgerHeightMutex.Unlock()
	if fake.UpdateLedgerHeightStub != nil {
		fake.UpdateLedgerHeightStub(height, chainID)
	}
}

func (fake *Gossip) UpdateLedgerHeightCallCount() int {
	fake.updateLedgerHeightMutex.RLock()
	defer fake.updateLedgerHeightMutex.RUnlock()
	return len(fake.updateLedgerHeightArgsForCall)
}

func (fake *Gossip) UpdateLedgerHeightArgsForCall(i int) (uint64, common.ChainID) {
	fake.updateLedgerHeightMutex.RLock()
	defer fake.updateLedgerHeightMutex.RUnlock()
	return fake.updateLedgerHeightArgsForCall[i].height, fake.updateLedgerHeightArgsForCall[i].chainID
}

func (fake *Gossip) UpdateChaincodes(chaincode []*proto.Chaincode, chainID common.ChainID) {
	var chaincodeCopy []*proto.Chaincode
	if chaincode != nil {
		chaincodeCopy = make([]*proto.Chaincode, len(chaincode))
		copy(chaincodeCopy, chaincode)
	}
	fake.updateChaincodesMutex.Lock()
	fake.updateChaincodesArgsForCall = append(fake.updateChaincodesArgsForCall, struct {
		chaincode []*proto.Chaincode
		chainID   common.ChainID
	}{chaincodeCopy, chainID})
	fake.recordInvocation("UpdateChaincodes", []interface{}{chaincodeCopy, chainID})
	fake.updateChaincodesMutex.Unlock()
	if fake.UpdateChaincodesStub != nil {
		fake.UpdateChaincodesStub(chaincode, chainID)
	}
}

func (fake *Gossip) UpdateChaincodesCallCount() int {
	fake.updateChaincodesMutex.RLock()
	defer fake.updateChaincodesMutex.RUnlock()
	return len(fake.updateChaincodesArgsForCall)
}

func (fake *Gossip) UpdateChaincodesArgsForCall(i int) ([]*proto.Chaincode, common.ChainID) {
	fake.updateChaincodesMutex.RLock()
	defer fake.updateChaincodesMutex.RUnlock()
	return fake.updateChaincodesArgsForCall[i].chaincode, fake.updateChaincodesArgsForCall[i].chainID
}

func (fake *Gossip) Gossip(msg *proto.GossipMessage) {
	fake.gossipMutex.Lock()
	fake.gossipArgsForCall = append(fake.gossipArgsForCall, struct {
		msg *proto.GossipMessage
	}{msg})
	fake.recordInvocation("Gossip", []interface{}{msg})
	fake.gossipMutex.Unlock()
	if fake.GossipStub != nil {
		fake.GossipStub(msg)
	}
}

func (fake *Gossip) GossipCallCount() int {
	fake.gossipMutex.RLock()
	defer fake.gossipMutex.RUnlock()
	return len(fake.gossipArgsForCall)
}

func (fake *Gossip) GossipArgsForCall(i int) *proto.GossipMessage {
	fake.gossipMutex.RLock()
	defer fake.gossipMutex.RUnlock()
	return fake.gossipArgsForCall[i].msg
}

func (fake *Gossip) PeerFilter(channel common.ChainID, messagePredicate api.SubChannelSelectionCriteria) (filter.RoutingFilter, error) {
	fake.peerFilterMutex.Lock()
	ret, specificReturn := fake.peerFilterReturnsOnCall[len(fake.peerFilterArgsForCall)]
	fake.peerFilterArgsForCall = append(fake.peerFilterArgsForCall, struct {
		channel          common.ChainID
		messagePredicate api.SubChannelSelectionCriteria
	}{channel, messagePredicate})
	fake.recordInvocation("PeerFilter", []interface{}{channel, messagePredicate})
	fake.peerFilterMutex.Unlock()
	if fake.PeerFilterStub != nil {
		return fake.PeerFilterStub(channel, messagePredicate)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.peerFilterReturns.result1, fake.peerFilterReturns.result2
}

func (fake *Gossip) PeerFilterCallCount() int {
	fake.peerFilterMutex.RLock()
	defer fake.peerFilterMutex.RUnlock()
	return len(fake.peerFilterArgsForCall)
}

func (fake *Gossip) PeerFilterArgsForCall(i int) (common.ChainID, api.SubChannelSelectionCriteria) {
	fake.peerFilterMutex.RLock()
	defer fake.peerFilterMutex.RUnlock()
	return fake.peerFilterArgsForCall[i].channel, fake.peerFilterArgsForCall[i].messagePredicate
}

func (fake *Gossip) PeerFilterReturns(result1 filter.RoutingFilter, result2 error) {
	fake.PeerFilterStub = nil
	fake.peerFilterReturns = struct {
		result1 filter.RoutingFilter
		result2 error
	}{result1, result2}
}

func (fake *Gossip) PeerFilterReturnsOnCall(i int, result1 filter.RoutingFilter, result2 error) {
	fake.PeerFilterStub = nil
	if fake.peerFilterReturnsOnCall == nil {
		fake.peerFilterReturnsOnCall = make(map[int]struct {
			result1 filter.RoutingFilter
			result2 error
		})
	}
	fake.peerFilterReturnsOnCall[i] = struct {
		result1 filter.RoutingFilter
		result2 error
	}{result1, result2}
}

func (fake *Gossip) Accept(acceptor common.MessageAcceptor, passThrough bool) (<-chan *proto.GossipMessage, <-chan proto.ReceivedMessage) {
	fake.acceptMutex.Lock()
	ret, specificReturn := fake.acceptReturnsOnCall[len(fake.acceptArgsForCall)]
	fake.acceptArgsForCall = append(fake.acceptArgsForCall, struct {
		acceptor    common.MessageAcceptor
		passThrough bool
	}{acceptor, passThrough})
	fake.recordInvocation("Accept", []interface{}{acceptor, passThrough})
	fake.acceptMutex.Unlock()
	if fake.AcceptStub != nil {
		return fake.AcceptStub(acceptor, passThrough)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.acceptReturns.result1, fake.acceptReturns.result2
}

func (fake *Gossip) AcceptCallCount() int {
	fake.acceptMutex.RLock()
	defer fake.acceptMutex.RUnlock()
	return len(fake.acceptArgsForCall)
}

func (fake *Gossip) AcceptArgsForCall(i int) (common.MessageAcceptor, bool) {
	fake.acceptMutex.RLock()
	defer fake.acceptMutex.RUnlock()
	return fake.acceptArgsForCall[i].acceptor, fake.acceptArgsForCall[i].passThrough
}

func (fake *Gossip) AcceptReturns(result1 <-chan *proto.GossipMessage, result2 <-chan proto.ReceivedMessage) {
	fake.AcceptStub = nil
	fake.acceptReturns = struct {
		result1 <-chan *proto.GossipMessage
		result2 <-chan proto.ReceivedMessage
	}{result1, result2}
}

func (fake *Gossip) AcceptReturnsOnCall(i int, result1 <-chan *proto.GossipMessage, result2 <-chan proto.ReceivedMessage) {
	fake.AcceptStub = nil
	if fake.acceptReturnsOnCall == nil {
		fake.acceptReturnsOnCall = make(map[int]struct {
			result1 <-chan *proto.GossipMessage
			result2 <-chan proto.ReceivedMessage
		})
	}
	fake.acceptReturnsOnCall[i] = struct {
		result1 <-chan *proto.GossipMessage
		result2 <-chan proto.ReceivedMessage
	}{result1, result2}
}

func (fake *Gossip) JoinChan(joinMsg api.JoinChannelMessage, chainID common.ChainID) {
	fake.joinChanMutex.Lock()
	fake.joinChanArgsForCall = append(fake.joinChanArgsForCall, struct {
		joinMsg api.JoinChannelMessage
		chainID common.ChainID
	}{joinMsg, chainID})
	fake.recordInvocation("JoinChan", []interface{}{joinMsg, chainID})
	fake.joinChanMutex.Unlock()
	if fake.JoinChanStub != nil {
		fake.JoinChanStub(joinMsg, chainID)
	}
}

func (fake *Gossip) JoinChanCallCount() int {
	fake.joinChanMutex.RLock()
	defer fake.joinChanMutex.RUnlock()
	return len(fake.joinChanArgsForCall)
}

func (fake *Gossip) JoinChanArgsForCall(i int) (api.JoinChannelMessage, common.ChainID) {
	fake.joinChanMutex.RLock()
	defer fake.joinChanMutex.RUnlock()
	return fake.joinChanArgsForCall[i].joinMsg, fake.joinChanArgsForCall[i].chainID
}

func (fake *Gossip) LeaveChan(chainID common.ChainID) {
	fake.leaveChanMutex.Lock()
	fake.leaveChanArgsForCall = append(fake.leaveChanArgsForCall, struct {
		chainID common.ChainID
	}{chainID})
	fake.recordInvocation("LeaveChan", []interface{}{chainID})
	fake.leaveChanMutex.Unlock()
	if fake.LeaveChanStub != nil {
		fake.LeaveChanStub(chainID)
	}
}

func (fake *Gossip) LeaveChanCallCount() int {
	fake.leaveChanMutex.RLock()
	defer fake.leaveChanMutex.RUnlock()
	return len(fake.leaveChanArgsForCall)
}

func (fake *Gossip) LeaveChanArgsForCall(i int) common.ChainID {
	fake.leaveChanMutex.RLock()
	defer fake.leaveChanMutex.RUnlock()
	return fake.leaveChanArgsForCall[i].chainID
}

func (fake *Gossip) SuspectPeers(s api.PeerSuspector) {
	fake.suspectPeersMutex.Lock()
	fake.suspectPeersArgsForCall = append(fake.suspectPeersArgsForCall, struct {
		s api.PeerSuspector
	}{s})
	fake.recordInvocation("SuspectPeers", []interface{}{s})
	fake.suspectPeersMutex.Unlock()
	if fake.SuspectPeersStub != nil {
		fake.SuspectPeersStub(s)
	}
}

func (fake *Gossip) SuspectPeersCallCount() int {
	fake.suspectPeersMutex.RLock()
	defer fake.suspectPeersMutex.RUnlock()
	return len(fake.suspectPeersArgsForCall)
}

func (fake *Gossip) SuspectPeersArgsForCall(i int) api.PeerSuspector {
	fake.suspectPeersMutex.RLock()
	defer fake.suspectPeersMutex.RUnlock()
	return fake.suspectPeersArgsForCall[i].s
}

func (fake *Gossip) IdentityInfo() api.PeerIdentitySet {
	fake.identityInfoMutex.Lock()
	ret, specificReturn := fake.identityInfoReturnsOnCall[len(fake.identityInfoArgsForCall)]
	fake.identityInfoArgsForCall = append(fake.identityInfoArgsForCall, struct{}{})
	fake.recordInvocation("IdentityInfo", []interface{}{})
	fake.identityInfoMutex.Unlock()
	if fake.IdentityInfoStub != nil {
		return fake.IdentityInfoStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.identityInfoReturns.result1
}

func (fake *Gossip) IdentityInfoCallCount() int {
	fake.identityInfoMutex.RLock()
	defer fake.identityInfoMutex.RUnlock()
	return len(fake.identityInfoArgsForCall)
}

func (fake *Gossip) IdentityInfoReturns(result1 api.PeerIdentitySet) {
	fake.IdentityInfoStub = nil
	fake.identityInfoReturns = struct {
		result1 api.PeerIdentitySet
	}{result1}
}

func (fake *Gossip) IdentityInfoReturnsOnCall(i int, result1 api.PeerIdentitySet) {
	fake.IdentityInfoStub = nil
	if fake.identityInfoReturnsOnCall == nil {
		fake.identityInfoReturnsOnCall = make(map[int]struct {
			result1 api.PeerIdentitySet
		})
	}
	fake.identityInfoReturnsOnCall[i] = struct {
		result1 api.PeerIdentitySet
	}{result1}
}

func (fake *Gossip) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		fake.StopStub()
	}
}

func (fake *Gossip) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *Gossip) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.selfMembershipInfoMutex.RLock()
	defer fake.selfMembershipInfoMutex.RUnlock()
	fake.selfChannelInfoMutex.RLock()
	defer fake.selfChannelInfoMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	fake.sendByCriteriaMutex.RLock()
	defer fake.sendByCriteriaMutex.RUnlock()
	fake.peersMutex.RLock()
	defer fake.peersMutex.RUnlock()
	fake.peersOfChannelMutex.RLock()
	defer fake.peersOfChannelMutex.RUnlock()
	fake.updateMetadataMutex.RLock()
	defer fake.updateMetadataMutex.RUnlock()
	fake.updateLedgerHeightMutex.RLock()
	defer fake.updateLedgerHeightMutex.RUnlock()
	fake.updateChaincodesMutex.RLock()
	defer fake.updateChaincodesMutex.RUnlock()
	fake.gossipMutex.RLock()
	defer fake.gossipMutex.RUnlock()
	fake.peerFilterMutex.RLock()
	defer fake.peerFilterMutex.RUnlock()
	fake.acceptMutex.RLock()
	defer fake.acceptMutex.RUnlock()
	fake.joinChanMutex.RLock()
	defer fake.joinChanMutex.RUnlock()
	fake.leaveChanMutex.RLock()
	defer fake.leaveChanMutex.RUnlock()
	fake.suspectPeersMutex.RLock()
	defer fake.suspectPeersMutex.RUnlock()
	fake.identityInfoMutex.RLock()
	defer fake.identityInfoMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Gossip) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ gossip.Gossip = new(Gossip)
