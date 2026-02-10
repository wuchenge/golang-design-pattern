package chainofresponsibility

func Example() {
	projectManager := NewProjectManagerChain()
	depManager := NewDepManagerChain()
	generalManager := NewGeneralManagerChain()
	depManager.SetSuccessor(generalManager)
	projectManager.SetSuccessor(depManager)

	projectManager.HandleFeeRequest("bob", 100)
	projectManager.HandleFeeRequest("tom", 1000)
	projectManager.HandleFeeRequest("ada", 10000)
	// Output:
	// Project manager permit bob 100 fee request
	// Dep manager permit tom 1000 fee request
	// General manager permit ada 10000 fee requests
}
