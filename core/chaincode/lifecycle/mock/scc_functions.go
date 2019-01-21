
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
)

type SCCFunctions struct {
	InstallChaincodeStub        func(string, string, []byte) ([]byte, error)
	installChaincodeMutex       sync.RWMutex
	installChaincodeArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []byte
	}
	installChaincodeReturns struct {
		result1 []byte
		result2 error
	}
	installChaincodeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	QueryInstalledChaincodeStub        func(string, string) ([]byte, error)
	queryInstalledChaincodeMutex       sync.RWMutex
	queryInstalledChaincodeArgsForCall []struct {
		arg1 string
		arg2 string
	}
	queryInstalledChaincodeReturns struct {
		result1 []byte
		result2 error
	}
	queryInstalledChaincodeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SCCFunctions) InstallChaincode(arg1 string, arg2 string, arg3 []byte) ([]byte, error) {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.installChaincodeMutex.Lock()
	ret, specificReturn := fake.installChaincodeReturnsOnCall[len(fake.installChaincodeArgsForCall)]
	fake.installChaincodeArgsForCall = append(fake.installChaincodeArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("InstallChaincode", []interface{}{arg1, arg2, arg3Copy})
	fake.installChaincodeMutex.Unlock()
	if fake.InstallChaincodeStub != nil {
		return fake.InstallChaincodeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.installChaincodeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) InstallChaincodeCallCount() int {
	fake.installChaincodeMutex.RLock()
	defer fake.installChaincodeMutex.RUnlock()
	return len(fake.installChaincodeArgsForCall)
}

func (fake *SCCFunctions) InstallChaincodeCalls(stub func(string, string, []byte) ([]byte, error)) {
	fake.installChaincodeMutex.Lock()
	defer fake.installChaincodeMutex.Unlock()
	fake.InstallChaincodeStub = stub
}

func (fake *SCCFunctions) InstallChaincodeArgsForCall(i int) (string, string, []byte) {
	fake.installChaincodeMutex.RLock()
	defer fake.installChaincodeMutex.RUnlock()
	argsForCall := fake.installChaincodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SCCFunctions) InstallChaincodeReturns(result1 []byte, result2 error) {
	fake.installChaincodeMutex.Lock()
	defer fake.installChaincodeMutex.Unlock()
	fake.InstallChaincodeStub = nil
	fake.installChaincodeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) InstallChaincodeReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.installChaincodeMutex.Lock()
	defer fake.installChaincodeMutex.Unlock()
	fake.InstallChaincodeStub = nil
	if fake.installChaincodeReturnsOnCall == nil {
		fake.installChaincodeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.installChaincodeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryInstalledChaincode(arg1 string, arg2 string) ([]byte, error) {
	fake.queryInstalledChaincodeMutex.Lock()
	ret, specificReturn := fake.queryInstalledChaincodeReturnsOnCall[len(fake.queryInstalledChaincodeArgsForCall)]
	fake.queryInstalledChaincodeArgsForCall = append(fake.queryInstalledChaincodeArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("QueryInstalledChaincode", []interface{}{arg1, arg2})
	fake.queryInstalledChaincodeMutex.Unlock()
	if fake.QueryInstalledChaincodeStub != nil {
		return fake.QueryInstalledChaincodeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryInstalledChaincodeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryInstalledChaincodeCallCount() int {
	fake.queryInstalledChaincodeMutex.RLock()
	defer fake.queryInstalledChaincodeMutex.RUnlock()
	return len(fake.queryInstalledChaincodeArgsForCall)
}

func (fake *SCCFunctions) QueryInstalledChaincodeCalls(stub func(string, string) ([]byte, error)) {
	fake.queryInstalledChaincodeMutex.Lock()
	defer fake.queryInstalledChaincodeMutex.Unlock()
	fake.QueryInstalledChaincodeStub = stub
}

func (fake *SCCFunctions) QueryInstalledChaincodeArgsForCall(i int) (string, string) {
	fake.queryInstalledChaincodeMutex.RLock()
	defer fake.queryInstalledChaincodeMutex.RUnlock()
	argsForCall := fake.queryInstalledChaincodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SCCFunctions) QueryInstalledChaincodeReturns(result1 []byte, result2 error) {
	fake.queryInstalledChaincodeMutex.Lock()
	defer fake.queryInstalledChaincodeMutex.Unlock()
	fake.QueryInstalledChaincodeStub = nil
	fake.queryInstalledChaincodeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryInstalledChaincodeReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.queryInstalledChaincodeMutex.Lock()
	defer fake.queryInstalledChaincodeMutex.Unlock()
	fake.QueryInstalledChaincodeStub = nil
	if fake.queryInstalledChaincodeReturnsOnCall == nil {
		fake.queryInstalledChaincodeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.queryInstalledChaincodeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.installChaincodeMutex.RLock()
	defer fake.installChaincodeMutex.RUnlock()
	fake.queryInstalledChaincodeMutex.RLock()
	defer fake.queryInstalledChaincodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SCCFunctions) recordInvocation(key string, args []interface{}) {
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
