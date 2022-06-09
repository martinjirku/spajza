package models_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/martinjirku/zasobar/models"
)

var _ = Describe("User", func() {

	Describe("Creation of user", func() {
		Context("with valid password", func() {

			It("should create not be error", func() {
				_, err := models.NewUserWithPassword("martinjirku@gmail.com", "totalneDobreHeslo")
				Expect(err).To(BeNil())
			})
			It("should create new user with email", func() {
				user, _ := models.NewUserWithPassword("martinjirku@gmail.com", "totalneDobreHeslo")
				Expect(user.Email).To(Equal("martinjirku@gmail.com"))
			})
			It("should create new user with hashed password", func() {
				user, _ := models.NewUserWithPassword("martinjirku@gmail.com", "totalneDobreHeslo")
				Expect(user.Password).NotTo(BeNil())
				Expect(user.Password).NotTo(Equal("totalneDobreHeslo"))
			})
		})
	})

	Describe("Password validation", func() {
		var user models.User
		var password string

		BeforeEach(func() {
			password = "totalneDobreHeslo"
			user, _ = models.NewUserWithPassword("martinjirku@gmail.com", password)
		})

		It("should successfully validate password", func() {
			Expect(user.VerifyPassword(password)).To(BeTrue())
		})
		It("should not successfully validate password", func() {
			Expect(user.VerifyPassword(password + "asdf")).To(BeFalse())
		})
	})
})
