// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node PathKey) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PATHKEY")

	if node.PkEclass != nil {
		node.PkEclass.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.PkNullsFirst))
}
