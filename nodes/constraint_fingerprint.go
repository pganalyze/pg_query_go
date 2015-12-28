// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Constraint) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Constraint")

	if node.AccessMethod != nil {
		ctx.WriteString("access_method")
		ctx.WriteString(*node.AccessMethod)
	}

	if node.Conname != nil {
		ctx.WriteString("conname")
		ctx.WriteString(*node.Conname)
	}

	if int(node.Contype) != 0 {
		ctx.WriteString("contype")
		ctx.WriteString(strconv.Itoa(int(node.Contype)))
	}

	if node.CookedExpr != nil {
		ctx.WriteString("cooked_expr")
		ctx.WriteString(*node.CookedExpr)
	}

	if node.Deferrable {
		ctx.WriteString("deferrable")
		ctx.WriteString(strconv.FormatBool(node.Deferrable))
	}

	if len(node.Exclusions.Items) > 0 {
		ctx.WriteString("exclusions")
		node.Exclusions.Fingerprint(ctx, "Exclusions")
	}

	if len(node.FkAttrs.Items) > 0 {
		ctx.WriteString("fk_attrs")
		node.FkAttrs.Fingerprint(ctx, "FkAttrs")
	}

	if node.FkDelAction != 0 {
		ctx.WriteString("fk_del_action")
		ctx.WriteString(string(node.FkDelAction))

	}

	if node.FkMatchtype != 0 {
		ctx.WriteString("fk_matchtype")
		ctx.WriteString(string(node.FkMatchtype))

	}

	if node.FkUpdAction != 0 {
		ctx.WriteString("fk_upd_action")
		ctx.WriteString(string(node.FkUpdAction))

	}

	if node.Indexname != nil {
		ctx.WriteString("indexname")
		ctx.WriteString(*node.Indexname)
	}

	if node.Indexspace != nil {
		ctx.WriteString("indexspace")
		ctx.WriteString(*node.Indexspace)
	}

	if node.Initdeferred {
		ctx.WriteString("initdeferred")
		ctx.WriteString(strconv.FormatBool(node.Initdeferred))
	}

	if node.InitiallyValid {
		ctx.WriteString("initially_valid")
		ctx.WriteString(strconv.FormatBool(node.InitiallyValid))
	}

	if node.IsNoInherit {
		ctx.WriteString("is_no_inherit")
		ctx.WriteString(strconv.FormatBool(node.IsNoInherit))
	}

	if len(node.Keys.Items) > 0 {
		ctx.WriteString("keys")
		node.Keys.Fingerprint(ctx, "Keys")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if len(node.OldConpfeqop.Items) > 0 {
		ctx.WriteString("old_conpfeqop")
		node.OldConpfeqop.Fingerprint(ctx, "OldConpfeqop")
	}

	if node.OldPktableOid != 0 {
		ctx.WriteString("old_pktable_oid")
		ctx.WriteString(strconv.Itoa(int(node.OldPktableOid)))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if len(node.PkAttrs.Items) > 0 {
		ctx.WriteString("pk_attrs")
		node.PkAttrs.Fingerprint(ctx, "PkAttrs")
	}

	if node.Pktable != nil {
		ctx.WriteString("pktable")
		node.Pktable.Fingerprint(ctx, "Pktable")
	}

	if node.RawExpr != nil {
		ctx.WriteString("raw_expr")
		node.RawExpr.Fingerprint(ctx, "RawExpr")
	}

	if node.SkipValidation {
		ctx.WriteString("skip_validation")
		ctx.WriteString(strconv.FormatBool(node.SkipValidation))
	}

	if node.WhereClause != nil {
		ctx.WriteString("where_clause")
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}
}
