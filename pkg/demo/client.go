package mongodb

// As DoSomething is defined in the interface, DemoClient struct
// would also have to implement it. This interface can later be used for
// generating mocks and can be used for unit testing. It can also be used
// with dependency injection approach
type DemoInterface interface {
	DoSomething(id string) string
}

type DemoClient struct {
	args any
}

func NewMongoRepositor() DemoInterface {
	var args any

	client := &DemoClient{
		args: args,
	}

	return client
}
