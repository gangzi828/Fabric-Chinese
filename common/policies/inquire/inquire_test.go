
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


package inquire

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name       string
	policy     string
	expected   map[string]struct{}
	principals []*msp.MSPPrincipal
}

func createPrincipals(orgNames ...string) []*msp.MSPPrincipal {
	principals := make([]*msp.MSPPrincipal, 0)
	appendPrincipal := func(orgName string) {
		principals = append(principals, &msp.MSPPrincipal{
			PrincipalClassification: msp.MSPPrincipal_ROLE,
			Principal:               utils.MarshalOrPanic(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: orgName})})
	}
	for _, org := range orgNames {
		appendPrincipal(org)
	}
	return principals
}

var cases = []testCase{
	{
		name:   "orOfAnds",
		policy: "OR(AND('A.member', 'B.member'), 'C.member', AND('A.member', 'D.member'))",
		expected: map[string]struct{}{
			fmt.Sprintf("%v", []string{"A", "B"}): {},
			fmt.Sprintf("%v", []string{"C"}):      {},
			fmt.Sprintf("%v", []string{"A", "D"}): {},
		},
		principals: createPrincipals("A", "B", "C", "D", "A"),
	},
	{
		name:   "andOfOrs",
		policy: "AND('A.member', 'C.member', OR('B.member', 'D.member'))",
		expected: map[string]struct{}{
			fmt.Sprintf("%v", []string{"A", "C", "B"}): {},
			fmt.Sprintf("%v", []string{"A", "C", "D"}): {},
		},
		principals: createPrincipals("A", "C", "B", "D"),
	},
	{
		name:   "orOfOrs",
		policy: "OR('A.member', OR('B.member', 'C.member'))",
		expected: map[string]struct{}{
			fmt.Sprintf("%v", []string{"A"}): {},
			fmt.Sprintf("%v", []string{"B"}): {},
			fmt.Sprintf("%v", []string{"C"}): {},
		},
		principals: createPrincipals("A", "B", "C"),
	},
	{
		name:   "andOfAnds",
		policy: "AND('A.member', AND('B.member', 'C.member'), AND('D.member','A.member'))",
		expected: map[string]struct{}{
			fmt.Sprintf("%v", []string{"A", "B", "C", "D", "A"}): {},
		},
		principals: createPrincipals("A", "B", "C", "D"),
	},
}

func TestSatisfiedBy(t *testing.T) {

	mspId := func(principal *msp.MSPPrincipal) string {
		role := &msp.MSPRole{}
		proto.Unmarshal(principal.Principal, role)
		return role.MspIdentifier
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			p, err := cauthdsl.FromString(test.policy)
			assert.NoError(t, err)

			ip := NewInquireableSignaturePolicy(p)
			satisfiedBy := ip.SatisfiedBy()

			actual := make(map[string]struct{})
			for _, ps := range satisfiedBy {
				var principals []string
				for _, principal := range ps {
					principals = append(principals, mspId(principal))
				}
				actual[fmt.Sprintf("%v", principals)] = struct{}{}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}
