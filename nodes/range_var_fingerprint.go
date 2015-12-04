// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RangeVar) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RANGEVAR")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	if node.Catalogname != nil {
		io.WriteString(ctx.hash, *node.Catalogname)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.InhOpt)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Relname != nil {
		io.WriteString(ctx.hash, *node.Relname)
	}

	if node.Schemaname != nil {
		io.WriteString(ctx.hash, *node.Schemaname)
	}
}
