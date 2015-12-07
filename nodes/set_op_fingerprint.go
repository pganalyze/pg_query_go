// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SetOp) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SETOP")
	ctx.WriteString(strconv.Itoa(int(node.Cmd)))
	ctx.WriteString(strconv.Itoa(int(node.Strategy)))
}
