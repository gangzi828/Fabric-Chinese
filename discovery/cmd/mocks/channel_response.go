
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import client "github.com/hyperledger/fabric/discovery/client"
import discovery "github.com/hyperledger/fabric/protos/discovery"
import mock "github.com/stretchr/testify/mock"

//channelResponse是针对channelResponse类型自动生成的模拟类型
type ChannelResponse struct {
	mock.Mock
}

//config提供了一个具有给定字段的模拟函数：
func (_m *ChannelResponse) Config() (*discovery.ConfigResult, error) {
	ret := _m.Called()

	var r0 *discovery.ConfigResult
	if rf, ok := ret.Get(0).(func() *discovery.ConfigResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discovery.ConfigResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//背书人为给定字段提供模拟函数：invocationChain，f
func (_m *ChannelResponse) Endorsers(invocationChain client.InvocationChain, f client.Filter) (client.Endorsers, error) {
	ret := _m.Called(invocationChain, f)

	var r0 client.Endorsers
	if rf, ok := ret.Get(0).(func(client.InvocationChain, client.Filter) client.Endorsers); ok {
		r0 = rf(invocationChain, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Endorsers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.InvocationChain, client.Filter) error); ok {
		r1 = rf(invocationChain, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//对等端为给定字段提供模拟函数：invocationChain
func (_m *ChannelResponse) Peers(invocationChain ...*discovery.ChaincodeCall) ([]*client.Peer, error) {
	_va := make([]interface{}, len(invocationChain))
	for _i := range invocationChain {
		_va[_i] = invocationChain[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*client.Peer
	if rf, ok := ret.Get(0).(func(...*discovery.ChaincodeCall) []*client.Peer); ok {
		r0 = rf(invocationChain...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*client.Peer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*discovery.ChaincodeCall) error); ok {
		r1 = rf(invocationChain...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
