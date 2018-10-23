package main

import "strings"

var sqlTypeToGoType = map[string]string{
	"TINYINT":          "int8",
	"SMALLINT":         "int16",
	"INT":              "int32",
	"BIGINT":           "int64",
	"REAL":             "float32",
	"FLOAT":            "float64",
	"VARBINARY":        "[]byte",
	"VARCHAR":          "string",
	"NVARCHAR":         "string",
	"BIT":              "bool",
	"DECIMAL":          "string",
	"SMALLMONEY":       "string",
	"MONEY":            "string",
	"SMALLDATETIME":    "time.Time",
	"DATETIME":         "time.Time",
	"DATETIME2":        "time.Time",
	"DATE":             "time.Time",
	"TIME":             "time.Time",
	"DATETIMEOFFSET":   "time.Time",
	"CHAR":             "byte",
	"NCHAR":            "string",
	"UNIQUEIDENTIFIER": "mssql.UniqueIdentifier",
	"XML":              "string",
	"TEXT":             "string",
	"NTEXT":            "string",
	"IMAGE":            "[]byte",
	"SQL_VARIANT":      "[]byte",
	"BINARY":           "[]byte",
}

func SQLTypeToGoType(typeName string) string {
	if t, ok := sqlTypeToGoType[strings.ToUpper(typeName)]; ok {
		return t
	}

	// Must be a custom type (table only), just return it
	return "*" + goName(typeName)
}
