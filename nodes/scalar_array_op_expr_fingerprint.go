// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ScalarArrayOpExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ScalarArrayOpExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Opfuncid)))
	ctx.WriteString(strconv.Itoa(int(node.Opno)))
	ctx.WriteString(strconv.FormatBool(node.UseOr))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
