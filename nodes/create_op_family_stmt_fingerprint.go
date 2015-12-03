// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateOpFamilyStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateOpFamilyStmt")

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx)
	}
}
