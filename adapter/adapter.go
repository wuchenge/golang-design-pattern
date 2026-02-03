package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

type adapteeImpl struct{}

func (a *adapteeImpl) SpecificRequest() string {
	// Implementation of the specific request
	return "Specific Request from Adapter"
}

func NewAdapter(adaptee Adaptee) Target {
	return &targetImpl{adapter: adaptee}
}

type targetImpl struct {
	adapter Adaptee
}

func (t *targetImpl) Request() string {
	return t.adapter.SpecificRequest()
}
