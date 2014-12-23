package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestProducts(t *testing.T) {
	langs := []string{"en", "ru"}
	for _, lang := range langs {
		fake.SetLang(lang)

		v := fake.Brand()
		if v == "" {
			t.Errorf("Brand failed with lang %s", lang)
		}

		v = fake.ProductName()
		if v == "" {
			t.Errorf("ProductName failed with lang %s", lang)
		}

		v = fake.Product()
		if v == "" {
			t.Errorf("Product failed with lang %s", lang)
		}

		v = fake.Model()
		if v == "" {
			t.Errorf("Model failed with lang %s", lang)
		}
	}
}
