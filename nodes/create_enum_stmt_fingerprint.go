// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateEnumStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateEnumStmt")

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Vals {
		subNode.Fingerprint(ctx)
	}
}
