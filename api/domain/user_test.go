package domain_test

import (
	"github.com/martinjirku/zasobar/domain"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {

	Describe("Creation of user", func() {
		Context("with valid password", func() {

			It("should create not be error", func() {
				_, err := domain.NewUserWithPassword("martinjirku@gmail.com", "totalneDobreHeslo")
				Expect(err).To(BeNil())
			})
			It("should create new user with email", func() {
				user, _ := domain.NewUserWithPassword("martinjirku@gmail.com", "totalneDobreHeslo")
				Expect(user.Email).To(Equal("martinjirku@gmail.com"))
			})
			It("should create new user with hashed password", func() {
				user, _ := domain.NewUserWithPassword("martinjirku@gmail.com", "totalneDobreHeslo")
				Expect(user.Password).NotTo(BeNil())
				Expect(user.Password).NotTo(Equal("totalneDobreHeslo"))
			})
		})
	})

	Describe("Password validation", func() {
		var user domain.User
		var password string

		BeforeEach(func() {
			password = "totalneDobreHeslo"
			user, _ = domain.NewUserWithPassword("martinjirku@gmail.com", password)
		})

		It("should successfully validate password", func() {
			Expect(user.VerifyPassword(password)).To(BeTrue())
		})
		It("should not successfully validate password", func() {
			Expect(user.VerifyPassword(password + "asdf")).To(BeFalse())
		})
	})
})
