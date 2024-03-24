package authentication

import (
	. "authorizer/core/autherror"
	"github.com/hashicorp/errwrap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authentication", func() {
	DescribeTable("Request Validatation", func(req Request, expectedErr error) {
		err := req.Validate()
		if expectedErr == nil {
			Expect(err).ToNot(HaveOccurred())
		} else {
			Expect(errwrap.Get(err, expectedErr.Error())).To(Equal(expectedErr))
		}
	},
		Entry("Valid Request", Request{
			Scope:        "scope",
			ClientID:     "client_id",
			ResponseType: "code",
			RedirectURI:  "redirect_uri",
			State:        "state",
			Nonce:        "nonce",
		}, nil),
		Entry("Invalid Request", Request{
			ClientID:     "client_id",
			ResponseType: "code",
			RedirectURI:  "redirect_uri",
			State:        "state",
			Nonce:        "nonce",
		}, ErrInvalidRequest),
		Entry("Invalid Request", Request{
			ClientID:     "client_id",
			ResponseType: "type esquisite",
			RedirectURI:  "redirect_uri",
			State:        "state",
			Nonce:        "nonce",
		}, ErrInvalidRequest),
	)
})
