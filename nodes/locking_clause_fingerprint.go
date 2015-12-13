// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node LockingClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("LockingClause")

	for _, subNode := range node.LockedRels {
		subNode.Fingerprint(ctx, "LockedRels")
	}

	ctx.WriteString(strconv.FormatBool(node.NoWait))
	ctx.WriteString(strconv.Itoa(int(node.Strength)))
}
