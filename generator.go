package main

import (
	"bytes"
	"log"
	"regexp"
	"strings"

	"github.com/alecthomas/template"
	"github.com/iancoleman/strcase"
)

func goName(s string) string {
	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)
	s = strings.TrimPrefix(s, "@")
	s = strcase.ToSnake(s)
	return uppercaseIds(strcase.ToCamel(s))
}

var funcMap = template.FuncMap{
	"GoNameLower": func(s string) string {
		s = strings.Replace(s, "[", "", -1)
		s = strings.Replace(s, "]", "", -1)
		s = strings.TrimPrefix(s, "@")
		s = strcase.ToSnake(s)
		return uppercaseIds(strcase.ToLowerCamel(s))
	},
	"GoName": goName,
	"GoDataType": func(s string) string {
		if strings.Contains(s, "(") {
			s = s[0:strings.Index(s, "(")]
		}
		return SQLTypeToGoType(s)
	},
	"Clean": func(s string) string {
		s = strings.TrimPrefix(s, "@")
		return stripQuotes(s)
	},
}

func Generate(packageName string, result *ParseResult) (string, error) {
	tmpl, err := template.New(".").Funcs(funcMap).Parse(tFile)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, map[string]interface{}{
		"packageName":  packageName,
		"tables":       result.Tables,
		"normalTables": result.NormalTables,
		"procedures":   result.Procedures,
	}); err != nil {
		log.Fatal(err)
	}

	return string(buf.Bytes()), nil
}

func stripQuotes(s string) string {
	s = strings.Replace(s, `"`, "", -1)
	return strings.Replace(s, `'`, "", -1)
}

func uppercaseIds(s string) string {
	var re = regexp.MustCompile(`(?m)(Id)(?:[A-Z]|$)`)

	for _, match := range re.FindAllStringIndex(s, -1) {
		s = s[0:match[0]] + "ID" + s[match[1]:]
	}
	return s
}
