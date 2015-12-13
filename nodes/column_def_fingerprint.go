// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ColumnDef) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ColumnDef")

	if node.CollClause != nil {
		node.CollClause.Fingerprint(ctx, "CollClause")
	}

	ctx.WriteString(strconv.Itoa(int(node.CollOid)))

	if node.Colname != nil {
		ctx.WriteString(*node.Colname)
	}

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx, "Constraints")
	}

	if node.CookedDefault != nil {
		node.CookedDefault.Fingerprint(ctx, "CookedDefault")
	}

	for _, subNode := range node.Fdwoptions {
		subNode.Fingerprint(ctx, "Fdwoptions")
	}

	ctx.WriteString(strconv.Itoa(int(node.Inhcount)))
	ctx.WriteString(strconv.FormatBool(node.IsFromType))
	ctx.WriteString(strconv.FormatBool(node.IsLocal))
	ctx.WriteString(strconv.FormatBool(node.IsNotNull))
	// Intentionally ignoring node.Location for fingerprinting

	if node.RawDefault != nil {
		node.RawDefault.Fingerprint(ctx, "RawDefault")
	}

	ctx.WriteString(string(node.Storage))

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx, "TypeName")
	}
}
