// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Param) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PARAM")
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Paramkind)))
}
