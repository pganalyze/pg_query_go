// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node NullTest) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("NULLTEST")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Argisrow))
	ctx.WriteString(strconv.Itoa(int(node.Nulltesttype)))
}
