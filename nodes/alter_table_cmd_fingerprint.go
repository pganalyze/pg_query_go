// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterTableCmd) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterTableCmd")
	if node.Def != nil {
		node.Def.Fingerprint(ctx)
	}
}
