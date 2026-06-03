package main

import (
	"fmt"
	"reflect"
)

type Container struct {
	Label    string
	Capacity int
	Contents []string
}

func experiment_reflect_type_and_kind() {
	var container = Container{"test1", 5, []string{}}
	t1 := reflect.TypeOf(container)
	k1 := t1.Kind()
	values := reflect.ValueOf(container)
	fmt.Printf("TypeOf Result: %v\n", t1)       //Prints the Struct type, as well as its origin module as main.Container
	fmt.Printf("Kind Result: %v\n", k1)         //Prints what kind of type it is, so in this case 'struct'
	fmt.Printf("ValueOf Result: %v\n", values)  //The contents of the struct
	fmt.Printf("ValueOf Result: %#v\n", values) //The field-value pairs of the struct
	//=> The reflect.Value struct contains key-value information on

	for i := range values.NumField() {
		typ := t1.Field(i) //Retrieves a single struct field by index.
		val := values.Field(i)
		fmt.Printf("Field: %v, Value %#v\n", typ.Name, val) //This specifically retrieves the field name attribute of the struct.
		fmt.Printf("Field struct: %v, Value %#v\n", typ, val)
		//This returns output like: "{Capacity  int  16 [1] false}", aka Fieldname, field type, byte offset from start of struct, index in struct & lastly... anonimity?

	}

}
