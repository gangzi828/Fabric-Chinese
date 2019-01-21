
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package aclmgmt

import (
	"github.com/hyperledger/fabric/common/flogging"
)

var aclLogger = flogging.MustGetLogger("aclmgmt")

type ACLProvider interface {
//checkacl使用
//IDFIN。IDinfo是一个对象，如SignedProposal，其中
//可以提取ID以根据策略进行测试
	CheckACL(resName string, channelID string, idinfo interface{}) error
}
