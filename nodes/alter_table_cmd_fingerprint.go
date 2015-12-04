// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterTableCmd) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTER TABLE CMD")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Behavior)))

	if node.Def != nil {
		node.Def.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Subtype)))
}
