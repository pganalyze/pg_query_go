// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ScalarArrayOpExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ScalarArrayOpExpr")
	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Opfuncid)))
	ctx.WriteString(strconv.Itoa(int(node.Opno)))
	ctx.WriteString(strconv.FormatBool(node.UseOr))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
