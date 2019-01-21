
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。

package mocks

import cluster "github.com/hyperledger/fabric/orderer/common/cluster"
import mock "github.com/stretchr/testify/mock"

//LedgerFactory是LedgerFactory类型的自动生成的模拟类型
type LedgerFactory struct {
	mock.Mock
}

//GetorCreate提供了一个具有给定字段的模拟函数：chainID
func (_m *LedgerFactory) GetOrCreate(chainID string) (cluster.LedgerWriter, error) {
	ret := _m.Called(chainID)

	var r0 cluster.LedgerWriter
	if rf, ok := ret.Get(0).(func(string) cluster.LedgerWriter); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cluster.LedgerWriter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
