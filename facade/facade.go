package facade

import (
	"fmt"
)

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (api *apiImpl) Test() string {
	aRet := api.a.TestA()
	bRet := api.b.TestB()
	return fmt.Sprintf("Facade Test:\n%s\n%s", aRet, bRet)
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleAPIImpl{}
}

type AModuleAPI interface {
	TestA() string
}

type aModuleAPIImpl struct{}

func (a *aModuleAPIImpl) TestA() string {
	return "AModuleAPI: TestA called"
}
func NewBModuleAPI() BModuleAPI {
	return &bModuleAPIImpl{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleAPIImpl struct{}

func (b *bModuleAPIImpl) TestB() string {
	return "BModuleAPI: TestB called"
}
