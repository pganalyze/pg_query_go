// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node XmlSerialize) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("XmlSerialize")

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx, "Expr")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx, "TypeName")
	}

	ctx.WriteString(strconv.Itoa(int(node.Xmloption)))
}
