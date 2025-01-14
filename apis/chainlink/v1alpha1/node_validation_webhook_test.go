package v1alpha1

import (
	"fmt"

	"github.com/kotalco/kotal/apis/shared"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

var _ = Describe("Chainlink node validation", func() {
	createCases := []struct {
		Title  string
		Node   *Node
		Errors field.ErrorList
	}{}

	updateCases := []struct {
		Title   string
		OldNode *Node
		NewNode *Node
		Errors  field.ErrorList
	}{
		{
			Title: "updating ethereum chain ID",
			OldNode: &Node{
				ObjectMeta: v1.ObjectMeta{
					Name: "my-node",
				},
				Spec: NodeSpec{
					EthereumChainId: 111,
				},
			},
			NewNode: &Node{
				ObjectMeta: v1.ObjectMeta{
					Name: "my-node",
				},
				Spec: NodeSpec{
					EthereumChainId: 222,
				},
			},
			Errors: field.ErrorList{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.ethereumChainId",
					BadValue: "222",
					Detail:   "field is immutable",
				},
			},
		},
		{
			Title: "updating LINK contract address",
			OldNode: &Node{
				ObjectMeta: v1.ObjectMeta{
					Name: "my-node",
				},
				Spec: NodeSpec{
					LinkContractAddress: "0x514910771af9ca656af840dff83e8264ecf986ca",
				},
			},
			NewNode: &Node{
				ObjectMeta: v1.ObjectMeta{
					Name: "my-node",
				},
				Spec: NodeSpec{
					LinkContractAddress: "0x326c977e6efc84e512bb9c30f76e30c160ed06fb",
				},
			},
			Errors: field.ErrorList{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.linkContractAddress",
					BadValue: "0x326c977e6efc84e512bb9c30f76e30c160ed06fb",
					Detail:   "field is immutable",
				},
			},
		},
	}

	Context("While creating node", func() {
		for _, c := range createCases {
			func() {
				cc := c
				It(fmt.Sprintf("Should validate %s", cc.Title), func() {
					cc.Node.Default()
					err := cc.Node.ValidateCreate()

					errStatus := err.(*errors.StatusError)

					causes := shared.ErrorsToCauses(cc.Errors)

					Expect(errStatus.ErrStatus.Details.Causes).To(ContainElements(causes))
				})
			}()
		}
	})

	Context("While updating node", func() {
		for _, c := range updateCases {
			func() {
				cc := c
				It(fmt.Sprintf("Should validate %s", cc.Title), func() {
					cc.OldNode.Default()
					cc.NewNode.Default()
					err := cc.NewNode.ValidateUpdate(cc.OldNode)

					errStatus := err.(*errors.StatusError)

					causes := shared.ErrorsToCauses(cc.Errors)

					Expect(errStatus.ErrStatus.Details.Causes).To(ContainElements(causes))
				})
			}()
		}
	})

})
