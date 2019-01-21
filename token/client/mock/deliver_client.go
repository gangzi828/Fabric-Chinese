
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
	"context"
	"crypto/tls"
	"sync"

	"github.com/hyperledger/fabric/token/client"
	"google.golang.org/grpc"
)

type DeliverClient struct {
	NewDeliverFilteredStub        func(ctx context.Context, opts ...grpc.CallOption) (client.DeliverFiltered, error)
	newDeliverFilteredMutex       sync.RWMutex
	newDeliverFilteredArgsForCall []struct {
		ctx  context.Context
		opts []grpc.CallOption
	}
	newDeliverFilteredReturns struct {
		result1 client.DeliverFiltered
		result2 error
	}
	newDeliverFilteredReturnsOnCall map[int]struct {
		result1 client.DeliverFiltered
		result2 error
	}
	CertificateStub        func() *tls.Certificate
	certificateMutex       sync.RWMutex
	certificateArgsForCall []struct{}
	certificateReturns     struct {
		result1 *tls.Certificate
	}
	certificateReturnsOnCall map[int]struct {
		result1 *tls.Certificate
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeliverClient) NewDeliverFiltered(ctx context.Context, opts ...grpc.CallOption) (client.DeliverFiltered, error) {
	fake.newDeliverFilteredMutex.Lock()
	ret, specificReturn := fake.newDeliverFilteredReturnsOnCall[len(fake.newDeliverFilteredArgsForCall)]
	fake.newDeliverFilteredArgsForCall = append(fake.newDeliverFilteredArgsForCall, struct {
		ctx  context.Context
		opts []grpc.CallOption
	}{ctx, opts})
	fake.recordInvocation("NewDeliverFiltered", []interface{}{ctx, opts})
	fake.newDeliverFilteredMutex.Unlock()
	if fake.NewDeliverFilteredStub != nil {
		return fake.NewDeliverFilteredStub(ctx, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newDeliverFilteredReturns.result1, fake.newDeliverFilteredReturns.result2
}

func (fake *DeliverClient) NewDeliverFilteredCallCount() int {
	fake.newDeliverFilteredMutex.RLock()
	defer fake.newDeliverFilteredMutex.RUnlock()
	return len(fake.newDeliverFilteredArgsForCall)
}

func (fake *DeliverClient) NewDeliverFilteredArgsForCall(i int) (context.Context, []grpc.CallOption) {
	fake.newDeliverFilteredMutex.RLock()
	defer fake.newDeliverFilteredMutex.RUnlock()
	return fake.newDeliverFilteredArgsForCall[i].ctx, fake.newDeliverFilteredArgsForCall[i].opts
}

func (fake *DeliverClient) NewDeliverFilteredReturns(result1 client.DeliverFiltered, result2 error) {
	fake.NewDeliverFilteredStub = nil
	fake.newDeliverFilteredReturns = struct {
		result1 client.DeliverFiltered
		result2 error
	}{result1, result2}
}

func (fake *DeliverClient) NewDeliverFilteredReturnsOnCall(i int, result1 client.DeliverFiltered, result2 error) {
	fake.NewDeliverFilteredStub = nil
	if fake.newDeliverFilteredReturnsOnCall == nil {
		fake.newDeliverFilteredReturnsOnCall = make(map[int]struct {
			result1 client.DeliverFiltered
			result2 error
		})
	}
	fake.newDeliverFilteredReturnsOnCall[i] = struct {
		result1 client.DeliverFiltered
		result2 error
	}{result1, result2}
}

func (fake *DeliverClient) Certificate() *tls.Certificate {
	fake.certificateMutex.Lock()
	ret, specificReturn := fake.certificateReturnsOnCall[len(fake.certificateArgsForCall)]
	fake.certificateArgsForCall = append(fake.certificateArgsForCall, struct{}{})
	fake.recordInvocation("Certificate", []interface{}{})
	fake.certificateMutex.Unlock()
	if fake.CertificateStub != nil {
		return fake.CertificateStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.certificateReturns.result1
}

func (fake *DeliverClient) CertificateCallCount() int {
	fake.certificateMutex.RLock()
	defer fake.certificateMutex.RUnlock()
	return len(fake.certificateArgsForCall)
}

func (fake *DeliverClient) CertificateReturns(result1 *tls.Certificate) {
	fake.CertificateStub = nil
	fake.certificateReturns = struct {
		result1 *tls.Certificate
	}{result1}
}

func (fake *DeliverClient) CertificateReturnsOnCall(i int, result1 *tls.Certificate) {
	fake.CertificateStub = nil
	if fake.certificateReturnsOnCall == nil {
		fake.certificateReturnsOnCall = make(map[int]struct {
			result1 *tls.Certificate
		})
	}
	fake.certificateReturnsOnCall[i] = struct {
		result1 *tls.Certificate
	}{result1}
}

func (fake *DeliverClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newDeliverFilteredMutex.RLock()
	defer fake.newDeliverFilteredMutex.RUnlock()
	fake.certificateMutex.RLock()
	defer fake.certificateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeliverClient) recordInvocation(key string, args []interface{}) {
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

var _ client.DeliverClient = new(DeliverClient)
