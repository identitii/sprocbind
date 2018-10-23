package main_test

import (
	"io/ioutil"
	"log"
	"testing"

	main "github.com/identitii/sprocbind"
	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	source, err := main.Generate("mypackage", &main.ParseResult{
		Tables: []main.Table{
			main.Table{
				Name: "MyAmazingTable",
				Columns: []main.Column{
					main.Column{
						Name:     "UPPERCASE_VARCHAR",
						DataType: "VARCHAR(500)",
					},
					main.Column{
						Name:     "lowercase_varchar",
						DataType: "varchar(500)",
					},
					main.Column{
						Name:     `"what_a_big_int"`,
						DataType: "BIGINT",
					},
				},
			},
		}})
	require.NoError(t, err)
	log.Println(source)
}

func TestGenerator1(t *testing.T) {
	tsql, err := ioutil.ReadFile("testdata/test1.sql")
	require.NoError(t, err)

	result, err := main.Parse(string(tsql), true)

	log.Printf("result: %+v", result)

	source, err := main.Generate("mypackage", result)
	require.NoError(t, err)

	log.Println(source)
	panic(123)
}

func TestGenerator2(t *testing.T) {
	tsql, err := ioutil.ReadFile("testdata/test2.sql")
	require.NoError(t, err)

	result, err := main.Parse(string(tsql), true)

	log.Printf("result: %+v", result)

	source, err := main.Generate("main", result)
	require.NoError(t, err)

	log.Println(source)
}
