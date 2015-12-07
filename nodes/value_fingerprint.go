package pg_query

import "strconv"

func (node Value) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString(strconv.Itoa(int(node.Type)))
	ctx.WriteString(node.Str)
	ctx.WriteString(strconv.Itoa(node.Ival))
}
