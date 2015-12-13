// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node NullTest) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("NullTest")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Argisrow))
	ctx.WriteString(strconv.Itoa(int(node.Nulltesttype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
