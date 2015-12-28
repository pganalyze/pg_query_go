// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ColumnDef) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ColumnDef")

	if node.CollClause != nil {
		ctx.WriteString("collClause")
		node.CollClause.Fingerprint(ctx, "CollClause")
	}

	if node.CollOid != 0 {
		ctx.WriteString("collOid")
		ctx.WriteString(strconv.Itoa(int(node.CollOid)))
	}

	if node.Colname != nil {
		ctx.WriteString("colname")
		ctx.WriteString(*node.Colname)
	}

	if len(node.Constraints.Items) > 0 {
		ctx.WriteString("constraints")
		node.Constraints.Fingerprint(ctx, "Constraints")
	}

	if node.CookedDefault != nil {
		ctx.WriteString("cooked_default")
		node.CookedDefault.Fingerprint(ctx, "CookedDefault")
	}

	if len(node.Fdwoptions.Items) > 0 {
		ctx.WriteString("fdwoptions")
		node.Fdwoptions.Fingerprint(ctx, "Fdwoptions")
	}

	if node.Inhcount != 0 {
		ctx.WriteString("inhcount")
		ctx.WriteString(strconv.Itoa(int(node.Inhcount)))
	}

	if node.IsFromType {
		ctx.WriteString("is_from_type")
		ctx.WriteString(strconv.FormatBool(node.IsFromType))
	}

	if node.IsLocal {
		ctx.WriteString("is_local")
		ctx.WriteString(strconv.FormatBool(node.IsLocal))
	}

	if node.IsNotNull {
		ctx.WriteString("is_not_null")
		ctx.WriteString(strconv.FormatBool(node.IsNotNull))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.RawDefault != nil {
		ctx.WriteString("raw_default")
		node.RawDefault.Fingerprint(ctx, "RawDefault")
	}

	if node.Storage != 0 {
		ctx.WriteString("storage")
		ctx.WriteString(string(node.Storage))

	}

	if node.TypeName != nil {
		ctx.WriteString("typeName")
		node.TypeName.Fingerprint(ctx, "TypeName")
	}
}
