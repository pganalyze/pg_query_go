// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Constraint) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CONSTRAINT")

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
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.FkAttrs {
		subNode.Fingerprint(ctx)
	}

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
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.OldConpfeqop {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.PkAttrs {
		subNode.Fingerprint(ctx)
	}

	if node.Pktable != nil {
		node.Pktable.Fingerprint(ctx)
	}

	if node.RawExpr != nil {
		node.RawExpr.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.SkipValidation))

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}
}
