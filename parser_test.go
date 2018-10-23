package main_test

import (
	"testing"

	main "github.com/identitii/sprocbind"
	"github.com/stretchr/testify/require"
)

func TestParseTable(t *testing.T) {

	result, err := main.Parse(`
		CREATE TYPE MyAmazingTable AS TABLE   
		( 
			UPPERCASE_VARCHAR   VARCHAR(500), 
			lowercase_varchar   varchar(500),
			"what_a_big_int"    BIGINT
		);  
		GO 
	`, false)
	require.NoError(t, err)
	require.Empty(t, result.Procedures)
	require.Equal(t, []main.Table{
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
	}, result.Tables)

}

func TestParseProcedure(t *testing.T) {

	// From: https://docs.microsoft.com/en-us/sql/t-sql/statements/create-procedure-transact-sql?view=sql-server-2017
	result, err := main.Parse(`
		CREATE PROCEDURE dbo.usp_add_kitchen @dept_id int, @kitchen_count int NOT NULL  
		WITH EXECUTE AS OWNER, SCHEMABINDING, NATIVE_COMPILATION  
		AS  
		BEGIN ATOMIC WITH (TRANSACTION ISOLATION LEVEL = SNAPSHOT, LANGUAGE = N'us_english')  
		
		UPDATE dbo.Departments  
		SET kitchen_count = ISNULL(kitchen_count, 0) + @kitchen_count  
		WHERE id = @dept_id  
		END;  
		GO  
	`, false)
	require.NoError(t, err)
	require.Empty(t, result.Tables)
	require.Equal(t, []main.Procedure{
		main.Procedure{
			Schema: "dbo",
			Name:   "usp_add_kitchen",
			Params: []main.ProcedureParam{
				main.ProcedureParam{
					Name:     "@dept_id",
					DataType: "int",
				},
				main.ProcedureParam{
					Name:     "@kitchen_count",
					DataType: "int",
				},
			},
		},
	}, result.Procedures)

}
