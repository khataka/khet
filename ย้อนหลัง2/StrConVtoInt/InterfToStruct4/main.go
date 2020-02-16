package main

import "fmt"

type Customer struct {
	Name string `json:"name"`
}

func main() {

	receivedData := somefunc()

	receivedCustomer := getCustomerFromDTO(receivedData)
	fmt.Println(receivedCustomer)
}

func getCustomerFromDTO(data interface{}) Customer {
	m := data.(map[string]interface{})
	customer := Customer{}
	if name, ok := m["name"].(string); ok {
		customer.Name = name
	}
	return customer
}
