package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct { // JSON数据修饰符
	ID         string      `json:"id"`
	Name       string      `json:"name,omitempty"` // 省略空值
	TotalPrice float64     `json:"total_price"`
	Items      []OrderItem `json:"items"`
}

func main() {
	order := Order{
		ID:         "0",
		Name:       "learn go",
		TotalPrice: 30,
		Items: []OrderItem{
			{"0", "Book00", 10},
			{"1", "Book01", 10},
			{"2", "Book02", 10}},
	}

	// 序列化
	data, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)

	// 反序列化
	// 也可以反序列化为"map[string]interface{}"，再通过Type Assertion进行类型转换
	err = json.Unmarshal([]byte("{\"id\":\"0\",\"name\":\"learn go\",\"total_price\":30,\"items\":[{\"id\":\"0\",\"name\":\"Book00\",\"price\":10},{\"id\":\"0\",\"name\":\"Book00\",\"price\":10},{\"id\":\"0\",\"name\":\"Book00\",\"price\":10}]}"), &order)
	if err != nil {
		panic(err)
	}
	fmt.Println(order)
}
