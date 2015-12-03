// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node LockRows) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "LockRows")

	for _, subNode := range node.RowMarks {
		subNode.Fingerprint(ctx)
	}
}
