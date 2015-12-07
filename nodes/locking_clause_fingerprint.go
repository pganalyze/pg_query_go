// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node LockingClause) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("LOCKINGCLAUSE")

	for _, subNode := range node.LockedRels {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.NoWait))
	ctx.WriteString(strconv.Itoa(int(node.Strength)))
}
