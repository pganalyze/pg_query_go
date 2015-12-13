// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ArrayExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ArrayExpr")
	ctx.WriteString(strconv.Itoa(int(node.ArrayCollid)))
	ctx.WriteString(strconv.Itoa(int(node.ArrayTypeid)))
	ctx.WriteString(strconv.Itoa(int(node.ElementTypeid)))

	for _, subNode := range node.Elements {
		subNode.Fingerprint(ctx, "Elements")
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.FormatBool(node.Multidims))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
