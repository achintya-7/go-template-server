package mongodb

// As GetUserName is defined in the interface, MongoRepository struct
// would also have to implement it. This interface can later be used for 
// generating mocks and can be used for unit testing. It can also be used
// with dependency injection approach
type MongoRepositoryInterface interface {
	GetUserName(id string) string
}

type MongoRepository struct {
	something any
}

func NewMongoRepositor() MongoRepositoryInterface {
	client := &MongoRepository{}

	return client
}
