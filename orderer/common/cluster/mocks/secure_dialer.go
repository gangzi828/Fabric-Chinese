
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
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"

//SecureDialer是SecureDialer类型的自动生成的模拟类型
type SecureDialer struct {
	mock.Mock
}

//dial提供了一个具有给定字段的模拟函数：address、verifyfunc
func (_m *SecureDialer) Dial(address string, verifyFunc cluster.RemoteVerifier) (*grpc.ClientConn, error) {
	ret := _m.Called(address, verifyFunc)

	var r0 *grpc.ClientConn
	if rf, ok := ret.Get(0).(func(string, cluster.RemoteVerifier) *grpc.ClientConn); ok {
		r0 = rf(address, verifyFunc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.ClientConn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, cluster.RemoteVerifier) error); ok {
		r1 = rf(address, verifyFunc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
