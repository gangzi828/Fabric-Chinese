
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
	sync "sync"

	privdata "github.com/hyperledger/fabric/core/common/privdata"
	ledger "github.com/hyperledger/fabric/core/ledger"
	common "github.com/hyperledger/fabric/protos/common"
	peer "github.com/hyperledger/fabric/protos/peer"
)

type CollectionStore struct {
	AccessFilterStub        func(string, *common.CollectionPolicyConfig) (privdata.Filter, error)
	accessFilterMutex       sync.RWMutex
	accessFilterArgsForCall []struct {
		arg1 string
		arg2 *common.CollectionPolicyConfig
	}
	accessFilterReturns struct {
		result1 privdata.Filter
		result2 error
	}
	accessFilterReturnsOnCall map[int]struct {
		result1 privdata.Filter
		result2 error
	}
	HasReadAccessStub        func(common.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) (bool, error)
	hasReadAccessMutex       sync.RWMutex
	hasReadAccessArgsForCall []struct {
		arg1 common.CollectionCriteria
		arg2 *peer.SignedProposal
		arg3 ledger.QueryExecutor
	}
	hasReadAccessReturns struct {
		result1 bool
		result2 error
	}
	hasReadAccessReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	RetrieveCollectionStub        func(common.CollectionCriteria) (privdata.Collection, error)
	retrieveCollectionMutex       sync.RWMutex
	retrieveCollectionArgsForCall []struct {
		arg1 common.CollectionCriteria
	}
	retrieveCollectionReturns struct {
		result1 privdata.Collection
		result2 error
	}
	retrieveCollectionReturnsOnCall map[int]struct {
		result1 privdata.Collection
		result2 error
	}
	RetrieveCollectionAccessPolicyStub        func(common.CollectionCriteria) (privdata.CollectionAccessPolicy, error)
	retrieveCollectionAccessPolicyMutex       sync.RWMutex
	retrieveCollectionAccessPolicyArgsForCall []struct {
		arg1 common.CollectionCriteria
	}
	retrieveCollectionAccessPolicyReturns struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}
	retrieveCollectionAccessPolicyReturnsOnCall map[int]struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}
	RetrieveCollectionConfigPackageStub        func(common.CollectionCriteria) (*common.CollectionConfigPackage, error)
	retrieveCollectionConfigPackageMutex       sync.RWMutex
	retrieveCollectionConfigPackageArgsForCall []struct {
		arg1 common.CollectionCriteria
	}
	retrieveCollectionConfigPackageReturns struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}
	retrieveCollectionConfigPackageReturnsOnCall map[int]struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}
	RetrieveCollectionPersistenceConfigsStub        func(common.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error)
	retrieveCollectionPersistenceConfigsMutex       sync.RWMutex
	retrieveCollectionPersistenceConfigsArgsForCall []struct {
		arg1 common.CollectionCriteria
	}
	retrieveCollectionPersistenceConfigsReturns struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}
	retrieveCollectionPersistenceConfigsReturnsOnCall map[int]struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CollectionStore) AccessFilter(arg1 string, arg2 *common.CollectionPolicyConfig) (privdata.Filter, error) {
	fake.accessFilterMutex.Lock()
	ret, specificReturn := fake.accessFilterReturnsOnCall[len(fake.accessFilterArgsForCall)]
	fake.accessFilterArgsForCall = append(fake.accessFilterArgsForCall, struct {
		arg1 string
		arg2 *common.CollectionPolicyConfig
	}{arg1, arg2})
	fake.recordInvocation("AccessFilter", []interface{}{arg1, arg2})
	fake.accessFilterMutex.Unlock()
	if fake.AccessFilterStub != nil {
		return fake.AccessFilterStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.accessFilterReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) AccessFilterCallCount() int {
	fake.accessFilterMutex.RLock()
	defer fake.accessFilterMutex.RUnlock()
	return len(fake.accessFilterArgsForCall)
}

func (fake *CollectionStore) AccessFilterCalls(stub func(string, *common.CollectionPolicyConfig) (privdata.Filter, error)) {
	fake.accessFilterMutex.Lock()
	defer fake.accessFilterMutex.Unlock()
	fake.AccessFilterStub = stub
}

func (fake *CollectionStore) AccessFilterArgsForCall(i int) (string, *common.CollectionPolicyConfig) {
	fake.accessFilterMutex.RLock()
	defer fake.accessFilterMutex.RUnlock()
	argsForCall := fake.accessFilterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CollectionStore) AccessFilterReturns(result1 privdata.Filter, result2 error) {
	fake.accessFilterMutex.Lock()
	defer fake.accessFilterMutex.Unlock()
	fake.AccessFilterStub = nil
	fake.accessFilterReturns = struct {
		result1 privdata.Filter
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) AccessFilterReturnsOnCall(i int, result1 privdata.Filter, result2 error) {
	fake.accessFilterMutex.Lock()
	defer fake.accessFilterMutex.Unlock()
	fake.AccessFilterStub = nil
	if fake.accessFilterReturnsOnCall == nil {
		fake.accessFilterReturnsOnCall = make(map[int]struct {
			result1 privdata.Filter
			result2 error
		})
	}
	fake.accessFilterReturnsOnCall[i] = struct {
		result1 privdata.Filter
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) HasReadAccess(arg1 common.CollectionCriteria, arg2 *peer.SignedProposal, arg3 ledger.QueryExecutor) (bool, error) {
	fake.hasReadAccessMutex.Lock()
	ret, specificReturn := fake.hasReadAccessReturnsOnCall[len(fake.hasReadAccessArgsForCall)]
	fake.hasReadAccessArgsForCall = append(fake.hasReadAccessArgsForCall, struct {
		arg1 common.CollectionCriteria
		arg2 *peer.SignedProposal
		arg3 ledger.QueryExecutor
	}{arg1, arg2, arg3})
	fake.recordInvocation("HasReadAccess", []interface{}{arg1, arg2, arg3})
	fake.hasReadAccessMutex.Unlock()
	if fake.HasReadAccessStub != nil {
		return fake.HasReadAccessStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.hasReadAccessReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) HasReadAccessCallCount() int {
	fake.hasReadAccessMutex.RLock()
	defer fake.hasReadAccessMutex.RUnlock()
	return len(fake.hasReadAccessArgsForCall)
}

func (fake *CollectionStore) HasReadAccessCalls(stub func(common.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) (bool, error)) {
	fake.hasReadAccessMutex.Lock()
	defer fake.hasReadAccessMutex.Unlock()
	fake.HasReadAccessStub = stub
}

func (fake *CollectionStore) HasReadAccessArgsForCall(i int) (common.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) {
	fake.hasReadAccessMutex.RLock()
	defer fake.hasReadAccessMutex.RUnlock()
	argsForCall := fake.hasReadAccessArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CollectionStore) HasReadAccessReturns(result1 bool, result2 error) {
	fake.hasReadAccessMutex.Lock()
	defer fake.hasReadAccessMutex.Unlock()
	fake.HasReadAccessStub = nil
	fake.hasReadAccessReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) HasReadAccessReturnsOnCall(i int, result1 bool, result2 error) {
	fake.hasReadAccessMutex.Lock()
	defer fake.hasReadAccessMutex.Unlock()
	fake.HasReadAccessStub = nil
	if fake.hasReadAccessReturnsOnCall == nil {
		fake.hasReadAccessReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasReadAccessReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollection(arg1 common.CollectionCriteria) (privdata.Collection, error) {
	fake.retrieveCollectionMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionReturnsOnCall[len(fake.retrieveCollectionArgsForCall)]
	fake.retrieveCollectionArgsForCall = append(fake.retrieveCollectionArgsForCall, struct {
		arg1 common.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollection", []interface{}{arg1})
	fake.retrieveCollectionMutex.Unlock()
	if fake.RetrieveCollectionStub != nil {
		return fake.RetrieveCollectionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveCollectionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionCallCount() int {
	fake.retrieveCollectionMutex.RLock()
	defer fake.retrieveCollectionMutex.RUnlock()
	return len(fake.retrieveCollectionArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionCalls(stub func(common.CollectionCriteria) (privdata.Collection, error)) {
	fake.retrieveCollectionMutex.Lock()
	defer fake.retrieveCollectionMutex.Unlock()
	fake.RetrieveCollectionStub = stub
}

func (fake *CollectionStore) RetrieveCollectionArgsForCall(i int) common.CollectionCriteria {
	fake.retrieveCollectionMutex.RLock()
	defer fake.retrieveCollectionMutex.RUnlock()
	argsForCall := fake.retrieveCollectionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CollectionStore) RetrieveCollectionReturns(result1 privdata.Collection, result2 error) {
	fake.retrieveCollectionMutex.Lock()
	defer fake.retrieveCollectionMutex.Unlock()
	fake.RetrieveCollectionStub = nil
	fake.retrieveCollectionReturns = struct {
		result1 privdata.Collection
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionReturnsOnCall(i int, result1 privdata.Collection, result2 error) {
	fake.retrieveCollectionMutex.Lock()
	defer fake.retrieveCollectionMutex.Unlock()
	fake.RetrieveCollectionStub = nil
	if fake.retrieveCollectionReturnsOnCall == nil {
		fake.retrieveCollectionReturnsOnCall = make(map[int]struct {
			result1 privdata.Collection
			result2 error
		})
	}
	fake.retrieveCollectionReturnsOnCall[i] = struct {
		result1 privdata.Collection
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicy(arg1 common.CollectionCriteria) (privdata.CollectionAccessPolicy, error) {
	fake.retrieveCollectionAccessPolicyMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionAccessPolicyReturnsOnCall[len(fake.retrieveCollectionAccessPolicyArgsForCall)]
	fake.retrieveCollectionAccessPolicyArgsForCall = append(fake.retrieveCollectionAccessPolicyArgsForCall, struct {
		arg1 common.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionAccessPolicy", []interface{}{arg1})
	fake.retrieveCollectionAccessPolicyMutex.Unlock()
	if fake.RetrieveCollectionAccessPolicyStub != nil {
		return fake.RetrieveCollectionAccessPolicyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveCollectionAccessPolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyCallCount() int {
	fake.retrieveCollectionAccessPolicyMutex.RLock()
	defer fake.retrieveCollectionAccessPolicyMutex.RUnlock()
	return len(fake.retrieveCollectionAccessPolicyArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyCalls(stub func(common.CollectionCriteria) (privdata.CollectionAccessPolicy, error)) {
	fake.retrieveCollectionAccessPolicyMutex.Lock()
	defer fake.retrieveCollectionAccessPolicyMutex.Unlock()
	fake.RetrieveCollectionAccessPolicyStub = stub
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyArgsForCall(i int) common.CollectionCriteria {
	fake.retrieveCollectionAccessPolicyMutex.RLock()
	defer fake.retrieveCollectionAccessPolicyMutex.RUnlock()
	argsForCall := fake.retrieveCollectionAccessPolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyReturns(result1 privdata.CollectionAccessPolicy, result2 error) {
	fake.retrieveCollectionAccessPolicyMutex.Lock()
	defer fake.retrieveCollectionAccessPolicyMutex.Unlock()
	fake.RetrieveCollectionAccessPolicyStub = nil
	fake.retrieveCollectionAccessPolicyReturns = struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyReturnsOnCall(i int, result1 privdata.CollectionAccessPolicy, result2 error) {
	fake.retrieveCollectionAccessPolicyMutex.Lock()
	defer fake.retrieveCollectionAccessPolicyMutex.Unlock()
	fake.RetrieveCollectionAccessPolicyStub = nil
	if fake.retrieveCollectionAccessPolicyReturnsOnCall == nil {
		fake.retrieveCollectionAccessPolicyReturnsOnCall = make(map[int]struct {
			result1 privdata.CollectionAccessPolicy
			result2 error
		})
	}
	fake.retrieveCollectionAccessPolicyReturnsOnCall[i] = struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionConfigPackage(arg1 common.CollectionCriteria) (*common.CollectionConfigPackage, error) {
	fake.retrieveCollectionConfigPackageMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionConfigPackageReturnsOnCall[len(fake.retrieveCollectionConfigPackageArgsForCall)]
	fake.retrieveCollectionConfigPackageArgsForCall = append(fake.retrieveCollectionConfigPackageArgsForCall, struct {
		arg1 common.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionConfigPackage", []interface{}{arg1})
	fake.retrieveCollectionConfigPackageMutex.Unlock()
	if fake.RetrieveCollectionConfigPackageStub != nil {
		return fake.RetrieveCollectionConfigPackageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveCollectionConfigPackageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageCallCount() int {
	fake.retrieveCollectionConfigPackageMutex.RLock()
	defer fake.retrieveCollectionConfigPackageMutex.RUnlock()
	return len(fake.retrieveCollectionConfigPackageArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageCalls(stub func(common.CollectionCriteria) (*common.CollectionConfigPackage, error)) {
	fake.retrieveCollectionConfigPackageMutex.Lock()
	defer fake.retrieveCollectionConfigPackageMutex.Unlock()
	fake.RetrieveCollectionConfigPackageStub = stub
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageArgsForCall(i int) common.CollectionCriteria {
	fake.retrieveCollectionConfigPackageMutex.RLock()
	defer fake.retrieveCollectionConfigPackageMutex.RUnlock()
	argsForCall := fake.retrieveCollectionConfigPackageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageReturns(result1 *common.CollectionConfigPackage, result2 error) {
	fake.retrieveCollectionConfigPackageMutex.Lock()
	defer fake.retrieveCollectionConfigPackageMutex.Unlock()
	fake.RetrieveCollectionConfigPackageStub = nil
	fake.retrieveCollectionConfigPackageReturns = struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageReturnsOnCall(i int, result1 *common.CollectionConfigPackage, result2 error) {
	fake.retrieveCollectionConfigPackageMutex.Lock()
	defer fake.retrieveCollectionConfigPackageMutex.Unlock()
	fake.RetrieveCollectionConfigPackageStub = nil
	if fake.retrieveCollectionConfigPackageReturnsOnCall == nil {
		fake.retrieveCollectionConfigPackageReturnsOnCall = make(map[int]struct {
			result1 *common.CollectionConfigPackage
			result2 error
		})
	}
	fake.retrieveCollectionConfigPackageReturnsOnCall[i] = struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigs(arg1 common.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error) {
	fake.retrieveCollectionPersistenceConfigsMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionPersistenceConfigsReturnsOnCall[len(fake.retrieveCollectionPersistenceConfigsArgsForCall)]
	fake.retrieveCollectionPersistenceConfigsArgsForCall = append(fake.retrieveCollectionPersistenceConfigsArgsForCall, struct {
		arg1 common.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionPersistenceConfigs", []interface{}{arg1})
	fake.retrieveCollectionPersistenceConfigsMutex.Unlock()
	if fake.RetrieveCollectionPersistenceConfigsStub != nil {
		return fake.RetrieveCollectionPersistenceConfigsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveCollectionPersistenceConfigsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsCallCount() int {
	fake.retrieveCollectionPersistenceConfigsMutex.RLock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.RUnlock()
	return len(fake.retrieveCollectionPersistenceConfigsArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsCalls(stub func(common.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error)) {
	fake.retrieveCollectionPersistenceConfigsMutex.Lock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.Unlock()
	fake.RetrieveCollectionPersistenceConfigsStub = stub
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsArgsForCall(i int) common.CollectionCriteria {
	fake.retrieveCollectionPersistenceConfigsMutex.RLock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.RUnlock()
	argsForCall := fake.retrieveCollectionPersistenceConfigsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsReturns(result1 privdata.CollectionPersistenceConfigs, result2 error) {
	fake.retrieveCollectionPersistenceConfigsMutex.Lock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.Unlock()
	fake.RetrieveCollectionPersistenceConfigsStub = nil
	fake.retrieveCollectionPersistenceConfigsReturns = struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsReturnsOnCall(i int, result1 privdata.CollectionPersistenceConfigs, result2 error) {
	fake.retrieveCollectionPersistenceConfigsMutex.Lock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.Unlock()
	fake.RetrieveCollectionPersistenceConfigsStub = nil
	if fake.retrieveCollectionPersistenceConfigsReturnsOnCall == nil {
		fake.retrieveCollectionPersistenceConfigsReturnsOnCall = make(map[int]struct {
			result1 privdata.CollectionPersistenceConfigs
			result2 error
		})
	}
	fake.retrieveCollectionPersistenceConfigsReturnsOnCall[i] = struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.accessFilterMutex.RLock()
	defer fake.accessFilterMutex.RUnlock()
	fake.hasReadAccessMutex.RLock()
	defer fake.hasReadAccessMutex.RUnlock()
	fake.retrieveCollectionMutex.RLock()
	defer fake.retrieveCollectionMutex.RUnlock()
	fake.retrieveCollectionAccessPolicyMutex.RLock()
	defer fake.retrieveCollectionAccessPolicyMutex.RUnlock()
	fake.retrieveCollectionConfigPackageMutex.RLock()
	defer fake.retrieveCollectionConfigPackageMutex.RUnlock()
	fake.retrieveCollectionPersistenceConfigsMutex.RLock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CollectionStore) recordInvocation(key string, args []interface{}) {
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
