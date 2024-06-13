package pg_query

type Visitor interface {
	// Enter is called before children nodes are visited.
	// skipChildren returns true means children nodes should be skipped,
	// this is useful when work is done in Enter and there is no need to visit children.
	Enter(n IsNode_Node) (skipChildren bool)
}

func (n *Node_Alias) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Alias != nil {
		if !n.Alias.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RangeVar) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RangeVar != nil {
		if !n.RangeVar.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_TableFunc) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.TableFunc != nil {
		if !n.TableFunc.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_Expr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Expr != nil {
		if !n.Expr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_Var) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Var != nil {
		if !n.Var.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_Param) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Param != nil {
		if !n.Param.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_Aggref) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Aggref != nil {
		if !n.Aggref.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_GroupingFunc) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.GroupingFunc != nil {
		if !n.GroupingFunc.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_WindowFunc) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.WindowFunc != nil {
		if !n.WindowFunc.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SubscriptingRef) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SubscriptingRef != nil {
		if !n.SubscriptingRef.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_FuncExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.FuncExpr != nil {
		if !n.FuncExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_NamedArgExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.NamedArgExpr != nil {
		if !n.NamedArgExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_OpExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.OpExpr != nil {
		if !n.OpExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DistinctExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DistinctExpr != nil {
		if !n.DistinctExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_NullIfExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.NullIfExpr != nil {
		if !n.NullIfExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ScalarArrayOpExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ScalarArrayOpExpr != nil {
		if !n.ScalarArrayOpExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_BoolExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.BoolExpr != nil {
		if !n.BoolExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SubLink) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SubLink != nil {
		if !n.SubLink.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SubPlan) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SubPlan != nil {
		if !n.SubPlan.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlternativeSubPlan) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlternativeSubPlan != nil {
		if !n.AlternativeSubPlan.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_FieldSelect) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.FieldSelect != nil {
		if !n.FieldSelect.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_FieldStore) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.FieldStore != nil {
		if !n.FieldStore.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RelabelType) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RelabelType != nil {
		if !n.RelabelType.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CoerceViaIo) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CoerceViaIo != nil {
		if !n.CoerceViaIo.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ArrayCoerceExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ArrayCoerceExpr != nil {
		if !n.ArrayCoerceExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ConvertRowtypeExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ConvertRowtypeExpr != nil {
		if !n.ConvertRowtypeExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CollateExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CollateExpr != nil {
		if !n.CollateExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CaseExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CaseExpr != nil {
		if !n.CaseExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CaseWhen) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CaseWhen != nil {
		if !n.CaseWhen.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CaseTestExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CaseTestExpr != nil {
		if !n.CaseTestExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ArrayExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ArrayExpr != nil {
		if !n.ArrayExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RowExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RowExpr != nil {
		if !n.RowExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RowCompareExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RowCompareExpr != nil {
		if !n.RowCompareExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CoalesceExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CoalesceExpr != nil {
		if !n.CoalesceExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_MinMaxExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.MinMaxExpr != nil {
		if !n.MinMaxExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SqlvalueFunction) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SqlvalueFunction != nil {
		if !n.SqlvalueFunction.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_XmlExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.XmlExpr != nil {
		if !n.XmlExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_NullTest) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.NullTest != nil {
		if !n.NullTest.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_BooleanTest) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.BooleanTest != nil {
		if !n.BooleanTest.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CoerceToDomain) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CoerceToDomain != nil {
		if !n.CoerceToDomain.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CoerceToDomainValue) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CoerceToDomainValue != nil {
		if !n.CoerceToDomainValue.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SetToDefault) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SetToDefault != nil {
		if !n.SetToDefault.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CurrentOfExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CurrentOfExpr != nil {
		if !n.CurrentOfExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_NextValueExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.NextValueExpr != nil {
		if !n.NextValueExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_InferenceElem) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.InferenceElem != nil {
		if !n.InferenceElem.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_TargetEntry) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.TargetEntry != nil {
		if !n.TargetEntry.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RangeTblRef) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RangeTblRef != nil {
		if !n.RangeTblRef.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_JoinExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.JoinExpr != nil {
		if !n.JoinExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_FromExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.FromExpr != nil {
		if !n.FromExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_OnConflictExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.OnConflictExpr != nil {
		if !n.OnConflictExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_IntoClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.IntoClause != nil {
		if !n.IntoClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RawStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RawStmt != nil {
		if !n.RawStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_Query) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Query != nil {
		if !n.Query.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_InsertStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.InsertStmt != nil {
		if !n.InsertStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DeleteStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DeleteStmt != nil {
		if !n.DeleteStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_UpdateStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.UpdateStmt != nil {
		if !n.UpdateStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SelectStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SelectStmt != nil {
		if !n.SelectStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterTableStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterTableStmt != nil {
		if !n.AlterTableStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterTableCmd) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterTableCmd != nil {
		if !n.AlterTableCmd.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterDomainStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterDomainStmt != nil {
		if !n.AlterDomainStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SetOperationStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SetOperationStmt != nil {
		if !n.SetOperationStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_GrantStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.GrantStmt != nil {
		if !n.GrantStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_GrantRoleStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.GrantRoleStmt != nil {
		if !n.GrantRoleStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterDefaultPrivilegesStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterDefaultPrivilegesStmt != nil {
		if !n.AlterDefaultPrivilegesStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ClosePortalStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ClosePortalStmt != nil {
		if !n.ClosePortalStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ClusterStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ClusterStmt != nil {
		if !n.ClusterStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CopyStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CopyStmt != nil {
		if !n.CopyStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateStmt != nil {
		if !n.CreateStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DefineStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DefineStmt != nil {
		if !n.DefineStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DropStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DropStmt != nil {
		if !n.DropStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_TruncateStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.TruncateStmt != nil {
		if !n.TruncateStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CommentStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CommentStmt != nil {
		if !n.CommentStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_FetchStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.FetchStmt != nil {
		if !n.FetchStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_IndexStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.IndexStmt != nil {
		if !n.IndexStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateFunctionStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateFunctionStmt != nil {
		if !n.CreateFunctionStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterFunctionStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterFunctionStmt != nil {
		if !n.AlterFunctionStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DoStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DoStmt != nil {
		if !n.DoStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RenameStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RenameStmt != nil {
		if !n.RenameStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RuleStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RuleStmt != nil {
		if !n.RuleStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_NotifyStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.NotifyStmt != nil {
		if !n.NotifyStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ListenStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ListenStmt != nil {
		if !n.ListenStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_UnlistenStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.UnlistenStmt != nil {
		if !n.UnlistenStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_TransactionStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.TransactionStmt != nil {
		if !n.TransactionStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ViewStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ViewStmt != nil {
		if !n.ViewStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_LoadStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.LoadStmt != nil {
		if !n.LoadStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateDomainStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateDomainStmt != nil {
		if !n.CreateDomainStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreatedbStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreatedbStmt != nil {
		if !n.CreatedbStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DropdbStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DropdbStmt != nil {
		if !n.DropdbStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_VacuumStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.VacuumStmt != nil {
		if !n.VacuumStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ExplainStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ExplainStmt != nil {
		if !n.ExplainStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateTableAsStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateTableAsStmt != nil {
		if !n.CreateTableAsStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateSeqStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateSeqStmt != nil {
		if !n.CreateSeqStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterSeqStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterSeqStmt != nil {
		if !n.AlterSeqStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_VariableSetStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.VariableSetStmt != nil {
		if !n.VariableSetStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_VariableShowStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.VariableShowStmt != nil {
		if !n.VariableShowStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DiscardStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DiscardStmt != nil {
		if !n.DiscardStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateTrigStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateTrigStmt != nil {
		if !n.CreateTrigStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreatePlangStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreatePlangStmt != nil {
		if !n.CreatePlangStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateRoleStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateRoleStmt != nil {
		if !n.CreateRoleStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterRoleStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterRoleStmt != nil {
		if !n.AlterRoleStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DropRoleStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DropRoleStmt != nil {
		if !n.DropRoleStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_LockStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.LockStmt != nil {
		if !n.LockStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ConstraintsSetStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ConstraintsSetStmt != nil {
		if !n.ConstraintsSetStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ReindexStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ReindexStmt != nil {
		if !n.ReindexStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CheckPointStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CheckPointStmt != nil {
		if !n.CheckPointStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateSchemaStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateSchemaStmt != nil {
		if !n.CreateSchemaStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterDatabaseStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterDatabaseStmt != nil {
		if !n.AlterDatabaseStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterDatabaseSetStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterDatabaseSetStmt != nil {
		if !n.AlterDatabaseSetStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterRoleSetStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterRoleSetStmt != nil {
		if !n.AlterRoleSetStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateConversionStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateConversionStmt != nil {
		if !n.CreateConversionStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateCastStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateCastStmt != nil {
		if !n.CreateCastStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateOpClassStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateOpClassStmt != nil {
		if !n.CreateOpClassStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateOpFamilyStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateOpFamilyStmt != nil {
		if !n.CreateOpFamilyStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterOpFamilyStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterOpFamilyStmt != nil {
		if !n.AlterOpFamilyStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_PrepareStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.PrepareStmt != nil {
		if !n.PrepareStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ExecuteStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ExecuteStmt != nil {
		if !n.ExecuteStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DeallocateStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DeallocateStmt != nil {
		if !n.DeallocateStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DeclareCursorStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DeclareCursorStmt != nil {
		if !n.DeclareCursorStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateTableSpaceStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateTableSpaceStmt != nil {
		if !n.CreateTableSpaceStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DropTableSpaceStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DropTableSpaceStmt != nil {
		if !n.DropTableSpaceStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterObjectDependsStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterObjectDependsStmt != nil {
		if !n.AlterObjectDependsStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterObjectSchemaStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterObjectSchemaStmt != nil {
		if !n.AlterObjectSchemaStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterOwnerStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterOwnerStmt != nil {
		if !n.AlterOwnerStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterOperatorStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterOperatorStmt != nil {
		if !n.AlterOperatorStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterTypeStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterTypeStmt != nil {
		if !n.AlterTypeStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DropOwnedStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DropOwnedStmt != nil {
		if !n.DropOwnedStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ReassignOwnedStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ReassignOwnedStmt != nil {
		if !n.ReassignOwnedStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CompositeTypeStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CompositeTypeStmt != nil {
		if !n.CompositeTypeStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateEnumStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateEnumStmt != nil {
		if !n.CreateEnumStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateRangeStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateRangeStmt != nil {
		if !n.CreateRangeStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterEnumStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterEnumStmt != nil {
		if !n.AlterEnumStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterTsdictionaryStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterTsdictionaryStmt != nil {
		if !n.AlterTsdictionaryStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterTsconfigurationStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterTsconfigurationStmt != nil {
		if !n.AlterTsconfigurationStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateFdwStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateFdwStmt != nil {
		if !n.CreateFdwStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterFdwStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterFdwStmt != nil {
		if !n.AlterFdwStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateForeignServerStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateForeignServerStmt != nil {
		if !n.CreateForeignServerStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterForeignServerStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterForeignServerStmt != nil {
		if !n.AlterForeignServerStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateUserMappingStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateUserMappingStmt != nil {
		if !n.CreateUserMappingStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterUserMappingStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterUserMappingStmt != nil {
		if !n.AlterUserMappingStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DropUserMappingStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DropUserMappingStmt != nil {
		if !n.DropUserMappingStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterTableSpaceOptionsStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterTableSpaceOptionsStmt != nil {
		if !n.AlterTableSpaceOptionsStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterTableMoveAllStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterTableMoveAllStmt != nil {
		if !n.AlterTableMoveAllStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SecLabelStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SecLabelStmt != nil {
		if !n.SecLabelStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateForeignTableStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateForeignTableStmt != nil {
		if !n.CreateForeignTableStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ImportForeignSchemaStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ImportForeignSchemaStmt != nil {
		if !n.ImportForeignSchemaStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateExtensionStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateExtensionStmt != nil {
		if !n.CreateExtensionStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterExtensionStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterExtensionStmt != nil {
		if !n.AlterExtensionStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterExtensionContentsStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterExtensionContentsStmt != nil {
		if !n.AlterExtensionContentsStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateEventTrigStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateEventTrigStmt != nil {
		if !n.CreateEventTrigStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterEventTrigStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterEventTrigStmt != nil {
		if !n.AlterEventTrigStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RefreshMatViewStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RefreshMatViewStmt != nil {
		if !n.RefreshMatViewStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ReplicaIdentityStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ReplicaIdentityStmt != nil {
		if !n.ReplicaIdentityStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterSystemStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterSystemStmt != nil {
		if !n.AlterSystemStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreatePolicyStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreatePolicyStmt != nil {
		if !n.CreatePolicyStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterPolicyStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterPolicyStmt != nil {
		if !n.AlterPolicyStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateTransformStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateTransformStmt != nil {
		if !n.CreateTransformStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateAmStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateAmStmt != nil {
		if !n.CreateAmStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreatePublicationStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreatePublicationStmt != nil {
		if !n.CreatePublicationStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterPublicationStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterPublicationStmt != nil {
		if !n.AlterPublicationStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateSubscriptionStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateSubscriptionStmt != nil {
		if !n.CreateSubscriptionStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterSubscriptionStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterSubscriptionStmt != nil {
		if !n.AlterSubscriptionStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DropSubscriptionStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DropSubscriptionStmt != nil {
		if !n.DropSubscriptionStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateStatsStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateStatsStmt != nil {
		if !n.CreateStatsStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterCollationStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterCollationStmt != nil {
		if !n.AlterCollationStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CallStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CallStmt != nil {
		if !n.CallStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AlterStatsStmt) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AlterStatsStmt != nil {
		if !n.AlterStatsStmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AExpr != nil {
		if !n.AExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ColumnRef) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ColumnRef != nil {
		if !n.ColumnRef.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ParamRef) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ParamRef != nil {
		if !n.ParamRef.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AConst) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AConst != nil {
		if !n.AConst.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_FuncCall) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.FuncCall != nil {
		if !n.FuncCall.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AStar) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AStar != nil {
		if !n.AStar.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AIndices) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AIndices != nil {
		if !n.AIndices.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AIndirection) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AIndirection != nil {
		if !n.AIndirection.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AArrayExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AArrayExpr != nil {
		if !n.AArrayExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ResTarget) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ResTarget != nil {
		if !n.ResTarget.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_MultiAssignRef) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.MultiAssignRef != nil {
		if !n.MultiAssignRef.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_TypeCast) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.TypeCast != nil {
		if !n.TypeCast.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CollateClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CollateClause != nil {
		if !n.CollateClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SortBy) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SortBy != nil {
		if !n.SortBy.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_WindowDef) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.WindowDef != nil {
		if !n.WindowDef.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RangeSubselect) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RangeSubselect != nil {
		if !n.RangeSubselect.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RangeFunction) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RangeFunction != nil {
		if !n.RangeFunction.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RangeTableSample) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RangeTableSample != nil {
		if !n.RangeTableSample.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RangeTableFunc) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RangeTableFunc != nil {
		if !n.RangeTableFunc.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RangeTableFuncCol) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RangeTableFuncCol != nil {
		if !n.RangeTableFuncCol.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_TypeName) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.TypeName != nil {
		if !n.TypeName.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ColumnDef) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ColumnDef != nil {
		if !n.ColumnDef.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_IndexElem) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.IndexElem != nil {
		if !n.IndexElem.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_Constraint) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Constraint != nil {
		if !n.Constraint.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_DefElem) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.DefElem != nil {
		if !n.DefElem.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RangeTblEntry) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RangeTblEntry != nil {
		if !n.RangeTblEntry.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RangeTblFunction) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RangeTblFunction != nil {
		if !n.RangeTblFunction.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_TableSampleClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.TableSampleClause != nil {
		if !n.TableSampleClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_WithCheckOption) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.WithCheckOption != nil {
		if !n.WithCheckOption.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_SortGroupClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.SortGroupClause != nil {
		if !n.SortGroupClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_GroupingSet) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.GroupingSet != nil {
		if !n.GroupingSet.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_WindowClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.WindowClause != nil {
		if !n.WindowClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_ObjectWithArgs) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.ObjectWithArgs != nil {
		if !n.ObjectWithArgs.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_AccessPriv) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.AccessPriv != nil {
		if !n.AccessPriv.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CreateOpClassItem) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CreateOpClassItem != nil {
		if !n.CreateOpClassItem.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_TableLikeClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.TableLikeClause != nil {
		if !n.TableLikeClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_FunctionParameter) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.FunctionParameter != nil {
		if !n.FunctionParameter.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_LockingClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.LockingClause != nil {
		if !n.LockingClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RowMarkClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RowMarkClause != nil {
		if !n.RowMarkClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_XmlSerialize) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.XmlSerialize != nil {
		if !n.XmlSerialize.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_WithClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.WithClause != nil {
		if !n.WithClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_InferClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.InferClause != nil {
		if !n.InferClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_OnConflictClause) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.OnConflictClause != nil {
		if !n.OnConflictClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CommonTableExpr) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CommonTableExpr != nil {
		if !n.CommonTableExpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_RoleSpec) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.RoleSpec != nil {
		if !n.RoleSpec.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_TriggerTransition) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.TriggerTransition != nil {
		if !n.TriggerTransition.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_PartitionElem) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.PartitionElem != nil {
		if !n.PartitionElem.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_PartitionSpec) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.PartitionSpec != nil {
		if !n.PartitionSpec.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_PartitionBoundSpec) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.PartitionBoundSpec != nil {
		if !n.PartitionBoundSpec.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_PartitionRangeDatum) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.PartitionRangeDatum != nil {
		if !n.PartitionRangeDatum.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_PartitionCmd) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.PartitionCmd != nil {
		if !n.PartitionCmd.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_VacuumRelation) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.VacuumRelation != nil {
		if !n.VacuumRelation.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_InlineCodeBlock) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.InlineCodeBlock != nil {
		if !n.InlineCodeBlock.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_CallContext) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.CallContext != nil {
		if !n.CallContext.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_Integer) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Integer != nil {
		if !n.Integer.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_Float) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Float != nil {
		if !n.Float.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_String_) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.String_ != nil {
		if !n.String_.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_BitString) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.BitString != nil {
		if !n.BitString.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_Null) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.Null != nil {
		if !n.Null.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_List) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.List != nil {
		if !n.List.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_IntList) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.IntList != nil {
		if !n.IntList.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Node_OidList) Accept(v Visitor) bool {
	if v.Enter(n) {
		return false
	}

	if n.OidList != nil {
		if !n.OidList.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Integer) Accept(v Visitor) bool {
	return true
}

func (n *Float) Accept(v Visitor) bool {
	return true
}

func (n *String) Accept(v Visitor) bool {
	return true
}

func (n *BitString) Accept(v Visitor) bool {
	return true
}

func (n *Null) Accept(v Visitor) bool {
	return true
}

func (n *List) Accept(v Visitor) bool {
	for _, item := range n.Items {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *OidList) Accept(v Visitor) bool {
	for _, item := range n.Items {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *IntList) Accept(v Visitor) bool {
	for _, item := range n.Items {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Alias) Accept(v Visitor) bool {
	for _, item := range n.Colnames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RangeVar) Accept(v Visitor) bool {
	if n.Alias != nil {
		if !n.Alias.Accept(v) {
			return false
		}
	}
	return true
}

func (n *TableFunc) Accept(v Visitor) bool {
	for _, item := range n.NsUris {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.NsNames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Docexpr != nil {
		if !n.Docexpr.Node.Accept(v) {
			return false
		}
	}
	if n.Rowexpr != nil {
		if !n.Rowexpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Colnames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Coltypes {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Coltypmods {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Colcollations {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Colexprs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Coldefexprs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Expr) Accept(v Visitor) bool {
	return true
}

func (n *AlterTableMoveAllStmt) Accept(v Visitor) bool {
	for _, item := range n.Roles {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Var) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Param) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Aggref) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Aggargtypes {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Aggdirectargs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Aggorder {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Aggdistinct {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Aggfilter != nil {
		if !n.Aggfilter.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *GroupingFunc) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Refs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Cols {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *WindowFunc) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Aggfilter != nil {
		if !n.Aggfilter.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SubscriptingRef) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Refupperindexpr {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Reflowerindexpr {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Refexpr != nil {
		if !n.Refexpr.Node.Accept(v) {
			return false
		}
	}
	if n.Refassgnexpr != nil {
		if !n.Refassgnexpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *FuncExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *NamedArgExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *OpExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DistinctExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *NullIfExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ScalarArrayOpExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *BoolExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SubLink) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Testexpr != nil {
		if !n.Testexpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.OperName {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Subselect != nil {
		if !n.Subselect.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SubPlan) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Testexpr != nil {
		if !n.Testexpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ParamIds {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.SetParam {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ParParam {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlternativeSubPlan) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Subplans {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *FieldSelect) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *FieldStore) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Newvals {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Fieldnums {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RelabelType) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CoerceViaIO) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ArrayCoerceExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	if n.Elemexpr != nil {
		if !n.Elemexpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ConvertRowtypeExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CollateExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CaseExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Defresult != nil {
		if !n.Defresult.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CaseWhen) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Expr != nil {
		if !n.Expr.Node.Accept(v) {
			return false
		}
	}
	if n.Result != nil {
		if !n.Result.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CaseTestExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ArrayExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Elements {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RowExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Colnames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RowCompareExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Opnos {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Opfamilies {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Inputcollids {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Largs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Rargs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CoalesceExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *MinMaxExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SQLValueFunction) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *XmlExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.NamedArgs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ArgNames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *NullTest) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *BooleanTest) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CoerceToDomain) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CoerceToDomainValue) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SetToDefault) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CurrentOfExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *NextValueExpr) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *InferenceElem) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Expr != nil {
		if !n.Expr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *TargetEntry) Accept(v Visitor) bool {
	if n.Xpr != nil {
		if !n.Xpr.Node.Accept(v) {
			return false
		}
	}
	if n.Expr != nil {
		if !n.Expr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RangeTblRef) Accept(v Visitor) bool {
	return true
}

func (n *JoinExpr) Accept(v Visitor) bool {
	if n.Larg != nil {
		if !n.Larg.Node.Accept(v) {
			return false
		}
	}
	if n.Rarg != nil {
		if !n.Rarg.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.UsingClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Quals != nil {
		if !n.Quals.Node.Accept(v) {
			return false
		}
	}
	if n.Alias != nil {
		if !n.Alias.Accept(v) {
			return false
		}
	}
	return true
}

func (n *FromExpr) Accept(v Visitor) bool {
	for _, item := range n.Fromlist {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Quals != nil {
		if !n.Quals.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *OnConflictExpr) Accept(v Visitor) bool {
	for _, item := range n.ArbiterElems {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.ArbiterWhere != nil {
		if !n.ArbiterWhere.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.OnConflictSet {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.OnConflictWhere != nil {
		if !n.OnConflictWhere.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ExclRelTlist {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *IntoClause) Accept(v Visitor) bool {
	if n.Rel != nil {
		if !n.Rel.Accept(v) {
			return false
		}
	}
	for _, item := range n.ColNames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.ViewQuery != nil {
		if !n.ViewQuery.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RawStmt) Accept(v Visitor) bool {
	if n.Stmt != nil {
		if !n.Stmt.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Query) Accept(v Visitor) bool {
	if n.UtilityStmt != nil {
		if !n.UtilityStmt.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.CteList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Rtable {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Jointree != nil {
		if !n.Jointree.Accept(v) {
			return false
		}
	}
	for _, item := range n.TargetList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.OnConflict != nil {
		if !n.OnConflict.Accept(v) {
			return false
		}
	}
	for _, item := range n.ReturningList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.GroupClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.GroupingSets {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.HavingQual != nil {
		if !n.HavingQual.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.WindowClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.DistinctClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.SortClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.LimitOffset != nil {
		if !n.LimitOffset.Node.Accept(v) {
			return false
		}
	}
	if n.LimitCount != nil {
		if !n.LimitCount.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.RowMarks {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.SetOperations != nil {
		if !n.SetOperations.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ConstraintDeps {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.WithCheckOptions {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *InsertStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	for _, item := range n.Cols {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.SelectStmt != nil {
		if !n.SelectStmt.Node.Accept(v) {
			return false
		}
	}
	if n.OnConflictClause != nil {
		if !n.OnConflictClause.Accept(v) {
			return false
		}
	}
	for _, item := range n.ReturningList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WithClause != nil {
		if !n.WithClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DeleteStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	for _, item := range n.UsingClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WhereClause != nil {
		if !n.WhereClause.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ReturningList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WithClause != nil {
		if !n.WithClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *UpdateStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	for _, item := range n.TargetList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WhereClause != nil {
		if !n.WhereClause.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.FromClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ReturningList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WithClause != nil {
		if !n.WithClause.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SelectStmt) Accept(v Visitor) bool {
	for _, item := range n.DistinctClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.IntoClause != nil {
		if !n.IntoClause.Accept(v) {
			return false
		}
	}
	for _, item := range n.TargetList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.FromClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WhereClause != nil {
		if !n.WhereClause.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.GroupClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.HavingClause != nil {
		if !n.HavingClause.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.WindowClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ValuesLists {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.SortClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.LimitOffset != nil {
		if !n.LimitOffset.Node.Accept(v) {
			return false
		}
	}
	if n.LimitCount != nil {
		if !n.LimitCount.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.LockingClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WithClause != nil {
		if !n.WithClause.Accept(v) {
			return false
		}
	}
	if n.Larg != nil {
		if !n.Larg.Accept(v) {
			return false
		}
	}
	if n.Rarg != nil {
		if !n.Rarg.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterTableStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	for _, item := range n.Cmds {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterTableCmd) Accept(v Visitor) bool {
	if n.Newowner != nil {
		if !n.Newowner.Accept(v) {
			return false
		}
	}
	if n.Def != nil {
		if !n.Def.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterDomainStmt) Accept(v Visitor) bool {
	for _, item := range n.TypeName {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Def != nil {
		if !n.Def.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SetOperationStmt) Accept(v Visitor) bool {
	if n.Larg != nil {
		if !n.Larg.Node.Accept(v) {
			return false
		}
	}
	if n.Rarg != nil {
		if !n.Rarg.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ColTypes {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ColTypmods {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ColCollations {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.GroupClauses {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *GrantStmt) Accept(v Visitor) bool {
	for _, item := range n.Objects {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Privileges {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Grantees {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *GrantRoleStmt) Accept(v Visitor) bool {
	for _, item := range n.GrantedRoles {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.GranteeRoles {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Grantor != nil {
		if !n.Grantor.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterDefaultPrivilegesStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Action != nil {
		if !n.Action.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ClosePortalStmt) Accept(v Visitor) bool {
	return true
}

func (n *ClusterStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CopyStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	if n.Query != nil {
		if !n.Query.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Attlist {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WhereClause != nil {
		if !n.WhereClause.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	for _, item := range n.TableElts {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.InhRelations {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Partbound != nil {
		if !n.Partbound.Accept(v) {
			return false
		}
	}
	if n.Partspec != nil {
		if !n.Partspec.Accept(v) {
			return false
		}
	}
	if n.OfTypename != nil {
		if !n.OfTypename.Accept(v) {
			return false
		}
	}
	for _, item := range n.Constraints {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DefineStmt) Accept(v Visitor) bool {
	for _, item := range n.Defnames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Definition {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DropStmt) Accept(v Visitor) bool {
	for _, item := range n.Objects {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *TruncateStmt) Accept(v Visitor) bool {
	for _, item := range n.Relations {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CommentStmt) Accept(v Visitor) bool {
	if n.Object != nil {
		if !n.Object.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *FetchStmt) Accept(v Visitor) bool {
	return true
}

func (n *IndexStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	for _, item := range n.IndexParams {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.IndexIncludingParams {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WhereClause != nil {
		if !n.WhereClause.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ExcludeOpNames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateFunctionStmt) Accept(v Visitor) bool {
	for _, item := range n.Funcname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Parameters {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.ReturnType != nil {
		if !n.ReturnType.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterFunctionStmt) Accept(v Visitor) bool {
	if n.Func != nil {
		if !n.Func.Accept(v) {
			return false
		}
	}
	for _, item := range n.Actions {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DoStmt) Accept(v Visitor) bool {
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RenameStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	if n.Object != nil {
		if !n.Object.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RuleStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	if n.WhereClause != nil {
		if !n.WhereClause.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Actions {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *NotifyStmt) Accept(v Visitor) bool {
	return true
}

func (n *ListenStmt) Accept(v Visitor) bool {
	return true
}

func (n *UnlistenStmt) Accept(v Visitor) bool {
	return true
}

func (n *TransactionStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ViewStmt) Accept(v Visitor) bool {
	if n.View != nil {
		if !n.View.Accept(v) {
			return false
		}
	}
	for _, item := range n.Aliases {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Query != nil {
		if !n.Query.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *LoadStmt) Accept(v Visitor) bool {
	return true
}

func (n *CreateDomainStmt) Accept(v Visitor) bool {
	for _, item := range n.Domainname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.TypeName != nil {
		if !n.TypeName.Accept(v) {
			return false
		}
	}
	if n.CollClause != nil {
		if !n.CollClause.Accept(v) {
			return false
		}
	}
	for _, item := range n.Constraints {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreatedbStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DropdbStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *VacuumStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Rels {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ExplainStmt) Accept(v Visitor) bool {
	if n.Query != nil {
		if !n.Query.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateTableAsStmt) Accept(v Visitor) bool {
	if n.Query != nil {
		if !n.Query.Node.Accept(v) {
			return false
		}
	}
	if n.Into != nil {
		if !n.Into.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateSeqStmt) Accept(v Visitor) bool {
	if n.Sequence != nil {
		if !n.Sequence.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterSeqStmt) Accept(v Visitor) bool {
	if n.Sequence != nil {
		if !n.Sequence.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *VariableSetStmt) Accept(v Visitor) bool {
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *VariableShowStmt) Accept(v Visitor) bool {
	return true
}

func (n *DiscardStmt) Accept(v Visitor) bool {
	return true
}

func (n *CreateTrigStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	for _, item := range n.Funcname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Columns {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WhenClause != nil {
		if !n.WhenClause.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.TransitionRels {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Constrrel != nil {
		if !n.Constrrel.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreatePLangStmt) Accept(v Visitor) bool {
	for _, item := range n.Plhandler {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Plinline {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Plvalidator {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateRoleStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterRoleStmt) Accept(v Visitor) bool {
	if n.Role != nil {
		if !n.Role.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DropRoleStmt) Accept(v Visitor) bool {
	for _, item := range n.Roles {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *LockStmt) Accept(v Visitor) bool {
	for _, item := range n.Relations {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ConstraintsSetStmt) Accept(v Visitor) bool {
	for _, item := range n.Constraints {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ReindexStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CheckPointStmt) Accept(v Visitor) bool {
	return true
}

func (n *CreateSchemaStmt) Accept(v Visitor) bool {
	if n.Authrole != nil {
		if !n.Authrole.Accept(v) {
			return false
		}
	}
	for _, item := range n.SchemaElts {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterDatabaseStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterDatabaseSetStmt) Accept(v Visitor) bool {
	if n.Setstmt != nil {
		if !n.Setstmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterRoleSetStmt) Accept(v Visitor) bool {
	if n.Role != nil {
		if !n.Role.Accept(v) {
			return false
		}
	}
	if n.Setstmt != nil {
		if !n.Setstmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateConversionStmt) Accept(v Visitor) bool {
	for _, item := range n.ConversionName {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.FuncName {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateCastStmt) Accept(v Visitor) bool {
	if n.Sourcetype != nil {
		if !n.Sourcetype.Accept(v) {
			return false
		}
	}
	if n.Targettype != nil {
		if !n.Targettype.Accept(v) {
			return false
		}
	}
	if n.Func != nil {
		if !n.Func.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateOpClassStmt) Accept(v Visitor) bool {
	for _, item := range n.Opclassname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Opfamilyname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Datatype != nil {
		if !n.Datatype.Accept(v) {
			return false
		}
	}
	for _, item := range n.Items {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateOpFamilyStmt) Accept(v Visitor) bool {
	for _, item := range n.Opfamilyname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterOpFamilyStmt) Accept(v Visitor) bool {
	for _, item := range n.Opfamilyname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Items {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *PrepareStmt) Accept(v Visitor) bool {
	for _, item := range n.Argtypes {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Query != nil {
		if !n.Query.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ExecuteStmt) Accept(v Visitor) bool {
	for _, item := range n.Params {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DeallocateStmt) Accept(v Visitor) bool {
	return true
}

func (n *DeclareCursorStmt) Accept(v Visitor) bool {
	if n.Query != nil {
		if !n.Query.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateTableSpaceStmt) Accept(v Visitor) bool {
	if n.Owner != nil {
		if !n.Owner.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DropTableSpaceStmt) Accept(v Visitor) bool {
	return true
}

func (n *AlterObjectDependsStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	if n.Object != nil {
		if !n.Object.Node.Accept(v) {
			return false
		}
	}
	if n.Extname != nil {
		if !n.Extname.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterObjectSchemaStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	if n.Object != nil {
		if !n.Object.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterOwnerStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	if n.Object != nil {
		if !n.Object.Node.Accept(v) {
			return false
		}
	}
	if n.Newowner != nil {
		if !n.Newowner.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterOperatorStmt) Accept(v Visitor) bool {
	if n.Opername != nil {
		if !n.Opername.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterTypeStmt) Accept(v Visitor) bool {
	for _, item := range n.TypeName {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DropOwnedStmt) Accept(v Visitor) bool {
	for _, item := range n.Roles {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ReassignOwnedStmt) Accept(v Visitor) bool {
	for _, item := range n.Roles {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Newrole != nil {
		if !n.Newrole.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CompositeTypeStmt) Accept(v Visitor) bool {
	if n.Typevar != nil {
		if !n.Typevar.Accept(v) {
			return false
		}
	}
	for _, item := range n.Coldeflist {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateEnumStmt) Accept(v Visitor) bool {
	for _, item := range n.TypeName {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Vals {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateRangeStmt) Accept(v Visitor) bool {
	for _, item := range n.TypeName {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Params {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterEnumStmt) Accept(v Visitor) bool {
	for _, item := range n.TypeName {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterTSDictionaryStmt) Accept(v Visitor) bool {
	for _, item := range n.Dictname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterTSConfigurationStmt) Accept(v Visitor) bool {
	for _, item := range n.Cfgname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Tokentype {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Dicts {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateFdwStmt) Accept(v Visitor) bool {
	for _, item := range n.FuncOptions {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterFdwStmt) Accept(v Visitor) bool {
	for _, item := range n.FuncOptions {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateForeignServerStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterForeignServerStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateUserMappingStmt) Accept(v Visitor) bool {
	if n.User != nil {
		if !n.User.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterUserMappingStmt) Accept(v Visitor) bool {
	if n.User != nil {
		if !n.User.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DropUserMappingStmt) Accept(v Visitor) bool {
	if n.User != nil {
		if !n.User.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterTableSpaceOptionsStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SecLabelStmt) Accept(v Visitor) bool {
	if n.Object != nil {
		if !n.Object.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateForeignTableStmt) Accept(v Visitor) bool {
	if n.BaseStmt != nil {
		if !n.BaseStmt.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ImportForeignSchemaStmt) Accept(v Visitor) bool {
	for _, item := range n.TableList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateExtensionStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterExtensionStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterExtensionContentsStmt) Accept(v Visitor) bool {
	if n.Object != nil {
		if !n.Object.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateEventTrigStmt) Accept(v Visitor) bool {
	for _, item := range n.Whenclause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Funcname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterEventTrigStmt) Accept(v Visitor) bool {
	return true
}

func (n *RefreshMatViewStmt) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ReplicaIdentityStmt) Accept(v Visitor) bool {
	return true
}

func (n *AlterSystemStmt) Accept(v Visitor) bool {
	if n.Setstmt != nil {
		if !n.Setstmt.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreatePolicyStmt) Accept(v Visitor) bool {
	if n.Table != nil {
		if !n.Table.Accept(v) {
			return false
		}
	}
	for _, item := range n.Roles {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Qual != nil {
		if !n.Qual.Node.Accept(v) {
			return false
		}
	}
	if n.WithCheck != nil {
		if !n.WithCheck.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterPolicyStmt) Accept(v Visitor) bool {
	if n.Table != nil {
		if !n.Table.Accept(v) {
			return false
		}
	}
	for _, item := range n.Roles {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Qual != nil {
		if !n.Qual.Node.Accept(v) {
			return false
		}
	}
	if n.WithCheck != nil {
		if !n.WithCheck.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateTransformStmt) Accept(v Visitor) bool {
	if n.TypeName != nil {
		if !n.TypeName.Accept(v) {
			return false
		}
	}
	if n.Fromsql != nil {
		if !n.Fromsql.Accept(v) {
			return false
		}
	}
	if n.Tosql != nil {
		if !n.Tosql.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateAmStmt) Accept(v Visitor) bool {
	for _, item := range n.HandlerName {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreatePublicationStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Tables {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterPublicationStmt) Accept(v Visitor) bool {
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Tables {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateSubscriptionStmt) Accept(v Visitor) bool {
	for _, item := range n.Publication {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterSubscriptionStmt) Accept(v Visitor) bool {
	for _, item := range n.Publication {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DropSubscriptionStmt) Accept(v Visitor) bool {
	return true
}

func (n *CreateStatsStmt) Accept(v Visitor) bool {
	for _, item := range n.Defnames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.StatTypes {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Exprs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Relations {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterCollationStmt) Accept(v Visitor) bool {
	for _, item := range n.Collname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CallStmt) Accept(v Visitor) bool {
	if n.Funccall != nil {
		if !n.Funccall.Accept(v) {
			return false
		}
	}
	if n.Funcexpr != nil {
		if !n.Funcexpr.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AlterStatsStmt) Accept(v Visitor) bool {
	for _, item := range n.Defnames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *A_Expr) Accept(v Visitor) bool {
	for _, item := range n.Name {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Lexpr != nil {
		if !n.Lexpr.Node.Accept(v) {
			return false
		}
	}
	if n.Rexpr != nil {
		if !n.Rexpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ColumnRef) Accept(v Visitor) bool {
	for _, item := range n.Fields {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ParamRef) Accept(v Visitor) bool {
	return true
}

func (n *A_Const) Accept(v Visitor) bool {
	if n.Val != nil {
		if !n.Val.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *FuncCall) Accept(v Visitor) bool {
	for _, item := range n.Funcname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.AggOrder {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.AggFilter != nil {
		if !n.AggFilter.Node.Accept(v) {
			return false
		}
	}
	if n.Over != nil {
		if !n.Over.Accept(v) {
			return false
		}
	}
	return true
}

func (n *A_Star) Accept(v Visitor) bool {
	return true
}

func (n *A_Indices) Accept(v Visitor) bool {
	if n.Lidx != nil {
		if !n.Lidx.Node.Accept(v) {
			return false
		}
	}
	if n.Uidx != nil {
		if !n.Uidx.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *A_Indirection) Accept(v Visitor) bool {
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Indirection {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *A_ArrayExpr) Accept(v Visitor) bool {
	for _, item := range n.Elements {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ResTarget) Accept(v Visitor) bool {
	for _, item := range n.Indirection {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Val != nil {
		if !n.Val.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *MultiAssignRef) Accept(v Visitor) bool {
	if n.Source != nil {
		if !n.Source.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *TypeCast) Accept(v Visitor) bool {
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	if n.TypeName != nil {
		if !n.TypeName.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CollateClause) Accept(v Visitor) bool {
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Collname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SortBy) Accept(v Visitor) bool {
	if n.Node != nil {
		if !n.Node.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.UseOp {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *WindowDef) Accept(v Visitor) bool {
	for _, item := range n.PartitionClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.OrderClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.StartOffset != nil {
		if !n.StartOffset.Node.Accept(v) {
			return false
		}
	}
	if n.EndOffset != nil {
		if !n.EndOffset.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RangeSubselect) Accept(v Visitor) bool {
	if n.Subquery != nil {
		if !n.Subquery.Node.Accept(v) {
			return false
		}
	}
	if n.Alias != nil {
		if !n.Alias.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RangeFunction) Accept(v Visitor) bool {
	for _, item := range n.Functions {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Alias != nil {
		if !n.Alias.Accept(v) {
			return false
		}
	}
	for _, item := range n.Coldeflist {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RangeTableSample) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Method {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Repeatable != nil {
		if !n.Repeatable.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RangeTableFunc) Accept(v Visitor) bool {
	if n.Docexpr != nil {
		if !n.Docexpr.Node.Accept(v) {
			return false
		}
	}
	if n.Rowexpr != nil {
		if !n.Rowexpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Namespaces {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Columns {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Alias != nil {
		if !n.Alias.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RangeTableFuncCol) Accept(v Visitor) bool {
	if n.TypeName != nil {
		if !n.TypeName.Accept(v) {
			return false
		}
	}
	if n.Colexpr != nil {
		if !n.Colexpr.Node.Accept(v) {
			return false
		}
	}
	if n.Coldefexpr != nil {
		if !n.Coldefexpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *TypeName) Accept(v Visitor) bool {
	for _, item := range n.Names {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Typmods {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ArrayBounds {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ColumnDef) Accept(v Visitor) bool {
	if n.TypeName != nil {
		if !n.TypeName.Accept(v) {
			return false
		}
	}
	if n.RawDefault != nil {
		if !n.RawDefault.Node.Accept(v) {
			return false
		}
	}
	if n.CookedDefault != nil {
		if !n.CookedDefault.Node.Accept(v) {
			return false
		}
	}
	if n.IdentitySequence != nil {
		if !n.IdentitySequence.Accept(v) {
			return false
		}
	}
	if n.CollClause != nil {
		if !n.CollClause.Accept(v) {
			return false
		}
	}
	for _, item := range n.Constraints {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Fdwoptions {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *IndexElem) Accept(v Visitor) bool {
	if n.Expr != nil {
		if !n.Expr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Collation {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Opclass {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Opclassopts {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *Constraint) Accept(v Visitor) bool {
	if n.RawExpr != nil {
		if !n.RawExpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Keys {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Including {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Exclusions {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Options {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WhereClause != nil {
		if !n.WhereClause.Node.Accept(v) {
			return false
		}
	}
	if n.Pktable != nil {
		if !n.Pktable.Accept(v) {
			return false
		}
	}
	for _, item := range n.FkAttrs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.PkAttrs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.OldConpfeqop {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *DefElem) Accept(v Visitor) bool {
	if n.Arg != nil {
		if !n.Arg.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RangeTblEntry) Accept(v Visitor) bool {
	if n.Tablesample != nil {
		if !n.Tablesample.Accept(v) {
			return false
		}
	}
	if n.Subquery != nil {
		if !n.Subquery.Accept(v) {
			return false
		}
	}
	for _, item := range n.Joinaliasvars {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Joinleftcols {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Joinrightcols {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Functions {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Tablefunc != nil {
		if !n.Tablefunc.Accept(v) {
			return false
		}
	}
	for _, item := range n.ValuesLists {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Coltypes {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Coltypmods {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Colcollations {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Alias != nil {
		if !n.Alias.Accept(v) {
			return false
		}
	}
	if n.Eref != nil {
		if !n.Eref.Accept(v) {
			return false
		}
	}
	for _, item := range n.SecurityQuals {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RangeTblFunction) Accept(v Visitor) bool {
	if n.Funcexpr != nil {
		if !n.Funcexpr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Funccolnames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Funccoltypes {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Funccoltypmods {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Funccolcollations {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *TableSampleClause) Accept(v Visitor) bool {
	for _, item := range n.Args {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Repeatable != nil {
		if !n.Repeatable.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *WithCheckOption) Accept(v Visitor) bool {
	if n.Qual != nil {
		if !n.Qual.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *SortGroupClause) Accept(v Visitor) bool {
	return true
}

func (n *GroupingSet) Accept(v Visitor) bool {
	for _, item := range n.Content {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *WindowClause) Accept(v Visitor) bool {
	for _, item := range n.PartitionClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.OrderClause {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.StartOffset != nil {
		if !n.StartOffset.Node.Accept(v) {
			return false
		}
	}
	if n.EndOffset != nil {
		if !n.EndOffset.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *ObjectWithArgs) Accept(v Visitor) bool {
	for _, item := range n.Objname {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Objargs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *AccessPriv) Accept(v Visitor) bool {
	for _, item := range n.Cols {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CreateOpClassItem) Accept(v Visitor) bool {
	if n.Name != nil {
		if !n.Name.Accept(v) {
			return false
		}
	}
	for _, item := range n.OrderFamily {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.ClassArgs {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Storedtype != nil {
		if !n.Storedtype.Accept(v) {
			return false
		}
	}
	return true
}

func (n *TableLikeClause) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	return true
}

func (n *FunctionParameter) Accept(v Visitor) bool {
	if n.ArgType != nil {
		if !n.ArgType.Accept(v) {
			return false
		}
	}
	if n.Defexpr != nil {
		if !n.Defexpr.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *LockingClause) Accept(v Visitor) bool {
	for _, item := range n.LockedRels {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RowMarkClause) Accept(v Visitor) bool {
	return true
}

func (n *XmlSerialize) Accept(v Visitor) bool {
	if n.Expr != nil {
		if !n.Expr.Node.Accept(v) {
			return false
		}
	}
	if n.TypeName != nil {
		if !n.TypeName.Accept(v) {
			return false
		}
	}
	return true
}

func (n *WithClause) Accept(v Visitor) bool {
	for _, item := range n.Ctes {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *InferClause) Accept(v Visitor) bool {
	for _, item := range n.IndexElems {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WhereClause != nil {
		if !n.WhereClause.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *OnConflictClause) Accept(v Visitor) bool {
	if n.Infer != nil {
		if !n.Infer.Accept(v) {
			return false
		}
	}
	for _, item := range n.TargetList {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.WhereClause != nil {
		if !n.WhereClause.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *CommonTableExpr) Accept(v Visitor) bool {
	for _, item := range n.Aliascolnames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	if n.Ctequery != nil {
		if !n.Ctequery.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Ctecolnames {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Ctecoltypes {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Ctecoltypmods {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Ctecolcollations {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *RoleSpec) Accept(v Visitor) bool {
	return true
}

func (n *TriggerTransition) Accept(v Visitor) bool {
	return true
}

func (n *PartitionElem) Accept(v Visitor) bool {
	if n.Expr != nil {
		if !n.Expr.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Collation {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Opclass {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *PartitionSpec) Accept(v Visitor) bool {
	for _, item := range n.PartParams {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *PartitionBoundSpec) Accept(v Visitor) bool {
	for _, item := range n.Listdatums {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Lowerdatums {
		if !item.Node.Accept(v) {
			return false
		}
	}
	for _, item := range n.Upperdatums {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *PartitionRangeDatum) Accept(v Visitor) bool {
	if n.Value != nil {
		if !n.Value.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *PartitionCmd) Accept(v Visitor) bool {
	if n.Name != nil {
		if !n.Name.Accept(v) {
			return false
		}
	}
	if n.Bound != nil {
		if !n.Bound.Accept(v) {
			return false
		}
	}
	return true
}

func (n *VacuumRelation) Accept(v Visitor) bool {
	if n.Relation != nil {
		if !n.Relation.Accept(v) {
			return false
		}
	}
	for _, item := range n.VaCols {
		if !item.Node.Accept(v) {
			return false
		}
	}
	return true
}

func (n *InlineCodeBlock) Accept(v Visitor) bool {
	return true
}

func (n *CallContext) Accept(v Visitor) bool {
	return true
}

func (n *ScanToken) Accept(v Visitor) bool {
	return true
}
