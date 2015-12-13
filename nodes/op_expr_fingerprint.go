// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node OpExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("OpExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Opcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Opfuncid)))
	ctx.WriteString(strconv.Itoa(int(node.Opno)))
	ctx.WriteString(strconv.Itoa(int(node.Opresulttype)))
	ctx.WriteString(strconv.FormatBool(node.Opretset))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
