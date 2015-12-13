// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node XmlSerialize) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("XmlSerialize")

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Xmloption)))
}
