package main

import (
	"fmt"
	"github.com/FDKevin0/converter/table2struct"
)

func main() {
	t2t := table2struct.NewTable2Struct()

	err := t2t.
		SavePath("/home/go/project/model/model.go").
		Dsn("root:root@tcp(localhost:3306)/test?charset=utf8").
		Run()
	fmt.Println(err)
}
