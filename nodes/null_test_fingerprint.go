// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node NullTest) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("NullTest")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	ctx.WriteString(strconv.FormatBool(node.Argisrow))
	ctx.WriteString(strconv.Itoa(int(node.Nulltesttype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
