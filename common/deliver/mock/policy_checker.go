
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

	deliver "github.com/hyperledger/fabric/common/deliver"
	common "github.com/hyperledger/fabric/protos/common"
)

type PolicyChecker struct {
	CheckPolicyStub        func(*common.Envelope, string) error
	checkPolicyMutex       sync.RWMutex
	checkPolicyArgsForCall []struct {
		arg1 *common.Envelope
		arg2 string
	}
	checkPolicyReturns struct {
		result1 error
	}
	checkPolicyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyChecker) CheckPolicy(arg1 *common.Envelope, arg2 string) error {
	fake.checkPolicyMutex.Lock()
	ret, specificReturn := fake.checkPolicyReturnsOnCall[len(fake.checkPolicyArgsForCall)]
	fake.checkPolicyArgsForCall = append(fake.checkPolicyArgsForCall, struct {
		arg1 *common.Envelope
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CheckPolicy", []interface{}{arg1, arg2})
	fake.checkPolicyMutex.Unlock()
	if fake.CheckPolicyStub != nil {
		return fake.CheckPolicyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkPolicyReturns
	return fakeReturns.result1
}

func (fake *PolicyChecker) CheckPolicyCallCount() int {
	fake.checkPolicyMutex.RLock()
	defer fake.checkPolicyMutex.RUnlock()
	return len(fake.checkPolicyArgsForCall)
}

func (fake *PolicyChecker) CheckPolicyCalls(stub func(*common.Envelope, string) error) {
	fake.checkPolicyMutex.Lock()
	defer fake.checkPolicyMutex.Unlock()
	fake.CheckPolicyStub = stub
}

func (fake *PolicyChecker) CheckPolicyArgsForCall(i int) (*common.Envelope, string) {
	fake.checkPolicyMutex.RLock()
	defer fake.checkPolicyMutex.RUnlock()
	argsForCall := fake.checkPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *PolicyChecker) CheckPolicyReturns(result1 error) {
	fake.checkPolicyMutex.Lock()
	defer fake.checkPolicyMutex.Unlock()
	fake.CheckPolicyStub = nil
	fake.checkPolicyReturns = struct {
		result1 error
	}{result1}
}

func (fake *PolicyChecker) CheckPolicyReturnsOnCall(i int, result1 error) {
	fake.checkPolicyMutex.Lock()
	defer fake.checkPolicyMutex.Unlock()
	fake.CheckPolicyStub = nil
	if fake.checkPolicyReturnsOnCall == nil {
		fake.checkPolicyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkPolicyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PolicyChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkPolicyMutex.RLock()
	defer fake.checkPolicyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PolicyChecker) recordInvocation(key string, args []interface{}) {
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

var _ deliver.PolicyChecker = new(PolicyChecker)
