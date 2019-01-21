
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import mock "github.com/stretchr/testify/mock"
import orderer "github.com/hyperledger/fabric/protos/orderer"

//处理程序是处理程序类型的自动生成的模拟类型
type Handler struct {
	mock.Mock
}

//onstep提供了一个具有给定字段的模拟函数：channel、sender、req
func (_m *Handler) OnStep(channel string, sender uint64, req *orderer.StepRequest) (*orderer.StepResponse, error) {
	ret := _m.Called(channel, sender, req)

	var r0 *orderer.StepResponse
	if rf, ok := ret.Get(0).(func(string, uint64, *orderer.StepRequest) *orderer.StepResponse); ok {
		r0 = rf(channel, sender, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orderer.StepResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint64, *orderer.StepRequest) error); ok {
		r1 = rf(channel, sender, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//onsubmit提供了一个具有给定字段的模拟函数：channel、sender、req
func (_m *Handler) OnSubmit(channel string, sender uint64, req *orderer.SubmitRequest) (*orderer.SubmitResponse, error) {
	ret := _m.Called(channel, sender, req)

	var r0 *orderer.SubmitResponse
	if rf, ok := ret.Get(0).(func(string, uint64, *orderer.SubmitRequest) *orderer.SubmitResponse); ok {
		r0 = rf(channel, sender, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orderer.SubmitResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint64, *orderer.SubmitRequest) error); ok {
		r1 = rf(channel, sender, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
