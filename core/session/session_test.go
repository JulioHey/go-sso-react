package session

import (
	. "authorizer/core/autherror"
	"github.com/hashicorp/errwrap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session", func() {
	DescribeTable("Validate", func(s *Session, expectedErr error) {
		err := s.Validate()
		if expectedErr == nil {
			Expect(err).ToNot(HaveOccurred())
		} else {
			Expect(errwrap.Get(err, expectedErr.Error())).To(Equal(expectedErr))
		}
	},
		Entry("Valid Session", &Session{
			State: "state",
		}, nil),
		Entry("Invalid Session", &Session{}, ErrInvalidRequest),
	)
})