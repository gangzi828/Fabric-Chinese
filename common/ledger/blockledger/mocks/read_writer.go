
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import blockledger "github.com/hyperledger/fabric/common/ledger/blockledger"
import common "github.com/hyperledger/fabric/protos/common"
import mock "github.com/stretchr/testify/mock"
import orderer "github.com/hyperledger/fabric/protos/orderer"

//
type ReadWriter struct {
	mock.Mock
}

//append提供了一个具有给定字段的模拟函数：block
func (_m *ReadWriter) Append(block *common.Block) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Block) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//Height provides a mock function with given fields:
func (_m *ReadWriter) Height() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

//
func (_m *ReadWriter) Iterator(startType *orderer.SeekPosition) (blockledger.Iterator, uint64) {
	ret := _m.Called(startType)

	var r0 blockledger.Iterator
	if rf, ok := ret.Get(0).(func(*orderer.SeekPosition) blockledger.Iterator); ok {
		r0 = rf(startType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blockledger.Iterator)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(*orderer.SeekPosition) uint64); ok {
		r1 = rf(startType)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}
