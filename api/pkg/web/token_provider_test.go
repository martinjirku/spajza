package web_test

import (
	"testing"
	"time"

	"github.com/martinjirku/zasobar/pkg/web"
)

func TestProvideValidToken(t *testing.T) {
	tokenProvider := web.NewTokenProvider("secret", 10, "zasobar")
	token, err := tokenProvider.GetToken("username", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Error(err)
	}
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6YXNvYmFyIiwic3ViIjoidXNlcm5hbWUiLCJleHAiOjE1NDYzMDE0MDB9.xRJ1SeWS9J8WlFx1W4Hy-RcCOAxaBEqXP-UMVyKthUg"
	if token != expected {
		t.Errorf("Expected %s, but got %s", expected, token)
	}
}
