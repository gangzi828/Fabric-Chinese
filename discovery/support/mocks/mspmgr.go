
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

	"github.com/hyperledger/fabric/msp"
	mspprotos "github.com/hyperledger/fabric/protos/msp"
)

type MSPManager struct {
	DeserializeIdentityStub        func(serializedIdentity []byte) (msp.Identity, error)
	deserializeIdentityMutex       sync.RWMutex
	deserializeIdentityArgsForCall []struct {
		serializedIdentity []byte
	}
	deserializeIdentityReturns struct {
		result1 msp.Identity
		result2 error
	}
	deserializeIdentityReturnsOnCall map[int]struct {
		result1 msp.Identity
		result2 error
	}
	IsWellFormedStub        func(identity *mspprotos.SerializedIdentity) error
	isWellFormedMutex       sync.RWMutex
	isWellFormedArgsForCall []struct {
		identity *mspprotos.SerializedIdentity
	}
	isWellFormedReturns struct {
		result1 error
	}
	isWellFormedReturnsOnCall map[int]struct {
		result1 error
	}
	SetupStub        func(msps []msp.MSP) error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		msps []msp.MSP
	}
	setupReturns struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	GetMSPsStub        func() (map[string]msp.MSP, error)
	getMSPsMutex       sync.RWMutex
	getMSPsArgsForCall []struct{}
	getMSPsReturns     struct {
		result1 map[string]msp.MSP
		result2 error
	}
	getMSPsReturnsOnCall map[int]struct {
		result1 map[string]msp.MSP
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MSPManager) DeserializeIdentity(serializedIdentity []byte) (msp.Identity, error) {
	var serializedIdentityCopy []byte
	if serializedIdentity != nil {
		serializedIdentityCopy = make([]byte, len(serializedIdentity))
		copy(serializedIdentityCopy, serializedIdentity)
	}
	fake.deserializeIdentityMutex.Lock()
	ret, specificReturn := fake.deserializeIdentityReturnsOnCall[len(fake.deserializeIdentityArgsForCall)]
	fake.deserializeIdentityArgsForCall = append(fake.deserializeIdentityArgsForCall, struct {
		serializedIdentity []byte
	}{serializedIdentityCopy})
	fake.recordInvocation("DeserializeIdentity", []interface{}{serializedIdentityCopy})
	fake.deserializeIdentityMutex.Unlock()
	if fake.DeserializeIdentityStub != nil {
		return fake.DeserializeIdentityStub(serializedIdentity)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deserializeIdentityReturns.result1, fake.deserializeIdentityReturns.result2
}

func (fake *MSPManager) DeserializeIdentityCallCount() int {
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	return len(fake.deserializeIdentityArgsForCall)
}

func (fake *MSPManager) DeserializeIdentityArgsForCall(i int) []byte {
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	return fake.deserializeIdentityArgsForCall[i].serializedIdentity
}

func (fake *MSPManager) DeserializeIdentityReturns(result1 msp.Identity, result2 error) {
	fake.DeserializeIdentityStub = nil
	fake.deserializeIdentityReturns = struct {
		result1 msp.Identity
		result2 error
	}{result1, result2}
}

func (fake *MSPManager) DeserializeIdentityReturnsOnCall(i int, result1 msp.Identity, result2 error) {
	fake.DeserializeIdentityStub = nil
	if fake.deserializeIdentityReturnsOnCall == nil {
		fake.deserializeIdentityReturnsOnCall = make(map[int]struct {
			result1 msp.Identity
			result2 error
		})
	}
	fake.deserializeIdentityReturnsOnCall[i] = struct {
		result1 msp.Identity
		result2 error
	}{result1, result2}
}

func (fake *MSPManager) IsWellFormed(identity *mspprotos.SerializedIdentity) error {
	fake.isWellFormedMutex.Lock()
	ret, specificReturn := fake.isWellFormedReturnsOnCall[len(fake.isWellFormedArgsForCall)]
	fake.isWellFormedArgsForCall = append(fake.isWellFormedArgsForCall, struct {
		identity *mspprotos.SerializedIdentity
	}{identity})
	fake.recordInvocation("IsWellFormed", []interface{}{identity})
	fake.isWellFormedMutex.Unlock()
	if fake.IsWellFormedStub != nil {
		return fake.IsWellFormedStub(identity)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isWellFormedReturns.result1
}

func (fake *MSPManager) IsWellFormedCallCount() int {
	fake.isWellFormedMutex.RLock()
	defer fake.isWellFormedMutex.RUnlock()
	return len(fake.isWellFormedArgsForCall)
}

func (fake *MSPManager) IsWellFormedArgsForCall(i int) *mspprotos.SerializedIdentity {
	fake.isWellFormedMutex.RLock()
	defer fake.isWellFormedMutex.RUnlock()
	return fake.isWellFormedArgsForCall[i].identity
}

func (fake *MSPManager) IsWellFormedReturns(result1 error) {
	fake.IsWellFormedStub = nil
	fake.isWellFormedReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSPManager) IsWellFormedReturnsOnCall(i int, result1 error) {
	fake.IsWellFormedStub = nil
	if fake.isWellFormedReturnsOnCall == nil {
		fake.isWellFormedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.isWellFormedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSPManager) Setup(msps []msp.MSP) error {
	var mspsCopy []msp.MSP
	if msps != nil {
		mspsCopy = make([]msp.MSP, len(msps))
		copy(mspsCopy, msps)
	}
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		msps []msp.MSP
	}{mspsCopy})
	fake.recordInvocation("Setup", []interface{}{mspsCopy})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(msps)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setupReturns.result1
}

func (fake *MSPManager) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *MSPManager) SetupArgsForCall(i int) []msp.MSP {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return fake.setupArgsForCall[i].msps
}

func (fake *MSPManager) SetupReturns(result1 error) {
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSPManager) SetupReturnsOnCall(i int, result1 error) {
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSPManager) GetMSPs() (map[string]msp.MSP, error) {
	fake.getMSPsMutex.Lock()
	ret, specificReturn := fake.getMSPsReturnsOnCall[len(fake.getMSPsArgsForCall)]
	fake.getMSPsArgsForCall = append(fake.getMSPsArgsForCall, struct{}{})
	fake.recordInvocation("GetMSPs", []interface{}{})
	fake.getMSPsMutex.Unlock()
	if fake.GetMSPsStub != nil {
		return fake.GetMSPsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getMSPsReturns.result1, fake.getMSPsReturns.result2
}

func (fake *MSPManager) GetMSPsCallCount() int {
	fake.getMSPsMutex.RLock()
	defer fake.getMSPsMutex.RUnlock()
	return len(fake.getMSPsArgsForCall)
}

func (fake *MSPManager) GetMSPsReturns(result1 map[string]msp.MSP, result2 error) {
	fake.GetMSPsStub = nil
	fake.getMSPsReturns = struct {
		result1 map[string]msp.MSP
		result2 error
	}{result1, result2}
}

func (fake *MSPManager) GetMSPsReturnsOnCall(i int, result1 map[string]msp.MSP, result2 error) {
	fake.GetMSPsStub = nil
	if fake.getMSPsReturnsOnCall == nil {
		fake.getMSPsReturnsOnCall = make(map[int]struct {
			result1 map[string]msp.MSP
			result2 error
		})
	}
	fake.getMSPsReturnsOnCall[i] = struct {
		result1 map[string]msp.MSP
		result2 error
	}{result1, result2}
}

func (fake *MSPManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	fake.isWellFormedMutex.RLock()
	defer fake.isWellFormedMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.getMSPsMutex.RLock()
	defer fake.getMSPsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MSPManager) recordInvocation(key string, args []interface{}) {
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

var _ msp.MSPManager = new(MSPManager)
