// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateOpFamilyStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATEOPFAMILYSTMT")

	if node.Amname != nil {
		io.WriteString(ctx.hash, *node.Amname)
	}

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx)
	}
}
