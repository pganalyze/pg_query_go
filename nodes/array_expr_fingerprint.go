// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ArrayExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ArrayExpr")

	if node.ArrayCollid != 0 {
		ctx.WriteString("array_collid")
		ctx.WriteString(strconv.Itoa(int(node.ArrayCollid)))
	}

	if node.ArrayTypeid != 0 {
		ctx.WriteString("array_typeid")
		ctx.WriteString(strconv.Itoa(int(node.ArrayTypeid)))
	}

	if node.ElementTypeid != 0 {
		ctx.WriteString("element_typeid")
		ctx.WriteString(strconv.Itoa(int(node.ElementTypeid)))
	}

	if len(node.Elements.Items) > 0 {
		ctx.WriteString("elements")
		node.Elements.Fingerprint(ctx, "Elements")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Multidims {
		ctx.WriteString("multidims")
		ctx.WriteString(strconv.FormatBool(node.Multidims))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
