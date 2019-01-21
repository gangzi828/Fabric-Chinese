
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

	"github.com/hyperledger/fabric/bccsp/idemix/handlers"
)

type NymSignatureScheme struct {
	SignStub        func(sk handlers.Big, Nym handlers.Ecp, RNym handlers.Big, ipk handlers.IssuerPublicKey, digest []byte) ([]byte, error)
	signMutex       sync.RWMutex
	signArgsForCall []struct {
		sk     handlers.Big
		Nym    handlers.Ecp
		RNym   handlers.Big
		ipk    handlers.IssuerPublicKey
		digest []byte
	}
	signReturns struct {
		result1 []byte
		result2 error
	}
	signReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	VerifyStub        func(pk handlers.IssuerPublicKey, Nym handlers.Ecp, signature, digest []byte) error
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		pk        handlers.IssuerPublicKey
		Nym       handlers.Ecp
		signature []byte
		digest    []byte
	}
	verifyReturns struct {
		result1 error
	}
	verifyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NymSignatureScheme) Sign(sk handlers.Big, Nym handlers.Ecp, RNym handlers.Big, ipk handlers.IssuerPublicKey, digest []byte) ([]byte, error) {
	var digestCopy []byte
	if digest != nil {
		digestCopy = make([]byte, len(digest))
		copy(digestCopy, digest)
	}
	fake.signMutex.Lock()
	ret, specificReturn := fake.signReturnsOnCall[len(fake.signArgsForCall)]
	fake.signArgsForCall = append(fake.signArgsForCall, struct {
		sk     handlers.Big
		Nym    handlers.Ecp
		RNym   handlers.Big
		ipk    handlers.IssuerPublicKey
		digest []byte
	}{sk, Nym, RNym, ipk, digestCopy})
	fake.recordInvocation("Sign", []interface{}{sk, Nym, RNym, ipk, digestCopy})
	fake.signMutex.Unlock()
	if fake.SignStub != nil {
		return fake.SignStub(sk, Nym, RNym, ipk, digest)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.signReturns.result1, fake.signReturns.result2
}

func (fake *NymSignatureScheme) SignCallCount() int {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return len(fake.signArgsForCall)
}

func (fake *NymSignatureScheme) SignArgsForCall(i int) (handlers.Big, handlers.Ecp, handlers.Big, handlers.IssuerPublicKey, []byte) {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return fake.signArgsForCall[i].sk, fake.signArgsForCall[i].Nym, fake.signArgsForCall[i].RNym, fake.signArgsForCall[i].ipk, fake.signArgsForCall[i].digest
}

func (fake *NymSignatureScheme) SignReturns(result1 []byte, result2 error) {
	fake.SignStub = nil
	fake.signReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *NymSignatureScheme) SignReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.SignStub = nil
	if fake.signReturnsOnCall == nil {
		fake.signReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.signReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *NymSignatureScheme) Verify(pk handlers.IssuerPublicKey, Nym handlers.Ecp, signature []byte, digest []byte) error {
	var signatureCopy []byte
	if signature != nil {
		signatureCopy = make([]byte, len(signature))
		copy(signatureCopy, signature)
	}
	var digestCopy []byte
	if digest != nil {
		digestCopy = make([]byte, len(digest))
		copy(digestCopy, digest)
	}
	fake.verifyMutex.Lock()
	ret, specificReturn := fake.verifyReturnsOnCall[len(fake.verifyArgsForCall)]
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		pk        handlers.IssuerPublicKey
		Nym       handlers.Ecp
		signature []byte
		digest    []byte
	}{pk, Nym, signatureCopy, digestCopy})
	fake.recordInvocation("Verify", []interface{}{pk, Nym, signatureCopy, digestCopy})
	fake.verifyMutex.Unlock()
	if fake.VerifyStub != nil {
		return fake.VerifyStub(pk, Nym, signature, digest)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.verifyReturns.result1
}

func (fake *NymSignatureScheme) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *NymSignatureScheme) VerifyArgsForCall(i int) (handlers.IssuerPublicKey, handlers.Ecp, []byte, []byte) {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return fake.verifyArgsForCall[i].pk, fake.verifyArgsForCall[i].Nym, fake.verifyArgsForCall[i].signature, fake.verifyArgsForCall[i].digest
}

func (fake *NymSignatureScheme) VerifyReturns(result1 error) {
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 error
	}{result1}
}

func (fake *NymSignatureScheme) VerifyReturnsOnCall(i int, result1 error) {
	fake.VerifyStub = nil
	if fake.verifyReturnsOnCall == nil {
		fake.verifyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NymSignatureScheme) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NymSignatureScheme) recordInvocation(key string, args []interface{}) {
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

var _ handlers.NymSignatureScheme = new(NymSignatureScheme)
