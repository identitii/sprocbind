package main

type Procedure struct {
	Database string
	Schema   string
	Name     string

	Params []ProcedureParam
}

// func (p *Procedure) NameGo() string {
// 	n := strings.ToLower(stripQuotes(p.Name))
// 	return uppercaseIds(strcase.ToCamel(n))
// }

type ProcedureParam struct {
	DefaultValue string
	DataType     string
	Name         string
}

// func (p *ProcedureParam) NameGo() string {
// 	n := strings.TrimPrefix(p.Name, "@")
// 	n = strings.ToLower(stripQuotes(n))
// 	return uppercaseIds(strcase.ToLowerCamel(n))
// }

// func (p *ProcedureParam) NameClean() string {
// 	n := strings.TrimPrefix(p.Name, "@")
// 	return strings.ToLower(stripQuotes(n))
// }

// func (p *ProcedureParam) DataTypeGo() string {
// 	return goDataType(p.DataType)
// }

type Table struct {
	Name    string
	Columns []Column
}

// func (p *Table) NameGo() string {
// 	n := strings.ToLower(stripQuotes(p.Name))
// 	return uppercaseIds(strcase.ToCamel(n))
// }

type Column struct {
	DataType string
	Name     string
}

// func (p *Column) NameGo() string {
// 	n := strings.ToLower(stripQuotes(p.Name))
// 	return uppercaseIds(strcase.ToCamel(n))
// }

// func (c *Column) NameClean() string {
// 	return stripQuotes(c.Name)
// }

// func (c *Column) DataTypeGo() string {
// 	return goDataType(c.DataType)
// }

// func goDataType(t string) string {
// 	if strings.Contains(t, "(") {
// 		t = t[0:strings.Index(t, "(")]
// 	}
// 	return SQLTypeToGoType(t)
// }

// func stripQuotes(s string) string {
// 	s = strings.Replace(s, `"`, "", -1)
// 	return strings.Replace(s, `'`, "", -1)
// }

// func uppercaseIds(s string) string {
// 	var re = regexp.MustCompile(`(?m)(Id)(?:[A-Z]|$)`)

// 	for _, match := range re.FindAllStringIndex(s, -1) {
// 		s = s[0:match[0]] + "ID" + s[match[1]:]
// 	}
// 	return s
// }
