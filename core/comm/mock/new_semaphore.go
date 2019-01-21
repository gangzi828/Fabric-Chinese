
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

	comm "github.com/hyperledger/fabric/core/comm"
)

type NewSemaphore struct {
	Stub        func(int) comm.Semaphore
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 int
	}
	returns struct {
		result1 comm.Semaphore
	}
	returnsOnCall map[int]struct {
		result1 comm.Semaphore
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NewSemaphore) Spy(arg1 int) comm.Semaphore {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("NewSemaphoreFunc", []interface{}{arg1})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.returns.result1
}

func (fake *NewSemaphore) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *NewSemaphore) Calls(stub func(int) comm.Semaphore) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *NewSemaphore) ArgsForCall(i int) int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1
}

func (fake *NewSemaphore) Returns(result1 comm.Semaphore) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 comm.Semaphore
	}{result1}
}

func (fake *NewSemaphore) ReturnsOnCall(i int, result1 comm.Semaphore) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 comm.Semaphore
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 comm.Semaphore
	}{result1}
}

func (fake *NewSemaphore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NewSemaphore) recordInvocation(key string, args []interface{}) {
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

var _ comm.NewSemaphoreFunc = new(NewSemaphore).Spy
