package structs

import "testing"

func TestEmptyProduct(t *testing.T) {

	var p = Product{}
	if p.Name != "" {
		t.Errorf("Name should be empty")
	}
}

func TestInitializedProduct(t *testing.T) {
	watch := Product{
		Name:     "AppleWatch",
		Price:    float64(349.96),
		Currency: "Euro",
	}
	if watch.Name != "AppleWatch" {
		t.Errorf("Error ")
	}
	if watch.Price != float64(349.96) {
		t.Errorf("Error in Price")
	}
}

func TestInitializedProductNoTrailingComma(t *testing.T) {
	watch := Product{"AppleWatch", float64(349.96), "Euro"}
	if watch.Name != "AppleWatch" {
		t.Errorf("Error ")
	}
	if watch.Price != float64(349.96) {
		t.Errorf("Error in Price")
	}
	if watch.Currency != "Euro" {
		t.Errorf("Error in initializing currency")
	}
}
