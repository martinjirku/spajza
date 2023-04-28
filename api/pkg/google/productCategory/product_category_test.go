package productCategory_test

import (
	"strings"
	"testing"

	"github.com/martinjirku/zasobar/pkg/google/productCategory"
)

func TestGetProductCategories(t *testing.T) {
	text := `# Google_Product_Taxonomy_Version: 2021-09-21
5181 - Batožina a tašky
100 - Batožina a tašky > Batohy
108 - Batožina a tašky > Kozmetické a toaletné tašky
6553 - Batožina a tašky > Kozmetické kufríky
107 - Batožina a tašky > Kufre`
	categories := productCategory.GetProductCategories(strings.NewReader(text))
	if categories[0].Id != 5181 {
		t.Errorf("Expected %d, but received the %d", 5181, categories[0].Id)
	}
	if categories[0].Name != "Batožina a tašky" {
		t.Errorf("Expected %q, but received the %q", "Batožina a tašky", categories[0].Name)
	}
	if categories[4].Id != 107 {
		t.Errorf("Expected %d, but received the %d", 5181, categories[0].Id)
	}
	if categories[4].Name != "Kufre" {
		t.Errorf("Expected %q, but received the %q", "Kufre", categories[0].Name)
	}

}
