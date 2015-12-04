// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node PrivGrantee) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PRIVGRANTEE")

	if node.Rolname != nil {
		io.WriteString(ctx.hash, *node.Rolname)
	}
}
