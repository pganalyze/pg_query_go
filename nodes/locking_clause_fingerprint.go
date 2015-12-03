// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node LockingClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "LockingClause")

	for _, subNode := range node.LockedRels {
		subNode.Fingerprint(ctx)
	}
}
