// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node NullTest) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("NullTest")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if node.Argisrow {
		ctx.WriteString("argisrow")
		ctx.WriteString(strconv.FormatBool(node.Argisrow))
	}

	if int(node.Nulltesttype) != 0 {
		ctx.WriteString("nulltesttype")
		ctx.WriteString(strconv.Itoa(int(node.Nulltesttype)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
