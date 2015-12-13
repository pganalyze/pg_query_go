// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Constraint) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Constraint")

	if node.AccessMethod != nil {
		ctx.WriteString(*node.AccessMethod)
	}

	if node.Conname != nil {
		ctx.WriteString(*node.Conname)
	}

	ctx.WriteString(strconv.Itoa(int(node.Contype)))

	if node.CookedExpr != nil {
		ctx.WriteString(*node.CookedExpr)
	}

	ctx.WriteString(strconv.FormatBool(node.Deferrable))

	for _, subNode := range node.Exclusions {
		subNode.Fingerprint(ctx, "Exclusions")
	}

	for _, subNode := range node.FkAttrs {
		subNode.Fingerprint(ctx, "FkAttrs")
	}

	ctx.WriteString(string(node.FkDelAction))
	ctx.WriteString(string(node.FkMatchtype))
	ctx.WriteString(string(node.FkUpdAction))

	if node.Indexname != nil {
		ctx.WriteString(*node.Indexname)
	}

	if node.Indexspace != nil {
		ctx.WriteString(*node.Indexspace)
	}

	ctx.WriteString(strconv.FormatBool(node.Initdeferred))
	ctx.WriteString(strconv.FormatBool(node.InitiallyValid))
	ctx.WriteString(strconv.FormatBool(node.IsNoInherit))

	for _, subNode := range node.Keys {
		subNode.Fingerprint(ctx, "Keys")
	}

	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.OldConpfeqop {
		subNode.Fingerprint(ctx, "OldConpfeqop")
	}

	ctx.WriteString(strconv.Itoa(int(node.OldPktableOid)))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}

	for _, subNode := range node.PkAttrs {
		subNode.Fingerprint(ctx, "PkAttrs")
	}

	if node.Pktable != nil {
		node.Pktable.Fingerprint(ctx, "Pktable")
	}

	if node.RawExpr != nil {
		node.RawExpr.Fingerprint(ctx, "RawExpr")
	}

	ctx.WriteString(strconv.FormatBool(node.SkipValidation))

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}
}
