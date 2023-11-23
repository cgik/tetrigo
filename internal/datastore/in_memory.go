package datastore

type InMemoryInterface struct {
	client map[string]interface{}
}

func NewInMemoryInterface() *InMemoryInterface {
	return &InMemoryInterface{
		client: make(map[string]interface{}),
	}
}

func (i *InMemoryInterface) Insert(collection string, item interface{}) error {
	i.client[collection] = item
	return nil
}

func (i *InMemoryInterface) FindById(collection string, document string, id int) error {
	return nil
}
