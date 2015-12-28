// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node OpExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("OpExpr")
	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if node.Inputcollid != 0 {
		ctx.WriteString("inputcollid")
		ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Opcollid != 0 {
		ctx.WriteString("opcollid")
		ctx.WriteString(strconv.Itoa(int(node.Opcollid)))
	}

	if node.Opfuncid != 0 {
		ctx.WriteString("opfuncid")
		ctx.WriteString(strconv.Itoa(int(node.Opfuncid)))
	}

	if node.Opno != 0 {
		ctx.WriteString("opno")
		ctx.WriteString(strconv.Itoa(int(node.Opno)))
	}

	if node.Opresulttype != 0 {
		ctx.WriteString("opresulttype")
		ctx.WriteString(strconv.Itoa(int(node.Opresulttype)))
	}

	if node.Opretset {
		ctx.WriteString("opretset")
		ctx.WriteString(strconv.FormatBool(node.Opretset))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
