
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import common "github.com/hyperledger/fabric/protos/common"
import mock "github.com/stretchr/testify/mock"

//policyEvaluator是policyEvaluator类型的自动生成的模拟类型
type PolicyEvaluator struct {
	mock.Mock
}

//Evaluate提供了一个具有给定字段的模拟函数：policyBytes、signatureSet
func (_m *PolicyEvaluator) Evaluate(policyBytes []byte, signatureSet []*common.SignedData) error {
	ret := _m.Called(policyBytes, signatureSet)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []*common.SignedData) error); ok {
		r0 = rf(policyBytes, signatureSet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
