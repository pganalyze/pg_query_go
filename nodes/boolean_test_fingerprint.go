// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node BooleanTest) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BooleanTest")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Booltesttype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
