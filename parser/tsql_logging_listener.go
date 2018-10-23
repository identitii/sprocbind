package parser // tsql

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// LoggingTsqlListener is a complete listener for a parse tree produced by tsqlParser.
type LoggingTsqlListener struct{}

var _ tsqlListener = &LoggingTsqlListener{}

// VisitTerminal is called when a terminal node is visited
func (s *LoggingTsqlListener) VisitTerminal(node antlr.TerminalNode) {
	//log.Println("VisitTerminal")
}

// VisitErrorNode is called when an error node is visited.
func (s *LoggingTsqlListener) VisitErrorNode(node antlr.ErrorNode) {
	log.Println("VisitErrorNode")
}

// EnterEveryRule is called when any rule is entered.
func (s *LoggingTsqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//log.Println("EnterEveryRule")
}

// ExitEveryRule is called when any rule is exited.
func (s *LoggingTsqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//log.Println("ExitEveryRule")
}

// EnterTsql_file is called when production tsql_file is entered.
func (s *LoggingTsqlListener) EnterTsql_file(ctx *Tsql_fileContext) {
	log.Println("EnterTsql_file")
}

// ExitTsql_file is called when production tsql_file is exited.
func (s *LoggingTsqlListener) ExitTsql_file(ctx *Tsql_fileContext) {
	log.Println("ExitTsql_file")
}

// EnterBatch is called when production batch is entered.
func (s *LoggingTsqlListener) EnterBatch(ctx *BatchContext) {
	log.Println("EnterBatch")
}

// ExitBatch is called when production batch is exited.
func (s *LoggingTsqlListener) ExitBatch(ctx *BatchContext) {
	log.Println("ExitBatch")
}

// EnterSql_clauses is called when production sql_clauses is entered.
func (s *LoggingTsqlListener) EnterSql_clauses(ctx *Sql_clausesContext) {
	log.Println("EnterSql_clauses")
}

// ExitSql_clauses is called when production sql_clauses is exited.
func (s *LoggingTsqlListener) ExitSql_clauses(ctx *Sql_clausesContext) {
	log.Println("ExitSql_clauses")
}

// EnterSql_clause is called when production sql_clause is entered.
func (s *LoggingTsqlListener) EnterSql_clause(ctx *Sql_clauseContext) {
	log.Println("EnterSql_clause")
}

// ExitSql_clause is called when production sql_clause is exited.
func (s *LoggingTsqlListener) ExitSql_clause(ctx *Sql_clauseContext) {
	log.Println("ExitSql_clause")
}

// EnterDml_clause is called when production dml_clause is entered.
func (s *LoggingTsqlListener) EnterDml_clause(ctx *Dml_clauseContext) {
	log.Println("EnterDml_clause")
}

// ExitDml_clause is called when production dml_clause is exited.
func (s *LoggingTsqlListener) ExitDml_clause(ctx *Dml_clauseContext) {
	log.Println("ExitDml_clause")
}

// EnterDdl_clause is called when production ddl_clause is entered.
func (s *LoggingTsqlListener) EnterDdl_clause(ctx *Ddl_clauseContext) {
	log.Println("EnterDdl_clause")
}

// ExitDdl_clause is called when production ddl_clause is exited.
func (s *LoggingTsqlListener) ExitDdl_clause(ctx *Ddl_clauseContext) {
	log.Println("ExitDdl_clause")
}

// EnterBlock_statement is called when production block_statement is entered.
func (s *LoggingTsqlListener) EnterBlock_statement(ctx *Block_statementContext) {
	log.Println("EnterBlock_statement")
}

// ExitBlock_statement is called when production block_statement is exited.
func (s *LoggingTsqlListener) ExitBlock_statement(ctx *Block_statementContext) {
	log.Println("ExitBlock_statement")
}

// EnterBreak_statement is called when production break_statement is entered.
func (s *LoggingTsqlListener) EnterBreak_statement(ctx *Break_statementContext) {
	log.Println("EnterBreak_statement")
}

// ExitBreak_statement is called when production break_statement is exited.
func (s *LoggingTsqlListener) ExitBreak_statement(ctx *Break_statementContext) {
	log.Println("ExitBreak_statement")
}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *LoggingTsqlListener) EnterContinue_statement(ctx *Continue_statementContext) {
	log.Println("EnterContinue_statement")
}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *LoggingTsqlListener) ExitContinue_statement(ctx *Continue_statementContext) {
	log.Println("ExitContinue_statement")
}

// EnterGoto_statement is called when production goto_statement is entered.
func (s *LoggingTsqlListener) EnterGoto_statement(ctx *Goto_statementContext) {
	log.Println("EnterGoto_statement")
}

// ExitGoto_statement is called when production goto_statement is exited.
func (s *LoggingTsqlListener) ExitGoto_statement(ctx *Goto_statementContext) {
	log.Println("ExitGoto_statement")
}

// EnterIf_statement is called when production if_statement is entered.
func (s *LoggingTsqlListener) EnterIf_statement(ctx *If_statementContext) {
	log.Println("EnterIf_statement")
}

// ExitIf_statement is called when production if_statement is exited.
func (s *LoggingTsqlListener) ExitIf_statement(ctx *If_statementContext) {
	log.Println("ExitIf_statement")
}

// EnterReturn_statement is called when production return_statement is entered.
func (s *LoggingTsqlListener) EnterReturn_statement(ctx *Return_statementContext) {
	log.Println("EnterReturn_statement")
}

// ExitReturn_statement is called when production return_statement is exited.
func (s *LoggingTsqlListener) ExitReturn_statement(ctx *Return_statementContext) {
	log.Println("ExitReturn_statement")
}

// EnterThrow_statement is called when production throw_statement is entered.
func (s *LoggingTsqlListener) EnterThrow_statement(ctx *Throw_statementContext) {
	log.Println("EnterThrow_statement")
}

// ExitThrow_statement is called when production throw_statement is exited.
func (s *LoggingTsqlListener) ExitThrow_statement(ctx *Throw_statementContext) {
	log.Println("ExitThrow_statement")
}

// EnterTry_catch_statement is called when production try_catch_statement is entered.
func (s *LoggingTsqlListener) EnterTry_catch_statement(ctx *Try_catch_statementContext) {
	log.Println("EnterTry_catch_statement")
}

// ExitTry_catch_statement is called when production try_catch_statement is exited.
func (s *LoggingTsqlListener) ExitTry_catch_statement(ctx *Try_catch_statementContext) {
	log.Println("ExitTry_catch_statement")
}

// EnterWaitfor_statement is called when production waitfor_statement is entered.
func (s *LoggingTsqlListener) EnterWaitfor_statement(ctx *Waitfor_statementContext) {
	log.Println("EnterWaitfor_statement")
}

// ExitWaitfor_statement is called when production waitfor_statement is exited.
func (s *LoggingTsqlListener) ExitWaitfor_statement(ctx *Waitfor_statementContext) {
	log.Println("ExitWaitfor_statement")
}

// EnterWhile_statement is called when production while_statement is entered.
func (s *LoggingTsqlListener) EnterWhile_statement(ctx *While_statementContext) {
	log.Println("EnterWhile_statement")
}

// ExitWhile_statement is called when production while_statement is exited.
func (s *LoggingTsqlListener) ExitWhile_statement(ctx *While_statementContext) {
	log.Println("ExitWhile_statement")
}

// EnterPrint_statement is called when production print_statement is entered.
func (s *LoggingTsqlListener) EnterPrint_statement(ctx *Print_statementContext) {
	log.Println("EnterPrint_statement")
}

// ExitPrint_statement is called when production print_statement is exited.
func (s *LoggingTsqlListener) ExitPrint_statement(ctx *Print_statementContext) {
	log.Println("ExitPrint_statement")
}

// EnterRaiseerror_statement is called when production raiseerror_statement is entered.
func (s *LoggingTsqlListener) EnterRaiseerror_statement(ctx *Raiseerror_statementContext) {
	log.Println("EnterRaiseerror_statement")
}

// ExitRaiseerror_statement is called when production raiseerror_statement is exited.
func (s *LoggingTsqlListener) ExitRaiseerror_statement(ctx *Raiseerror_statementContext) {
	log.Println("ExitRaiseerror_statement")
}

// EnterAnother_statement is called when production another_statement is entered.
func (s *LoggingTsqlListener) EnterAnother_statement(ctx *Another_statementContext) {
	log.Println("EnterAnother_statement")
}

// ExitAnother_statement is called when production another_statement is exited.
func (s *LoggingTsqlListener) ExitAnother_statement(ctx *Another_statementContext) {
	log.Println("ExitAnother_statement")
}

// EnterDelete_statement is called when production delete_statement is entered.
func (s *LoggingTsqlListener) EnterDelete_statement(ctx *Delete_statementContext) {
	log.Println("EnterDelete_statement")
}

// ExitDelete_statement is called when production delete_statement is exited.
func (s *LoggingTsqlListener) ExitDelete_statement(ctx *Delete_statementContext) {
	log.Println("ExitDelete_statement")
}

// EnterDelete_statement_from is called when production delete_statement_from is entered.
func (s *LoggingTsqlListener) EnterDelete_statement_from(ctx *Delete_statement_fromContext) {
	log.Println("EnterDelete_statement_from")
}

// ExitDelete_statement_from is called when production delete_statement_from is exited.
func (s *LoggingTsqlListener) ExitDelete_statement_from(ctx *Delete_statement_fromContext) {
	log.Println("ExitDelete_statement_from")
}

// EnterInsert_statement is called when production insert_statement is entered.
func (s *LoggingTsqlListener) EnterInsert_statement(ctx *Insert_statementContext) {
	log.Println("EnterInsert_statement")
}

// ExitInsert_statement is called when production insert_statement is exited.
func (s *LoggingTsqlListener) ExitInsert_statement(ctx *Insert_statementContext) {
	log.Println("ExitInsert_statement")
}

// EnterInsert_statement_value is called when production insert_statement_value is entered.
func (s *LoggingTsqlListener) EnterInsert_statement_value(ctx *Insert_statement_valueContext) {
	log.Println("EnterInsert_statement_value")
}

// ExitInsert_statement_value is called when production insert_statement_value is exited.
func (s *LoggingTsqlListener) ExitInsert_statement_value(ctx *Insert_statement_valueContext) {
	log.Println("ExitInsert_statement_value")
}

// EnterSelect_statement is called when production select_statement is entered.
func (s *LoggingTsqlListener) EnterSelect_statement(ctx *Select_statementContext) {
	log.Println("EnterSelect_statement")
}

// ExitSelect_statement is called when production select_statement is exited.
func (s *LoggingTsqlListener) ExitSelect_statement(ctx *Select_statementContext) {
	log.Println("ExitSelect_statement")
}

// EnterUpdate_statement is called when production update_statement is entered.
func (s *LoggingTsqlListener) EnterUpdate_statement(ctx *Update_statementContext) {
	log.Println("EnterUpdate_statement")
}

// ExitUpdate_statement is called when production update_statement is exited.
func (s *LoggingTsqlListener) ExitUpdate_statement(ctx *Update_statementContext) {
	log.Println("ExitUpdate_statement")
}

// EnterWhere_clause_dml is called when production where_clause_dml is entered.
func (s *LoggingTsqlListener) EnterWhere_clause_dml(ctx *Where_clause_dmlContext) {
	log.Println("EnterWhere_clause_dml")
}

// ExitWhere_clause_dml is called when production where_clause_dml is exited.
func (s *LoggingTsqlListener) ExitWhere_clause_dml(ctx *Where_clause_dmlContext) {
	log.Println("ExitWhere_clause_dml")
}

// EnterOutput_clause is called when production output_clause is entered.
func (s *LoggingTsqlListener) EnterOutput_clause(ctx *Output_clauseContext) {
	log.Println("EnterOutput_clause")
}

// ExitOutput_clause is called when production output_clause is exited.
func (s *LoggingTsqlListener) ExitOutput_clause(ctx *Output_clauseContext) {
	log.Println("ExitOutput_clause")
}

// EnterOutput_dml_list_elem is called when production output_dml_list_elem is entered.
func (s *LoggingTsqlListener) EnterOutput_dml_list_elem(ctx *Output_dml_list_elemContext) {
	log.Println("EnterOutput_dml_list_elem")
}

// ExitOutput_dml_list_elem is called when production output_dml_list_elem is exited.
func (s *LoggingTsqlListener) ExitOutput_dml_list_elem(ctx *Output_dml_list_elemContext) {
	log.Println("ExitOutput_dml_list_elem")
}

// EnterOutput_column_name is called when production output_column_name is entered.
func (s *LoggingTsqlListener) EnterOutput_column_name(ctx *Output_column_nameContext) {
	log.Println("EnterOutput_column_name")
}

// ExitOutput_column_name is called when production output_column_name is exited.
func (s *LoggingTsqlListener) ExitOutput_column_name(ctx *Output_column_nameContext) {
	log.Println("ExitOutput_column_name")
}

// EnterCreate_database is called when production create_database is entered.
func (s *LoggingTsqlListener) EnterCreate_database(ctx *Create_databaseContext) {
	log.Println("EnterCreate_database")
}

// ExitCreate_database is called when production create_database is exited.
func (s *LoggingTsqlListener) ExitCreate_database(ctx *Create_databaseContext) {
	log.Println("ExitCreate_database")
}

// EnterCreate_index is called when production create_index is entered.
func (s *LoggingTsqlListener) EnterCreate_index(ctx *Create_indexContext) {
	log.Println("EnterCreate_index")
}

// ExitCreate_index is called when production create_index is exited.
func (s *LoggingTsqlListener) ExitCreate_index(ctx *Create_indexContext) {
	log.Println("ExitCreate_index")
}

// EnterCreate_procedure is called when production create_procedure is entered.
func (s *LoggingTsqlListener) EnterCreate_procedure(ctx *Create_procedureContext) {
	log.Println("EnterCreate_procedure")
}

// ExitCreate_procedure is called when production create_procedure is exited.
func (s *LoggingTsqlListener) ExitCreate_procedure(ctx *Create_procedureContext) {
	log.Println("ExitCreate_procedure")
}

// EnterProcedure_param is called when production procedure_param is entered.
func (s *LoggingTsqlListener) EnterProcedure_param(ctx *Procedure_paramContext) {
	log.Println("EnterProcedure_param")
}

// ExitProcedure_param is called when production procedure_param is exited.
func (s *LoggingTsqlListener) ExitProcedure_param(ctx *Procedure_paramContext) {
	log.Println("ExitProcedure_param")
}

// EnterProcedure_option is called when production procedure_option is entered.
func (s *LoggingTsqlListener) EnterProcedure_option(ctx *Procedure_optionContext) {
	log.Println("EnterProcedure_option")
}

// ExitProcedure_option is called when production procedure_option is exited.
func (s *LoggingTsqlListener) ExitProcedure_option(ctx *Procedure_optionContext) {
	log.Println("ExitProcedure_option")
}

// EnterCreate_statistics is called when production create_statistics is entered.
func (s *LoggingTsqlListener) EnterCreate_statistics(ctx *Create_statisticsContext) {
	log.Println("EnterCreate_statistics")
}

// ExitCreate_statistics is called when production create_statistics is exited.
func (s *LoggingTsqlListener) ExitCreate_statistics(ctx *Create_statisticsContext) {
	log.Println("ExitCreate_statistics")
}

// EnterCreate_table is called when production create_table is entered.
func (s *LoggingTsqlListener) EnterCreate_table(ctx *Create_tableContext) {
	log.Println("EnterCreate_table")
}

// ExitCreate_table is called when production create_table is exited.
func (s *LoggingTsqlListener) ExitCreate_table(ctx *Create_tableContext) {
	log.Println("ExitCreate_table")
}

// EnterCreate_view is called when production create_view is entered.
func (s *LoggingTsqlListener) EnterCreate_view(ctx *Create_viewContext) {
	log.Println("EnterCreate_view")
}

// ExitCreate_view is called when production create_view is exited.
func (s *LoggingTsqlListener) ExitCreate_view(ctx *Create_viewContext) {
	log.Println("ExitCreate_view")
}

// EnterView_attribute is called when production view_attribute is entered.
func (s *LoggingTsqlListener) EnterView_attribute(ctx *View_attributeContext) {
	log.Println("EnterView_attribute")
}

// ExitView_attribute is called when production view_attribute is exited.
func (s *LoggingTsqlListener) ExitView_attribute(ctx *View_attributeContext) {
	log.Println("ExitView_attribute")
}

// EnterAlter_table is called when production alter_table is entered.
func (s *LoggingTsqlListener) EnterAlter_table(ctx *Alter_tableContext) {
	log.Println("EnterAlter_table")
}

// ExitAlter_table is called when production alter_table is exited.
func (s *LoggingTsqlListener) ExitAlter_table(ctx *Alter_tableContext) {
	log.Println("ExitAlter_table")
}

// EnterAlter_database is called when production alter_database is entered.
func (s *LoggingTsqlListener) EnterAlter_database(ctx *Alter_databaseContext) {
	log.Println("EnterAlter_database")
}

// ExitAlter_database is called when production alter_database is exited.
func (s *LoggingTsqlListener) ExitAlter_database(ctx *Alter_databaseContext) {
	log.Println("ExitAlter_database")
}

// EnterDatabase_optionspec is called when production database_optionspec is entered.
func (s *LoggingTsqlListener) EnterDatabase_optionspec(ctx *Database_optionspecContext) {
	log.Println("EnterDatabase_optionspec")
}

// ExitDatabase_optionspec is called when production database_optionspec is exited.
func (s *LoggingTsqlListener) ExitDatabase_optionspec(ctx *Database_optionspecContext) {
	log.Println("ExitDatabase_optionspec")
}

// EnterAuto_option is called when production auto_option is entered.
func (s *LoggingTsqlListener) EnterAuto_option(ctx *Auto_optionContext) {
	log.Println("EnterAuto_option")
}

// ExitAuto_option is called when production auto_option is exited.
func (s *LoggingTsqlListener) ExitAuto_option(ctx *Auto_optionContext) {
	log.Println("ExitAuto_option")
}

// EnterChange_tracking_option is called when production change_tracking_option is entered.
func (s *LoggingTsqlListener) EnterChange_tracking_option(ctx *Change_tracking_optionContext) {
	log.Println("EnterChange_tracking_option")
}

// ExitChange_tracking_option is called when production change_tracking_option is exited.
func (s *LoggingTsqlListener) ExitChange_tracking_option(ctx *Change_tracking_optionContext) {
	log.Println("ExitChange_tracking_option")
}

// EnterChange_tracking_option_list is called when production change_tracking_option_list is entered.
func (s *LoggingTsqlListener) EnterChange_tracking_option_list(ctx *Change_tracking_option_listContext) {
	log.Println("EnterChange_tracking_option_list")
}

// ExitChange_tracking_option_list is called when production change_tracking_option_list is exited.
func (s *LoggingTsqlListener) ExitChange_tracking_option_list(ctx *Change_tracking_option_listContext) {
	log.Println("ExitChange_tracking_option_list")
}

// EnterContainment_option is called when production containment_option is entered.
func (s *LoggingTsqlListener) EnterContainment_option(ctx *Containment_optionContext) {
	log.Println("EnterContainment_option")
}

// ExitContainment_option is called when production containment_option is exited.
func (s *LoggingTsqlListener) ExitContainment_option(ctx *Containment_optionContext) {
	log.Println("ExitContainment_option")
}

// EnterCursor_option is called when production cursor_option is entered.
func (s *LoggingTsqlListener) EnterCursor_option(ctx *Cursor_optionContext) {
	log.Println("EnterCursor_option")
}

// ExitCursor_option is called when production cursor_option is exited.
func (s *LoggingTsqlListener) ExitCursor_option(ctx *Cursor_optionContext) {
	log.Println("ExitCursor_option")
}

// EnterDate_correlation_optimization_option is called when production date_correlation_optimization_option is entered.
func (s *LoggingTsqlListener) EnterDate_correlation_optimization_option(ctx *Date_correlation_optimization_optionContext) {
	log.Println("EnterDate_correlation_optimization_option")
}

// ExitDate_correlation_optimization_option is called when production date_correlation_optimization_option is exited.
func (s *LoggingTsqlListener) ExitDate_correlation_optimization_option(ctx *Date_correlation_optimization_optionContext) {
	log.Println("ExitDate_correlation_optimization_option")
}

// EnterDb_encryption_option is called when production db_encryption_option is entered.
func (s *LoggingTsqlListener) EnterDb_encryption_option(ctx *Db_encryption_optionContext) {
	log.Println("EnterDb_encryption_option")
}

// ExitDb_encryption_option is called when production db_encryption_option is exited.
func (s *LoggingTsqlListener) ExitDb_encryption_option(ctx *Db_encryption_optionContext) {
	log.Println("ExitDb_encryption_option")
}

// EnterDb_state_option is called when production db_state_option is entered.
func (s *LoggingTsqlListener) EnterDb_state_option(ctx *Db_state_optionContext) {
	log.Println("EnterDb_state_option")
}

// ExitDb_state_option is called when production db_state_option is exited.
func (s *LoggingTsqlListener) ExitDb_state_option(ctx *Db_state_optionContext) {
	log.Println("ExitDb_state_option")
}

// EnterDb_update_option is called when production db_update_option is entered.
func (s *LoggingTsqlListener) EnterDb_update_option(ctx *Db_update_optionContext) {
	log.Println("EnterDb_update_option")
}

// ExitDb_update_option is called when production db_update_option is exited.
func (s *LoggingTsqlListener) ExitDb_update_option(ctx *Db_update_optionContext) {
	log.Println("ExitDb_update_option")
}

// EnterDb_user_access_option is called when production db_user_access_option is entered.
func (s *LoggingTsqlListener) EnterDb_user_access_option(ctx *Db_user_access_optionContext) {
	log.Println("EnterDb_user_access_option")
}

// ExitDb_user_access_option is called when production db_user_access_option is exited.
func (s *LoggingTsqlListener) ExitDb_user_access_option(ctx *Db_user_access_optionContext) {
	log.Println("ExitDb_user_access_option")
}

// EnterDelayed_durability_option is called when production delayed_durability_option is entered.
func (s *LoggingTsqlListener) EnterDelayed_durability_option(ctx *Delayed_durability_optionContext) {
	log.Println("EnterDelayed_durability_option")
}

// ExitDelayed_durability_option is called when production delayed_durability_option is exited.
func (s *LoggingTsqlListener) ExitDelayed_durability_option(ctx *Delayed_durability_optionContext) {
	log.Println("ExitDelayed_durability_option")
}

// EnterExternal_access_option is called when production external_access_option is entered.
func (s *LoggingTsqlListener) EnterExternal_access_option(ctx *External_access_optionContext) {
	log.Println("EnterExternal_access_option")
}

// ExitExternal_access_option is called when production external_access_option is exited.
func (s *LoggingTsqlListener) ExitExternal_access_option(ctx *External_access_optionContext) {
	log.Println("ExitExternal_access_option")
}

// EnterMixed_page_allocation_option is called when production mixed_page_allocation_option is entered.
func (s *LoggingTsqlListener) EnterMixed_page_allocation_option(ctx *Mixed_page_allocation_optionContext) {
	log.Println("EnterMixed_page_allocation_option")
}

// ExitMixed_page_allocation_option is called when production mixed_page_allocation_option is exited.
func (s *LoggingTsqlListener) ExitMixed_page_allocation_option(ctx *Mixed_page_allocation_optionContext) {
	log.Println("ExitMixed_page_allocation_option")
}

// EnterParameterization_option is called when production parameterization_option is entered.
func (s *LoggingTsqlListener) EnterParameterization_option(ctx *Parameterization_optionContext) {
	log.Println("EnterParameterization_option")
}

// ExitParameterization_option is called when production parameterization_option is exited.
func (s *LoggingTsqlListener) ExitParameterization_option(ctx *Parameterization_optionContext) {
	log.Println("ExitParameterization_option")
}

// EnterRecovery_option is called when production recovery_option is entered.
func (s *LoggingTsqlListener) EnterRecovery_option(ctx *Recovery_optionContext) {
	log.Println("EnterRecovery_option")
}

// ExitRecovery_option is called when production recovery_option is exited.
func (s *LoggingTsqlListener) ExitRecovery_option(ctx *Recovery_optionContext) {
	log.Println("ExitRecovery_option")
}

// EnterService_broker_option is called when production service_broker_option is entered.
func (s *LoggingTsqlListener) EnterService_broker_option(ctx *Service_broker_optionContext) {
	log.Println("EnterService_broker_option")
}

// ExitService_broker_option is called when production service_broker_option is exited.
func (s *LoggingTsqlListener) ExitService_broker_option(ctx *Service_broker_optionContext) {
	log.Println("ExitService_broker_option")
}

// EnterSnapshot_option is called when production snapshot_option is entered.
func (s *LoggingTsqlListener) EnterSnapshot_option(ctx *Snapshot_optionContext) {
	log.Println("EnterSnapshot_option")
}

// ExitSnapshot_option is called when production snapshot_option is exited.
func (s *LoggingTsqlListener) ExitSnapshot_option(ctx *Snapshot_optionContext) {
	log.Println("ExitSnapshot_option")
}

// EnterSql_option is called when production sql_option is entered.
func (s *LoggingTsqlListener) EnterSql_option(ctx *Sql_optionContext) {
	log.Println("EnterSql_option")
}

// ExitSql_option is called when production sql_option is exited.
func (s *LoggingTsqlListener) ExitSql_option(ctx *Sql_optionContext) {
	log.Println("ExitSql_option")
}

// EnterTarget_recovery_time_option is called when production target_recovery_time_option is entered.
func (s *LoggingTsqlListener) EnterTarget_recovery_time_option(ctx *Target_recovery_time_optionContext) {
	log.Println("EnterTarget_recovery_time_option")
}

// ExitTarget_recovery_time_option is called when production target_recovery_time_option is exited.
func (s *LoggingTsqlListener) ExitTarget_recovery_time_option(ctx *Target_recovery_time_optionContext) {
	log.Println("ExitTarget_recovery_time_option")
}

// EnterTermination is called when production termination is entered.
func (s *LoggingTsqlListener) EnterTermination(ctx *TerminationContext) {
	log.Println("EnterTermination")
}

// ExitTermination is called when production termination is exited.
func (s *LoggingTsqlListener) ExitTermination(ctx *TerminationContext) {
	log.Println("ExitTermination")
}

// EnterDrop_index is called when production drop_index is entered.
func (s *LoggingTsqlListener) EnterDrop_index(ctx *Drop_indexContext) {
	log.Println("EnterDrop_index")
}

// ExitDrop_index is called when production drop_index is exited.
func (s *LoggingTsqlListener) ExitDrop_index(ctx *Drop_indexContext) {
	log.Println("ExitDrop_index")
}

// EnterDrop_procedure is called when production drop_procedure is entered.
func (s *LoggingTsqlListener) EnterDrop_procedure(ctx *Drop_procedureContext) {
	log.Println("EnterDrop_procedure")
}

// ExitDrop_procedure is called when production drop_procedure is exited.
func (s *LoggingTsqlListener) ExitDrop_procedure(ctx *Drop_procedureContext) {
	log.Println("ExitDrop_procedure")
}

// EnterDrop_statistics is called when production drop_statistics is entered.
func (s *LoggingTsqlListener) EnterDrop_statistics(ctx *Drop_statisticsContext) {
	log.Println("EnterDrop_statistics")
}

// ExitDrop_statistics is called when production drop_statistics is exited.
func (s *LoggingTsqlListener) ExitDrop_statistics(ctx *Drop_statisticsContext) {
	log.Println("ExitDrop_statistics")
}

// EnterDrop_table is called when production drop_table is entered.
func (s *LoggingTsqlListener) EnterDrop_table(ctx *Drop_tableContext) {
	log.Println("EnterDrop_table")
}

// ExitDrop_table is called when production drop_table is exited.
func (s *LoggingTsqlListener) ExitDrop_table(ctx *Drop_tableContext) {
	log.Println("ExitDrop_table")
}

// EnterDrop_view is called when production drop_view is entered.
func (s *LoggingTsqlListener) EnterDrop_view(ctx *Drop_viewContext) {
	log.Println("EnterDrop_view")
}

// ExitDrop_view is called when production drop_view is exited.
func (s *LoggingTsqlListener) ExitDrop_view(ctx *Drop_viewContext) {
	log.Println("ExitDrop_view")
}

// EnterCreate_type is called when production create_type is entered.
func (s *LoggingTsqlListener) EnterCreate_type(ctx *Create_typeContext) {
	log.Println("EnterCreate_type")
}

// ExitCreate_type is called when production create_type is exited.
func (s *LoggingTsqlListener) ExitCreate_type(ctx *Create_typeContext) {
	log.Println("ExitCreate_type")
}

// EnterDrop_type is called when production drop_type is entered.
func (s *LoggingTsqlListener) EnterDrop_type(ctx *Drop_typeContext) {
	log.Println("EnterDrop_type")
}

// ExitDrop_type is called when production drop_type is exited.
func (s *LoggingTsqlListener) ExitDrop_type(ctx *Drop_typeContext) {
	log.Println("ExitDrop_type")
}

// EnterRowset_function_limited is called when production rowset_function_limited is entered.
func (s *LoggingTsqlListener) EnterRowset_function_limited(ctx *Rowset_function_limitedContext) {
	log.Println("EnterRowset_function_limited")
}

// ExitRowset_function_limited is called when production rowset_function_limited is exited.
func (s *LoggingTsqlListener) ExitRowset_function_limited(ctx *Rowset_function_limitedContext) {
	log.Println("ExitRowset_function_limited")
}

// EnterOpenquery is called when production openquery is entered.
func (s *LoggingTsqlListener) EnterOpenquery(ctx *OpenqueryContext) {
	log.Println("EnterOpenquery")
}

// ExitOpenquery is called when production openquery is exited.
func (s *LoggingTsqlListener) ExitOpenquery(ctx *OpenqueryContext) {
	log.Println("ExitOpenquery")
}

// EnterOpendatasource is called when production opendatasource is entered.
func (s *LoggingTsqlListener) EnterOpendatasource(ctx *OpendatasourceContext) {
	log.Println("EnterOpendatasource")
}

// ExitOpendatasource is called when production opendatasource is exited.
func (s *LoggingTsqlListener) ExitOpendatasource(ctx *OpendatasourceContext) {
	log.Println("ExitOpendatasource")
}

// EnterDeclare_statement is called when production declare_statement is entered.
func (s *LoggingTsqlListener) EnterDeclare_statement(ctx *Declare_statementContext) {
	log.Println("EnterDeclare_statement")
}

// ExitDeclare_statement is called when production declare_statement is exited.
func (s *LoggingTsqlListener) ExitDeclare_statement(ctx *Declare_statementContext) {
	log.Println("ExitDeclare_statement")
}

// EnterCursor_statement is called when production cursor_statement is entered.
func (s *LoggingTsqlListener) EnterCursor_statement(ctx *Cursor_statementContext) {
	log.Println("EnterCursor_statement")
}

// ExitCursor_statement is called when production cursor_statement is exited.
func (s *LoggingTsqlListener) ExitCursor_statement(ctx *Cursor_statementContext) {
	log.Println("ExitCursor_statement")
}

// EnterExecute_statement is called when production execute_statement is entered.
func (s *LoggingTsqlListener) EnterExecute_statement(ctx *Execute_statementContext) {
	log.Println("EnterExecute_statement")
}

// ExitExecute_statement is called when production execute_statement is exited.
func (s *LoggingTsqlListener) ExitExecute_statement(ctx *Execute_statementContext) {
	log.Println("ExitExecute_statement")
}

// EnterExecute_statement_arg is called when production execute_statement_arg is entered.
func (s *LoggingTsqlListener) EnterExecute_statement_arg(ctx *Execute_statement_argContext) {
	log.Println("EnterExecute_statement_arg")
}

// ExitExecute_statement_arg is called when production execute_statement_arg is exited.
func (s *LoggingTsqlListener) ExitExecute_statement_arg(ctx *Execute_statement_argContext) {
	log.Println("ExitExecute_statement_arg")
}

// EnterExecute_var_string is called when production execute_var_string is entered.
func (s *LoggingTsqlListener) EnterExecute_var_string(ctx *Execute_var_stringContext) {
	log.Println("EnterExecute_var_string")
}

// ExitExecute_var_string is called when production execute_var_string is exited.
func (s *LoggingTsqlListener) ExitExecute_var_string(ctx *Execute_var_stringContext) {
	log.Println("ExitExecute_var_string")
}

// EnterSecurity_statement is called when production security_statement is entered.
func (s *LoggingTsqlListener) EnterSecurity_statement(ctx *Security_statementContext) {
	log.Println("EnterSecurity_statement")
}

// ExitSecurity_statement is called when production security_statement is exited.
func (s *LoggingTsqlListener) ExitSecurity_statement(ctx *Security_statementContext) {
	log.Println("ExitSecurity_statement")
}

// EnterGrant_permission is called when production grant_permission is entered.
func (s *LoggingTsqlListener) EnterGrant_permission(ctx *Grant_permissionContext) {
	log.Println("EnterGrant_permission")
}

// ExitGrant_permission is called when production grant_permission is exited.
func (s *LoggingTsqlListener) ExitGrant_permission(ctx *Grant_permissionContext) {
	log.Println("ExitGrant_permission")
}

// EnterSet_statement is called when production set_statement is entered.
func (s *LoggingTsqlListener) EnterSet_statement(ctx *Set_statementContext) {
	log.Println("EnterSet_statement")
}

// ExitSet_statement is called when production set_statement is exited.
func (s *LoggingTsqlListener) ExitSet_statement(ctx *Set_statementContext) {
	log.Println("ExitSet_statement")
}

// EnterTransaction_statement is called when production transaction_statement is entered.
func (s *LoggingTsqlListener) EnterTransaction_statement(ctx *Transaction_statementContext) {
	log.Println("EnterTransaction_statement")
}

// ExitTransaction_statement is called when production transaction_statement is exited.
func (s *LoggingTsqlListener) ExitTransaction_statement(ctx *Transaction_statementContext) {
	log.Println("ExitTransaction_statement")
}

// EnterGo_statement is called when production go_statement is entered.
func (s *LoggingTsqlListener) EnterGo_statement(ctx *Go_statementContext) {
	log.Println("EnterGo_statement")
}

// ExitGo_statement is called when production go_statement is exited.
func (s *LoggingTsqlListener) ExitGo_statement(ctx *Go_statementContext) {
	log.Println("ExitGo_statement")
}

// EnterUse_statement is called when production use_statement is entered.
func (s *LoggingTsqlListener) EnterUse_statement(ctx *Use_statementContext) {
	log.Println("EnterUse_statement")
}

// ExitUse_statement is called when production use_statement is exited.
func (s *LoggingTsqlListener) ExitUse_statement(ctx *Use_statementContext) {
	log.Println("ExitUse_statement")
}

// EnterExecute_clause is called when production execute_clause is entered.
func (s *LoggingTsqlListener) EnterExecute_clause(ctx *Execute_clauseContext) {
	log.Println("EnterExecute_clause")
}

// ExitExecute_clause is called when production execute_clause is exited.
func (s *LoggingTsqlListener) ExitExecute_clause(ctx *Execute_clauseContext) {
	log.Println("ExitExecute_clause")
}

// EnterDeclare_local is called when production declare_local is entered.
func (s *LoggingTsqlListener) EnterDeclare_local(ctx *Declare_localContext) {
	log.Println("EnterDeclare_local")
}

// ExitDeclare_local is called when production declare_local is exited.
func (s *LoggingTsqlListener) ExitDeclare_local(ctx *Declare_localContext) {
	log.Println("ExitDeclare_local")
}

// EnterTable_type_definition is called when production table_type_definition is entered.
func (s *LoggingTsqlListener) EnterTable_type_definition(ctx *Table_type_definitionContext) {
	log.Println("EnterTable_type_definition")
}

// ExitTable_type_definition is called when production table_type_definition is exited.
func (s *LoggingTsqlListener) ExitTable_type_definition(ctx *Table_type_definitionContext) {
	log.Println("ExitTable_type_definition")
}

// EnterColumn_def_table_constraints is called when production column_def_table_constraints is entered.
func (s *LoggingTsqlListener) EnterColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) {
	log.Println("EnterColumn_def_table_constraints")
}

// ExitColumn_def_table_constraints is called when production column_def_table_constraints is exited.
func (s *LoggingTsqlListener) ExitColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) {
	log.Println("ExitColumn_def_table_constraints")
}

// EnterColumn_def_table_constraint is called when production column_def_table_constraint is entered.
func (s *LoggingTsqlListener) EnterColumn_def_table_constraint(ctx *Column_def_table_constraintContext) {
	log.Println("EnterColumn_def_table_constraint")
}

// ExitColumn_def_table_constraint is called when production column_def_table_constraint is exited.
func (s *LoggingTsqlListener) ExitColumn_def_table_constraint(ctx *Column_def_table_constraintContext) {
	log.Println("ExitColumn_def_table_constraint")
}

// EnterColumn_definition is called when production column_definition is entered.
func (s *LoggingTsqlListener) EnterColumn_definition(ctx *Column_definitionContext) {
	log.Println("EnterColumn_definition")
}

// ExitColumn_definition is called when production column_definition is exited.
func (s *LoggingTsqlListener) ExitColumn_definition(ctx *Column_definitionContext) {
	log.Println("ExitColumn_definition")
}

// EnterColumn_constraint is called when production column_constraint is entered.
func (s *LoggingTsqlListener) EnterColumn_constraint(ctx *Column_constraintContext) {
	log.Println("EnterColumn_constraint")
}

// ExitColumn_constraint is called when production column_constraint is exited.
func (s *LoggingTsqlListener) ExitColumn_constraint(ctx *Column_constraintContext) {
	log.Println("ExitColumn_constraint")
}

// EnterTable_constraint is called when production table_constraint is entered.
func (s *LoggingTsqlListener) EnterTable_constraint(ctx *Table_constraintContext) {
	log.Println("EnterTable_constraint")
}

// ExitTable_constraint is called when production table_constraint is exited.
func (s *LoggingTsqlListener) ExitTable_constraint(ctx *Table_constraintContext) {
	log.Println("ExitTable_constraint")
}

// EnterIndex_options is called when production index_options is entered.
func (s *LoggingTsqlListener) EnterIndex_options(ctx *Index_optionsContext) {
	log.Println("EnterIndex_options")
}

// ExitIndex_options is called when production index_options is exited.
func (s *LoggingTsqlListener) ExitIndex_options(ctx *Index_optionsContext) {
	log.Println("ExitIndex_options")
}

// EnterIndex_option is called when production index_option is entered.
func (s *LoggingTsqlListener) EnterIndex_option(ctx *Index_optionContext) {
	log.Println("EnterIndex_option")
}

// ExitIndex_option is called when production index_option is exited.
func (s *LoggingTsqlListener) ExitIndex_option(ctx *Index_optionContext) {
	log.Println("ExitIndex_option")
}

// EnterDeclare_cursor is called when production declare_cursor is entered.
func (s *LoggingTsqlListener) EnterDeclare_cursor(ctx *Declare_cursorContext) {
	log.Println("EnterDeclare_cursor")
}

// ExitDeclare_cursor is called when production declare_cursor is exited.
func (s *LoggingTsqlListener) ExitDeclare_cursor(ctx *Declare_cursorContext) {
	log.Println("ExitDeclare_cursor")
}

// EnterDeclare_set_cursor_common is called when production declare_set_cursor_common is entered.
func (s *LoggingTsqlListener) EnterDeclare_set_cursor_common(ctx *Declare_set_cursor_commonContext) {
	log.Println("EnterDeclare_set_cursor_common")
}

// ExitDeclare_set_cursor_common is called when production declare_set_cursor_common is exited.
func (s *LoggingTsqlListener) ExitDeclare_set_cursor_common(ctx *Declare_set_cursor_commonContext) {
	log.Println("ExitDeclare_set_cursor_common")
}

// EnterFetch_cursor is called when production fetch_cursor is entered.
func (s *LoggingTsqlListener) EnterFetch_cursor(ctx *Fetch_cursorContext) {
	log.Println("EnterFetch_cursor")
}

// ExitFetch_cursor is called when production fetch_cursor is exited.
func (s *LoggingTsqlListener) ExitFetch_cursor(ctx *Fetch_cursorContext) {
	log.Println("ExitFetch_cursor")
}

// EnterSet_special is called when production set_special is entered.
func (s *LoggingTsqlListener) EnterSet_special(ctx *Set_specialContext) {
	log.Println("EnterSet_special")
}

// ExitSet_special is called when production set_special is exited.
func (s *LoggingTsqlListener) ExitSet_special(ctx *Set_specialContext) {
	log.Println("ExitSet_special")
}

// EnterConstant_LOCAL_ID is called when production constant_LOCAL_ID is entered.
func (s *LoggingTsqlListener) EnterConstant_LOCAL_ID(ctx *Constant_LOCAL_IDContext) {
	log.Println("EnterConstant_LOCAL_ID")
}

// ExitConstant_LOCAL_ID is called when production constant_LOCAL_ID is exited.
func (s *LoggingTsqlListener) ExitConstant_LOCAL_ID(ctx *Constant_LOCAL_IDContext) {
	log.Println("ExitConstant_LOCAL_ID")
}

// EnterBinary_operator_expression is called when production binary_operator_expression is entered.
func (s *LoggingTsqlListener) EnterBinary_operator_expression(ctx *Binary_operator_expressionContext) {
	log.Println("EnterBinary_operator_expression")
}

// ExitBinary_operator_expression is called when production binary_operator_expression is exited.
func (s *LoggingTsqlListener) ExitBinary_operator_expression(ctx *Binary_operator_expressionContext) {
	log.Println("ExitBinary_operator_expression")
}

// EnterPrimitive_expression is called when production primitive_expression is entered.
func (s *LoggingTsqlListener) EnterPrimitive_expression(ctx *Primitive_expressionContext) {
	log.Println("EnterPrimitive_expression")
}

// ExitPrimitive_expression is called when production primitive_expression is exited.
func (s *LoggingTsqlListener) ExitPrimitive_expression(ctx *Primitive_expressionContext) {
	log.Println("ExitPrimitive_expression")
}

// EnterBracket_expression is called when production bracket_expression is entered.
func (s *LoggingTsqlListener) EnterBracket_expression(ctx *Bracket_expressionContext) {
	log.Println("EnterBracket_expression")
}

// ExitBracket_expression is called when production bracket_expression is exited.
func (s *LoggingTsqlListener) ExitBracket_expression(ctx *Bracket_expressionContext) {
	log.Println("ExitBracket_expression")
}

// EnterUnary_operator_expression is called when production unary_operator_expression is entered.
func (s *LoggingTsqlListener) EnterUnary_operator_expression(ctx *Unary_operator_expressionContext) {
	log.Println("EnterUnary_operator_expression")
}

// ExitUnary_operator_expression is called when production unary_operator_expression is exited.
func (s *LoggingTsqlListener) ExitUnary_operator_expression(ctx *Unary_operator_expressionContext) {
	log.Println("ExitUnary_operator_expression")
}

// EnterFunction_call_expression is called when production function_call_expression is entered.
func (s *LoggingTsqlListener) EnterFunction_call_expression(ctx *Function_call_expressionContext) {
	log.Println("EnterFunction_call_expression")
}

// ExitFunction_call_expression is called when production function_call_expression is exited.
func (s *LoggingTsqlListener) ExitFunction_call_expression(ctx *Function_call_expressionContext) {
	log.Println("ExitFunction_call_expression")
}

// EnterCase_expression is called when production case_expression is entered.
func (s *LoggingTsqlListener) EnterCase_expression(ctx *Case_expressionContext) {
	log.Println("EnterCase_expression")
}

// ExitCase_expression is called when production case_expression is exited.
func (s *LoggingTsqlListener) ExitCase_expression(ctx *Case_expressionContext) {
	log.Println("ExitCase_expression")
}

// EnterColumn_ref_expression is called when production column_ref_expression is entered.
func (s *LoggingTsqlListener) EnterColumn_ref_expression(ctx *Column_ref_expressionContext) {
	log.Println("EnterColumn_ref_expression")
}

// ExitColumn_ref_expression is called when production column_ref_expression is exited.
func (s *LoggingTsqlListener) ExitColumn_ref_expression(ctx *Column_ref_expressionContext) {
	log.Println("ExitColumn_ref_expression")
}

// EnterSubquery_expression is called when production subquery_expression is entered.
func (s *LoggingTsqlListener) EnterSubquery_expression(ctx *Subquery_expressionContext) {
	log.Println("EnterSubquery_expression")
}

// ExitSubquery_expression is called when production subquery_expression is exited.
func (s *LoggingTsqlListener) ExitSubquery_expression(ctx *Subquery_expressionContext) {
	log.Println("ExitSubquery_expression")
}

// EnterOver_clause_expression is called when production over_clause_expression is entered.
func (s *LoggingTsqlListener) EnterOver_clause_expression(ctx *Over_clause_expressionContext) {
	log.Println("EnterOver_clause_expression")
}

// ExitOver_clause_expression is called when production over_clause_expression is exited.
func (s *LoggingTsqlListener) ExitOver_clause_expression(ctx *Over_clause_expressionContext) {
	log.Println("ExitOver_clause_expression")
}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *LoggingTsqlListener) EnterConstant_expression(ctx *Constant_expressionContext) {
	log.Println("EnterConstant_expression")
}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *LoggingTsqlListener) ExitConstant_expression(ctx *Constant_expressionContext) {
	log.Println("ExitConstant_expression")
}

// EnterSubquery is called when production subquery is entered.
func (s *LoggingTsqlListener) EnterSubquery(ctx *SubqueryContext) {
	log.Println("EnterSubquery")
}

// ExitSubquery is called when production subquery is exited.
func (s *LoggingTsqlListener) ExitSubquery(ctx *SubqueryContext) {
	log.Println("ExitSubquery")
}

// EnterWith_expression is called when production with_expression is entered.
func (s *LoggingTsqlListener) EnterWith_expression(ctx *With_expressionContext) {
	log.Println("EnterWith_expression")
}

// ExitWith_expression is called when production with_expression is exited.
func (s *LoggingTsqlListener) ExitWith_expression(ctx *With_expressionContext) {
	log.Println("ExitWith_expression")
}

// EnterCommon_table_expression is called when production common_table_expression is entered.
func (s *LoggingTsqlListener) EnterCommon_table_expression(ctx *Common_table_expressionContext) {
	log.Println("EnterCommon_table_expression")
}

// ExitCommon_table_expression is called when production common_table_expression is exited.
func (s *LoggingTsqlListener) ExitCommon_table_expression(ctx *Common_table_expressionContext) {
	log.Println("ExitCommon_table_expression")
}

// EnterUpdate_elem is called when production update_elem is entered.
func (s *LoggingTsqlListener) EnterUpdate_elem(ctx *Update_elemContext) {
	log.Println("EnterUpdate_elem")
}

// ExitUpdate_elem is called when production update_elem is exited.
func (s *LoggingTsqlListener) ExitUpdate_elem(ctx *Update_elemContext) {
	log.Println("ExitUpdate_elem")
}

// EnterSearch_condition_list is called when production search_condition_list is entered.
func (s *LoggingTsqlListener) EnterSearch_condition_list(ctx *Search_condition_listContext) {
	log.Println("EnterSearch_condition_list")
}

// ExitSearch_condition_list is called when production search_condition_list is exited.
func (s *LoggingTsqlListener) ExitSearch_condition_list(ctx *Search_condition_listContext) {
	log.Println("ExitSearch_condition_list")
}

// EnterSearch_cond_or is called when production search_cond_or is entered.
func (s *LoggingTsqlListener) EnterSearch_cond_or(ctx *Search_cond_orContext) {
	log.Println("EnterSearch_cond_or")
}

// ExitSearch_cond_or is called when production search_cond_or is exited.
func (s *LoggingTsqlListener) ExitSearch_cond_or(ctx *Search_cond_orContext) {
	log.Println("ExitSearch_cond_or")
}

// EnterSearch_cond_pred is called when production search_cond_pred is entered.
func (s *LoggingTsqlListener) EnterSearch_cond_pred(ctx *Search_cond_predContext) {
	log.Println("EnterSearch_cond_pred")
}

// ExitSearch_cond_pred is called when production search_cond_pred is exited.
func (s *LoggingTsqlListener) ExitSearch_cond_pred(ctx *Search_cond_predContext) {
	log.Println("ExitSearch_cond_pred")
}

// EnterSearch_cond_and is called when production search_cond_and is entered.
func (s *LoggingTsqlListener) EnterSearch_cond_and(ctx *Search_cond_andContext) {
	log.Println("EnterSearch_cond_and")
}

// ExitSearch_cond_and is called when production search_cond_and is exited.
func (s *LoggingTsqlListener) ExitSearch_cond_and(ctx *Search_cond_andContext) {
	log.Println("ExitSearch_cond_and")
}

// EnterUnary_operator_expression3 is called when production unary_operator_expression3 is entered.
func (s *LoggingTsqlListener) EnterUnary_operator_expression3(ctx *Unary_operator_expression3Context) {
	log.Println("EnterUnary_operator_expression3")
}

// ExitUnary_operator_expression3 is called when production unary_operator_expression3 is exited.
func (s *LoggingTsqlListener) ExitUnary_operator_expression3(ctx *Unary_operator_expression3Context) {
	log.Println("ExitUnary_operator_expression3")
}

// EnterUnary_operator_expression2 is called when production unary_operator_expression2 is entered.
func (s *LoggingTsqlListener) EnterUnary_operator_expression2(ctx *Unary_operator_expression2Context) {
	log.Println("EnterUnary_operator_expression2")
}

// ExitUnary_operator_expression2 is called when production unary_operator_expression2 is exited.
func (s *LoggingTsqlListener) ExitUnary_operator_expression2(ctx *Unary_operator_expression2Context) {
	log.Println("ExitUnary_operator_expression2")
}

// EnterBinary_operator_expression2 is called when production binary_operator_expression2 is entered.
func (s *LoggingTsqlListener) EnterBinary_operator_expression2(ctx *Binary_operator_expression2Context) {
	log.Println("EnterBinary_operator_expression2")
}

// ExitBinary_operator_expression2 is called when production binary_operator_expression2 is exited.
func (s *LoggingTsqlListener) ExitBinary_operator_expression2(ctx *Binary_operator_expression2Context) {
	log.Println("ExitBinary_operator_expression2")
}

// EnterSublink_expression is called when production sublink_expression is entered.
func (s *LoggingTsqlListener) EnterSublink_expression(ctx *Sublink_expressionContext) {
	log.Println("EnterSublink_expression")
}

// ExitSublink_expression is called when production sublink_expression is exited.
func (s *LoggingTsqlListener) ExitSublink_expression(ctx *Sublink_expressionContext) {
	log.Println("ExitSublink_expression")
}

// EnterBinary_mod_expression is called when production binary_mod_expression is entered.
func (s *LoggingTsqlListener) EnterBinary_mod_expression(ctx *Binary_mod_expressionContext) {
	log.Println("EnterBinary_mod_expression")
}

// ExitBinary_mod_expression is called when production binary_mod_expression is exited.
func (s *LoggingTsqlListener) ExitBinary_mod_expression(ctx *Binary_mod_expressionContext) {
	log.Println("ExitBinary_mod_expression")
}

// EnterBinary_in_expression is called when production binary_in_expression is entered.
func (s *LoggingTsqlListener) EnterBinary_in_expression(ctx *Binary_in_expressionContext) {
	log.Println("EnterBinary_in_expression")
}

// ExitBinary_in_expression is called when production binary_in_expression is exited.
func (s *LoggingTsqlListener) ExitBinary_in_expression(ctx *Binary_in_expressionContext) {
	log.Println("ExitBinary_in_expression")
}

// EnterBracket_search_expression is called when production bracket_search_expression is entered.
func (s *LoggingTsqlListener) EnterBracket_search_expression(ctx *Bracket_search_expressionContext) {
	log.Println("EnterBracket_search_expression")
}

// ExitBracket_search_expression is called when production bracket_search_expression is exited.
func (s *LoggingTsqlListener) ExitBracket_search_expression(ctx *Bracket_search_expressionContext) {
	log.Println("ExitBracket_search_expression")
}

// EnterDecimal_expression is called when production decimal_expression is entered.
func (s *LoggingTsqlListener) EnterDecimal_expression(ctx *Decimal_expressionContext) {
	log.Println("EnterDecimal_expression")
}

// ExitDecimal_expression is called when production decimal_expression is exited.
func (s *LoggingTsqlListener) ExitDecimal_expression(ctx *Decimal_expressionContext) {
	log.Println("ExitDecimal_expression")
}

// EnterBracket_query_expression is called when production bracket_query_expression is entered.
func (s *LoggingTsqlListener) EnterBracket_query_expression(ctx *Bracket_query_expressionContext) {
	log.Println("EnterBracket_query_expression")
}

// ExitBracket_query_expression is called when production bracket_query_expression is exited.
func (s *LoggingTsqlListener) ExitBracket_query_expression(ctx *Bracket_query_expressionContext) {
	log.Println("ExitBracket_query_expression")
}

// EnterQuery_specification_expression is called when production query_specification_expression is entered.
func (s *LoggingTsqlListener) EnterQuery_specification_expression(ctx *Query_specification_expressionContext) {
	log.Println("EnterQuery_specification_expression")
}

// ExitQuery_specification_expression is called when production query_specification_expression is exited.
func (s *LoggingTsqlListener) ExitQuery_specification_expression(ctx *Query_specification_expressionContext) {
	log.Println("ExitQuery_specification_expression")
}

// EnterUnion_query_expression is called when production union_query_expression is entered.
func (s *LoggingTsqlListener) EnterUnion_query_expression(ctx *Union_query_expressionContext) {
	log.Println("EnterUnion_query_expression")
}

// ExitUnion_query_expression is called when production union_query_expression is exited.
func (s *LoggingTsqlListener) ExitUnion_query_expression(ctx *Union_query_expressionContext) {
	log.Println("ExitUnion_query_expression")
}

// EnterUnion_op is called when production union_op is entered.
func (s *LoggingTsqlListener) EnterUnion_op(ctx *Union_opContext) {
	log.Println("EnterUnion_op")
}

// ExitUnion_op is called when production union_op is exited.
func (s *LoggingTsqlListener) ExitUnion_op(ctx *Union_opContext) {
	log.Println("ExitUnion_op")
}

// EnterQuery_specification is called when production query_specification is entered.
func (s *LoggingTsqlListener) EnterQuery_specification(ctx *Query_specificationContext) {
	log.Println("EnterQuery_specification")
}

// ExitQuery_specification is called when production query_specification is exited.
func (s *LoggingTsqlListener) ExitQuery_specification(ctx *Query_specificationContext) {
	log.Println("ExitQuery_specification")
}

// EnterTop_clause is called when production top_clause is entered.
func (s *LoggingTsqlListener) EnterTop_clause(ctx *Top_clauseContext) {
	log.Println("EnterTop_clause")
}

// ExitTop_clause is called when production top_clause is exited.
func (s *LoggingTsqlListener) ExitTop_clause(ctx *Top_clauseContext) {
	log.Println("ExitTop_clause")
}

// EnterTop_clause_dm is called when production top_clause_dm is entered.
func (s *LoggingTsqlListener) EnterTop_clause_dm(ctx *Top_clause_dmContext) {
	log.Println("EnterTop_clause_dm")
}

// ExitTop_clause_dm is called when production top_clause_dm is exited.
func (s *LoggingTsqlListener) ExitTop_clause_dm(ctx *Top_clause_dmContext) {
	log.Println("ExitTop_clause_dm")
}

// EnterOrder_by_clause is called when production order_by_clause is entered.
func (s *LoggingTsqlListener) EnterOrder_by_clause(ctx *Order_by_clauseContext) {
	log.Println("EnterOrder_by_clause")
}

// ExitOrder_by_clause is called when production order_by_clause is exited.
func (s *LoggingTsqlListener) ExitOrder_by_clause(ctx *Order_by_clauseContext) {
	log.Println("ExitOrder_by_clause")
}

// EnterFetch_expression is called when production fetch_expression is entered.
func (s *LoggingTsqlListener) EnterFetch_expression(ctx *Fetch_expressionContext) {
	log.Println("EnterFetch_expression")
}

// ExitFetch_expression is called when production fetch_expression is exited.
func (s *LoggingTsqlListener) ExitFetch_expression(ctx *Fetch_expressionContext) {
	log.Println("ExitFetch_expression")
}

// EnterFor_clause is called when production for_clause is entered.
func (s *LoggingTsqlListener) EnterFor_clause(ctx *For_clauseContext) {
	log.Println("EnterFor_clause")
}

// ExitFor_clause is called when production for_clause is exited.
func (s *LoggingTsqlListener) ExitFor_clause(ctx *For_clauseContext) {
	log.Println("ExitFor_clause")
}

// EnterXml_common_directives is called when production xml_common_directives is entered.
func (s *LoggingTsqlListener) EnterXml_common_directives(ctx *Xml_common_directivesContext) {
	log.Println("EnterXml_common_directives")
}

// ExitXml_common_directives is called when production xml_common_directives is exited.
func (s *LoggingTsqlListener) ExitXml_common_directives(ctx *Xml_common_directivesContext) {
	log.Println("ExitXml_common_directives")
}

// EnterOrder_by_expression is called when production order_by_expression is entered.
func (s *LoggingTsqlListener) EnterOrder_by_expression(ctx *Order_by_expressionContext) {
	log.Println("EnterOrder_by_expression")
}

// ExitOrder_by_expression is called when production order_by_expression is exited.
func (s *LoggingTsqlListener) ExitOrder_by_expression(ctx *Order_by_expressionContext) {
	log.Println("ExitOrder_by_expression")
}

// EnterGroup_by_item is called when production group_by_item is entered.
func (s *LoggingTsqlListener) EnterGroup_by_item(ctx *Group_by_itemContext) {
	log.Println("EnterGroup_by_item")
}

// ExitGroup_by_item is called when production group_by_item is exited.
func (s *LoggingTsqlListener) ExitGroup_by_item(ctx *Group_by_itemContext) {
	log.Println("ExitGroup_by_item")
}

// EnterOption_clause is called when production option_clause is entered.
func (s *LoggingTsqlListener) EnterOption_clause(ctx *Option_clauseContext) {
	log.Println("EnterOption_clause")
}

// ExitOption_clause is called when production option_clause is exited.
func (s *LoggingTsqlListener) ExitOption_clause(ctx *Option_clauseContext) {
	log.Println("ExitOption_clause")
}

// EnterOption is called when production option is entered.
func (s *LoggingTsqlListener) EnterOption(ctx *OptionContext) {
	log.Println("EnterOption")
}

// ExitOption is called when production option is exited.
func (s *LoggingTsqlListener) ExitOption(ctx *OptionContext) {
	log.Println("ExitOption")
}

// EnterOptimize_for_arg is called when production optimize_for_arg is entered.
func (s *LoggingTsqlListener) EnterOptimize_for_arg(ctx *Optimize_for_argContext) {
	log.Println("EnterOptimize_for_arg")
}

// ExitOptimize_for_arg is called when production optimize_for_arg is exited.
func (s *LoggingTsqlListener) ExitOptimize_for_arg(ctx *Optimize_for_argContext) {
	log.Println("ExitOptimize_for_arg")
}

// EnterSelect_list is called when production select_list is entered.
func (s *LoggingTsqlListener) EnterSelect_list(ctx *Select_listContext) {
	log.Println("EnterSelect_list")
}

// ExitSelect_list is called when production select_list is exited.
func (s *LoggingTsqlListener) ExitSelect_list(ctx *Select_listContext) {
	log.Println("ExitSelect_list")
}

// EnterSelect_list_elem is called when production select_list_elem is entered.
func (s *LoggingTsqlListener) EnterSelect_list_elem(ctx *Select_list_elemContext) {
	log.Println("EnterSelect_list_elem")
}

// ExitSelect_list_elem is called when production select_list_elem is exited.
func (s *LoggingTsqlListener) ExitSelect_list_elem(ctx *Select_list_elemContext) {
	log.Println("ExitSelect_list_elem")
}

// EnterTable_sources is called when production table_sources is entered.
func (s *LoggingTsqlListener) EnterTable_sources(ctx *Table_sourcesContext) {
	log.Println("EnterTable_sources")
}

// ExitTable_sources is called when production table_sources is exited.
func (s *LoggingTsqlListener) ExitTable_sources(ctx *Table_sourcesContext) {
	log.Println("ExitTable_sources")
}

// EnterCross_join is called when production cross_join is entered.
func (s *LoggingTsqlListener) EnterCross_join(ctx *Cross_joinContext) {
	log.Println("EnterCross_join")
}

// ExitCross_join is called when production cross_join is exited.
func (s *LoggingTsqlListener) ExitCross_join(ctx *Cross_joinContext) {
	log.Println("ExitCross_join")
}

// EnterTable_source_item_join is called when production table_source_item_join is entered.
func (s *LoggingTsqlListener) EnterTable_source_item_join(ctx *Table_source_item_joinContext) {
	log.Println("EnterTable_source_item_join")
}

// ExitTable_source_item_join is called when production table_source_item_join is exited.
func (s *LoggingTsqlListener) ExitTable_source_item_join(ctx *Table_source_item_joinContext) {
	log.Println("ExitTable_source_item_join")
}

// EnterStandard_join is called when production standard_join is entered.
func (s *LoggingTsqlListener) EnterStandard_join(ctx *Standard_joinContext) {
	log.Println("EnterStandard_join")
}

// ExitStandard_join is called when production standard_join is exited.
func (s *LoggingTsqlListener) ExitStandard_join(ctx *Standard_joinContext) {
	log.Println("ExitStandard_join")
}

// EnterApply_join is called when production apply_join is entered.
func (s *LoggingTsqlListener) EnterApply_join(ctx *Apply_joinContext) {
	log.Println("EnterApply_join")
}

// ExitApply_join is called when production apply_join is exited.
func (s *LoggingTsqlListener) ExitApply_join(ctx *Apply_joinContext) {
	log.Println("ExitApply_join")
}

// EnterBracket_table_source is called when production bracket_table_source is entered.
func (s *LoggingTsqlListener) EnterBracket_table_source(ctx *Bracket_table_sourceContext) {
	log.Println("EnterBracket_table_source")
}

// ExitBracket_table_source is called when production bracket_table_source is exited.
func (s *LoggingTsqlListener) ExitBracket_table_source(ctx *Bracket_table_sourceContext) {
	log.Println("ExitBracket_table_source")
}

// EnterTable_source_item_name is called when production table_source_item_name is entered.
func (s *LoggingTsqlListener) EnterTable_source_item_name(ctx *Table_source_item_nameContext) {
	log.Println("EnterTable_source_item_name")
}

// ExitTable_source_item_name is called when production table_source_item_name is exited.
func (s *LoggingTsqlListener) ExitTable_source_item_name(ctx *Table_source_item_nameContext) {
	log.Println("ExitTable_source_item_name")
}

// EnterTable_source_item_simple is called when production table_source_item_simple is entered.
func (s *LoggingTsqlListener) EnterTable_source_item_simple(ctx *Table_source_item_simpleContext) {
	log.Println("EnterTable_source_item_simple")
}

// ExitTable_source_item_simple is called when production table_source_item_simple is exited.
func (s *LoggingTsqlListener) ExitTable_source_item_simple(ctx *Table_source_item_simpleContext) {
	log.Println("ExitTable_source_item_simple")
}

// EnterTable_source_item_complex is called when production table_source_item_complex is entered.
func (s *LoggingTsqlListener) EnterTable_source_item_complex(ctx *Table_source_item_complexContext) {
	log.Println("EnterTable_source_item_complex")
}

// ExitTable_source_item_complex is called when production table_source_item_complex is exited.
func (s *LoggingTsqlListener) ExitTable_source_item_complex(ctx *Table_source_item_complexContext) {
	log.Println("ExitTable_source_item_complex")
}

// EnterTable_alias is called when production table_alias is entered.
func (s *LoggingTsqlListener) EnterTable_alias(ctx *Table_aliasContext) {
	log.Println("EnterTable_alias")
}

// ExitTable_alias is called when production table_alias is exited.
func (s *LoggingTsqlListener) ExitTable_alias(ctx *Table_aliasContext) {
	log.Println("ExitTable_alias")
}

// EnterChange_table is called when production change_table is entered.
func (s *LoggingTsqlListener) EnterChange_table(ctx *Change_tableContext) {
	log.Println("EnterChange_table")
}

// ExitChange_table is called when production change_table is exited.
func (s *LoggingTsqlListener) ExitChange_table(ctx *Change_tableContext) {
	log.Println("ExitChange_table")
}

// EnterJoin_type is called when production join_type is entered.
func (s *LoggingTsqlListener) EnterJoin_type(ctx *Join_typeContext) {
	log.Println("EnterJoin_type")
}

// ExitJoin_type is called when production join_type is exited.
func (s *LoggingTsqlListener) ExitJoin_type(ctx *Join_typeContext) {
	log.Println("ExitJoin_type")
}

// EnterTable_name_with_hint is called when production table_name_with_hint is entered.
func (s *LoggingTsqlListener) EnterTable_name_with_hint(ctx *Table_name_with_hintContext) {
	log.Println("EnterTable_name_with_hint")
}

// ExitTable_name_with_hint is called when production table_name_with_hint is exited.
func (s *LoggingTsqlListener) ExitTable_name_with_hint(ctx *Table_name_with_hintContext) {
	log.Println("ExitTable_name_with_hint")
}

// EnterRowset_function is called when production rowset_function is entered.
func (s *LoggingTsqlListener) EnterRowset_function(ctx *Rowset_functionContext) {
	log.Println("EnterRowset_function")
}

// ExitRowset_function is called when production rowset_function is exited.
func (s *LoggingTsqlListener) ExitRowset_function(ctx *Rowset_functionContext) {
	log.Println("ExitRowset_function")
}

// EnterBulk_option is called when production bulk_option is entered.
func (s *LoggingTsqlListener) EnterBulk_option(ctx *Bulk_optionContext) {
	log.Println("EnterBulk_option")
}

// ExitBulk_option is called when production bulk_option is exited.
func (s *LoggingTsqlListener) ExitBulk_option(ctx *Bulk_optionContext) {
	log.Println("ExitBulk_option")
}

// EnterDerived_table is called when production derived_table is entered.
func (s *LoggingTsqlListener) EnterDerived_table(ctx *Derived_tableContext) {
	log.Println("EnterDerived_table")
}

// ExitDerived_table is called when production derived_table is exited.
func (s *LoggingTsqlListener) ExitDerived_table(ctx *Derived_tableContext) {
	log.Println("ExitDerived_table")
}

// EnterRank_call is called when production rank_call is entered.
func (s *LoggingTsqlListener) EnterRank_call(ctx *Rank_callContext) {
	log.Println("EnterRank_call")
}

// ExitRank_call is called when production rank_call is exited.
func (s *LoggingTsqlListener) ExitRank_call(ctx *Rank_callContext) {
	log.Println("ExitRank_call")
}

// EnterAggregate_call is called when production aggregate_call is entered.
func (s *LoggingTsqlListener) EnterAggregate_call(ctx *Aggregate_callContext) {
	log.Println("EnterAggregate_call")
}

// ExitAggregate_call is called when production aggregate_call is exited.
func (s *LoggingTsqlListener) ExitAggregate_call(ctx *Aggregate_callContext) {
	log.Println("ExitAggregate_call")
}

// EnterStandard_call is called when production standard_call is entered.
func (s *LoggingTsqlListener) EnterStandard_call(ctx *Standard_callContext) {
	log.Println("EnterStandard_call")
}

// ExitStandard_call is called when production standard_call is exited.
func (s *LoggingTsqlListener) ExitStandard_call(ctx *Standard_callContext) {
	log.Println("ExitStandard_call")
}

// EnterNvf_call is called when production nvf_call is entered.
func (s *LoggingTsqlListener) EnterNvf_call(ctx *Nvf_callContext) {
	log.Println("EnterNvf_call")
}

// ExitNvf_call is called when production nvf_call is exited.
func (s *LoggingTsqlListener) ExitNvf_call(ctx *Nvf_callContext) {
	log.Println("ExitNvf_call")
}

// EnterCast_call is called when production cast_call is entered.
func (s *LoggingTsqlListener) EnterCast_call(ctx *Cast_callContext) {
	log.Println("EnterCast_call")
}

// ExitCast_call is called when production cast_call is exited.
func (s *LoggingTsqlListener) ExitCast_call(ctx *Cast_callContext) {
	log.Println("ExitCast_call")
}

// EnterSimple_call is called when production simple_call is entered.
func (s *LoggingTsqlListener) EnterSimple_call(ctx *Simple_callContext) {
	log.Println("EnterSimple_call")
}

// ExitSimple_call is called when production simple_call is exited.
func (s *LoggingTsqlListener) ExitSimple_call(ctx *Simple_callContext) {
	log.Println("ExitSimple_call")
}

// EnterSwitch_section is called when production switch_section is entered.
func (s *LoggingTsqlListener) EnterSwitch_section(ctx *Switch_sectionContext) {
	log.Println("EnterSwitch_section")
}

// ExitSwitch_section is called when production switch_section is exited.
func (s *LoggingTsqlListener) ExitSwitch_section(ctx *Switch_sectionContext) {
	log.Println("ExitSwitch_section")
}

// EnterSwitch_search_condition_section is called when production switch_search_condition_section is entered.
func (s *LoggingTsqlListener) EnterSwitch_search_condition_section(ctx *Switch_search_condition_sectionContext) {
	log.Println("EnterSwitch_search_condition_section")
}

// ExitSwitch_search_condition_section is called when production switch_search_condition_section is exited.
func (s *LoggingTsqlListener) ExitSwitch_search_condition_section(ctx *Switch_search_condition_sectionContext) {
	log.Println("ExitSwitch_search_condition_section")
}

// EnterWith_table_hints is called when production with_table_hints is entered.
func (s *LoggingTsqlListener) EnterWith_table_hints(ctx *With_table_hintsContext) {
	log.Println("EnterWith_table_hints")
}

// ExitWith_table_hints is called when production with_table_hints is exited.
func (s *LoggingTsqlListener) ExitWith_table_hints(ctx *With_table_hintsContext) {
	log.Println("ExitWith_table_hints")
}

// EnterInsert_with_table_hints is called when production insert_with_table_hints is entered.
func (s *LoggingTsqlListener) EnterInsert_with_table_hints(ctx *Insert_with_table_hintsContext) {
	log.Println("EnterInsert_with_table_hints")
}

// ExitInsert_with_table_hints is called when production insert_with_table_hints is exited.
func (s *LoggingTsqlListener) ExitInsert_with_table_hints(ctx *Insert_with_table_hintsContext) {
	log.Println("ExitInsert_with_table_hints")
}

// EnterTable_hint is called when production table_hint is entered.
func (s *LoggingTsqlListener) EnterTable_hint(ctx *Table_hintContext) {
	log.Println("EnterTable_hint")
}

// ExitTable_hint is called when production table_hint is exited.
func (s *LoggingTsqlListener) ExitTable_hint(ctx *Table_hintContext) {
	log.Println("ExitTable_hint")
}

// EnterIndex_value is called when production index_value is entered.
func (s *LoggingTsqlListener) EnterIndex_value(ctx *Index_valueContext) {
	log.Println("EnterIndex_value")
}

// ExitIndex_value is called when production index_value is exited.
func (s *LoggingTsqlListener) ExitIndex_value(ctx *Index_valueContext) {
	log.Println("ExitIndex_value")
}

// EnterColumn_alias_list is called when production column_alias_list is entered.
func (s *LoggingTsqlListener) EnterColumn_alias_list(ctx *Column_alias_listContext) {
	log.Println("EnterColumn_alias_list")
}

// ExitColumn_alias_list is called when production column_alias_list is exited.
func (s *LoggingTsqlListener) ExitColumn_alias_list(ctx *Column_alias_listContext) {
	log.Println("ExitColumn_alias_list")
}

// EnterColumn_alias is called when production column_alias is entered.
func (s *LoggingTsqlListener) EnterColumn_alias(ctx *Column_aliasContext) {
	log.Println("EnterColumn_alias")
}

// ExitColumn_alias is called when production column_alias is exited.
func (s *LoggingTsqlListener) ExitColumn_alias(ctx *Column_aliasContext) {
	log.Println("ExitColumn_alias")
}

// EnterA_star is called when production a_star is entered.
func (s *LoggingTsqlListener) EnterA_star(ctx *A_starContext) {
	log.Println("EnterA_star")
}

// ExitA_star is called when production a_star is exited.
func (s *LoggingTsqlListener) ExitA_star(ctx *A_starContext) {
	log.Println("ExitA_star")
}

// EnterTable_value_constructor is called when production table_value_constructor is entered.
func (s *LoggingTsqlListener) EnterTable_value_constructor(ctx *Table_value_constructorContext) {
	log.Println("EnterTable_value_constructor")
}

// ExitTable_value_constructor is called when production table_value_constructor is exited.
func (s *LoggingTsqlListener) ExitTable_value_constructor(ctx *Table_value_constructorContext) {
	log.Println("ExitTable_value_constructor")
}

// EnterExpression_list is called when production expression_list is entered.
func (s *LoggingTsqlListener) EnterExpression_list(ctx *Expression_listContext) {
	log.Println("EnterExpression_list")
}

// ExitExpression_list is called when production expression_list is exited.
func (s *LoggingTsqlListener) ExitExpression_list(ctx *Expression_listContext) {
	log.Println("ExitExpression_list")
}

// EnterValue_list is called when production value_list is entered.
func (s *LoggingTsqlListener) EnterValue_list(ctx *Value_listContext) {
	log.Println("EnterValue_list")
}

// ExitValue_list is called when production value_list is exited.
func (s *LoggingTsqlListener) ExitValue_list(ctx *Value_listContext) {
	log.Println("ExitValue_list")
}

// EnterNext_value_for is called when production next_value_for is entered.
func (s *LoggingTsqlListener) EnterNext_value_for(ctx *Next_value_forContext) {
	log.Println("EnterNext_value_for")
}

// ExitNext_value_for is called when production next_value_for is exited.
func (s *LoggingTsqlListener) ExitNext_value_for(ctx *Next_value_forContext) {
	log.Println("ExitNext_value_for")
}

// EnterNext_value_for_function is called when production next_value_for_function is entered.
func (s *LoggingTsqlListener) EnterNext_value_for_function(ctx *Next_value_for_functionContext) {
	log.Println("EnterNext_value_for_function")
}

// ExitNext_value_for_function is called when production next_value_for_function is exited.
func (s *LoggingTsqlListener) ExitNext_value_for_function(ctx *Next_value_for_functionContext) {
	log.Println("ExitNext_value_for_function")
}

// EnterRanking_windowed_function is called when production ranking_windowed_function is entered.
func (s *LoggingTsqlListener) EnterRanking_windowed_function(ctx *Ranking_windowed_functionContext) {
	log.Println("EnterRanking_windowed_function")
}

// ExitRanking_windowed_function is called when production ranking_windowed_function is exited.
func (s *LoggingTsqlListener) ExitRanking_windowed_function(ctx *Ranking_windowed_functionContext) {
	log.Println("ExitRanking_windowed_function")
}

// EnterAggregate_windowed_function is called when production aggregate_windowed_function is entered.
func (s *LoggingTsqlListener) EnterAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) {
	log.Println("EnterAggregate_windowed_function")
}

// ExitAggregate_windowed_function is called when production aggregate_windowed_function is exited.
func (s *LoggingTsqlListener) ExitAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) {
	log.Println("ExitAggregate_windowed_function")
}

// EnterAll_distinct is called when production all_distinct is entered.
func (s *LoggingTsqlListener) EnterAll_distinct(ctx *All_distinctContext) {
	log.Println("EnterAll_distinct")
}

// ExitAll_distinct is called when production all_distinct is exited.
func (s *LoggingTsqlListener) ExitAll_distinct(ctx *All_distinctContext) {
	log.Println("ExitAll_distinct")
}

// EnterOver_clause is called when production over_clause is entered.
func (s *LoggingTsqlListener) EnterOver_clause(ctx *Over_clauseContext) {
	log.Println("EnterOver_clause")
}

// ExitOver_clause is called when production over_clause is exited.
func (s *LoggingTsqlListener) ExitOver_clause(ctx *Over_clauseContext) {
	log.Println("ExitOver_clause")
}

// EnterRow_or_range_clause is called when production row_or_range_clause is entered.
func (s *LoggingTsqlListener) EnterRow_or_range_clause(ctx *Row_or_range_clauseContext) {
	log.Println("EnterRow_or_range_clause")
}

// ExitRow_or_range_clause is called when production row_or_range_clause is exited.
func (s *LoggingTsqlListener) ExitRow_or_range_clause(ctx *Row_or_range_clauseContext) {
	log.Println("ExitRow_or_range_clause")
}

// EnterWindow_frame_extent is called when production window_frame_extent is entered.
func (s *LoggingTsqlListener) EnterWindow_frame_extent(ctx *Window_frame_extentContext) {
	log.Println("EnterWindow_frame_extent")
}

// ExitWindow_frame_extent is called when production window_frame_extent is exited.
func (s *LoggingTsqlListener) ExitWindow_frame_extent(ctx *Window_frame_extentContext) {
	log.Println("ExitWindow_frame_extent")
}

// EnterWindow_frame_bound is called when production window_frame_bound is entered.
func (s *LoggingTsqlListener) EnterWindow_frame_bound(ctx *Window_frame_boundContext) {
	log.Println("EnterWindow_frame_bound")
}

// ExitWindow_frame_bound is called when production window_frame_bound is exited.
func (s *LoggingTsqlListener) ExitWindow_frame_bound(ctx *Window_frame_boundContext) {
	log.Println("ExitWindow_frame_bound")
}

// EnterWindow_frame_preceding is called when production window_frame_preceding is entered.
func (s *LoggingTsqlListener) EnterWindow_frame_preceding(ctx *Window_frame_precedingContext) {
	log.Println("EnterWindow_frame_preceding")
}

// ExitWindow_frame_preceding is called when production window_frame_preceding is exited.
func (s *LoggingTsqlListener) ExitWindow_frame_preceding(ctx *Window_frame_precedingContext) {
	log.Println("ExitWindow_frame_preceding")
}

// EnterWindow_frame_following is called when production window_frame_following is entered.
func (s *LoggingTsqlListener) EnterWindow_frame_following(ctx *Window_frame_followingContext) {
	log.Println("EnterWindow_frame_following")
}

// ExitWindow_frame_following is called when production window_frame_following is exited.
func (s *LoggingTsqlListener) ExitWindow_frame_following(ctx *Window_frame_followingContext) {
	log.Println("ExitWindow_frame_following")
}

// EnterCreate_database_option is called when production create_database_option is entered.
func (s *LoggingTsqlListener) EnterCreate_database_option(ctx *Create_database_optionContext) {
	log.Println("EnterCreate_database_option")
}

// ExitCreate_database_option is called when production create_database_option is exited.
func (s *LoggingTsqlListener) ExitCreate_database_option(ctx *Create_database_optionContext) {
	log.Println("ExitCreate_database_option")
}

// EnterDatabase_filestream_option is called when production database_filestream_option is entered.
func (s *LoggingTsqlListener) EnterDatabase_filestream_option(ctx *Database_filestream_optionContext) {
	log.Println("EnterDatabase_filestream_option")
}

// ExitDatabase_filestream_option is called when production database_filestream_option is exited.
func (s *LoggingTsqlListener) ExitDatabase_filestream_option(ctx *Database_filestream_optionContext) {
	log.Println("ExitDatabase_filestream_option")
}

// EnterDatabase_file_spec is called when production database_file_spec is entered.
func (s *LoggingTsqlListener) EnterDatabase_file_spec(ctx *Database_file_specContext) {
	log.Println("EnterDatabase_file_spec")
}

// ExitDatabase_file_spec is called when production database_file_spec is exited.
func (s *LoggingTsqlListener) ExitDatabase_file_spec(ctx *Database_file_specContext) {
	log.Println("ExitDatabase_file_spec")
}

// EnterFile_group is called when production file_group is entered.
func (s *LoggingTsqlListener) EnterFile_group(ctx *File_groupContext) {
	log.Println("EnterFile_group")
}

// ExitFile_group is called when production file_group is exited.
func (s *LoggingTsqlListener) ExitFile_group(ctx *File_groupContext) {
	log.Println("ExitFile_group")
}

// EnterFile_spec is called when production file_spec is entered.
func (s *LoggingTsqlListener) EnterFile_spec(ctx *File_specContext) {
	log.Println("EnterFile_spec")
}

// ExitFile_spec is called when production file_spec is exited.
func (s *LoggingTsqlListener) ExitFile_spec(ctx *File_specContext) {
	log.Println("ExitFile_spec")
}

// EnterFull_table_name is called when production full_table_name is entered.
func (s *LoggingTsqlListener) EnterFull_table_name(ctx *Full_table_nameContext) {
	log.Println("EnterFull_table_name")
}

// ExitFull_table_name is called when production full_table_name is exited.
func (s *LoggingTsqlListener) ExitFull_table_name(ctx *Full_table_nameContext) {
	log.Println("ExitFull_table_name")
}

// EnterTable_name is called when production table_name is entered.
func (s *LoggingTsqlListener) EnterTable_name(ctx *Table_nameContext) {
	log.Println("EnterTable_name")
}

// ExitTable_name is called when production table_name is exited.
func (s *LoggingTsqlListener) ExitTable_name(ctx *Table_nameContext) {
	log.Println("ExitTable_name")
}

// EnterSimple_name is called when production simple_name is entered.
func (s *LoggingTsqlListener) EnterSimple_name(ctx *Simple_nameContext) {
	log.Println("EnterSimple_name")
}

// ExitSimple_name is called when production simple_name is exited.
func (s *LoggingTsqlListener) ExitSimple_name(ctx *Simple_nameContext) {
	log.Println("ExitSimple_name")
}

// EnterFunc_proc_name is called when production func_proc_name is entered.
func (s *LoggingTsqlListener) EnterFunc_proc_name(ctx *Func_proc_nameContext) {
	log.Println("EnterFunc_proc_name")
}

// ExitFunc_proc_name is called when production func_proc_name is exited.
func (s *LoggingTsqlListener) ExitFunc_proc_name(ctx *Func_proc_nameContext) {
	log.Println("ExitFunc_proc_name")
}

// EnterDdl_object is called when production ddl_object is entered.
func (s *LoggingTsqlListener) EnterDdl_object(ctx *Ddl_objectContext) {
	log.Println("EnterDdl_object")
}

// ExitDdl_object is called when production ddl_object is exited.
func (s *LoggingTsqlListener) ExitDdl_object(ctx *Ddl_objectContext) {
	log.Println("ExitDdl_object")
}

// EnterFull_column_name is called when production full_column_name is entered.
func (s *LoggingTsqlListener) EnterFull_column_name(ctx *Full_column_nameContext) {
	log.Println("EnterFull_column_name")
}

// ExitFull_column_name is called when production full_column_name is exited.
func (s *LoggingTsqlListener) ExitFull_column_name(ctx *Full_column_nameContext) {
	log.Println("ExitFull_column_name")
}

// EnterColumn_name_list is called when production column_name_list is entered.
func (s *LoggingTsqlListener) EnterColumn_name_list(ctx *Column_name_listContext) {
	log.Println("EnterColumn_name_list")
}

// ExitColumn_name_list is called when production column_name_list is exited.
func (s *LoggingTsqlListener) ExitColumn_name_list(ctx *Column_name_listContext) {
	log.Println("ExitColumn_name_list")
}

// EnterCursor_name is called when production cursor_name is entered.
func (s *LoggingTsqlListener) EnterCursor_name(ctx *Cursor_nameContext) {
	log.Println("EnterCursor_name")
}

// ExitCursor_name is called when production cursor_name is exited.
func (s *LoggingTsqlListener) ExitCursor_name(ctx *Cursor_nameContext) {
	log.Println("ExitCursor_name")
}

// EnterOn_off is called when production on_off is entered.
func (s *LoggingTsqlListener) EnterOn_off(ctx *On_offContext) {
	log.Println("EnterOn_off")
}

// ExitOn_off is called when production on_off is exited.
func (s *LoggingTsqlListener) ExitOn_off(ctx *On_offContext) {
	log.Println("ExitOn_off")
}

// EnterClustered is called when production clustered is entered.
func (s *LoggingTsqlListener) EnterClustered(ctx *ClusteredContext) {
	log.Println("EnterClustered")
}

// ExitClustered is called when production clustered is exited.
func (s *LoggingTsqlListener) ExitClustered(ctx *ClusteredContext) {
	log.Println("ExitClustered")
}

// EnterNull_notnull is called when production null_notnull is entered.
func (s *LoggingTsqlListener) EnterNull_notnull(ctx *Null_notnullContext) {
	log.Println("EnterNull_notnull")
}

// ExitNull_notnull is called when production null_notnull is exited.
func (s *LoggingTsqlListener) ExitNull_notnull(ctx *Null_notnullContext) {
	log.Println("ExitNull_notnull")
}

// EnterScalar_function_name is called when production scalar_function_name is entered.
func (s *LoggingTsqlListener) EnterScalar_function_name(ctx *Scalar_function_nameContext) {
	log.Println("EnterScalar_function_name")
}

// ExitScalar_function_name is called when production scalar_function_name is exited.
func (s *LoggingTsqlListener) ExitScalar_function_name(ctx *Scalar_function_nameContext) {
	log.Println("ExitScalar_function_name")
}

// EnterData_type is called when production data_type is entered.
func (s *LoggingTsqlListener) EnterData_type(ctx *Data_typeContext) {
	log.Println("EnterData_type")
}

// ExitData_type is called when production data_type is exited.
func (s *LoggingTsqlListener) ExitData_type(ctx *Data_typeContext) {
	log.Println("ExitData_type")
}

// EnterDefault_value is called when production default_value is entered.
func (s *LoggingTsqlListener) EnterDefault_value(ctx *Default_valueContext) {
	log.Println("EnterDefault_value")
}

// ExitDefault_value is called when production default_value is exited.
func (s *LoggingTsqlListener) ExitDefault_value(ctx *Default_valueContext) {
	log.Println("ExitDefault_value")
}

// EnterConstant is called when production constant is entered.
func (s *LoggingTsqlListener) EnterConstant(ctx *ConstantContext) {
	log.Println("EnterConstant")
}

// ExitConstant is called when production constant is exited.
func (s *LoggingTsqlListener) ExitConstant(ctx *ConstantContext) {
	log.Println("ExitConstant")
}

// EnterSign is called when production sign is entered.
func (s *LoggingTsqlListener) EnterSign(ctx *SignContext) {
	log.Println("EnterSign")
}

// ExitSign is called when production sign is exited.
func (s *LoggingTsqlListener) ExitSign(ctx *SignContext) {
	log.Println("ExitSign")
}

// EnterR_id is called when production r_id is entered.
func (s *LoggingTsqlListener) EnterR_id(ctx *R_idContext) {
	log.Println("EnterR_id")
}

// ExitR_id is called when production r_id is exited.
func (s *LoggingTsqlListener) ExitR_id(ctx *R_idContext) {
	log.Println("ExitR_id")
}

// EnterSimple_id is called when production simple_id is entered.
func (s *LoggingTsqlListener) EnterSimple_id(ctx *Simple_idContext) {
	log.Println("EnterSimple_id")
}

// ExitSimple_id is called when production simple_id is exited.
func (s *LoggingTsqlListener) ExitSimple_id(ctx *Simple_idContext) {
	log.Println("ExitSimple_id")
}

// EnterComparison_operator is called when production comparison_operator is entered.
func (s *LoggingTsqlListener) EnterComparison_operator(ctx *Comparison_operatorContext) {
	log.Println("EnterComparison_operator")
}

// ExitComparison_operator is called when production comparison_operator is exited.
func (s *LoggingTsqlListener) ExitComparison_operator(ctx *Comparison_operatorContext) {
	log.Println("ExitComparison_operator")
}

// EnterAssignment_operator is called when production assignment_operator is entered.
func (s *LoggingTsqlListener) EnterAssignment_operator(ctx *Assignment_operatorContext) {
	log.Println("EnterAssignment_operator")
}

// ExitAssignment_operator is called when production assignment_operator is exited.
func (s *LoggingTsqlListener) ExitAssignment_operator(ctx *Assignment_operatorContext) {
	log.Println("ExitAssignment_operator")
}

// EnterFile_size is called when production file_size is entered.
func (s *LoggingTsqlListener) EnterFile_size(ctx *File_sizeContext) {
	log.Println("EnterFile_size")
}

// ExitFile_size is called when production file_size is exited.
func (s *LoggingTsqlListener) ExitFile_size(ctx *File_sizeContext) {
	log.Println("ExitFile_size")
}
