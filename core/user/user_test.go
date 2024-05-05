package user

import (
	"errors"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Context("User", func() {
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

	Context("UserPassword", func() {
		userID := uuid.New()
		DescribeTable("Validate", func(input *Password, expectedErr error) {
			err := input.Validate()
			if expectedErr == nil {
				Expect(err).ToNot(HaveOccurred())
			} else {
				Expect(err).To(HaveOccurred())
			}
		},
			Entry("Valid UserPassword", &Password{
				Password: "123",
				UserID:   uuid.New(),
			}, nil),
			Entry("Invalid UserPassword", &Password{
				Password: "123",
			}, errors.New("Teste")),
		)

		DescribeTable("HashPassword and CheckPassword",
			func(password, userID string, expectedResult *Password) {
				result, err := HashPassword(password, userID)
				if expectedResult == nil {
					Expect(err).To(HaveOccurred())
					Expect(result).To(BeNil())
				} else {
					Expect(err).ToNot(HaveOccurred())
					Expect(CheckPasswordHash(password, result.Password)).To(Equal(true))
					Expect(CheckPasswordHash("password", result.Password)).To(Equal(false))
					Expect(result.UserID).To(Equal(expectedResult.UserID))
				}
			},
			Entry("Valid Password", "123", userID.String(), &Password{
				UserID: userID,
			}),
			Entry("Valid Password", "123", "123", nil),
		)
	})
})
