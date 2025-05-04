package main

import (
	_ "embed"
	"fmt"

	"github.com/mikerybka/brass"
)

//go:embed icon.png
var icon []byte

func main() {
	dir := "/home/mike/src/github.com/mikerybka/schemacafe"
	app := &brass.App{
		Repo:     "github.com/mikerybka/schemacafe",
		Name:     "Schema.cafe",
		Icon:     icon,
		CoreType: "schema",
		Types:    map[string]brass.Type{},
	}
	err := app.GenerateSourceCode(dir)
	if err != nil {
		fmt.Println(err)
	}
}
