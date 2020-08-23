package products

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

type AMQPEndpointFactory struct {
	productService ProductService
}

func NewAmqpEndpointFactory(productService ProductService) *AMQPEndpointFactory {
	return &AMQPEndpointFactory{productService: productService}
}

func(fac *AMQPEndpointFactory) GetProductByIdAMQPEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := GetProductByIdCommand{}
		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return &amqp.Message{}
		}
		resp, err := cmd.Exec(fac.productService)
		if err != nil {
			return &amqp.Message{}
		}
		return OK(resp)
	}
}

func OK(d interface{}) *amqp.Message {
	data, _ := json.Marshal(d)
	return &amqp.Message{Body: data}
}


