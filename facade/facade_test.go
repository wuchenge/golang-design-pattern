package facade

import "testing"

var expect = `Facade Test:
AModuleAPI: TestA called
BModuleAPI: TestB called`

func TestFacade(t *testing.T) {
	api := NewAPI()
	result := api.Test()
	if result != expect {
		t.Errorf("Expected: %s, Got: %s", expect, result)
	}
}
