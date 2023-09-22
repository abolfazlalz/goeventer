package lib

import "fmt"

type Print struct {
}

func (p Print) Name() string {
	return "print"
}

func (p Print) Call(params ...interface{}) interface{} {
	fmt.Println(params...)
	return nil
}
