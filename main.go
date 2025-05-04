package main

import (
	_ "embed"
	"fmt"

	"github.com/mikerybka/brass"
	"github.com/mikerybka/english"
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
		Types: map[string]brass.Type{
			"schema": brass.Type{
				Name:       english.NewName("Schema"),
				PluralName: english.NewName("Schemas"),
				IsStruct:   true,
				Fields: []brass.Field{
					{
						Name:   english.NewName("Fields"),
						TypeID: "[]field",
					},
				},
			},
			"field": brass.Type{
				Name:       english.NewName("Field"),
				PluralName: english.NewName("Fields"),
				IsStruct:   true,
				Fields: []brass.Field{
					{
						Name:   english.NewName("Name"),
						TypeID: "string",
					},
					{
						Name:   english.NewName("Type"),
						TypeID: "string",
					},
				},
			},
		},
	}
	err := app.GenerateSourceCode(dir)
	if err != nil {
		fmt.Println(err)
	}
}
