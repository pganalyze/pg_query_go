// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node XmlSerialize) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("XmlSerialize")

	if node.Expr != nil {
		ctx.WriteString("expr")
		node.Expr.Fingerprint(ctx, "Expr")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.TypeName != nil {
		ctx.WriteString("typeName")
		node.TypeName.Fingerprint(ctx, "TypeName")
	}

	if int(node.Xmloption) != 0 {
		ctx.WriteString("xmloption")
		ctx.WriteString(strconv.Itoa(int(node.Xmloption)))
	}
}
