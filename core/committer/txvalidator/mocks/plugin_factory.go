
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
import validation "github.com/hyperledger/fabric/core/handlers/validation/api"

//PlugInfectory是PlugInfectory类型的自动生成的模拟类型
type PluginFactory struct {
	mock.Mock
}

//new为给定字段提供模拟函数：
func (_m *PluginFactory) New() validation.Plugin {
	ret := _m.Called()

	var r0 validation.Plugin
	if rf, ok := ret.Get(0).(func() validation.Plugin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Plugin)
		}
	}

	return r0
}
