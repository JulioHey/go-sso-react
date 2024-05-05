package user

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Context("Validate", func() {

		It("Valid User", func() {
			user := &User{
				Username:  "test",
				Email:     "test@gmail.com",
				FirstName: "test",
				LastName:  "test",
			}
			err := user.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Invalid Email", func() {
			user := &User{
				Username:  "test",
				Email:     "test",
				FirstName: "test",
				LastName:  "test",
			}
			err := user.Validate()
			Expect(err).To(HaveOccurred())
		})

		It("Invalid User", func() {
			user := &User{}
			err := user.Validate()
			Expect(err).To(HaveOccurred())
		})
	})
})
