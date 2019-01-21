
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/token/client"
)

type DeliverFiltered struct {
	SendStub        func(*common.Envelope) error
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		arg1 *common.Envelope
	}
	sendReturns struct {
		result1 error
	}
	sendReturnsOnCall map[int]struct {
		result1 error
	}
	RecvStub        func() (*pb.DeliverResponse, error)
	recvMutex       sync.RWMutex
	recvArgsForCall []struct{}
	recvReturns     struct {
		result1 *pb.DeliverResponse
		result2 error
	}
	recvReturnsOnCall map[int]struct {
		result1 *pb.DeliverResponse
		result2 error
	}
	CloseSendStub        func() error
	closeSendMutex       sync.RWMutex
	closeSendArgsForCall []struct{}
	closeSendReturns     struct {
		result1 error
	}
	closeSendReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeliverFiltered) Send(arg1 *common.Envelope) error {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		arg1 *common.Envelope
	}{arg1})
	fake.recordInvocation("Send", []interface{}{arg1})
	fake.sendMutex.Unlock()
	if fake.SendStub != nil {
		return fake.SendStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendReturns.result1
}

func (fake *DeliverFiltered) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *DeliverFiltered) SendArgsForCall(i int) *common.Envelope {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return fake.sendArgsForCall[i].arg1
}

func (fake *DeliverFiltered) SendReturns(result1 error) {
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 error
	}{result1}
}

func (fake *DeliverFiltered) SendReturnsOnCall(i int, result1 error) {
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DeliverFiltered) Recv() (*pb.DeliverResponse, error) {
	fake.recvMutex.Lock()
	ret, specificReturn := fake.recvReturnsOnCall[len(fake.recvArgsForCall)]
	fake.recvArgsForCall = append(fake.recvArgsForCall, struct{}{})
	fake.recordInvocation("Recv", []interface{}{})
	fake.recvMutex.Unlock()
	if fake.RecvStub != nil {
		return fake.RecvStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.recvReturns.result1, fake.recvReturns.result2
}

func (fake *DeliverFiltered) RecvCallCount() int {
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	return len(fake.recvArgsForCall)
}

func (fake *DeliverFiltered) RecvReturns(result1 *pb.DeliverResponse, result2 error) {
	fake.RecvStub = nil
	fake.recvReturns = struct {
		result1 *pb.DeliverResponse
		result2 error
	}{result1, result2}
}

func (fake *DeliverFiltered) RecvReturnsOnCall(i int, result1 *pb.DeliverResponse, result2 error) {
	fake.RecvStub = nil
	if fake.recvReturnsOnCall == nil {
		fake.recvReturnsOnCall = make(map[int]struct {
			result1 *pb.DeliverResponse
			result2 error
		})
	}
	fake.recvReturnsOnCall[i] = struct {
		result1 *pb.DeliverResponse
		result2 error
	}{result1, result2}
}

func (fake *DeliverFiltered) CloseSend() error {
	fake.closeSendMutex.Lock()
	ret, specificReturn := fake.closeSendReturnsOnCall[len(fake.closeSendArgsForCall)]
	fake.closeSendArgsForCall = append(fake.closeSendArgsForCall, struct{}{})
	fake.recordInvocation("CloseSend", []interface{}{})
	fake.closeSendMutex.Unlock()
	if fake.CloseSendStub != nil {
		return fake.CloseSendStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeSendReturns.result1
}

func (fake *DeliverFiltered) CloseSendCallCount() int {
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	return len(fake.closeSendArgsForCall)
}

func (fake *DeliverFiltered) CloseSendReturns(result1 error) {
	fake.CloseSendStub = nil
	fake.closeSendReturns = struct {
		result1 error
	}{result1}
}

func (fake *DeliverFiltered) CloseSendReturnsOnCall(i int, result1 error) {
	fake.CloseSendStub = nil
	if fake.closeSendReturnsOnCall == nil {
		fake.closeSendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeSendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DeliverFiltered) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeliverFiltered) recordInvocation(key string, args []interface{}) {
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

var _ client.DeliverFiltered = new(DeliverFiltered)
