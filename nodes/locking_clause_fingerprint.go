// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node LockingClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("LockingClause")
	if len(node.LockedRels.Items) > 0 {
		ctx.WriteString("lockedRels")
		node.LockedRels.Fingerprint(ctx, "LockedRels")
	}

	if node.NoWait {
		ctx.WriteString("noWait")
		ctx.WriteString(strconv.FormatBool(node.NoWait))
	}

	if int(node.Strength) != 0 {
		ctx.WriteString("strength")
		ctx.WriteString(strconv.Itoa(int(node.Strength)))
	}
}
