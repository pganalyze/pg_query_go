// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node XmlSerialize) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "XMLSERIALIZE")

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Xmloption)))
}
