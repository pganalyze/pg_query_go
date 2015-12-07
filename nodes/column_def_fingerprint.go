// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ColumnDef) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("COLUMNDEF")

	if node.CollClause != nil {
		node.CollClause.Fingerprint(ctx)
	}

	if node.Colname != nil {
		ctx.WriteString(*node.Colname)
	}

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}

	if node.CookedDefault != nil {
		node.CookedDefault.Fingerprint(ctx)
	}

	for _, subNode := range node.Fdwoptions {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.IsFromType))
	ctx.WriteString(strconv.FormatBool(node.IsLocal))
	ctx.WriteString(strconv.FormatBool(node.IsNotNull))
	// Intentionally ignoring node.Location for fingerprinting

	if node.RawDefault != nil {
		node.RawDefault.Fingerprint(ctx)
	}

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}
}
