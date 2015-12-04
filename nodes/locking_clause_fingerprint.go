// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node LockingClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "LOCKINGCLAUSE")

	for _, subNode := range node.LockedRels {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.NoWait))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Strength)))
}
