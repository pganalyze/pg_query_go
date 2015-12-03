// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterOpFamilyStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterOpFamilyStmt")

	for _, subNode := range node.Items {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx)
	}
}
