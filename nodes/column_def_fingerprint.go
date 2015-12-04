// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ColumnDef) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "COLUMNDEF")

	if node.CollClause != nil {
		node.CollClause.Fingerprint(ctx)
	}

	if node.Colname != nil {
		io.WriteString(ctx.hash, *node.Colname)
	}

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}

	if node.CookedDefault != nil {
		node.CookedDefault.Fingerprint(ctx)
	}

	for _, subNode := range node.Fdwoptions {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsFromType))
	io.WriteString(ctx.hash, strconv.FormatBool(node.IsLocal))
	io.WriteString(ctx.hash, strconv.FormatBool(node.IsNotNull))
	// Intentionally ignoring node.Location for fingerprinting

	if node.RawDefault != nil {
		node.RawDefault.Fingerprint(ctx)
	}

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}
}
