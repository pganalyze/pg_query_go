// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Hash) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("HASH")
	ctx.WriteString(strconv.FormatBool(node.SkewInherit))
}
