// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node MergeScanSelCache) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("MERGESCANSELCACHE")
	ctx.WriteString(strconv.FormatBool(node.NullsFirst))
}
