package main

import (
	"fmt"
	"log"
	"protojsonDemo/todo"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	value, err := structpb.NewValue(map[string]interface{}{
		"firstName": "John",
		"lastName":  "Smith",
		"isAlive":   true,
		"age":       27,
		"address": map[string]interface{}{
			"streetAddress": "21 2nd Street",
			"city":          "New York",
			"state":         "NY",
			"postalCode":    "10021-3100",
		},
		"phoneNumbers": []interface{}{
			map[string]interface{}{
				"type":   "home",
				"number": "212 555-1234",
			},
			map[string]interface{}{
				"type":   "office",
				"number": "646 555-4567",
			},
		},
		"children": []interface{}{},
		"spouse":   nil,
	})
	if err != nil {
		log.Panic(err)
	}
	exampleTodo := &todo.Todo{
		TodoId:    uuid.New().String(),
		Status:    todo.Status_STATUS_CREATED,
		Content:   "Do stuff",
		CreatedAt: timestamppb.Now(),
		Data:      value,
	}

	m := protojson.MarshalOptions{
		Indent:          "  ",
		EmitUnpopulated: true,
	}
	jsonBytes, _ := m.Marshal(exampleTodo)
	fmt.Println(string(jsonBytes))
}
