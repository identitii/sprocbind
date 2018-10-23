// Code generated from tsql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tsql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// tsqlListener is a complete listener for a parse tree produced by tsqlParser.
type tsqlListener interface {
	antlr.ParseTreeListener

	// EnterTsql_file is called when entering the tsql_file production.
	EnterTsql_file(c *Tsql_fileContext)

	// EnterBatch is called when entering the batch production.
	EnterBatch(c *BatchContext)

	// EnterSql_clauses is called when entering the sql_clauses production.
	EnterSql_clauses(c *Sql_clausesContext)

	// EnterSql_clause is called when entering the sql_clause production.
	EnterSql_clause(c *Sql_clauseContext)

	// EnterDml_clause is called when entering the dml_clause production.
	EnterDml_clause(c *Dml_clauseContext)

	// EnterDdl_clause is called when entering the ddl_clause production.
	EnterDdl_clause(c *Ddl_clauseContext)

	// EnterBlock_statement is called when entering the block_statement production.
	EnterBlock_statement(c *Block_statementContext)

	// EnterBreak_statement is called when entering the break_statement production.
	EnterBreak_statement(c *Break_statementContext)

	// EnterContinue_statement is called when entering the continue_statement production.
	EnterContinue_statement(c *Continue_statementContext)

	// EnterGoto_statement is called when entering the goto_statement production.
	EnterGoto_statement(c *Goto_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterThrow_statement is called when entering the throw_statement production.
	EnterThrow_statement(c *Throw_statementContext)

	// EnterTry_catch_statement is called when entering the try_catch_statement production.
	EnterTry_catch_statement(c *Try_catch_statementContext)

	// EnterWaitfor_statement is called when entering the waitfor_statement production.
	EnterWaitfor_statement(c *Waitfor_statementContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterPrint_statement is called when entering the print_statement production.
	EnterPrint_statement(c *Print_statementContext)

	// EnterRaiseerror_statement is called when entering the raiseerror_statement production.
	EnterRaiseerror_statement(c *Raiseerror_statementContext)

	// EnterAnother_statement is called when entering the another_statement production.
	EnterAnother_statement(c *Another_statementContext)

	// EnterDelete_statement is called when entering the delete_statement production.
	EnterDelete_statement(c *Delete_statementContext)

	// EnterDelete_statement_from is called when entering the delete_statement_from production.
	EnterDelete_statement_from(c *Delete_statement_fromContext)

	// EnterInsert_statement is called when entering the insert_statement production.
	EnterInsert_statement(c *Insert_statementContext)

	// EnterInsert_statement_value is called when entering the insert_statement_value production.
	EnterInsert_statement_value(c *Insert_statement_valueContext)

	// EnterSelect_statement is called when entering the select_statement production.
	EnterSelect_statement(c *Select_statementContext)

	// EnterUpdate_statement is called when entering the update_statement production.
	EnterUpdate_statement(c *Update_statementContext)

	// EnterWhere_clause_dml is called when entering the where_clause_dml production.
	EnterWhere_clause_dml(c *Where_clause_dmlContext)

	// EnterOutput_clause is called when entering the output_clause production.
	EnterOutput_clause(c *Output_clauseContext)

	// EnterOutput_dml_list_elem is called when entering the output_dml_list_elem production.
	EnterOutput_dml_list_elem(c *Output_dml_list_elemContext)

	// EnterOutput_column_name is called when entering the output_column_name production.
	EnterOutput_column_name(c *Output_column_nameContext)

	// EnterCreate_database is called when entering the create_database production.
	EnterCreate_database(c *Create_databaseContext)

	// EnterCreate_index is called when entering the create_index production.
	EnterCreate_index(c *Create_indexContext)

	// EnterCreate_procedure is called when entering the create_procedure production.
	EnterCreate_procedure(c *Create_procedureContext)

	// EnterProcedure_param is called when entering the procedure_param production.
	EnterProcedure_param(c *Procedure_paramContext)

	// EnterProcedure_option is called when entering the procedure_option production.
	EnterProcedure_option(c *Procedure_optionContext)

	// EnterCreate_statistics is called when entering the create_statistics production.
	EnterCreate_statistics(c *Create_statisticsContext)

	// EnterCreate_table is called when entering the create_table production.
	EnterCreate_table(c *Create_tableContext)

	// EnterCreate_view is called when entering the create_view production.
	EnterCreate_view(c *Create_viewContext)

	// EnterView_attribute is called when entering the view_attribute production.
	EnterView_attribute(c *View_attributeContext)

	// EnterAlter_table is called when entering the alter_table production.
	EnterAlter_table(c *Alter_tableContext)

	// EnterAlter_database is called when entering the alter_database production.
	EnterAlter_database(c *Alter_databaseContext)

	// EnterDatabase_optionspec is called when entering the database_optionspec production.
	EnterDatabase_optionspec(c *Database_optionspecContext)

	// EnterAuto_option is called when entering the auto_option production.
	EnterAuto_option(c *Auto_optionContext)

	// EnterChange_tracking_option is called when entering the change_tracking_option production.
	EnterChange_tracking_option(c *Change_tracking_optionContext)

	// EnterChange_tracking_option_list is called when entering the change_tracking_option_list production.
	EnterChange_tracking_option_list(c *Change_tracking_option_listContext)

	// EnterContainment_option is called when entering the containment_option production.
	EnterContainment_option(c *Containment_optionContext)

	// EnterCursor_option is called when entering the cursor_option production.
	EnterCursor_option(c *Cursor_optionContext)

	// EnterDate_correlation_optimization_option is called when entering the date_correlation_optimization_option production.
	EnterDate_correlation_optimization_option(c *Date_correlation_optimization_optionContext)

	// EnterDb_encryption_option is called when entering the db_encryption_option production.
	EnterDb_encryption_option(c *Db_encryption_optionContext)

	// EnterDb_state_option is called when entering the db_state_option production.
	EnterDb_state_option(c *Db_state_optionContext)

	// EnterDb_update_option is called when entering the db_update_option production.
	EnterDb_update_option(c *Db_update_optionContext)

	// EnterDb_user_access_option is called when entering the db_user_access_option production.
	EnterDb_user_access_option(c *Db_user_access_optionContext)

	// EnterDelayed_durability_option is called when entering the delayed_durability_option production.
	EnterDelayed_durability_option(c *Delayed_durability_optionContext)

	// EnterExternal_access_option is called when entering the external_access_option production.
	EnterExternal_access_option(c *External_access_optionContext)

	// EnterMixed_page_allocation_option is called when entering the mixed_page_allocation_option production.
	EnterMixed_page_allocation_option(c *Mixed_page_allocation_optionContext)

	// EnterParameterization_option is called when entering the parameterization_option production.
	EnterParameterization_option(c *Parameterization_optionContext)

	// EnterRecovery_option is called when entering the recovery_option production.
	EnterRecovery_option(c *Recovery_optionContext)

	// EnterService_broker_option is called when entering the service_broker_option production.
	EnterService_broker_option(c *Service_broker_optionContext)

	// EnterSnapshot_option is called when entering the snapshot_option production.
	EnterSnapshot_option(c *Snapshot_optionContext)

	// EnterSql_option is called when entering the sql_option production.
	EnterSql_option(c *Sql_optionContext)

	// EnterTarget_recovery_time_option is called when entering the target_recovery_time_option production.
	EnterTarget_recovery_time_option(c *Target_recovery_time_optionContext)

	// EnterTermination is called when entering the termination production.
	EnterTermination(c *TerminationContext)

	// EnterDrop_index is called when entering the drop_index production.
	EnterDrop_index(c *Drop_indexContext)

	// EnterDrop_procedure is called when entering the drop_procedure production.
	EnterDrop_procedure(c *Drop_procedureContext)

	// EnterDrop_statistics is called when entering the drop_statistics production.
	EnterDrop_statistics(c *Drop_statisticsContext)

	// EnterDrop_table is called when entering the drop_table production.
	EnterDrop_table(c *Drop_tableContext)

	// EnterDrop_view is called when entering the drop_view production.
	EnterDrop_view(c *Drop_viewContext)

	// EnterCreate_type is called when entering the create_type production.
	EnterCreate_type(c *Create_typeContext)

	// EnterDrop_type is called when entering the drop_type production.
	EnterDrop_type(c *Drop_typeContext)

	// EnterRowset_function_limited is called when entering the rowset_function_limited production.
	EnterRowset_function_limited(c *Rowset_function_limitedContext)

	// EnterOpenquery is called when entering the openquery production.
	EnterOpenquery(c *OpenqueryContext)

	// EnterOpendatasource is called when entering the opendatasource production.
	EnterOpendatasource(c *OpendatasourceContext)

	// EnterDeclare_statement is called when entering the declare_statement production.
	EnterDeclare_statement(c *Declare_statementContext)

	// EnterCursor_statement is called when entering the cursor_statement production.
	EnterCursor_statement(c *Cursor_statementContext)

	// EnterExecute_statement is called when entering the execute_statement production.
	EnterExecute_statement(c *Execute_statementContext)

	// EnterExecute_statement_arg is called when entering the execute_statement_arg production.
	EnterExecute_statement_arg(c *Execute_statement_argContext)

	// EnterExecute_var_string is called when entering the execute_var_string production.
	EnterExecute_var_string(c *Execute_var_stringContext)

	// EnterSecurity_statement is called when entering the security_statement production.
	EnterSecurity_statement(c *Security_statementContext)

	// EnterGrant_permission is called when entering the grant_permission production.
	EnterGrant_permission(c *Grant_permissionContext)

	// EnterSet_statement is called when entering the set_statement production.
	EnterSet_statement(c *Set_statementContext)

	// EnterTransaction_statement is called when entering the transaction_statement production.
	EnterTransaction_statement(c *Transaction_statementContext)

	// EnterGo_statement is called when entering the go_statement production.
	EnterGo_statement(c *Go_statementContext)

	// EnterUse_statement is called when entering the use_statement production.
	EnterUse_statement(c *Use_statementContext)

	// EnterExecute_clause is called when entering the execute_clause production.
	EnterExecute_clause(c *Execute_clauseContext)

	// EnterDeclare_local is called when entering the declare_local production.
	EnterDeclare_local(c *Declare_localContext)

	// EnterTable_type_definition is called when entering the table_type_definition production.
	EnterTable_type_definition(c *Table_type_definitionContext)

	// EnterColumn_def_table_constraints is called when entering the column_def_table_constraints production.
	EnterColumn_def_table_constraints(c *Column_def_table_constraintsContext)

	// EnterColumn_def_table_constraint is called when entering the column_def_table_constraint production.
	EnterColumn_def_table_constraint(c *Column_def_table_constraintContext)

	// EnterColumn_definition is called when entering the column_definition production.
	EnterColumn_definition(c *Column_definitionContext)

	// EnterColumn_constraint is called when entering the column_constraint production.
	EnterColumn_constraint(c *Column_constraintContext)

	// EnterTable_constraint is called when entering the table_constraint production.
	EnterTable_constraint(c *Table_constraintContext)

	// EnterIndex_options is called when entering the index_options production.
	EnterIndex_options(c *Index_optionsContext)

	// EnterIndex_option is called when entering the index_option production.
	EnterIndex_option(c *Index_optionContext)

	// EnterDeclare_cursor is called when entering the declare_cursor production.
	EnterDeclare_cursor(c *Declare_cursorContext)

	// EnterDeclare_set_cursor_common is called when entering the declare_set_cursor_common production.
	EnterDeclare_set_cursor_common(c *Declare_set_cursor_commonContext)

	// EnterFetch_cursor is called when entering the fetch_cursor production.
	EnterFetch_cursor(c *Fetch_cursorContext)

	// EnterSet_special is called when entering the set_special production.
	EnterSet_special(c *Set_specialContext)

	// EnterConstant_LOCAL_ID is called when entering the constant_LOCAL_ID production.
	EnterConstant_LOCAL_ID(c *Constant_LOCAL_IDContext)

	// EnterBinary_operator_expression is called when entering the binary_operator_expression production.
	EnterBinary_operator_expression(c *Binary_operator_expressionContext)

	// EnterPrimitive_expression is called when entering the primitive_expression production.
	EnterPrimitive_expression(c *Primitive_expressionContext)

	// EnterBracket_expression is called when entering the bracket_expression production.
	EnterBracket_expression(c *Bracket_expressionContext)

	// EnterUnary_operator_expression is called when entering the unary_operator_expression production.
	EnterUnary_operator_expression(c *Unary_operator_expressionContext)

	// EnterFunction_call_expression is called when entering the function_call_expression production.
	EnterFunction_call_expression(c *Function_call_expressionContext)

	// EnterCase_expression is called when entering the case_expression production.
	EnterCase_expression(c *Case_expressionContext)

	// EnterColumn_ref_expression is called when entering the column_ref_expression production.
	EnterColumn_ref_expression(c *Column_ref_expressionContext)

	// EnterSubquery_expression is called when entering the subquery_expression production.
	EnterSubquery_expression(c *Subquery_expressionContext)

	// EnterOver_clause_expression is called when entering the over_clause_expression production.
	EnterOver_clause_expression(c *Over_clause_expressionContext)

	// EnterConstant_expression is called when entering the constant_expression production.
	EnterConstant_expression(c *Constant_expressionContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterWith_expression is called when entering the with_expression production.
	EnterWith_expression(c *With_expressionContext)

	// EnterCommon_table_expression is called when entering the common_table_expression production.
	EnterCommon_table_expression(c *Common_table_expressionContext)

	// EnterUpdate_elem is called when entering the update_elem production.
	EnterUpdate_elem(c *Update_elemContext)

	// EnterSearch_condition_list is called when entering the search_condition_list production.
	EnterSearch_condition_list(c *Search_condition_listContext)

	// EnterSearch_cond_or is called when entering the search_cond_or production.
	EnterSearch_cond_or(c *Search_cond_orContext)

	// EnterSearch_cond_pred is called when entering the search_cond_pred production.
	EnterSearch_cond_pred(c *Search_cond_predContext)

	// EnterSearch_cond_and is called when entering the search_cond_and production.
	EnterSearch_cond_and(c *Search_cond_andContext)

	// EnterUnary_operator_expression3 is called when entering the unary_operator_expression3 production.
	EnterUnary_operator_expression3(c *Unary_operator_expression3Context)

	// EnterUnary_operator_expression2 is called when entering the unary_operator_expression2 production.
	EnterUnary_operator_expression2(c *Unary_operator_expression2Context)

	// EnterBinary_operator_expression2 is called when entering the binary_operator_expression2 production.
	EnterBinary_operator_expression2(c *Binary_operator_expression2Context)

	// EnterSublink_expression is called when entering the sublink_expression production.
	EnterSublink_expression(c *Sublink_expressionContext)

	// EnterBinary_mod_expression is called when entering the binary_mod_expression production.
	EnterBinary_mod_expression(c *Binary_mod_expressionContext)

	// EnterBinary_in_expression is called when entering the binary_in_expression production.
	EnterBinary_in_expression(c *Binary_in_expressionContext)

	// EnterBracket_search_expression is called when entering the bracket_search_expression production.
	EnterBracket_search_expression(c *Bracket_search_expressionContext)

	// EnterDecimal_expression is called when entering the decimal_expression production.
	EnterDecimal_expression(c *Decimal_expressionContext)

	// EnterBracket_query_expression is called when entering the bracket_query_expression production.
	EnterBracket_query_expression(c *Bracket_query_expressionContext)

	// EnterQuery_specification_expression is called when entering the query_specification_expression production.
	EnterQuery_specification_expression(c *Query_specification_expressionContext)

	// EnterUnion_query_expression is called when entering the union_query_expression production.
	EnterUnion_query_expression(c *Union_query_expressionContext)

	// EnterUnion_op is called when entering the union_op production.
	EnterUnion_op(c *Union_opContext)

	// EnterQuery_specification is called when entering the query_specification production.
	EnterQuery_specification(c *Query_specificationContext)

	// EnterTop_clause is called when entering the top_clause production.
	EnterTop_clause(c *Top_clauseContext)

	// EnterTop_clause_dm is called when entering the top_clause_dm production.
	EnterTop_clause_dm(c *Top_clause_dmContext)

	// EnterOrder_by_clause is called when entering the order_by_clause production.
	EnterOrder_by_clause(c *Order_by_clauseContext)

	// EnterFetch_expression is called when entering the fetch_expression production.
	EnterFetch_expression(c *Fetch_expressionContext)

	// EnterFor_clause is called when entering the for_clause production.
	EnterFor_clause(c *For_clauseContext)

	// EnterXml_common_directives is called when entering the xml_common_directives production.
	EnterXml_common_directives(c *Xml_common_directivesContext)

	// EnterOrder_by_expression is called when entering the order_by_expression production.
	EnterOrder_by_expression(c *Order_by_expressionContext)

	// EnterGroup_by_item is called when entering the group_by_item production.
	EnterGroup_by_item(c *Group_by_itemContext)

	// EnterOption_clause is called when entering the option_clause production.
	EnterOption_clause(c *Option_clauseContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterOptimize_for_arg is called when entering the optimize_for_arg production.
	EnterOptimize_for_arg(c *Optimize_for_argContext)

	// EnterSelect_list is called when entering the select_list production.
	EnterSelect_list(c *Select_listContext)

	// EnterSelect_list_elem is called when entering the select_list_elem production.
	EnterSelect_list_elem(c *Select_list_elemContext)

	// EnterTable_sources is called when entering the table_sources production.
	EnterTable_sources(c *Table_sourcesContext)

	// EnterCross_join is called when entering the cross_join production.
	EnterCross_join(c *Cross_joinContext)

	// EnterTable_source_item_join is called when entering the table_source_item_join production.
	EnterTable_source_item_join(c *Table_source_item_joinContext)

	// EnterStandard_join is called when entering the standard_join production.
	EnterStandard_join(c *Standard_joinContext)

	// EnterApply_join is called when entering the apply_join production.
	EnterApply_join(c *Apply_joinContext)

	// EnterBracket_table_source is called when entering the bracket_table_source production.
	EnterBracket_table_source(c *Bracket_table_sourceContext)

	// EnterTable_source_item_name is called when entering the table_source_item_name production.
	EnterTable_source_item_name(c *Table_source_item_nameContext)

	// EnterTable_source_item_simple is called when entering the table_source_item_simple production.
	EnterTable_source_item_simple(c *Table_source_item_simpleContext)

	// EnterTable_source_item_complex is called when entering the table_source_item_complex production.
	EnterTable_source_item_complex(c *Table_source_item_complexContext)

	// EnterTable_alias is called when entering the table_alias production.
	EnterTable_alias(c *Table_aliasContext)

	// EnterChange_table is called when entering the change_table production.
	EnterChange_table(c *Change_tableContext)

	// EnterJoin_type is called when entering the join_type production.
	EnterJoin_type(c *Join_typeContext)

	// EnterTable_name_with_hint is called when entering the table_name_with_hint production.
	EnterTable_name_with_hint(c *Table_name_with_hintContext)

	// EnterRowset_function is called when entering the rowset_function production.
	EnterRowset_function(c *Rowset_functionContext)

	// EnterBulk_option is called when entering the bulk_option production.
	EnterBulk_option(c *Bulk_optionContext)

	// EnterDerived_table is called when entering the derived_table production.
	EnterDerived_table(c *Derived_tableContext)

	// EnterRank_call is called when entering the rank_call production.
	EnterRank_call(c *Rank_callContext)

	// EnterAggregate_call is called when entering the aggregate_call production.
	EnterAggregate_call(c *Aggregate_callContext)

	// EnterStandard_call is called when entering the standard_call production.
	EnterStandard_call(c *Standard_callContext)

	// EnterNvf_call is called when entering the nvf_call production.
	EnterNvf_call(c *Nvf_callContext)

	// EnterCast_call is called when entering the cast_call production.
	EnterCast_call(c *Cast_callContext)

	// EnterSimple_call is called when entering the simple_call production.
	EnterSimple_call(c *Simple_callContext)

	// EnterSwitch_section is called when entering the switch_section production.
	EnterSwitch_section(c *Switch_sectionContext)

	// EnterSwitch_search_condition_section is called when entering the switch_search_condition_section production.
	EnterSwitch_search_condition_section(c *Switch_search_condition_sectionContext)

	// EnterWith_table_hints is called when entering the with_table_hints production.
	EnterWith_table_hints(c *With_table_hintsContext)

	// EnterInsert_with_table_hints is called when entering the insert_with_table_hints production.
	EnterInsert_with_table_hints(c *Insert_with_table_hintsContext)

	// EnterTable_hint is called when entering the table_hint production.
	EnterTable_hint(c *Table_hintContext)

	// EnterIndex_value is called when entering the index_value production.
	EnterIndex_value(c *Index_valueContext)

	// EnterColumn_alias_list is called when entering the column_alias_list production.
	EnterColumn_alias_list(c *Column_alias_listContext)

	// EnterColumn_alias is called when entering the column_alias production.
	EnterColumn_alias(c *Column_aliasContext)

	// EnterA_star is called when entering the a_star production.
	EnterA_star(c *A_starContext)

	// EnterTable_value_constructor is called when entering the table_value_constructor production.
	EnterTable_value_constructor(c *Table_value_constructorContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterValue_list is called when entering the value_list production.
	EnterValue_list(c *Value_listContext)

	// EnterNext_value_for is called when entering the next_value_for production.
	EnterNext_value_for(c *Next_value_forContext)

	// EnterNext_value_for_function is called when entering the next_value_for_function production.
	EnterNext_value_for_function(c *Next_value_for_functionContext)

	// EnterRanking_windowed_function is called when entering the ranking_windowed_function production.
	EnterRanking_windowed_function(c *Ranking_windowed_functionContext)

	// EnterAggregate_windowed_function is called when entering the aggregate_windowed_function production.
	EnterAggregate_windowed_function(c *Aggregate_windowed_functionContext)

	// EnterAll_distinct is called when entering the all_distinct production.
	EnterAll_distinct(c *All_distinctContext)

	// EnterOver_clause is called when entering the over_clause production.
	EnterOver_clause(c *Over_clauseContext)

	// EnterRow_or_range_clause is called when entering the row_or_range_clause production.
	EnterRow_or_range_clause(c *Row_or_range_clauseContext)

	// EnterWindow_frame_extent is called when entering the window_frame_extent production.
	EnterWindow_frame_extent(c *Window_frame_extentContext)

	// EnterWindow_frame_bound is called when entering the window_frame_bound production.
	EnterWindow_frame_bound(c *Window_frame_boundContext)

	// EnterWindow_frame_preceding is called when entering the window_frame_preceding production.
	EnterWindow_frame_preceding(c *Window_frame_precedingContext)

	// EnterWindow_frame_following is called when entering the window_frame_following production.
	EnterWindow_frame_following(c *Window_frame_followingContext)

	// EnterCreate_database_option is called when entering the create_database_option production.
	EnterCreate_database_option(c *Create_database_optionContext)

	// EnterDatabase_filestream_option is called when entering the database_filestream_option production.
	EnterDatabase_filestream_option(c *Database_filestream_optionContext)

	// EnterDatabase_file_spec is called when entering the database_file_spec production.
	EnterDatabase_file_spec(c *Database_file_specContext)

	// EnterFile_group is called when entering the file_group production.
	EnterFile_group(c *File_groupContext)

	// EnterFile_spec is called when entering the file_spec production.
	EnterFile_spec(c *File_specContext)

	// EnterFull_table_name is called when entering the full_table_name production.
	EnterFull_table_name(c *Full_table_nameContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterSimple_name is called when entering the simple_name production.
	EnterSimple_name(c *Simple_nameContext)

	// EnterFunc_proc_name is called when entering the func_proc_name production.
	EnterFunc_proc_name(c *Func_proc_nameContext)

	// EnterDdl_object is called when entering the ddl_object production.
	EnterDdl_object(c *Ddl_objectContext)

	// EnterFull_column_name is called when entering the full_column_name production.
	EnterFull_column_name(c *Full_column_nameContext)

	// EnterColumn_name_list is called when entering the column_name_list production.
	EnterColumn_name_list(c *Column_name_listContext)

	// EnterCursor_name is called when entering the cursor_name production.
	EnterCursor_name(c *Cursor_nameContext)

	// EnterOn_off is called when entering the on_off production.
	EnterOn_off(c *On_offContext)

	// EnterClustered is called when entering the clustered production.
	EnterClustered(c *ClusteredContext)

	// EnterNull_notnull is called when entering the null_notnull production.
	EnterNull_notnull(c *Null_notnullContext)

	// EnterScalar_function_name is called when entering the scalar_function_name production.
	EnterScalar_function_name(c *Scalar_function_nameContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterDefault_value is called when entering the default_value production.
	EnterDefault_value(c *Default_valueContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterSign is called when entering the sign production.
	EnterSign(c *SignContext)

	// EnterR_id is called when entering the r_id production.
	EnterR_id(c *R_idContext)

	// EnterSimple_id is called when entering the simple_id production.
	EnterSimple_id(c *Simple_idContext)

	// EnterComparison_operator is called when entering the comparison_operator production.
	EnterComparison_operator(c *Comparison_operatorContext)

	// EnterAssignment_operator is called when entering the assignment_operator production.
	EnterAssignment_operator(c *Assignment_operatorContext)

	// EnterFile_size is called when entering the file_size production.
	EnterFile_size(c *File_sizeContext)

	// ExitTsql_file is called when exiting the tsql_file production.
	ExitTsql_file(c *Tsql_fileContext)

	// ExitBatch is called when exiting the batch production.
	ExitBatch(c *BatchContext)

	// ExitSql_clauses is called when exiting the sql_clauses production.
	ExitSql_clauses(c *Sql_clausesContext)

	// ExitSql_clause is called when exiting the sql_clause production.
	ExitSql_clause(c *Sql_clauseContext)

	// ExitDml_clause is called when exiting the dml_clause production.
	ExitDml_clause(c *Dml_clauseContext)

	// ExitDdl_clause is called when exiting the ddl_clause production.
	ExitDdl_clause(c *Ddl_clauseContext)

	// ExitBlock_statement is called when exiting the block_statement production.
	ExitBlock_statement(c *Block_statementContext)

	// ExitBreak_statement is called when exiting the break_statement production.
	ExitBreak_statement(c *Break_statementContext)

	// ExitContinue_statement is called when exiting the continue_statement production.
	ExitContinue_statement(c *Continue_statementContext)

	// ExitGoto_statement is called when exiting the goto_statement production.
	ExitGoto_statement(c *Goto_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitThrow_statement is called when exiting the throw_statement production.
	ExitThrow_statement(c *Throw_statementContext)

	// ExitTry_catch_statement is called when exiting the try_catch_statement production.
	ExitTry_catch_statement(c *Try_catch_statementContext)

	// ExitWaitfor_statement is called when exiting the waitfor_statement production.
	ExitWaitfor_statement(c *Waitfor_statementContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitPrint_statement is called when exiting the print_statement production.
	ExitPrint_statement(c *Print_statementContext)

	// ExitRaiseerror_statement is called when exiting the raiseerror_statement production.
	ExitRaiseerror_statement(c *Raiseerror_statementContext)

	// ExitAnother_statement is called when exiting the another_statement production.
	ExitAnother_statement(c *Another_statementContext)

	// ExitDelete_statement is called when exiting the delete_statement production.
	ExitDelete_statement(c *Delete_statementContext)

	// ExitDelete_statement_from is called when exiting the delete_statement_from production.
	ExitDelete_statement_from(c *Delete_statement_fromContext)

	// ExitInsert_statement is called when exiting the insert_statement production.
	ExitInsert_statement(c *Insert_statementContext)

	// ExitInsert_statement_value is called when exiting the insert_statement_value production.
	ExitInsert_statement_value(c *Insert_statement_valueContext)

	// ExitSelect_statement is called when exiting the select_statement production.
	ExitSelect_statement(c *Select_statementContext)

	// ExitUpdate_statement is called when exiting the update_statement production.
	ExitUpdate_statement(c *Update_statementContext)

	// ExitWhere_clause_dml is called when exiting the where_clause_dml production.
	ExitWhere_clause_dml(c *Where_clause_dmlContext)

	// ExitOutput_clause is called when exiting the output_clause production.
	ExitOutput_clause(c *Output_clauseContext)

	// ExitOutput_dml_list_elem is called when exiting the output_dml_list_elem production.
	ExitOutput_dml_list_elem(c *Output_dml_list_elemContext)

	// ExitOutput_column_name is called when exiting the output_column_name production.
	ExitOutput_column_name(c *Output_column_nameContext)

	// ExitCreate_database is called when exiting the create_database production.
	ExitCreate_database(c *Create_databaseContext)

	// ExitCreate_index is called when exiting the create_index production.
	ExitCreate_index(c *Create_indexContext)

	// ExitCreate_procedure is called when exiting the create_procedure production.
	ExitCreate_procedure(c *Create_procedureContext)

	// ExitProcedure_param is called when exiting the procedure_param production.
	ExitProcedure_param(c *Procedure_paramContext)

	// ExitProcedure_option is called when exiting the procedure_option production.
	ExitProcedure_option(c *Procedure_optionContext)

	// ExitCreate_statistics is called when exiting the create_statistics production.
	ExitCreate_statistics(c *Create_statisticsContext)

	// ExitCreate_table is called when exiting the create_table production.
	ExitCreate_table(c *Create_tableContext)

	// ExitCreate_view is called when exiting the create_view production.
	ExitCreate_view(c *Create_viewContext)

	// ExitView_attribute is called when exiting the view_attribute production.
	ExitView_attribute(c *View_attributeContext)

	// ExitAlter_table is called when exiting the alter_table production.
	ExitAlter_table(c *Alter_tableContext)

	// ExitAlter_database is called when exiting the alter_database production.
	ExitAlter_database(c *Alter_databaseContext)

	// ExitDatabase_optionspec is called when exiting the database_optionspec production.
	ExitDatabase_optionspec(c *Database_optionspecContext)

	// ExitAuto_option is called when exiting the auto_option production.
	ExitAuto_option(c *Auto_optionContext)

	// ExitChange_tracking_option is called when exiting the change_tracking_option production.
	ExitChange_tracking_option(c *Change_tracking_optionContext)

	// ExitChange_tracking_option_list is called when exiting the change_tracking_option_list production.
	ExitChange_tracking_option_list(c *Change_tracking_option_listContext)

	// ExitContainment_option is called when exiting the containment_option production.
	ExitContainment_option(c *Containment_optionContext)

	// ExitCursor_option is called when exiting the cursor_option production.
	ExitCursor_option(c *Cursor_optionContext)

	// ExitDate_correlation_optimization_option is called when exiting the date_correlation_optimization_option production.
	ExitDate_correlation_optimization_option(c *Date_correlation_optimization_optionContext)

	// ExitDb_encryption_option is called when exiting the db_encryption_option production.
	ExitDb_encryption_option(c *Db_encryption_optionContext)

	// ExitDb_state_option is called when exiting the db_state_option production.
	ExitDb_state_option(c *Db_state_optionContext)

	// ExitDb_update_option is called when exiting the db_update_option production.
	ExitDb_update_option(c *Db_update_optionContext)

	// ExitDb_user_access_option is called when exiting the db_user_access_option production.
	ExitDb_user_access_option(c *Db_user_access_optionContext)

	// ExitDelayed_durability_option is called when exiting the delayed_durability_option production.
	ExitDelayed_durability_option(c *Delayed_durability_optionContext)

	// ExitExternal_access_option is called when exiting the external_access_option production.
	ExitExternal_access_option(c *External_access_optionContext)

	// ExitMixed_page_allocation_option is called when exiting the mixed_page_allocation_option production.
	ExitMixed_page_allocation_option(c *Mixed_page_allocation_optionContext)

	// ExitParameterization_option is called when exiting the parameterization_option production.
	ExitParameterization_option(c *Parameterization_optionContext)

	// ExitRecovery_option is called when exiting the recovery_option production.
	ExitRecovery_option(c *Recovery_optionContext)

	// ExitService_broker_option is called when exiting the service_broker_option production.
	ExitService_broker_option(c *Service_broker_optionContext)

	// ExitSnapshot_option is called when exiting the snapshot_option production.
	ExitSnapshot_option(c *Snapshot_optionContext)

	// ExitSql_option is called when exiting the sql_option production.
	ExitSql_option(c *Sql_optionContext)

	// ExitTarget_recovery_time_option is called when exiting the target_recovery_time_option production.
	ExitTarget_recovery_time_option(c *Target_recovery_time_optionContext)

	// ExitTermination is called when exiting the termination production.
	ExitTermination(c *TerminationContext)

	// ExitDrop_index is called when exiting the drop_index production.
	ExitDrop_index(c *Drop_indexContext)

	// ExitDrop_procedure is called when exiting the drop_procedure production.
	ExitDrop_procedure(c *Drop_procedureContext)

	// ExitDrop_statistics is called when exiting the drop_statistics production.
	ExitDrop_statistics(c *Drop_statisticsContext)

	// ExitDrop_table is called when exiting the drop_table production.
	ExitDrop_table(c *Drop_tableContext)

	// ExitDrop_view is called when exiting the drop_view production.
	ExitDrop_view(c *Drop_viewContext)

	// ExitCreate_type is called when exiting the create_type production.
	ExitCreate_type(c *Create_typeContext)

	// ExitDrop_type is called when exiting the drop_type production.
	ExitDrop_type(c *Drop_typeContext)

	// ExitRowset_function_limited is called when exiting the rowset_function_limited production.
	ExitRowset_function_limited(c *Rowset_function_limitedContext)

	// ExitOpenquery is called when exiting the openquery production.
	ExitOpenquery(c *OpenqueryContext)

	// ExitOpendatasource is called when exiting the opendatasource production.
	ExitOpendatasource(c *OpendatasourceContext)

	// ExitDeclare_statement is called when exiting the declare_statement production.
	ExitDeclare_statement(c *Declare_statementContext)

	// ExitCursor_statement is called when exiting the cursor_statement production.
	ExitCursor_statement(c *Cursor_statementContext)

	// ExitExecute_statement is called when exiting the execute_statement production.
	ExitExecute_statement(c *Execute_statementContext)

	// ExitExecute_statement_arg is called when exiting the execute_statement_arg production.
	ExitExecute_statement_arg(c *Execute_statement_argContext)

	// ExitExecute_var_string is called when exiting the execute_var_string production.
	ExitExecute_var_string(c *Execute_var_stringContext)

	// ExitSecurity_statement is called when exiting the security_statement production.
	ExitSecurity_statement(c *Security_statementContext)

	// ExitGrant_permission is called when exiting the grant_permission production.
	ExitGrant_permission(c *Grant_permissionContext)

	// ExitSet_statement is called when exiting the set_statement production.
	ExitSet_statement(c *Set_statementContext)

	// ExitTransaction_statement is called when exiting the transaction_statement production.
	ExitTransaction_statement(c *Transaction_statementContext)

	// ExitGo_statement is called when exiting the go_statement production.
	ExitGo_statement(c *Go_statementContext)

	// ExitUse_statement is called when exiting the use_statement production.
	ExitUse_statement(c *Use_statementContext)

	// ExitExecute_clause is called when exiting the execute_clause production.
	ExitExecute_clause(c *Execute_clauseContext)

	// ExitDeclare_local is called when exiting the declare_local production.
	ExitDeclare_local(c *Declare_localContext)

	// ExitTable_type_definition is called when exiting the table_type_definition production.
	ExitTable_type_definition(c *Table_type_definitionContext)

	// ExitColumn_def_table_constraints is called when exiting the column_def_table_constraints production.
	ExitColumn_def_table_constraints(c *Column_def_table_constraintsContext)

	// ExitColumn_def_table_constraint is called when exiting the column_def_table_constraint production.
	ExitColumn_def_table_constraint(c *Column_def_table_constraintContext)

	// ExitColumn_definition is called when exiting the column_definition production.
	ExitColumn_definition(c *Column_definitionContext)

	// ExitColumn_constraint is called when exiting the column_constraint production.
	ExitColumn_constraint(c *Column_constraintContext)

	// ExitTable_constraint is called when exiting the table_constraint production.
	ExitTable_constraint(c *Table_constraintContext)

	// ExitIndex_options is called when exiting the index_options production.
	ExitIndex_options(c *Index_optionsContext)

	// ExitIndex_option is called when exiting the index_option production.
	ExitIndex_option(c *Index_optionContext)

	// ExitDeclare_cursor is called when exiting the declare_cursor production.
	ExitDeclare_cursor(c *Declare_cursorContext)

	// ExitDeclare_set_cursor_common is called when exiting the declare_set_cursor_common production.
	ExitDeclare_set_cursor_common(c *Declare_set_cursor_commonContext)

	// ExitFetch_cursor is called when exiting the fetch_cursor production.
	ExitFetch_cursor(c *Fetch_cursorContext)

	// ExitSet_special is called when exiting the set_special production.
	ExitSet_special(c *Set_specialContext)

	// ExitConstant_LOCAL_ID is called when exiting the constant_LOCAL_ID production.
	ExitConstant_LOCAL_ID(c *Constant_LOCAL_IDContext)

	// ExitBinary_operator_expression is called when exiting the binary_operator_expression production.
	ExitBinary_operator_expression(c *Binary_operator_expressionContext)

	// ExitPrimitive_expression is called when exiting the primitive_expression production.
	ExitPrimitive_expression(c *Primitive_expressionContext)

	// ExitBracket_expression is called when exiting the bracket_expression production.
	ExitBracket_expression(c *Bracket_expressionContext)

	// ExitUnary_operator_expression is called when exiting the unary_operator_expression production.
	ExitUnary_operator_expression(c *Unary_operator_expressionContext)

	// ExitFunction_call_expression is called when exiting the function_call_expression production.
	ExitFunction_call_expression(c *Function_call_expressionContext)

	// ExitCase_expression is called when exiting the case_expression production.
	ExitCase_expression(c *Case_expressionContext)

	// ExitColumn_ref_expression is called when exiting the column_ref_expression production.
	ExitColumn_ref_expression(c *Column_ref_expressionContext)

	// ExitSubquery_expression is called when exiting the subquery_expression production.
	ExitSubquery_expression(c *Subquery_expressionContext)

	// ExitOver_clause_expression is called when exiting the over_clause_expression production.
	ExitOver_clause_expression(c *Over_clause_expressionContext)

	// ExitConstant_expression is called when exiting the constant_expression production.
	ExitConstant_expression(c *Constant_expressionContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitWith_expression is called when exiting the with_expression production.
	ExitWith_expression(c *With_expressionContext)

	// ExitCommon_table_expression is called when exiting the common_table_expression production.
	ExitCommon_table_expression(c *Common_table_expressionContext)

	// ExitUpdate_elem is called when exiting the update_elem production.
	ExitUpdate_elem(c *Update_elemContext)

	// ExitSearch_condition_list is called when exiting the search_condition_list production.
	ExitSearch_condition_list(c *Search_condition_listContext)

	// ExitSearch_cond_or is called when exiting the search_cond_or production.
	ExitSearch_cond_or(c *Search_cond_orContext)

	// ExitSearch_cond_pred is called when exiting the search_cond_pred production.
	ExitSearch_cond_pred(c *Search_cond_predContext)

	// ExitSearch_cond_and is called when exiting the search_cond_and production.
	ExitSearch_cond_and(c *Search_cond_andContext)

	// ExitUnary_operator_expression3 is called when exiting the unary_operator_expression3 production.
	ExitUnary_operator_expression3(c *Unary_operator_expression3Context)

	// ExitUnary_operator_expression2 is called when exiting the unary_operator_expression2 production.
	ExitUnary_operator_expression2(c *Unary_operator_expression2Context)

	// ExitBinary_operator_expression2 is called when exiting the binary_operator_expression2 production.
	ExitBinary_operator_expression2(c *Binary_operator_expression2Context)

	// ExitSublink_expression is called when exiting the sublink_expression production.
	ExitSublink_expression(c *Sublink_expressionContext)

	// ExitBinary_mod_expression is called when exiting the binary_mod_expression production.
	ExitBinary_mod_expression(c *Binary_mod_expressionContext)

	// ExitBinary_in_expression is called when exiting the binary_in_expression production.
	ExitBinary_in_expression(c *Binary_in_expressionContext)

	// ExitBracket_search_expression is called when exiting the bracket_search_expression production.
	ExitBracket_search_expression(c *Bracket_search_expressionContext)

	// ExitDecimal_expression is called when exiting the decimal_expression production.
	ExitDecimal_expression(c *Decimal_expressionContext)

	// ExitBracket_query_expression is called when exiting the bracket_query_expression production.
	ExitBracket_query_expression(c *Bracket_query_expressionContext)

	// ExitQuery_specification_expression is called when exiting the query_specification_expression production.
	ExitQuery_specification_expression(c *Query_specification_expressionContext)

	// ExitUnion_query_expression is called when exiting the union_query_expression production.
	ExitUnion_query_expression(c *Union_query_expressionContext)

	// ExitUnion_op is called when exiting the union_op production.
	ExitUnion_op(c *Union_opContext)

	// ExitQuery_specification is called when exiting the query_specification production.
	ExitQuery_specification(c *Query_specificationContext)

	// ExitTop_clause is called when exiting the top_clause production.
	ExitTop_clause(c *Top_clauseContext)

	// ExitTop_clause_dm is called when exiting the top_clause_dm production.
	ExitTop_clause_dm(c *Top_clause_dmContext)

	// ExitOrder_by_clause is called when exiting the order_by_clause production.
	ExitOrder_by_clause(c *Order_by_clauseContext)

	// ExitFetch_expression is called when exiting the fetch_expression production.
	ExitFetch_expression(c *Fetch_expressionContext)

	// ExitFor_clause is called when exiting the for_clause production.
	ExitFor_clause(c *For_clauseContext)

	// ExitXml_common_directives is called when exiting the xml_common_directives production.
	ExitXml_common_directives(c *Xml_common_directivesContext)

	// ExitOrder_by_expression is called when exiting the order_by_expression production.
	ExitOrder_by_expression(c *Order_by_expressionContext)

	// ExitGroup_by_item is called when exiting the group_by_item production.
	ExitGroup_by_item(c *Group_by_itemContext)

	// ExitOption_clause is called when exiting the option_clause production.
	ExitOption_clause(c *Option_clauseContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitOptimize_for_arg is called when exiting the optimize_for_arg production.
	ExitOptimize_for_arg(c *Optimize_for_argContext)

	// ExitSelect_list is called when exiting the select_list production.
	ExitSelect_list(c *Select_listContext)

	// ExitSelect_list_elem is called when exiting the select_list_elem production.
	ExitSelect_list_elem(c *Select_list_elemContext)

	// ExitTable_sources is called when exiting the table_sources production.
	ExitTable_sources(c *Table_sourcesContext)

	// ExitCross_join is called when exiting the cross_join production.
	ExitCross_join(c *Cross_joinContext)

	// ExitTable_source_item_join is called when exiting the table_source_item_join production.
	ExitTable_source_item_join(c *Table_source_item_joinContext)

	// ExitStandard_join is called when exiting the standard_join production.
	ExitStandard_join(c *Standard_joinContext)

	// ExitApply_join is called when exiting the apply_join production.
	ExitApply_join(c *Apply_joinContext)

	// ExitBracket_table_source is called when exiting the bracket_table_source production.
	ExitBracket_table_source(c *Bracket_table_sourceContext)

	// ExitTable_source_item_name is called when exiting the table_source_item_name production.
	ExitTable_source_item_name(c *Table_source_item_nameContext)

	// ExitTable_source_item_simple is called when exiting the table_source_item_simple production.
	ExitTable_source_item_simple(c *Table_source_item_simpleContext)

	// ExitTable_source_item_complex is called when exiting the table_source_item_complex production.
	ExitTable_source_item_complex(c *Table_source_item_complexContext)

	// ExitTable_alias is called when exiting the table_alias production.
	ExitTable_alias(c *Table_aliasContext)

	// ExitChange_table is called when exiting the change_table production.
	ExitChange_table(c *Change_tableContext)

	// ExitJoin_type is called when exiting the join_type production.
	ExitJoin_type(c *Join_typeContext)

	// ExitTable_name_with_hint is called when exiting the table_name_with_hint production.
	ExitTable_name_with_hint(c *Table_name_with_hintContext)

	// ExitRowset_function is called when exiting the rowset_function production.
	ExitRowset_function(c *Rowset_functionContext)

	// ExitBulk_option is called when exiting the bulk_option production.
	ExitBulk_option(c *Bulk_optionContext)

	// ExitDerived_table is called when exiting the derived_table production.
	ExitDerived_table(c *Derived_tableContext)

	// ExitRank_call is called when exiting the rank_call production.
	ExitRank_call(c *Rank_callContext)

	// ExitAggregate_call is called when exiting the aggregate_call production.
	ExitAggregate_call(c *Aggregate_callContext)

	// ExitStandard_call is called when exiting the standard_call production.
	ExitStandard_call(c *Standard_callContext)

	// ExitNvf_call is called when exiting the nvf_call production.
	ExitNvf_call(c *Nvf_callContext)

	// ExitCast_call is called when exiting the cast_call production.
	ExitCast_call(c *Cast_callContext)

	// ExitSimple_call is called when exiting the simple_call production.
	ExitSimple_call(c *Simple_callContext)

	// ExitSwitch_section is called when exiting the switch_section production.
	ExitSwitch_section(c *Switch_sectionContext)

	// ExitSwitch_search_condition_section is called when exiting the switch_search_condition_section production.
	ExitSwitch_search_condition_section(c *Switch_search_condition_sectionContext)

	// ExitWith_table_hints is called when exiting the with_table_hints production.
	ExitWith_table_hints(c *With_table_hintsContext)

	// ExitInsert_with_table_hints is called when exiting the insert_with_table_hints production.
	ExitInsert_with_table_hints(c *Insert_with_table_hintsContext)

	// ExitTable_hint is called when exiting the table_hint production.
	ExitTable_hint(c *Table_hintContext)

	// ExitIndex_value is called when exiting the index_value production.
	ExitIndex_value(c *Index_valueContext)

	// ExitColumn_alias_list is called when exiting the column_alias_list production.
	ExitColumn_alias_list(c *Column_alias_listContext)

	// ExitColumn_alias is called when exiting the column_alias production.
	ExitColumn_alias(c *Column_aliasContext)

	// ExitA_star is called when exiting the a_star production.
	ExitA_star(c *A_starContext)

	// ExitTable_value_constructor is called when exiting the table_value_constructor production.
	ExitTable_value_constructor(c *Table_value_constructorContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitValue_list is called when exiting the value_list production.
	ExitValue_list(c *Value_listContext)

	// ExitNext_value_for is called when exiting the next_value_for production.
	ExitNext_value_for(c *Next_value_forContext)

	// ExitNext_value_for_function is called when exiting the next_value_for_function production.
	ExitNext_value_for_function(c *Next_value_for_functionContext)

	// ExitRanking_windowed_function is called when exiting the ranking_windowed_function production.
	ExitRanking_windowed_function(c *Ranking_windowed_functionContext)

	// ExitAggregate_windowed_function is called when exiting the aggregate_windowed_function production.
	ExitAggregate_windowed_function(c *Aggregate_windowed_functionContext)

	// ExitAll_distinct is called when exiting the all_distinct production.
	ExitAll_distinct(c *All_distinctContext)

	// ExitOver_clause is called when exiting the over_clause production.
	ExitOver_clause(c *Over_clauseContext)

	// ExitRow_or_range_clause is called when exiting the row_or_range_clause production.
	ExitRow_or_range_clause(c *Row_or_range_clauseContext)

	// ExitWindow_frame_extent is called when exiting the window_frame_extent production.
	ExitWindow_frame_extent(c *Window_frame_extentContext)

	// ExitWindow_frame_bound is called when exiting the window_frame_bound production.
	ExitWindow_frame_bound(c *Window_frame_boundContext)

	// ExitWindow_frame_preceding is called when exiting the window_frame_preceding production.
	ExitWindow_frame_preceding(c *Window_frame_precedingContext)

	// ExitWindow_frame_following is called when exiting the window_frame_following production.
	ExitWindow_frame_following(c *Window_frame_followingContext)

	// ExitCreate_database_option is called when exiting the create_database_option production.
	ExitCreate_database_option(c *Create_database_optionContext)

	// ExitDatabase_filestream_option is called when exiting the database_filestream_option production.
	ExitDatabase_filestream_option(c *Database_filestream_optionContext)

	// ExitDatabase_file_spec is called when exiting the database_file_spec production.
	ExitDatabase_file_spec(c *Database_file_specContext)

	// ExitFile_group is called when exiting the file_group production.
	ExitFile_group(c *File_groupContext)

	// ExitFile_spec is called when exiting the file_spec production.
	ExitFile_spec(c *File_specContext)

	// ExitFull_table_name is called when exiting the full_table_name production.
	ExitFull_table_name(c *Full_table_nameContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitSimple_name is called when exiting the simple_name production.
	ExitSimple_name(c *Simple_nameContext)

	// ExitFunc_proc_name is called when exiting the func_proc_name production.
	ExitFunc_proc_name(c *Func_proc_nameContext)

	// ExitDdl_object is called when exiting the ddl_object production.
	ExitDdl_object(c *Ddl_objectContext)

	// ExitFull_column_name is called when exiting the full_column_name production.
	ExitFull_column_name(c *Full_column_nameContext)

	// ExitColumn_name_list is called when exiting the column_name_list production.
	ExitColumn_name_list(c *Column_name_listContext)

	// ExitCursor_name is called when exiting the cursor_name production.
	ExitCursor_name(c *Cursor_nameContext)

	// ExitOn_off is called when exiting the on_off production.
	ExitOn_off(c *On_offContext)

	// ExitClustered is called when exiting the clustered production.
	ExitClustered(c *ClusteredContext)

	// ExitNull_notnull is called when exiting the null_notnull production.
	ExitNull_notnull(c *Null_notnullContext)

	// ExitScalar_function_name is called when exiting the scalar_function_name production.
	ExitScalar_function_name(c *Scalar_function_nameContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitDefault_value is called when exiting the default_value production.
	ExitDefault_value(c *Default_valueContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitSign is called when exiting the sign production.
	ExitSign(c *SignContext)

	// ExitR_id is called when exiting the r_id production.
	ExitR_id(c *R_idContext)

	// ExitSimple_id is called when exiting the simple_id production.
	ExitSimple_id(c *Simple_idContext)

	// ExitComparison_operator is called when exiting the comparison_operator production.
	ExitComparison_operator(c *Comparison_operatorContext)

	// ExitAssignment_operator is called when exiting the assignment_operator production.
	ExitAssignment_operator(c *Assignment_operatorContext)

	// ExitFile_size is called when exiting the file_size production.
	ExitFile_size(c *File_sizeContext)
}
