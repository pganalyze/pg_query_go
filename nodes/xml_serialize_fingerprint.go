// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node XmlSerialize) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "XmlSerialize")
	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}
}
