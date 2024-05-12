package postgres

import (
	"authorizer/core/authentication"
	"authorizer/core/user"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
)

var _ = Describe("AuthenticationManager", func() {
	var db *gorm.DB
	var mock sqlmock.Sqlmock
	repo := &AuthStore{}

	BeforeEach(func() {
		var err error
		var conn *sql.DB
		conn, mock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred(), "failed to open connection")
		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_0",
			DriverName:           "postgres",
			Conn:                 conn,
			PreferSimpleProtocol: true,
		})
		db, err = gorm.Open(dialector, &gorm.Config{})
		Expect(err).ShouldNot(HaveOccurred(), "failed to open connection")
	})

	Context("Authenticate", func() {

		It("Should authenticate a user", func() {
			mock.ExpectBegin()
			tx := db.Begin()
			email := "teste@test.com"
			userID := uuid.New()

			rows := sqlmock.NewRows([]string{"id"}).AddRow(userID)
			password := "123456"
			res, _ := user.HashPassword(password, userID.String())
			passwordRows := sqlmock.NewRows([]string{"password"}).AddRow(res.Password)

			query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`)

			mock.ExpectQuery(query).
				WithArgs(email, 1).
				WillReturnRows(rows)

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "passwords" WHERE user_id = $1 ORDER BY "passwords"."id" LIMIT $2`)).
				WithArgs(userID, 1).WillReturnRows(passwordRows)

			err := repo.Authenticate(context.Background(), tx, &authentication.Request{
				Extra: map[string]any{
					"email":    email,
					"password": password,
				},
			})

			Expect(err).ToNot(HaveOccurred())
		})

		It("should not authenticate a user if it dont exists", func() {
			mock.ExpectBegin()
			tx := db.Begin()
			email := "teste@test.com"

			password := "123456"

			query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`)

			mock.ExpectQuery(query).
				WithArgs(email, 1).WillReturnError(gorm.ErrRecordNotFound)

			err := repo.Authenticate(context.Background(), tx, &authentication.Request{
				Extra: map[string]any{
					"email":    email,
					"password": password,
				},
			})

			Expect(err).To(Equal(gorm.ErrRecordNotFound))
		})

		It("should not authenticate a user if password dont match", func() {
			mock.ExpectBegin()
			tx := db.Begin()
			email := "teste@test.com"
			userID := uuid.New()

			rows := sqlmock.NewRows([]string{"id"}).AddRow(userID)
			password := "123456"
			res, _ := user.HashPassword(password, userID.String())
			passwordRows := sqlmock.NewRows([]string{"password"}).AddRow(res.Password)

			query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`)

			mock.ExpectQuery(query).
				WithArgs(email, 1).
				WillReturnRows(rows)

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "passwords" WHERE user_id = $1 ORDER BY "passwords"."id" LIMIT $2`)).
				WithArgs(userID, 1).WillReturnRows(passwordRows)

			err := repo.Authenticate(context.Background(), tx, &authentication.Request{
				Extra: map[string]any{
					"email":    email,
					"password": "password",
				},
			})

			Expect(err).To(Equal(user.InvalidUserPasswordError))
		})
	})
})
