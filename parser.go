package main

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/fatih/color"

	"github.com/identitii/sprocbind/parser"
)

type fileParser struct {
	*parser.BasetsqlListener

	genNormalTables bool

	procedures   []Procedure
	tables       []Table
	normalTables []Table
}

// EnterCreate_or_alter_procedure is called when production create_or_alter_procedure is entered.
func (s *fileParser) EnterCreate_procedure(ctx *parser.Create_procedureContext) {
	log.Println("Starting procedure")

	nameCtx := ctx.GetProc()
	var procedure Procedure
	if nameCtx.GetDatabase() != nil {
		procedure.Database = nameCtx.GetDatabase().GetText()
	}
	if nameCtx.GetSchema() != nil {
		procedure.Schema = nameCtx.GetSchema().GetText()
	}
	if nameCtx.GetProcedure() != nil {
		procedure.Name = nameCtx.GetProcedure().GetText()
	}

	log.Printf("read procedure name. Database: %s Schema: %s Name: %s", color.YellowString(procedure.Database), color.YellowString(procedure.Schema), color.YellowString(procedure.Name))

	p := &procParser{
		proc: &procedure,
	}

	antlr.ParseTreeWalkerDefault.Walk(p, ctx)
	s.procedures = append(s.procedures, procedure)
}

func (s *fileParser) EnterCreate_type(ctx *parser.Create_typeContext) {
	log.Println("Starting type")

	if ctx.GetData_type_name() != nil {
		// Not a table. TODO: Support this?
		return
	}

	var table Table

	table.Name = ctx.GetName().GetText()

	p := &columnsParser{}

	antlr.ParseTreeWalkerDefault.Walk(p, ctx)
	table.Columns = p.columns

	s.tables = append(s.tables, table)

}

func (s *fileParser) EnterCreate_table(ctx *parser.Create_tableContext) {
	if !s.genNormalTables {
		return
	}
	log.Println("Starting table")

	var table Table

	table.Name = ctx.GetName().GetText()

	p := &columnsParser{}

	antlr.ParseTreeWalkerDefault.Walk(p, ctx)
	table.Columns = p.columns

	s.normalTables = append(s.normalTables, table)
}

type columnsParser struct {
	*parser.BasetsqlListener
	columns []Column
}

func (s *columnsParser) EnterColumn_definition(ctx *parser.Column_definitionContext) {
	var col Column
	col.Name = ctx.GetName().GetText()
	if ctx.GetData_type_name() != nil {
		col.DataType = ctx.GetData_type_name().GetText()
	}
	log.Printf("parsed column %+v", col)
	s.columns = append(s.columns, col)
}

type procParser struct {
	*parser.LoggingTsqlListener
	proc *Procedure
	done bool
}

func (s *procParser) EnterProcedure_param(ctx *parser.Procedure_paramContext) {
	if s.done {
		return
	}
	log.Printf("starting param")

	param := ProcedureParam{
		DataType: ctx.GetData_type_name().GetText(),
		Name:     ctx.GetParam_name().GetText(),
	}

	if ctx.GetDefault_val() != nil {
		param.DefaultValue = ctx.GetDefault_val().GetText()
	}

	log.Printf("read param DataType: %s Name: %s Defaultvalue: %s", color.YellowString(param.DataType), color.YellowString(param.Name), color.YellowString(param.DefaultValue))

	s.proc.Params = append(s.proc.Params, param)
}

func (s *procParser) EnterSql_clauses(c *parser.Sql_clausesContext) {
	s.done = true
}

type ParseResult struct {
	Procedures   []Procedure
	Tables       []Table
	NormalTables []Table
}

func Parse(tsql string, genNormalTables bool) (*ParseResult, error) {
	is := antlr.NewInputStream(tsql)
	lexer := parser.NewtsqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewtsqlParser(stream)

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	fileP := &fileParser{
		genNormalTables: genNormalTables,
	}

	antlr.ParseTreeWalkerDefault.Walk(fileP, p.Tsql_file())

	return &ParseResult{
		Tables:       fileP.tables,
		NormalTables: fileP.normalTables,
		Procedures:   fileP.procedures,
	}, nil
}
