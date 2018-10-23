// Code generated from tsql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tsql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetsqlListener is a complete listener for a parse tree produced by tsqlParser.
type BasetsqlListener struct{}

var _ tsqlListener = &BasetsqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetsqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetsqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetsqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetsqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTsql_file is called when production tsql_file is entered.
func (s *BasetsqlListener) EnterTsql_file(ctx *Tsql_fileContext) {}

// ExitTsql_file is called when production tsql_file is exited.
func (s *BasetsqlListener) ExitTsql_file(ctx *Tsql_fileContext) {}

// EnterBatch is called when production batch is entered.
func (s *BasetsqlListener) EnterBatch(ctx *BatchContext) {}

// ExitBatch is called when production batch is exited.
func (s *BasetsqlListener) ExitBatch(ctx *BatchContext) {}

// EnterSql_clauses is called when production sql_clauses is entered.
func (s *BasetsqlListener) EnterSql_clauses(ctx *Sql_clausesContext) {}

// ExitSql_clauses is called when production sql_clauses is exited.
func (s *BasetsqlListener) ExitSql_clauses(ctx *Sql_clausesContext) {}

// EnterSql_clause is called when production sql_clause is entered.
func (s *BasetsqlListener) EnterSql_clause(ctx *Sql_clauseContext) {}

// ExitSql_clause is called when production sql_clause is exited.
func (s *BasetsqlListener) ExitSql_clause(ctx *Sql_clauseContext) {}

// EnterDml_clause is called when production dml_clause is entered.
func (s *BasetsqlListener) EnterDml_clause(ctx *Dml_clauseContext) {}

// ExitDml_clause is called when production dml_clause is exited.
func (s *BasetsqlListener) ExitDml_clause(ctx *Dml_clauseContext) {}

// EnterDdl_clause is called when production ddl_clause is entered.
func (s *BasetsqlListener) EnterDdl_clause(ctx *Ddl_clauseContext) {}

// ExitDdl_clause is called when production ddl_clause is exited.
func (s *BasetsqlListener) ExitDdl_clause(ctx *Ddl_clauseContext) {}

// EnterBlock_statement is called when production block_statement is entered.
func (s *BasetsqlListener) EnterBlock_statement(ctx *Block_statementContext) {}

// ExitBlock_statement is called when production block_statement is exited.
func (s *BasetsqlListener) ExitBlock_statement(ctx *Block_statementContext) {}

// EnterBreak_statement is called when production break_statement is entered.
func (s *BasetsqlListener) EnterBreak_statement(ctx *Break_statementContext) {}

// ExitBreak_statement is called when production break_statement is exited.
func (s *BasetsqlListener) ExitBreak_statement(ctx *Break_statementContext) {}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *BasetsqlListener) EnterContinue_statement(ctx *Continue_statementContext) {}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *BasetsqlListener) ExitContinue_statement(ctx *Continue_statementContext) {}

// EnterGoto_statement is called when production goto_statement is entered.
func (s *BasetsqlListener) EnterGoto_statement(ctx *Goto_statementContext) {}

// ExitGoto_statement is called when production goto_statement is exited.
func (s *BasetsqlListener) ExitGoto_statement(ctx *Goto_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BasetsqlListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BasetsqlListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BasetsqlListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BasetsqlListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterThrow_statement is called when production throw_statement is entered.
func (s *BasetsqlListener) EnterThrow_statement(ctx *Throw_statementContext) {}

// ExitThrow_statement is called when production throw_statement is exited.
func (s *BasetsqlListener) ExitThrow_statement(ctx *Throw_statementContext) {}

// EnterTry_catch_statement is called when production try_catch_statement is entered.
func (s *BasetsqlListener) EnterTry_catch_statement(ctx *Try_catch_statementContext) {}

// ExitTry_catch_statement is called when production try_catch_statement is exited.
func (s *BasetsqlListener) ExitTry_catch_statement(ctx *Try_catch_statementContext) {}

// EnterWaitfor_statement is called when production waitfor_statement is entered.
func (s *BasetsqlListener) EnterWaitfor_statement(ctx *Waitfor_statementContext) {}

// ExitWaitfor_statement is called when production waitfor_statement is exited.
func (s *BasetsqlListener) ExitWaitfor_statement(ctx *Waitfor_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BasetsqlListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BasetsqlListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterPrint_statement is called when production print_statement is entered.
func (s *BasetsqlListener) EnterPrint_statement(ctx *Print_statementContext) {}

// ExitPrint_statement is called when production print_statement is exited.
func (s *BasetsqlListener) ExitPrint_statement(ctx *Print_statementContext) {}

// EnterRaiseerror_statement is called when production raiseerror_statement is entered.
func (s *BasetsqlListener) EnterRaiseerror_statement(ctx *Raiseerror_statementContext) {}

// ExitRaiseerror_statement is called when production raiseerror_statement is exited.
func (s *BasetsqlListener) ExitRaiseerror_statement(ctx *Raiseerror_statementContext) {}

// EnterAnother_statement is called when production another_statement is entered.
func (s *BasetsqlListener) EnterAnother_statement(ctx *Another_statementContext) {}

// ExitAnother_statement is called when production another_statement is exited.
func (s *BasetsqlListener) ExitAnother_statement(ctx *Another_statementContext) {}

// EnterDelete_statement is called when production delete_statement is entered.
func (s *BasetsqlListener) EnterDelete_statement(ctx *Delete_statementContext) {}

// ExitDelete_statement is called when production delete_statement is exited.
func (s *BasetsqlListener) ExitDelete_statement(ctx *Delete_statementContext) {}

// EnterDelete_statement_from is called when production delete_statement_from is entered.
func (s *BasetsqlListener) EnterDelete_statement_from(ctx *Delete_statement_fromContext) {}

// ExitDelete_statement_from is called when production delete_statement_from is exited.
func (s *BasetsqlListener) ExitDelete_statement_from(ctx *Delete_statement_fromContext) {}

// EnterInsert_statement is called when production insert_statement is entered.
func (s *BasetsqlListener) EnterInsert_statement(ctx *Insert_statementContext) {}

// ExitInsert_statement is called when production insert_statement is exited.
func (s *BasetsqlListener) ExitInsert_statement(ctx *Insert_statementContext) {}

// EnterInsert_statement_value is called when production insert_statement_value is entered.
func (s *BasetsqlListener) EnterInsert_statement_value(ctx *Insert_statement_valueContext) {}

// ExitInsert_statement_value is called when production insert_statement_value is exited.
func (s *BasetsqlListener) ExitInsert_statement_value(ctx *Insert_statement_valueContext) {}

// EnterSelect_statement is called when production select_statement is entered.
func (s *BasetsqlListener) EnterSelect_statement(ctx *Select_statementContext) {}

// ExitSelect_statement is called when production select_statement is exited.
func (s *BasetsqlListener) ExitSelect_statement(ctx *Select_statementContext) {}

// EnterUpdate_statement is called when production update_statement is entered.
func (s *BasetsqlListener) EnterUpdate_statement(ctx *Update_statementContext) {}

// ExitUpdate_statement is called when production update_statement is exited.
func (s *BasetsqlListener) ExitUpdate_statement(ctx *Update_statementContext) {}

// EnterWhere_clause_dml is called when production where_clause_dml is entered.
func (s *BasetsqlListener) EnterWhere_clause_dml(ctx *Where_clause_dmlContext) {}

// ExitWhere_clause_dml is called when production where_clause_dml is exited.
func (s *BasetsqlListener) ExitWhere_clause_dml(ctx *Where_clause_dmlContext) {}

// EnterOutput_clause is called when production output_clause is entered.
func (s *BasetsqlListener) EnterOutput_clause(ctx *Output_clauseContext) {}

// ExitOutput_clause is called when production output_clause is exited.
func (s *BasetsqlListener) ExitOutput_clause(ctx *Output_clauseContext) {}

// EnterOutput_dml_list_elem is called when production output_dml_list_elem is entered.
func (s *BasetsqlListener) EnterOutput_dml_list_elem(ctx *Output_dml_list_elemContext) {}

// ExitOutput_dml_list_elem is called when production output_dml_list_elem is exited.
func (s *BasetsqlListener) ExitOutput_dml_list_elem(ctx *Output_dml_list_elemContext) {}

// EnterOutput_column_name is called when production output_column_name is entered.
func (s *BasetsqlListener) EnterOutput_column_name(ctx *Output_column_nameContext) {}

// ExitOutput_column_name is called when production output_column_name is exited.
func (s *BasetsqlListener) ExitOutput_column_name(ctx *Output_column_nameContext) {}

// EnterCreate_database is called when production create_database is entered.
func (s *BasetsqlListener) EnterCreate_database(ctx *Create_databaseContext) {}

// ExitCreate_database is called when production create_database is exited.
func (s *BasetsqlListener) ExitCreate_database(ctx *Create_databaseContext) {}

// EnterCreate_index is called when production create_index is entered.
func (s *BasetsqlListener) EnterCreate_index(ctx *Create_indexContext) {}

// ExitCreate_index is called when production create_index is exited.
func (s *BasetsqlListener) ExitCreate_index(ctx *Create_indexContext) {}

// EnterCreate_procedure is called when production create_procedure is entered.
func (s *BasetsqlListener) EnterCreate_procedure(ctx *Create_procedureContext) {}

// ExitCreate_procedure is called when production create_procedure is exited.
func (s *BasetsqlListener) ExitCreate_procedure(ctx *Create_procedureContext) {}

// EnterProcedure_param is called when production procedure_param is entered.
func (s *BasetsqlListener) EnterProcedure_param(ctx *Procedure_paramContext) {}

// ExitProcedure_param is called when production procedure_param is exited.
func (s *BasetsqlListener) ExitProcedure_param(ctx *Procedure_paramContext) {}

// EnterProcedure_option is called when production procedure_option is entered.
func (s *BasetsqlListener) EnterProcedure_option(ctx *Procedure_optionContext) {}

// ExitProcedure_option is called when production procedure_option is exited.
func (s *BasetsqlListener) ExitProcedure_option(ctx *Procedure_optionContext) {}

// EnterCreate_statistics is called when production create_statistics is entered.
func (s *BasetsqlListener) EnterCreate_statistics(ctx *Create_statisticsContext) {}

// ExitCreate_statistics is called when production create_statistics is exited.
func (s *BasetsqlListener) ExitCreate_statistics(ctx *Create_statisticsContext) {}

// EnterCreate_table is called when production create_table is entered.
func (s *BasetsqlListener) EnterCreate_table(ctx *Create_tableContext) {}

// ExitCreate_table is called when production create_table is exited.
func (s *BasetsqlListener) ExitCreate_table(ctx *Create_tableContext) {}

// EnterCreate_view is called when production create_view is entered.
func (s *BasetsqlListener) EnterCreate_view(ctx *Create_viewContext) {}

// ExitCreate_view is called when production create_view is exited.
func (s *BasetsqlListener) ExitCreate_view(ctx *Create_viewContext) {}

// EnterView_attribute is called when production view_attribute is entered.
func (s *BasetsqlListener) EnterView_attribute(ctx *View_attributeContext) {}

// ExitView_attribute is called when production view_attribute is exited.
func (s *BasetsqlListener) ExitView_attribute(ctx *View_attributeContext) {}

// EnterAlter_table is called when production alter_table is entered.
func (s *BasetsqlListener) EnterAlter_table(ctx *Alter_tableContext) {}

// ExitAlter_table is called when production alter_table is exited.
func (s *BasetsqlListener) ExitAlter_table(ctx *Alter_tableContext) {}

// EnterAlter_database is called when production alter_database is entered.
func (s *BasetsqlListener) EnterAlter_database(ctx *Alter_databaseContext) {}

// ExitAlter_database is called when production alter_database is exited.
func (s *BasetsqlListener) ExitAlter_database(ctx *Alter_databaseContext) {}

// EnterDatabase_optionspec is called when production database_optionspec is entered.
func (s *BasetsqlListener) EnterDatabase_optionspec(ctx *Database_optionspecContext) {}

// ExitDatabase_optionspec is called when production database_optionspec is exited.
func (s *BasetsqlListener) ExitDatabase_optionspec(ctx *Database_optionspecContext) {}

// EnterAuto_option is called when production auto_option is entered.
func (s *BasetsqlListener) EnterAuto_option(ctx *Auto_optionContext) {}

// ExitAuto_option is called when production auto_option is exited.
func (s *BasetsqlListener) ExitAuto_option(ctx *Auto_optionContext) {}

// EnterChange_tracking_option is called when production change_tracking_option is entered.
func (s *BasetsqlListener) EnterChange_tracking_option(ctx *Change_tracking_optionContext) {}

// ExitChange_tracking_option is called when production change_tracking_option is exited.
func (s *BasetsqlListener) ExitChange_tracking_option(ctx *Change_tracking_optionContext) {}

// EnterChange_tracking_option_list is called when production change_tracking_option_list is entered.
func (s *BasetsqlListener) EnterChange_tracking_option_list(ctx *Change_tracking_option_listContext) {}

// ExitChange_tracking_option_list is called when production change_tracking_option_list is exited.
func (s *BasetsqlListener) ExitChange_tracking_option_list(ctx *Change_tracking_option_listContext) {}

// EnterContainment_option is called when production containment_option is entered.
func (s *BasetsqlListener) EnterContainment_option(ctx *Containment_optionContext) {}

// ExitContainment_option is called when production containment_option is exited.
func (s *BasetsqlListener) ExitContainment_option(ctx *Containment_optionContext) {}

// EnterCursor_option is called when production cursor_option is entered.
func (s *BasetsqlListener) EnterCursor_option(ctx *Cursor_optionContext) {}

// ExitCursor_option is called when production cursor_option is exited.
func (s *BasetsqlListener) ExitCursor_option(ctx *Cursor_optionContext) {}

// EnterDate_correlation_optimization_option is called when production date_correlation_optimization_option is entered.
func (s *BasetsqlListener) EnterDate_correlation_optimization_option(ctx *Date_correlation_optimization_optionContext) {
}

// ExitDate_correlation_optimization_option is called when production date_correlation_optimization_option is exited.
func (s *BasetsqlListener) ExitDate_correlation_optimization_option(ctx *Date_correlation_optimization_optionContext) {
}

// EnterDb_encryption_option is called when production db_encryption_option is entered.
func (s *BasetsqlListener) EnterDb_encryption_option(ctx *Db_encryption_optionContext) {}

// ExitDb_encryption_option is called when production db_encryption_option is exited.
func (s *BasetsqlListener) ExitDb_encryption_option(ctx *Db_encryption_optionContext) {}

// EnterDb_state_option is called when production db_state_option is entered.
func (s *BasetsqlListener) EnterDb_state_option(ctx *Db_state_optionContext) {}

// ExitDb_state_option is called when production db_state_option is exited.
func (s *BasetsqlListener) ExitDb_state_option(ctx *Db_state_optionContext) {}

// EnterDb_update_option is called when production db_update_option is entered.
func (s *BasetsqlListener) EnterDb_update_option(ctx *Db_update_optionContext) {}

// ExitDb_update_option is called when production db_update_option is exited.
func (s *BasetsqlListener) ExitDb_update_option(ctx *Db_update_optionContext) {}

// EnterDb_user_access_option is called when production db_user_access_option is entered.
func (s *BasetsqlListener) EnterDb_user_access_option(ctx *Db_user_access_optionContext) {}

// ExitDb_user_access_option is called when production db_user_access_option is exited.
func (s *BasetsqlListener) ExitDb_user_access_option(ctx *Db_user_access_optionContext) {}

// EnterDelayed_durability_option is called when production delayed_durability_option is entered.
func (s *BasetsqlListener) EnterDelayed_durability_option(ctx *Delayed_durability_optionContext) {}

// ExitDelayed_durability_option is called when production delayed_durability_option is exited.
func (s *BasetsqlListener) ExitDelayed_durability_option(ctx *Delayed_durability_optionContext) {}

// EnterExternal_access_option is called when production external_access_option is entered.
func (s *BasetsqlListener) EnterExternal_access_option(ctx *External_access_optionContext) {}

// ExitExternal_access_option is called when production external_access_option is exited.
func (s *BasetsqlListener) ExitExternal_access_option(ctx *External_access_optionContext) {}

// EnterMixed_page_allocation_option is called when production mixed_page_allocation_option is entered.
func (s *BasetsqlListener) EnterMixed_page_allocation_option(ctx *Mixed_page_allocation_optionContext) {
}

// ExitMixed_page_allocation_option is called when production mixed_page_allocation_option is exited.
func (s *BasetsqlListener) ExitMixed_page_allocation_option(ctx *Mixed_page_allocation_optionContext) {
}

// EnterParameterization_option is called when production parameterization_option is entered.
func (s *BasetsqlListener) EnterParameterization_option(ctx *Parameterization_optionContext) {}

// ExitParameterization_option is called when production parameterization_option is exited.
func (s *BasetsqlListener) ExitParameterization_option(ctx *Parameterization_optionContext) {}

// EnterRecovery_option is called when production recovery_option is entered.
func (s *BasetsqlListener) EnterRecovery_option(ctx *Recovery_optionContext) {}

// ExitRecovery_option is called when production recovery_option is exited.
func (s *BasetsqlListener) ExitRecovery_option(ctx *Recovery_optionContext) {}

// EnterService_broker_option is called when production service_broker_option is entered.
func (s *BasetsqlListener) EnterService_broker_option(ctx *Service_broker_optionContext) {}

// ExitService_broker_option is called when production service_broker_option is exited.
func (s *BasetsqlListener) ExitService_broker_option(ctx *Service_broker_optionContext) {}

// EnterSnapshot_option is called when production snapshot_option is entered.
func (s *BasetsqlListener) EnterSnapshot_option(ctx *Snapshot_optionContext) {}

// ExitSnapshot_option is called when production snapshot_option is exited.
func (s *BasetsqlListener) ExitSnapshot_option(ctx *Snapshot_optionContext) {}

// EnterSql_option is called when production sql_option is entered.
func (s *BasetsqlListener) EnterSql_option(ctx *Sql_optionContext) {}

// ExitSql_option is called when production sql_option is exited.
func (s *BasetsqlListener) ExitSql_option(ctx *Sql_optionContext) {}

// EnterTarget_recovery_time_option is called when production target_recovery_time_option is entered.
func (s *BasetsqlListener) EnterTarget_recovery_time_option(ctx *Target_recovery_time_optionContext) {}

// ExitTarget_recovery_time_option is called when production target_recovery_time_option is exited.
func (s *BasetsqlListener) ExitTarget_recovery_time_option(ctx *Target_recovery_time_optionContext) {}

// EnterTermination is called when production termination is entered.
func (s *BasetsqlListener) EnterTermination(ctx *TerminationContext) {}

// ExitTermination is called when production termination is exited.
func (s *BasetsqlListener) ExitTermination(ctx *TerminationContext) {}

// EnterDrop_index is called when production drop_index is entered.
func (s *BasetsqlListener) EnterDrop_index(ctx *Drop_indexContext) {}

// ExitDrop_index is called when production drop_index is exited.
func (s *BasetsqlListener) ExitDrop_index(ctx *Drop_indexContext) {}

// EnterDrop_procedure is called when production drop_procedure is entered.
func (s *BasetsqlListener) EnterDrop_procedure(ctx *Drop_procedureContext) {}

// ExitDrop_procedure is called when production drop_procedure is exited.
func (s *BasetsqlListener) ExitDrop_procedure(ctx *Drop_procedureContext) {}

// EnterDrop_statistics is called when production drop_statistics is entered.
func (s *BasetsqlListener) EnterDrop_statistics(ctx *Drop_statisticsContext) {}

// ExitDrop_statistics is called when production drop_statistics is exited.
func (s *BasetsqlListener) ExitDrop_statistics(ctx *Drop_statisticsContext) {}

// EnterDrop_table is called when production drop_table is entered.
func (s *BasetsqlListener) EnterDrop_table(ctx *Drop_tableContext) {}

// ExitDrop_table is called when production drop_table is exited.
func (s *BasetsqlListener) ExitDrop_table(ctx *Drop_tableContext) {}

// EnterDrop_view is called when production drop_view is entered.
func (s *BasetsqlListener) EnterDrop_view(ctx *Drop_viewContext) {}

// ExitDrop_view is called when production drop_view is exited.
func (s *BasetsqlListener) ExitDrop_view(ctx *Drop_viewContext) {}

// EnterCreate_type is called when production create_type is entered.
func (s *BasetsqlListener) EnterCreate_type(ctx *Create_typeContext) {}

// ExitCreate_type is called when production create_type is exited.
func (s *BasetsqlListener) ExitCreate_type(ctx *Create_typeContext) {}

// EnterDrop_type is called when production drop_type is entered.
func (s *BasetsqlListener) EnterDrop_type(ctx *Drop_typeContext) {}

// ExitDrop_type is called when production drop_type is exited.
func (s *BasetsqlListener) ExitDrop_type(ctx *Drop_typeContext) {}

// EnterRowset_function_limited is called when production rowset_function_limited is entered.
func (s *BasetsqlListener) EnterRowset_function_limited(ctx *Rowset_function_limitedContext) {}

// ExitRowset_function_limited is called when production rowset_function_limited is exited.
func (s *BasetsqlListener) ExitRowset_function_limited(ctx *Rowset_function_limitedContext) {}

// EnterOpenquery is called when production openquery is entered.
func (s *BasetsqlListener) EnterOpenquery(ctx *OpenqueryContext) {}

// ExitOpenquery is called when production openquery is exited.
func (s *BasetsqlListener) ExitOpenquery(ctx *OpenqueryContext) {}

// EnterOpendatasource is called when production opendatasource is entered.
func (s *BasetsqlListener) EnterOpendatasource(ctx *OpendatasourceContext) {}

// ExitOpendatasource is called when production opendatasource is exited.
func (s *BasetsqlListener) ExitOpendatasource(ctx *OpendatasourceContext) {}

// EnterDeclare_statement is called when production declare_statement is entered.
func (s *BasetsqlListener) EnterDeclare_statement(ctx *Declare_statementContext) {}

// ExitDeclare_statement is called when production declare_statement is exited.
func (s *BasetsqlListener) ExitDeclare_statement(ctx *Declare_statementContext) {}

// EnterCursor_statement is called when production cursor_statement is entered.
func (s *BasetsqlListener) EnterCursor_statement(ctx *Cursor_statementContext) {}

// ExitCursor_statement is called when production cursor_statement is exited.
func (s *BasetsqlListener) ExitCursor_statement(ctx *Cursor_statementContext) {}

// EnterExecute_statement is called when production execute_statement is entered.
func (s *BasetsqlListener) EnterExecute_statement(ctx *Execute_statementContext) {}

// ExitExecute_statement is called when production execute_statement is exited.
func (s *BasetsqlListener) ExitExecute_statement(ctx *Execute_statementContext) {}

// EnterExecute_statement_arg is called when production execute_statement_arg is entered.
func (s *BasetsqlListener) EnterExecute_statement_arg(ctx *Execute_statement_argContext) {}

// ExitExecute_statement_arg is called when production execute_statement_arg is exited.
func (s *BasetsqlListener) ExitExecute_statement_arg(ctx *Execute_statement_argContext) {}

// EnterExecute_var_string is called when production execute_var_string is entered.
func (s *BasetsqlListener) EnterExecute_var_string(ctx *Execute_var_stringContext) {}

// ExitExecute_var_string is called when production execute_var_string is exited.
func (s *BasetsqlListener) ExitExecute_var_string(ctx *Execute_var_stringContext) {}

// EnterSecurity_statement is called when production security_statement is entered.
func (s *BasetsqlListener) EnterSecurity_statement(ctx *Security_statementContext) {}

// ExitSecurity_statement is called when production security_statement is exited.
func (s *BasetsqlListener) ExitSecurity_statement(ctx *Security_statementContext) {}

// EnterGrant_permission is called when production grant_permission is entered.
func (s *BasetsqlListener) EnterGrant_permission(ctx *Grant_permissionContext) {}

// ExitGrant_permission is called when production grant_permission is exited.
func (s *BasetsqlListener) ExitGrant_permission(ctx *Grant_permissionContext) {}

// EnterSet_statement is called when production set_statement is entered.
func (s *BasetsqlListener) EnterSet_statement(ctx *Set_statementContext) {}

// ExitSet_statement is called when production set_statement is exited.
func (s *BasetsqlListener) ExitSet_statement(ctx *Set_statementContext) {}

// EnterTransaction_statement is called when production transaction_statement is entered.
func (s *BasetsqlListener) EnterTransaction_statement(ctx *Transaction_statementContext) {}

// ExitTransaction_statement is called when production transaction_statement is exited.
func (s *BasetsqlListener) ExitTransaction_statement(ctx *Transaction_statementContext) {}

// EnterGo_statement is called when production go_statement is entered.
func (s *BasetsqlListener) EnterGo_statement(ctx *Go_statementContext) {}

// ExitGo_statement is called when production go_statement is exited.
func (s *BasetsqlListener) ExitGo_statement(ctx *Go_statementContext) {}

// EnterUse_statement is called when production use_statement is entered.
func (s *BasetsqlListener) EnterUse_statement(ctx *Use_statementContext) {}

// ExitUse_statement is called when production use_statement is exited.
func (s *BasetsqlListener) ExitUse_statement(ctx *Use_statementContext) {}

// EnterExecute_clause is called when production execute_clause is entered.
func (s *BasetsqlListener) EnterExecute_clause(ctx *Execute_clauseContext) {}

// ExitExecute_clause is called when production execute_clause is exited.
func (s *BasetsqlListener) ExitExecute_clause(ctx *Execute_clauseContext) {}

// EnterDeclare_local is called when production declare_local is entered.
func (s *BasetsqlListener) EnterDeclare_local(ctx *Declare_localContext) {}

// ExitDeclare_local is called when production declare_local is exited.
func (s *BasetsqlListener) ExitDeclare_local(ctx *Declare_localContext) {}

// EnterTable_type_definition is called when production table_type_definition is entered.
func (s *BasetsqlListener) EnterTable_type_definition(ctx *Table_type_definitionContext) {}

// ExitTable_type_definition is called when production table_type_definition is exited.
func (s *BasetsqlListener) ExitTable_type_definition(ctx *Table_type_definitionContext) {}

// EnterColumn_def_table_constraints is called when production column_def_table_constraints is entered.
func (s *BasetsqlListener) EnterColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) {
}

// ExitColumn_def_table_constraints is called when production column_def_table_constraints is exited.
func (s *BasetsqlListener) ExitColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) {
}

// EnterColumn_def_table_constraint is called when production column_def_table_constraint is entered.
func (s *BasetsqlListener) EnterColumn_def_table_constraint(ctx *Column_def_table_constraintContext) {}

// ExitColumn_def_table_constraint is called when production column_def_table_constraint is exited.
func (s *BasetsqlListener) ExitColumn_def_table_constraint(ctx *Column_def_table_constraintContext) {}

// EnterColumn_definition is called when production column_definition is entered.
func (s *BasetsqlListener) EnterColumn_definition(ctx *Column_definitionContext) {}

// ExitColumn_definition is called when production column_definition is exited.
func (s *BasetsqlListener) ExitColumn_definition(ctx *Column_definitionContext) {}

// EnterColumn_constraint is called when production column_constraint is entered.
func (s *BasetsqlListener) EnterColumn_constraint(ctx *Column_constraintContext) {}

// ExitColumn_constraint is called when production column_constraint is exited.
func (s *BasetsqlListener) ExitColumn_constraint(ctx *Column_constraintContext) {}

// EnterTable_constraint is called when production table_constraint is entered.
func (s *BasetsqlListener) EnterTable_constraint(ctx *Table_constraintContext) {}

// ExitTable_constraint is called when production table_constraint is exited.
func (s *BasetsqlListener) ExitTable_constraint(ctx *Table_constraintContext) {}

// EnterIndex_options is called when production index_options is entered.
func (s *BasetsqlListener) EnterIndex_options(ctx *Index_optionsContext) {}

// ExitIndex_options is called when production index_options is exited.
func (s *BasetsqlListener) ExitIndex_options(ctx *Index_optionsContext) {}

// EnterIndex_option is called when production index_option is entered.
func (s *BasetsqlListener) EnterIndex_option(ctx *Index_optionContext) {}

// ExitIndex_option is called when production index_option is exited.
func (s *BasetsqlListener) ExitIndex_option(ctx *Index_optionContext) {}

// EnterDeclare_cursor is called when production declare_cursor is entered.
func (s *BasetsqlListener) EnterDeclare_cursor(ctx *Declare_cursorContext) {}

// ExitDeclare_cursor is called when production declare_cursor is exited.
func (s *BasetsqlListener) ExitDeclare_cursor(ctx *Declare_cursorContext) {}

// EnterDeclare_set_cursor_common is called when production declare_set_cursor_common is entered.
func (s *BasetsqlListener) EnterDeclare_set_cursor_common(ctx *Declare_set_cursor_commonContext) {}

// ExitDeclare_set_cursor_common is called when production declare_set_cursor_common is exited.
func (s *BasetsqlListener) ExitDeclare_set_cursor_common(ctx *Declare_set_cursor_commonContext) {}

// EnterFetch_cursor is called when production fetch_cursor is entered.
func (s *BasetsqlListener) EnterFetch_cursor(ctx *Fetch_cursorContext) {}

// ExitFetch_cursor is called when production fetch_cursor is exited.
func (s *BasetsqlListener) ExitFetch_cursor(ctx *Fetch_cursorContext) {}

// EnterSet_special is called when production set_special is entered.
func (s *BasetsqlListener) EnterSet_special(ctx *Set_specialContext) {}

// ExitSet_special is called when production set_special is exited.
func (s *BasetsqlListener) ExitSet_special(ctx *Set_specialContext) {}

// EnterConstant_LOCAL_ID is called when production constant_LOCAL_ID is entered.
func (s *BasetsqlListener) EnterConstant_LOCAL_ID(ctx *Constant_LOCAL_IDContext) {}

// ExitConstant_LOCAL_ID is called when production constant_LOCAL_ID is exited.
func (s *BasetsqlListener) ExitConstant_LOCAL_ID(ctx *Constant_LOCAL_IDContext) {}

// EnterBinary_operator_expression is called when production binary_operator_expression is entered.
func (s *BasetsqlListener) EnterBinary_operator_expression(ctx *Binary_operator_expressionContext) {}

// ExitBinary_operator_expression is called when production binary_operator_expression is exited.
func (s *BasetsqlListener) ExitBinary_operator_expression(ctx *Binary_operator_expressionContext) {}

// EnterPrimitive_expression is called when production primitive_expression is entered.
func (s *BasetsqlListener) EnterPrimitive_expression(ctx *Primitive_expressionContext) {}

// ExitPrimitive_expression is called when production primitive_expression is exited.
func (s *BasetsqlListener) ExitPrimitive_expression(ctx *Primitive_expressionContext) {}

// EnterBracket_expression is called when production bracket_expression is entered.
func (s *BasetsqlListener) EnterBracket_expression(ctx *Bracket_expressionContext) {}

// ExitBracket_expression is called when production bracket_expression is exited.
func (s *BasetsqlListener) ExitBracket_expression(ctx *Bracket_expressionContext) {}

// EnterUnary_operator_expression is called when production unary_operator_expression is entered.
func (s *BasetsqlListener) EnterUnary_operator_expression(ctx *Unary_operator_expressionContext) {}

// ExitUnary_operator_expression is called when production unary_operator_expression is exited.
func (s *BasetsqlListener) ExitUnary_operator_expression(ctx *Unary_operator_expressionContext) {}

// EnterFunction_call_expression is called when production function_call_expression is entered.
func (s *BasetsqlListener) EnterFunction_call_expression(ctx *Function_call_expressionContext) {}

// ExitFunction_call_expression is called when production function_call_expression is exited.
func (s *BasetsqlListener) ExitFunction_call_expression(ctx *Function_call_expressionContext) {}

// EnterCase_expression is called when production case_expression is entered.
func (s *BasetsqlListener) EnterCase_expression(ctx *Case_expressionContext) {}

// ExitCase_expression is called when production case_expression is exited.
func (s *BasetsqlListener) ExitCase_expression(ctx *Case_expressionContext) {}

// EnterColumn_ref_expression is called when production column_ref_expression is entered.
func (s *BasetsqlListener) EnterColumn_ref_expression(ctx *Column_ref_expressionContext) {}

// ExitColumn_ref_expression is called when production column_ref_expression is exited.
func (s *BasetsqlListener) ExitColumn_ref_expression(ctx *Column_ref_expressionContext) {}

// EnterSubquery_expression is called when production subquery_expression is entered.
func (s *BasetsqlListener) EnterSubquery_expression(ctx *Subquery_expressionContext) {}

// ExitSubquery_expression is called when production subquery_expression is exited.
func (s *BasetsqlListener) ExitSubquery_expression(ctx *Subquery_expressionContext) {}

// EnterOver_clause_expression is called when production over_clause_expression is entered.
func (s *BasetsqlListener) EnterOver_clause_expression(ctx *Over_clause_expressionContext) {}

// ExitOver_clause_expression is called when production over_clause_expression is exited.
func (s *BasetsqlListener) ExitOver_clause_expression(ctx *Over_clause_expressionContext) {}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *BasetsqlListener) EnterConstant_expression(ctx *Constant_expressionContext) {}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *BasetsqlListener) ExitConstant_expression(ctx *Constant_expressionContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BasetsqlListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BasetsqlListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterWith_expression is called when production with_expression is entered.
func (s *BasetsqlListener) EnterWith_expression(ctx *With_expressionContext) {}

// ExitWith_expression is called when production with_expression is exited.
func (s *BasetsqlListener) ExitWith_expression(ctx *With_expressionContext) {}

// EnterCommon_table_expression is called when production common_table_expression is entered.
func (s *BasetsqlListener) EnterCommon_table_expression(ctx *Common_table_expressionContext) {}

// ExitCommon_table_expression is called when production common_table_expression is exited.
func (s *BasetsqlListener) ExitCommon_table_expression(ctx *Common_table_expressionContext) {}

// EnterUpdate_elem is called when production update_elem is entered.
func (s *BasetsqlListener) EnterUpdate_elem(ctx *Update_elemContext) {}

// ExitUpdate_elem is called when production update_elem is exited.
func (s *BasetsqlListener) ExitUpdate_elem(ctx *Update_elemContext) {}

// EnterSearch_condition_list is called when production search_condition_list is entered.
func (s *BasetsqlListener) EnterSearch_condition_list(ctx *Search_condition_listContext) {}

// ExitSearch_condition_list is called when production search_condition_list is exited.
func (s *BasetsqlListener) ExitSearch_condition_list(ctx *Search_condition_listContext) {}

// EnterSearch_cond_or is called when production search_cond_or is entered.
func (s *BasetsqlListener) EnterSearch_cond_or(ctx *Search_cond_orContext) {}

// ExitSearch_cond_or is called when production search_cond_or is exited.
func (s *BasetsqlListener) ExitSearch_cond_or(ctx *Search_cond_orContext) {}

// EnterSearch_cond_pred is called when production search_cond_pred is entered.
func (s *BasetsqlListener) EnterSearch_cond_pred(ctx *Search_cond_predContext) {}

// ExitSearch_cond_pred is called when production search_cond_pred is exited.
func (s *BasetsqlListener) ExitSearch_cond_pred(ctx *Search_cond_predContext) {}

// EnterSearch_cond_and is called when production search_cond_and is entered.
func (s *BasetsqlListener) EnterSearch_cond_and(ctx *Search_cond_andContext) {}

// ExitSearch_cond_and is called when production search_cond_and is exited.
func (s *BasetsqlListener) ExitSearch_cond_and(ctx *Search_cond_andContext) {}

// EnterUnary_operator_expression3 is called when production unary_operator_expression3 is entered.
func (s *BasetsqlListener) EnterUnary_operator_expression3(ctx *Unary_operator_expression3Context) {}

// ExitUnary_operator_expression3 is called when production unary_operator_expression3 is exited.
func (s *BasetsqlListener) ExitUnary_operator_expression3(ctx *Unary_operator_expression3Context) {}

// EnterUnary_operator_expression2 is called when production unary_operator_expression2 is entered.
func (s *BasetsqlListener) EnterUnary_operator_expression2(ctx *Unary_operator_expression2Context) {}

// ExitUnary_operator_expression2 is called when production unary_operator_expression2 is exited.
func (s *BasetsqlListener) ExitUnary_operator_expression2(ctx *Unary_operator_expression2Context) {}

// EnterBinary_operator_expression2 is called when production binary_operator_expression2 is entered.
func (s *BasetsqlListener) EnterBinary_operator_expression2(ctx *Binary_operator_expression2Context) {}

// ExitBinary_operator_expression2 is called when production binary_operator_expression2 is exited.
func (s *BasetsqlListener) ExitBinary_operator_expression2(ctx *Binary_operator_expression2Context) {}

// EnterSublink_expression is called when production sublink_expression is entered.
func (s *BasetsqlListener) EnterSublink_expression(ctx *Sublink_expressionContext) {}

// ExitSublink_expression is called when production sublink_expression is exited.
func (s *BasetsqlListener) ExitSublink_expression(ctx *Sublink_expressionContext) {}

// EnterBinary_mod_expression is called when production binary_mod_expression is entered.
func (s *BasetsqlListener) EnterBinary_mod_expression(ctx *Binary_mod_expressionContext) {}

// ExitBinary_mod_expression is called when production binary_mod_expression is exited.
func (s *BasetsqlListener) ExitBinary_mod_expression(ctx *Binary_mod_expressionContext) {}

// EnterBinary_in_expression is called when production binary_in_expression is entered.
func (s *BasetsqlListener) EnterBinary_in_expression(ctx *Binary_in_expressionContext) {}

// ExitBinary_in_expression is called when production binary_in_expression is exited.
func (s *BasetsqlListener) ExitBinary_in_expression(ctx *Binary_in_expressionContext) {}

// EnterBracket_search_expression is called when production bracket_search_expression is entered.
func (s *BasetsqlListener) EnterBracket_search_expression(ctx *Bracket_search_expressionContext) {}

// ExitBracket_search_expression is called when production bracket_search_expression is exited.
func (s *BasetsqlListener) ExitBracket_search_expression(ctx *Bracket_search_expressionContext) {}

// EnterDecimal_expression is called when production decimal_expression is entered.
func (s *BasetsqlListener) EnterDecimal_expression(ctx *Decimal_expressionContext) {}

// ExitDecimal_expression is called when production decimal_expression is exited.
func (s *BasetsqlListener) ExitDecimal_expression(ctx *Decimal_expressionContext) {}

// EnterBracket_query_expression is called when production bracket_query_expression is entered.
func (s *BasetsqlListener) EnterBracket_query_expression(ctx *Bracket_query_expressionContext) {}

// ExitBracket_query_expression is called when production bracket_query_expression is exited.
func (s *BasetsqlListener) ExitBracket_query_expression(ctx *Bracket_query_expressionContext) {}

// EnterQuery_specification_expression is called when production query_specification_expression is entered.
func (s *BasetsqlListener) EnterQuery_specification_expression(ctx *Query_specification_expressionContext) {
}

// ExitQuery_specification_expression is called when production query_specification_expression is exited.
func (s *BasetsqlListener) ExitQuery_specification_expression(ctx *Query_specification_expressionContext) {
}

// EnterUnion_query_expression is called when production union_query_expression is entered.
func (s *BasetsqlListener) EnterUnion_query_expression(ctx *Union_query_expressionContext) {}

// ExitUnion_query_expression is called when production union_query_expression is exited.
func (s *BasetsqlListener) ExitUnion_query_expression(ctx *Union_query_expressionContext) {}

// EnterUnion_op is called when production union_op is entered.
func (s *BasetsqlListener) EnterUnion_op(ctx *Union_opContext) {}

// ExitUnion_op is called when production union_op is exited.
func (s *BasetsqlListener) ExitUnion_op(ctx *Union_opContext) {}

// EnterQuery_specification is called when production query_specification is entered.
func (s *BasetsqlListener) EnterQuery_specification(ctx *Query_specificationContext) {}

// ExitQuery_specification is called when production query_specification is exited.
func (s *BasetsqlListener) ExitQuery_specification(ctx *Query_specificationContext) {}

// EnterTop_clause is called when production top_clause is entered.
func (s *BasetsqlListener) EnterTop_clause(ctx *Top_clauseContext) {}

// ExitTop_clause is called when production top_clause is exited.
func (s *BasetsqlListener) ExitTop_clause(ctx *Top_clauseContext) {}

// EnterTop_clause_dm is called when production top_clause_dm is entered.
func (s *BasetsqlListener) EnterTop_clause_dm(ctx *Top_clause_dmContext) {}

// ExitTop_clause_dm is called when production top_clause_dm is exited.
func (s *BasetsqlListener) ExitTop_clause_dm(ctx *Top_clause_dmContext) {}

// EnterOrder_by_clause is called when production order_by_clause is entered.
func (s *BasetsqlListener) EnterOrder_by_clause(ctx *Order_by_clauseContext) {}

// ExitOrder_by_clause is called when production order_by_clause is exited.
func (s *BasetsqlListener) ExitOrder_by_clause(ctx *Order_by_clauseContext) {}

// EnterFetch_expression is called when production fetch_expression is entered.
func (s *BasetsqlListener) EnterFetch_expression(ctx *Fetch_expressionContext) {}

// ExitFetch_expression is called when production fetch_expression is exited.
func (s *BasetsqlListener) ExitFetch_expression(ctx *Fetch_expressionContext) {}

// EnterFor_clause is called when production for_clause is entered.
func (s *BasetsqlListener) EnterFor_clause(ctx *For_clauseContext) {}

// ExitFor_clause is called when production for_clause is exited.
func (s *BasetsqlListener) ExitFor_clause(ctx *For_clauseContext) {}

// EnterXml_common_directives is called when production xml_common_directives is entered.
func (s *BasetsqlListener) EnterXml_common_directives(ctx *Xml_common_directivesContext) {}

// ExitXml_common_directives is called when production xml_common_directives is exited.
func (s *BasetsqlListener) ExitXml_common_directives(ctx *Xml_common_directivesContext) {}

// EnterOrder_by_expression is called when production order_by_expression is entered.
func (s *BasetsqlListener) EnterOrder_by_expression(ctx *Order_by_expressionContext) {}

// ExitOrder_by_expression is called when production order_by_expression is exited.
func (s *BasetsqlListener) ExitOrder_by_expression(ctx *Order_by_expressionContext) {}

// EnterGroup_by_item is called when production group_by_item is entered.
func (s *BasetsqlListener) EnterGroup_by_item(ctx *Group_by_itemContext) {}

// ExitGroup_by_item is called when production group_by_item is exited.
func (s *BasetsqlListener) ExitGroup_by_item(ctx *Group_by_itemContext) {}

// EnterOption_clause is called when production option_clause is entered.
func (s *BasetsqlListener) EnterOption_clause(ctx *Option_clauseContext) {}

// ExitOption_clause is called when production option_clause is exited.
func (s *BasetsqlListener) ExitOption_clause(ctx *Option_clauseContext) {}

// EnterOption is called when production option is entered.
func (s *BasetsqlListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BasetsqlListener) ExitOption(ctx *OptionContext) {}

// EnterOptimize_for_arg is called when production optimize_for_arg is entered.
func (s *BasetsqlListener) EnterOptimize_for_arg(ctx *Optimize_for_argContext) {}

// ExitOptimize_for_arg is called when production optimize_for_arg is exited.
func (s *BasetsqlListener) ExitOptimize_for_arg(ctx *Optimize_for_argContext) {}

// EnterSelect_list is called when production select_list is entered.
func (s *BasetsqlListener) EnterSelect_list(ctx *Select_listContext) {}

// ExitSelect_list is called when production select_list is exited.
func (s *BasetsqlListener) ExitSelect_list(ctx *Select_listContext) {}

// EnterSelect_list_elem is called when production select_list_elem is entered.
func (s *BasetsqlListener) EnterSelect_list_elem(ctx *Select_list_elemContext) {}

// ExitSelect_list_elem is called when production select_list_elem is exited.
func (s *BasetsqlListener) ExitSelect_list_elem(ctx *Select_list_elemContext) {}

// EnterTable_sources is called when production table_sources is entered.
func (s *BasetsqlListener) EnterTable_sources(ctx *Table_sourcesContext) {}

// ExitTable_sources is called when production table_sources is exited.
func (s *BasetsqlListener) ExitTable_sources(ctx *Table_sourcesContext) {}

// EnterCross_join is called when production cross_join is entered.
func (s *BasetsqlListener) EnterCross_join(ctx *Cross_joinContext) {}

// ExitCross_join is called when production cross_join is exited.
func (s *BasetsqlListener) ExitCross_join(ctx *Cross_joinContext) {}

// EnterTable_source_item_join is called when production table_source_item_join is entered.
func (s *BasetsqlListener) EnterTable_source_item_join(ctx *Table_source_item_joinContext) {}

// ExitTable_source_item_join is called when production table_source_item_join is exited.
func (s *BasetsqlListener) ExitTable_source_item_join(ctx *Table_source_item_joinContext) {}

// EnterStandard_join is called when production standard_join is entered.
func (s *BasetsqlListener) EnterStandard_join(ctx *Standard_joinContext) {}

// ExitStandard_join is called when production standard_join is exited.
func (s *BasetsqlListener) ExitStandard_join(ctx *Standard_joinContext) {}

// EnterApply_join is called when production apply_join is entered.
func (s *BasetsqlListener) EnterApply_join(ctx *Apply_joinContext) {}

// ExitApply_join is called when production apply_join is exited.
func (s *BasetsqlListener) ExitApply_join(ctx *Apply_joinContext) {}

// EnterBracket_table_source is called when production bracket_table_source is entered.
func (s *BasetsqlListener) EnterBracket_table_source(ctx *Bracket_table_sourceContext) {}

// ExitBracket_table_source is called when production bracket_table_source is exited.
func (s *BasetsqlListener) ExitBracket_table_source(ctx *Bracket_table_sourceContext) {}

// EnterTable_source_item_name is called when production table_source_item_name is entered.
func (s *BasetsqlListener) EnterTable_source_item_name(ctx *Table_source_item_nameContext) {}

// ExitTable_source_item_name is called when production table_source_item_name is exited.
func (s *BasetsqlListener) ExitTable_source_item_name(ctx *Table_source_item_nameContext) {}

// EnterTable_source_item_simple is called when production table_source_item_simple is entered.
func (s *BasetsqlListener) EnterTable_source_item_simple(ctx *Table_source_item_simpleContext) {}

// ExitTable_source_item_simple is called when production table_source_item_simple is exited.
func (s *BasetsqlListener) ExitTable_source_item_simple(ctx *Table_source_item_simpleContext) {}

// EnterTable_source_item_complex is called when production table_source_item_complex is entered.
func (s *BasetsqlListener) EnterTable_source_item_complex(ctx *Table_source_item_complexContext) {}

// ExitTable_source_item_complex is called when production table_source_item_complex is exited.
func (s *BasetsqlListener) ExitTable_source_item_complex(ctx *Table_source_item_complexContext) {}

// EnterTable_alias is called when production table_alias is entered.
func (s *BasetsqlListener) EnterTable_alias(ctx *Table_aliasContext) {}

// ExitTable_alias is called when production table_alias is exited.
func (s *BasetsqlListener) ExitTable_alias(ctx *Table_aliasContext) {}

// EnterChange_table is called when production change_table is entered.
func (s *BasetsqlListener) EnterChange_table(ctx *Change_tableContext) {}

// ExitChange_table is called when production change_table is exited.
func (s *BasetsqlListener) ExitChange_table(ctx *Change_tableContext) {}

// EnterJoin_type is called when production join_type is entered.
func (s *BasetsqlListener) EnterJoin_type(ctx *Join_typeContext) {}

// ExitJoin_type is called when production join_type is exited.
func (s *BasetsqlListener) ExitJoin_type(ctx *Join_typeContext) {}

// EnterTable_name_with_hint is called when production table_name_with_hint is entered.
func (s *BasetsqlListener) EnterTable_name_with_hint(ctx *Table_name_with_hintContext) {}

// ExitTable_name_with_hint is called when production table_name_with_hint is exited.
func (s *BasetsqlListener) ExitTable_name_with_hint(ctx *Table_name_with_hintContext) {}

// EnterRowset_function is called when production rowset_function is entered.
func (s *BasetsqlListener) EnterRowset_function(ctx *Rowset_functionContext) {}

// ExitRowset_function is called when production rowset_function is exited.
func (s *BasetsqlListener) ExitRowset_function(ctx *Rowset_functionContext) {}

// EnterBulk_option is called when production bulk_option is entered.
func (s *BasetsqlListener) EnterBulk_option(ctx *Bulk_optionContext) {}

// ExitBulk_option is called when production bulk_option is exited.
func (s *BasetsqlListener) ExitBulk_option(ctx *Bulk_optionContext) {}

// EnterDerived_table is called when production derived_table is entered.
func (s *BasetsqlListener) EnterDerived_table(ctx *Derived_tableContext) {}

// ExitDerived_table is called when production derived_table is exited.
func (s *BasetsqlListener) ExitDerived_table(ctx *Derived_tableContext) {}

// EnterRank_call is called when production rank_call is entered.
func (s *BasetsqlListener) EnterRank_call(ctx *Rank_callContext) {}

// ExitRank_call is called when production rank_call is exited.
func (s *BasetsqlListener) ExitRank_call(ctx *Rank_callContext) {}

// EnterAggregate_call is called when production aggregate_call is entered.
func (s *BasetsqlListener) EnterAggregate_call(ctx *Aggregate_callContext) {}

// ExitAggregate_call is called when production aggregate_call is exited.
func (s *BasetsqlListener) ExitAggregate_call(ctx *Aggregate_callContext) {}

// EnterStandard_call is called when production standard_call is entered.
func (s *BasetsqlListener) EnterStandard_call(ctx *Standard_callContext) {}

// ExitStandard_call is called when production standard_call is exited.
func (s *BasetsqlListener) ExitStandard_call(ctx *Standard_callContext) {}

// EnterNvf_call is called when production nvf_call is entered.
func (s *BasetsqlListener) EnterNvf_call(ctx *Nvf_callContext) {}

// ExitNvf_call is called when production nvf_call is exited.
func (s *BasetsqlListener) ExitNvf_call(ctx *Nvf_callContext) {}

// EnterCast_call is called when production cast_call is entered.
func (s *BasetsqlListener) EnterCast_call(ctx *Cast_callContext) {}

// ExitCast_call is called when production cast_call is exited.
func (s *BasetsqlListener) ExitCast_call(ctx *Cast_callContext) {}

// EnterSimple_call is called when production simple_call is entered.
func (s *BasetsqlListener) EnterSimple_call(ctx *Simple_callContext) {}

// ExitSimple_call is called when production simple_call is exited.
func (s *BasetsqlListener) ExitSimple_call(ctx *Simple_callContext) {}

// EnterSwitch_section is called when production switch_section is entered.
func (s *BasetsqlListener) EnterSwitch_section(ctx *Switch_sectionContext) {}

// ExitSwitch_section is called when production switch_section is exited.
func (s *BasetsqlListener) ExitSwitch_section(ctx *Switch_sectionContext) {}

// EnterSwitch_search_condition_section is called when production switch_search_condition_section is entered.
func (s *BasetsqlListener) EnterSwitch_search_condition_section(ctx *Switch_search_condition_sectionContext) {
}

// ExitSwitch_search_condition_section is called when production switch_search_condition_section is exited.
func (s *BasetsqlListener) ExitSwitch_search_condition_section(ctx *Switch_search_condition_sectionContext) {
}

// EnterWith_table_hints is called when production with_table_hints is entered.
func (s *BasetsqlListener) EnterWith_table_hints(ctx *With_table_hintsContext) {}

// ExitWith_table_hints is called when production with_table_hints is exited.
func (s *BasetsqlListener) ExitWith_table_hints(ctx *With_table_hintsContext) {}

// EnterInsert_with_table_hints is called when production insert_with_table_hints is entered.
func (s *BasetsqlListener) EnterInsert_with_table_hints(ctx *Insert_with_table_hintsContext) {}

// ExitInsert_with_table_hints is called when production insert_with_table_hints is exited.
func (s *BasetsqlListener) ExitInsert_with_table_hints(ctx *Insert_with_table_hintsContext) {}

// EnterTable_hint is called when production table_hint is entered.
func (s *BasetsqlListener) EnterTable_hint(ctx *Table_hintContext) {}

// ExitTable_hint is called when production table_hint is exited.
func (s *BasetsqlListener) ExitTable_hint(ctx *Table_hintContext) {}

// EnterIndex_value is called when production index_value is entered.
func (s *BasetsqlListener) EnterIndex_value(ctx *Index_valueContext) {}

// ExitIndex_value is called when production index_value is exited.
func (s *BasetsqlListener) ExitIndex_value(ctx *Index_valueContext) {}

// EnterColumn_alias_list is called when production column_alias_list is entered.
func (s *BasetsqlListener) EnterColumn_alias_list(ctx *Column_alias_listContext) {}

// ExitColumn_alias_list is called when production column_alias_list is exited.
func (s *BasetsqlListener) ExitColumn_alias_list(ctx *Column_alias_listContext) {}

// EnterColumn_alias is called when production column_alias is entered.
func (s *BasetsqlListener) EnterColumn_alias(ctx *Column_aliasContext) {}

// ExitColumn_alias is called when production column_alias is exited.
func (s *BasetsqlListener) ExitColumn_alias(ctx *Column_aliasContext) {}

// EnterA_star is called when production a_star is entered.
func (s *BasetsqlListener) EnterA_star(ctx *A_starContext) {}

// ExitA_star is called when production a_star is exited.
func (s *BasetsqlListener) ExitA_star(ctx *A_starContext) {}

// EnterTable_value_constructor is called when production table_value_constructor is entered.
func (s *BasetsqlListener) EnterTable_value_constructor(ctx *Table_value_constructorContext) {}

// ExitTable_value_constructor is called when production table_value_constructor is exited.
func (s *BasetsqlListener) ExitTable_value_constructor(ctx *Table_value_constructorContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BasetsqlListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BasetsqlListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterValue_list is called when production value_list is entered.
func (s *BasetsqlListener) EnterValue_list(ctx *Value_listContext) {}

// ExitValue_list is called when production value_list is exited.
func (s *BasetsqlListener) ExitValue_list(ctx *Value_listContext) {}

// EnterNext_value_for is called when production next_value_for is entered.
func (s *BasetsqlListener) EnterNext_value_for(ctx *Next_value_forContext) {}

// ExitNext_value_for is called when production next_value_for is exited.
func (s *BasetsqlListener) ExitNext_value_for(ctx *Next_value_forContext) {}

// EnterNext_value_for_function is called when production next_value_for_function is entered.
func (s *BasetsqlListener) EnterNext_value_for_function(ctx *Next_value_for_functionContext) {}

// ExitNext_value_for_function is called when production next_value_for_function is exited.
func (s *BasetsqlListener) ExitNext_value_for_function(ctx *Next_value_for_functionContext) {}

// EnterRanking_windowed_function is called when production ranking_windowed_function is entered.
func (s *BasetsqlListener) EnterRanking_windowed_function(ctx *Ranking_windowed_functionContext) {}

// ExitRanking_windowed_function is called when production ranking_windowed_function is exited.
func (s *BasetsqlListener) ExitRanking_windowed_function(ctx *Ranking_windowed_functionContext) {}

// EnterAggregate_windowed_function is called when production aggregate_windowed_function is entered.
func (s *BasetsqlListener) EnterAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) {}

// ExitAggregate_windowed_function is called when production aggregate_windowed_function is exited.
func (s *BasetsqlListener) ExitAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) {}

// EnterAll_distinct is called when production all_distinct is entered.
func (s *BasetsqlListener) EnterAll_distinct(ctx *All_distinctContext) {}

// ExitAll_distinct is called when production all_distinct is exited.
func (s *BasetsqlListener) ExitAll_distinct(ctx *All_distinctContext) {}

// EnterOver_clause is called when production over_clause is entered.
func (s *BasetsqlListener) EnterOver_clause(ctx *Over_clauseContext) {}

// ExitOver_clause is called when production over_clause is exited.
func (s *BasetsqlListener) ExitOver_clause(ctx *Over_clauseContext) {}

// EnterRow_or_range_clause is called when production row_or_range_clause is entered.
func (s *BasetsqlListener) EnterRow_or_range_clause(ctx *Row_or_range_clauseContext) {}

// ExitRow_or_range_clause is called when production row_or_range_clause is exited.
func (s *BasetsqlListener) ExitRow_or_range_clause(ctx *Row_or_range_clauseContext) {}

// EnterWindow_frame_extent is called when production window_frame_extent is entered.
func (s *BasetsqlListener) EnterWindow_frame_extent(ctx *Window_frame_extentContext) {}

// ExitWindow_frame_extent is called when production window_frame_extent is exited.
func (s *BasetsqlListener) ExitWindow_frame_extent(ctx *Window_frame_extentContext) {}

// EnterWindow_frame_bound is called when production window_frame_bound is entered.
func (s *BasetsqlListener) EnterWindow_frame_bound(ctx *Window_frame_boundContext) {}

// ExitWindow_frame_bound is called when production window_frame_bound is exited.
func (s *BasetsqlListener) ExitWindow_frame_bound(ctx *Window_frame_boundContext) {}

// EnterWindow_frame_preceding is called when production window_frame_preceding is entered.
func (s *BasetsqlListener) EnterWindow_frame_preceding(ctx *Window_frame_precedingContext) {}

// ExitWindow_frame_preceding is called when production window_frame_preceding is exited.
func (s *BasetsqlListener) ExitWindow_frame_preceding(ctx *Window_frame_precedingContext) {}

// EnterWindow_frame_following is called when production window_frame_following is entered.
func (s *BasetsqlListener) EnterWindow_frame_following(ctx *Window_frame_followingContext) {}

// ExitWindow_frame_following is called when production window_frame_following is exited.
func (s *BasetsqlListener) ExitWindow_frame_following(ctx *Window_frame_followingContext) {}

// EnterCreate_database_option is called when production create_database_option is entered.
func (s *BasetsqlListener) EnterCreate_database_option(ctx *Create_database_optionContext) {}

// ExitCreate_database_option is called when production create_database_option is exited.
func (s *BasetsqlListener) ExitCreate_database_option(ctx *Create_database_optionContext) {}

// EnterDatabase_filestream_option is called when production database_filestream_option is entered.
func (s *BasetsqlListener) EnterDatabase_filestream_option(ctx *Database_filestream_optionContext) {}

// ExitDatabase_filestream_option is called when production database_filestream_option is exited.
func (s *BasetsqlListener) ExitDatabase_filestream_option(ctx *Database_filestream_optionContext) {}

// EnterDatabase_file_spec is called when production database_file_spec is entered.
func (s *BasetsqlListener) EnterDatabase_file_spec(ctx *Database_file_specContext) {}

// ExitDatabase_file_spec is called when production database_file_spec is exited.
func (s *BasetsqlListener) ExitDatabase_file_spec(ctx *Database_file_specContext) {}

// EnterFile_group is called when production file_group is entered.
func (s *BasetsqlListener) EnterFile_group(ctx *File_groupContext) {}

// ExitFile_group is called when production file_group is exited.
func (s *BasetsqlListener) ExitFile_group(ctx *File_groupContext) {}

// EnterFile_spec is called when production file_spec is entered.
func (s *BasetsqlListener) EnterFile_spec(ctx *File_specContext) {}

// ExitFile_spec is called when production file_spec is exited.
func (s *BasetsqlListener) ExitFile_spec(ctx *File_specContext) {}

// EnterFull_table_name is called when production full_table_name is entered.
func (s *BasetsqlListener) EnterFull_table_name(ctx *Full_table_nameContext) {}

// ExitFull_table_name is called when production full_table_name is exited.
func (s *BasetsqlListener) ExitFull_table_name(ctx *Full_table_nameContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BasetsqlListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BasetsqlListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterSimple_name is called when production simple_name is entered.
func (s *BasetsqlListener) EnterSimple_name(ctx *Simple_nameContext) {}

// ExitSimple_name is called when production simple_name is exited.
func (s *BasetsqlListener) ExitSimple_name(ctx *Simple_nameContext) {}

// EnterFunc_proc_name is called when production func_proc_name is entered.
func (s *BasetsqlListener) EnterFunc_proc_name(ctx *Func_proc_nameContext) {}

// ExitFunc_proc_name is called when production func_proc_name is exited.
func (s *BasetsqlListener) ExitFunc_proc_name(ctx *Func_proc_nameContext) {}

// EnterDdl_object is called when production ddl_object is entered.
func (s *BasetsqlListener) EnterDdl_object(ctx *Ddl_objectContext) {}

// ExitDdl_object is called when production ddl_object is exited.
func (s *BasetsqlListener) ExitDdl_object(ctx *Ddl_objectContext) {}

// EnterFull_column_name is called when production full_column_name is entered.
func (s *BasetsqlListener) EnterFull_column_name(ctx *Full_column_nameContext) {}

// ExitFull_column_name is called when production full_column_name is exited.
func (s *BasetsqlListener) ExitFull_column_name(ctx *Full_column_nameContext) {}

// EnterColumn_name_list is called when production column_name_list is entered.
func (s *BasetsqlListener) EnterColumn_name_list(ctx *Column_name_listContext) {}

// ExitColumn_name_list is called when production column_name_list is exited.
func (s *BasetsqlListener) ExitColumn_name_list(ctx *Column_name_listContext) {}

// EnterCursor_name is called when production cursor_name is entered.
func (s *BasetsqlListener) EnterCursor_name(ctx *Cursor_nameContext) {}

// ExitCursor_name is called when production cursor_name is exited.
func (s *BasetsqlListener) ExitCursor_name(ctx *Cursor_nameContext) {}

// EnterOn_off is called when production on_off is entered.
func (s *BasetsqlListener) EnterOn_off(ctx *On_offContext) {}

// ExitOn_off is called when production on_off is exited.
func (s *BasetsqlListener) ExitOn_off(ctx *On_offContext) {}

// EnterClustered is called when production clustered is entered.
func (s *BasetsqlListener) EnterClustered(ctx *ClusteredContext) {}

// ExitClustered is called when production clustered is exited.
func (s *BasetsqlListener) ExitClustered(ctx *ClusteredContext) {}

// EnterNull_notnull is called when production null_notnull is entered.
func (s *BasetsqlListener) EnterNull_notnull(ctx *Null_notnullContext) {}

// ExitNull_notnull is called when production null_notnull is exited.
func (s *BasetsqlListener) ExitNull_notnull(ctx *Null_notnullContext) {}

// EnterScalar_function_name is called when production scalar_function_name is entered.
func (s *BasetsqlListener) EnterScalar_function_name(ctx *Scalar_function_nameContext) {}

// ExitScalar_function_name is called when production scalar_function_name is exited.
func (s *BasetsqlListener) ExitScalar_function_name(ctx *Scalar_function_nameContext) {}

// EnterData_type is called when production data_type is entered.
func (s *BasetsqlListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BasetsqlListener) ExitData_type(ctx *Data_typeContext) {}

// EnterDefault_value is called when production default_value is entered.
func (s *BasetsqlListener) EnterDefault_value(ctx *Default_valueContext) {}

// ExitDefault_value is called when production default_value is exited.
func (s *BasetsqlListener) ExitDefault_value(ctx *Default_valueContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasetsqlListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasetsqlListener) ExitConstant(ctx *ConstantContext) {}

// EnterSign is called when production sign is entered.
func (s *BasetsqlListener) EnterSign(ctx *SignContext) {}

// ExitSign is called when production sign is exited.
func (s *BasetsqlListener) ExitSign(ctx *SignContext) {}

// EnterR_id is called when production r_id is entered.
func (s *BasetsqlListener) EnterR_id(ctx *R_idContext) {}

// ExitR_id is called when production r_id is exited.
func (s *BasetsqlListener) ExitR_id(ctx *R_idContext) {}

// EnterSimple_id is called when production simple_id is entered.
func (s *BasetsqlListener) EnterSimple_id(ctx *Simple_idContext) {}

// ExitSimple_id is called when production simple_id is exited.
func (s *BasetsqlListener) ExitSimple_id(ctx *Simple_idContext) {}

// EnterComparison_operator is called when production comparison_operator is entered.
func (s *BasetsqlListener) EnterComparison_operator(ctx *Comparison_operatorContext) {}

// ExitComparison_operator is called when production comparison_operator is exited.
func (s *BasetsqlListener) ExitComparison_operator(ctx *Comparison_operatorContext) {}

// EnterAssignment_operator is called when production assignment_operator is entered.
func (s *BasetsqlListener) EnterAssignment_operator(ctx *Assignment_operatorContext) {}

// ExitAssignment_operator is called when production assignment_operator is exited.
func (s *BasetsqlListener) ExitAssignment_operator(ctx *Assignment_operatorContext) {}

// EnterFile_size is called when production file_size is entered.
func (s *BasetsqlListener) EnterFile_size(ctx *File_sizeContext) {}

// ExitFile_size is called when production file_size is exited.
func (s *BasetsqlListener) ExitFile_size(ctx *File_sizeContext) {}
