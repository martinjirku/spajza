package float_test

import (
	"testing"

	"github.com/martinjirku/zasobar/pkg/float"
)

func Test_Compare(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		result := float.Compare(1.0, 2.0)
		if result != 1 {
			t.Errorf("Expected %d, received %d", 1, result)
		}
	})
	t.Run("should return -1", func(t *testing.T) {
		result := float.Compare(3.0, 2.0)
		if result != -1 {
			t.Errorf("Expected %d, received %d", -1, result)
		}
	})
	t.Run("should return 0", func(t *testing.T) {
		result := float.Compare(0.1, 0.1)
		if result != 0 {
			t.Errorf("Expected %d, received %d", 0, result)
		}
	})
}
