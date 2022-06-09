package services_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/martinjirku/zasobar/services"
)

var _ = Describe("TokenProvider", func() {
	It("should provide valid token", func() {
		tokenProvider := services.NewTokenProvider("secret", 10, "zasobar")
		token, err := tokenProvider.GetToken("username", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))
		Expect(err).To(BeNil())
		Expect(token).To(Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6YXNvYmFyIiwic3ViIjoidXNlcm5hbWUiLCJleHAiOjE1NDYzMDE0MDB9.xRJ1SeWS9J8WlFx1W4Hy-RcCOAxaBEqXP-UMVyKthUg"))
	})
})
