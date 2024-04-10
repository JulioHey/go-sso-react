package service

import (
	"authorizer/core/authentication"
	. "authorizer/core/autherror"
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

type mockAuthManager struct {
	mock.Mock
	AuthManager
}

type mockCodeManager struct {
	mock.Mock
	CodeManager
}

func (m *mockAuthManager) Authenticate(ctx context.Context,
	req *authentication.Request) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *mockCodeManager) CreateCode(ctx context.Context,
	req *authentication.Request) (*authentication.Response, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authentication.Response), args.Error(1)
}

var _ = Describe("Service Test Suite", Ordered, func() {
	var authManager *mockAuthManager
	var codeManager *mockCodeManager

	BeforeEach(func() {
		authManager = new(mockAuthManager)
		codeManager = new(mockCodeManager)
	})

	Context("AuthenticationSerice", Ordered, func() {
		var service AuthService

		BeforeEach(func() {
			service = AuthService{
				authManager: authManager,
				sessManager: codeManager,
			}
		})
		It("should authenticate and create code", func() {
			authManager.On("Authenticate", mock.Anything,
				mock.Anything).Return(nil)

			codeManager.On("CreateCode", mock.Anything,
				mock.Anything).Return(&authentication.Response{}, nil)

			res, err := service.AuthenticationFlow(context.Background(), &authentication.Request{})

			Expect(res).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not authenticate", func() {
			authManager.On("Authenticate", mock.Anything,
				mock.Anything).Return(ErrInvalidRequest)

			res, err := service.AuthenticationFlow(context.Background(), &authentication.Request{})

			Expect(res).To(BeNil())
			Expect(err).To(HaveOccurred())
		})

		It("should not create code", func() {
			authManager.On("Authenticate", mock.Anything,
				mock.Anything).Return(nil)

			codeManager.On("CreateCode", mock.Anything,
				mock.Anything).Return(&authentication.Response{}, ErrInvalidRequest)

			res, err := service.AuthenticationFlow(context.Background(), &authentication.Request{})

			Expect(res).To(BeNil())
			Expect(err).To(HaveOccurred())
		})
	})
})
