package grpc

type OrdersTransformer interface {
	FromAPIType(apiType interface{}) (serviceType []string, err error)
	ToAPIType(serviceType interface{}) (apiType interface{}, err error)
}

type Transformer struct{}
