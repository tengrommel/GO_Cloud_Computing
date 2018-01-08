package abstract_factory

import "fmt"

// OrderMainDAO为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO为订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory DAO抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// RDBMainDAO 为关系型数据库的OrderMainDAO实现
type RDBMainDAO struct {}

// SaveOrderMain ...
func (*RDBMainDAO)SaveOrderMain()  {
	fmt.Print("rdb main save\n")
}

// RDBDetailDAO为关系型数据库的OrderDetailDAO实现
type RDBDetailDAO struct {}

// SaveOrderDetail ...
func (*RDBDetailDAO)SaveOrderDetail()  {
	fmt.Print("rdb detail save\n")
}

// RDBDAOFactory是RDB抽象工厂实现
type RDBDAOFactory struct {}

func (*RDBDAOFactory)CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory)CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO XML存储
type XMLMainDAO struct {}

// SaveOrderMain ...
func (*XMLMainDAO)SaveOrderMain()  {
	fmt.Print("xml main save\n")
}

// XMLDetailDAO XML存储
type XMLDetailDAO struct {}

// SaveOrderDetail ...
func (*XMLDetailDAO)SaveOrderDetail()  {
	fmt.Print("xml detail save")
}

// XMLDAOFactory 是RDB抽象工厂实现
type XMLDAOFactory struct {}

func (*XMLDAOFactory)CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}