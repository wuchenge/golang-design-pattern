package abstractfactory

import "fmt"

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDetailDAO interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type RDBMainDAO struct{}

func (dao *RDBMainDAO) SaveOrderMain() {
	fmt.Println("Saving Order Main to RDB")
}

type RDBDetailDAO struct{}

func (dao *RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("Saving Order Detail to RDB")
}

type RDBDAOFactory struct{}

func (factory *RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (factory *RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

type XMLMainDAO struct{}

func (dao *XMLMainDAO) SaveOrderMain() {
	fmt.Println("Saving Order Main to XML")
}

type XMLDetailDAO struct{}

func (dao *XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("Saving Order Detail to XML")
}

type XMLDAOFactory struct{}

func (factory *XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (factory *XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
