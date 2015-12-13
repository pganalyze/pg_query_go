// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node BlockIdData) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BlockIdData")
	ctx.WriteString(strconv.Itoa(int(node.BiHi)))
	ctx.WriteString(strconv.Itoa(int(node.BiLo)))
}
