// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node BooleanTest) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("BooleanTest")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if int(node.Booltesttype) != 0 {
		ctx.WriteString("booltesttype")
		ctx.WriteString(strconv.Itoa(int(node.Booltesttype)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
