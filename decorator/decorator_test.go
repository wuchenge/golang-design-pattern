package decorator

import "fmt"

func Example() {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecporator(c, 8)
	res := c.Calc()
	fmt.Printf("res %d\n", res)
	// Output:
	// res 80
}
