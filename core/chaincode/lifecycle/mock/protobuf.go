
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

	"github.com/golang/protobuf/proto"
)

type Protobuf struct {
	MarshalStub        func(msg proto.Message) (marshaled []byte, err error)
	marshalMutex       sync.RWMutex
	marshalArgsForCall []struct {
		msg proto.Message
	}
	marshalReturns struct {
		result1 []byte
		result2 error
	}
	marshalReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	UnmarshalStub        func(marshaled []byte, msg proto.Message) error
	unmarshalMutex       sync.RWMutex
	unmarshalArgsForCall []struct {
		marshaled []byte
		msg       proto.Message
	}
	unmarshalReturns struct {
		result1 error
	}
	unmarshalReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Protobuf) Marshal(msg proto.Message) (marshaled []byte, err error) {
	fake.marshalMutex.Lock()
	ret, specificReturn := fake.marshalReturnsOnCall[len(fake.marshalArgsForCall)]
	fake.marshalArgsForCall = append(fake.marshalArgsForCall, struct {
		msg proto.Message
	}{msg})
	fake.recordInvocation("Marshal", []interface{}{msg})
	fake.marshalMutex.Unlock()
	if fake.MarshalStub != nil {
		return fake.MarshalStub(msg)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.marshalReturns.result1, fake.marshalReturns.result2
}

func (fake *Protobuf) MarshalCallCount() int {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	return len(fake.marshalArgsForCall)
}

func (fake *Protobuf) MarshalArgsForCall(i int) proto.Message {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	return fake.marshalArgsForCall[i].msg
}

func (fake *Protobuf) MarshalReturns(result1 []byte, result2 error) {
	fake.MarshalStub = nil
	fake.marshalReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Protobuf) MarshalReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.MarshalStub = nil
	if fake.marshalReturnsOnCall == nil {
		fake.marshalReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.marshalReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Protobuf) Unmarshal(marshaled []byte, msg proto.Message) error {
	var marshaledCopy []byte
	if marshaled != nil {
		marshaledCopy = make([]byte, len(marshaled))
		copy(marshaledCopy, marshaled)
	}
	fake.unmarshalMutex.Lock()
	ret, specificReturn := fake.unmarshalReturnsOnCall[len(fake.unmarshalArgsForCall)]
	fake.unmarshalArgsForCall = append(fake.unmarshalArgsForCall, struct {
		marshaled []byte
		msg       proto.Message
	}{marshaledCopy, msg})
	fake.recordInvocation("Unmarshal", []interface{}{marshaledCopy, msg})
	fake.unmarshalMutex.Unlock()
	if fake.UnmarshalStub != nil {
		return fake.UnmarshalStub(marshaled, msg)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unmarshalReturns.result1
}

func (fake *Protobuf) UnmarshalCallCount() int {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return len(fake.unmarshalArgsForCall)
}

func (fake *Protobuf) UnmarshalArgsForCall(i int) ([]byte, proto.Message) {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return fake.unmarshalArgsForCall[i].marshaled, fake.unmarshalArgsForCall[i].msg
}

func (fake *Protobuf) UnmarshalReturns(result1 error) {
	fake.UnmarshalStub = nil
	fake.unmarshalReturns = struct {
		result1 error
	}{result1}
}

func (fake *Protobuf) UnmarshalReturnsOnCall(i int, result1 error) {
	fake.UnmarshalStub = nil
	if fake.unmarshalReturnsOnCall == nil {
		fake.unmarshalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unmarshalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Protobuf) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Protobuf) recordInvocation(key string, args []interface{}) {
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
