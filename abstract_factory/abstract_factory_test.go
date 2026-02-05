package abstractfactory

import (
	"log"
	"runtime"
	"testing"
)

func runFuncName() string {
	pc := make([]uintptr, 1)
	n := runtime.Callers(2, pc)
	if n == 0 {
		return "unknown"
	}
	f := runtime.FuncForPC(pc[0])
	if f == nil {
		return "unknown"
	}
	return f.Name()
}

func getMainAndDetail(factory DAOFactory) {
	log.Printf("Running test 函数名: %s\n", runFuncName())
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestAbstractFactory(t *testing.T) {
	log.Printf("Running test 函数名: %s\n", runFuncName())
	getMainAndDetail(&RDBDAOFactory{})
	getMainAndDetail(&XMLDAOFactory{})
}
