
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package metricsfakes

import (
	"sync"

	"github.com/hyperledger/fabric/common/metrics"
)

type Provider struct {
	NewCounterStub        func(metrics.CounterOpts) metrics.Counter
	newCounterMutex       sync.RWMutex
	newCounterArgsForCall []struct {
		arg1 metrics.CounterOpts
	}
	newCounterReturns struct {
		result1 metrics.Counter
	}
	newCounterReturnsOnCall map[int]struct {
		result1 metrics.Counter
	}
	NewGaugeStub        func(metrics.GaugeOpts) metrics.Gauge
	newGaugeMutex       sync.RWMutex
	newGaugeArgsForCall []struct {
		arg1 metrics.GaugeOpts
	}
	newGaugeReturns struct {
		result1 metrics.Gauge
	}
	newGaugeReturnsOnCall map[int]struct {
		result1 metrics.Gauge
	}
	NewHistogramStub        func(metrics.HistogramOpts) metrics.Histogram
	newHistogramMutex       sync.RWMutex
	newHistogramArgsForCall []struct {
		arg1 metrics.HistogramOpts
	}
	newHistogramReturns struct {
		result1 metrics.Histogram
	}
	newHistogramReturnsOnCall map[int]struct {
		result1 metrics.Histogram
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Provider) NewCounter(arg1 metrics.CounterOpts) metrics.Counter {
	fake.newCounterMutex.Lock()
	ret, specificReturn := fake.newCounterReturnsOnCall[len(fake.newCounterArgsForCall)]
	fake.newCounterArgsForCall = append(fake.newCounterArgsForCall, struct {
		arg1 metrics.CounterOpts
	}{arg1})
	fake.recordInvocation("NewCounter", []interface{}{arg1})
	fake.newCounterMutex.Unlock()
	if fake.NewCounterStub != nil {
		return fake.NewCounterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newCounterReturns.result1
}

func (fake *Provider) NewCounterCallCount() int {
	fake.newCounterMutex.RLock()
	defer fake.newCounterMutex.RUnlock()
	return len(fake.newCounterArgsForCall)
}

func (fake *Provider) NewCounterArgsForCall(i int) metrics.CounterOpts {
	fake.newCounterMutex.RLock()
	defer fake.newCounterMutex.RUnlock()
	return fake.newCounterArgsForCall[i].arg1
}

func (fake *Provider) NewCounterReturns(result1 metrics.Counter) {
	fake.NewCounterStub = nil
	fake.newCounterReturns = struct {
		result1 metrics.Counter
	}{result1}
}

func (fake *Provider) NewCounterReturnsOnCall(i int, result1 metrics.Counter) {
	fake.NewCounterStub = nil
	if fake.newCounterReturnsOnCall == nil {
		fake.newCounterReturnsOnCall = make(map[int]struct {
			result1 metrics.Counter
		})
	}
	fake.newCounterReturnsOnCall[i] = struct {
		result1 metrics.Counter
	}{result1}
}

func (fake *Provider) NewGauge(arg1 metrics.GaugeOpts) metrics.Gauge {
	fake.newGaugeMutex.Lock()
	ret, specificReturn := fake.newGaugeReturnsOnCall[len(fake.newGaugeArgsForCall)]
	fake.newGaugeArgsForCall = append(fake.newGaugeArgsForCall, struct {
		arg1 metrics.GaugeOpts
	}{arg1})
	fake.recordInvocation("NewGauge", []interface{}{arg1})
	fake.newGaugeMutex.Unlock()
	if fake.NewGaugeStub != nil {
		return fake.NewGaugeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newGaugeReturns.result1
}

func (fake *Provider) NewGaugeCallCount() int {
	fake.newGaugeMutex.RLock()
	defer fake.newGaugeMutex.RUnlock()
	return len(fake.newGaugeArgsForCall)
}

func (fake *Provider) NewGaugeArgsForCall(i int) metrics.GaugeOpts {
	fake.newGaugeMutex.RLock()
	defer fake.newGaugeMutex.RUnlock()
	return fake.newGaugeArgsForCall[i].arg1
}

func (fake *Provider) NewGaugeReturns(result1 metrics.Gauge) {
	fake.NewGaugeStub = nil
	fake.newGaugeReturns = struct {
		result1 metrics.Gauge
	}{result1}
}

func (fake *Provider) NewGaugeReturnsOnCall(i int, result1 metrics.Gauge) {
	fake.NewGaugeStub = nil
	if fake.newGaugeReturnsOnCall == nil {
		fake.newGaugeReturnsOnCall = make(map[int]struct {
			result1 metrics.Gauge
		})
	}
	fake.newGaugeReturnsOnCall[i] = struct {
		result1 metrics.Gauge
	}{result1}
}

func (fake *Provider) NewHistogram(arg1 metrics.HistogramOpts) metrics.Histogram {
	fake.newHistogramMutex.Lock()
	ret, specificReturn := fake.newHistogramReturnsOnCall[len(fake.newHistogramArgsForCall)]
	fake.newHistogramArgsForCall = append(fake.newHistogramArgsForCall, struct {
		arg1 metrics.HistogramOpts
	}{arg1})
	fake.recordInvocation("NewHistogram", []interface{}{arg1})
	fake.newHistogramMutex.Unlock()
	if fake.NewHistogramStub != nil {
		return fake.NewHistogramStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newHistogramReturns.result1
}

func (fake *Provider) NewHistogramCallCount() int {
	fake.newHistogramMutex.RLock()
	defer fake.newHistogramMutex.RUnlock()
	return len(fake.newHistogramArgsForCall)
}

func (fake *Provider) NewHistogramArgsForCall(i int) metrics.HistogramOpts {
	fake.newHistogramMutex.RLock()
	defer fake.newHistogramMutex.RUnlock()
	return fake.newHistogramArgsForCall[i].arg1
}

func (fake *Provider) NewHistogramReturns(result1 metrics.Histogram) {
	fake.NewHistogramStub = nil
	fake.newHistogramReturns = struct {
		result1 metrics.Histogram
	}{result1}
}

func (fake *Provider) NewHistogramReturnsOnCall(i int, result1 metrics.Histogram) {
	fake.NewHistogramStub = nil
	if fake.newHistogramReturnsOnCall == nil {
		fake.newHistogramReturnsOnCall = make(map[int]struct {
			result1 metrics.Histogram
		})
	}
	fake.newHistogramReturnsOnCall[i] = struct {
		result1 metrics.Histogram
	}{result1}
}

func (fake *Provider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newCounterMutex.RLock()
	defer fake.newCounterMutex.RUnlock()
	fake.newGaugeMutex.RLock()
	defer fake.newGaugeMutex.RUnlock()
	fake.newHistogramMutex.RLock()
	defer fake.newHistogramMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Provider) recordInvocation(key string, args []interface{}) {
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

var _ metrics.Provider = new(Provider)
