
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import chaincode "github.com/hyperledger/fabric/common/chaincode"
import mock "github.com/stretchr/testify/mock"

//
type LifeCycleChangeListener struct {
	mock.Mock
}

//LifecycleChangeListener提供具有给定字段的模拟函数：channel、chaincodes
func (_m *LifeCycleChangeListener) LifeCycleChangeListener(channel string, chaincodes chaincode.MetadataSet) {
	_m.Called(channel, chaincodes)
}
