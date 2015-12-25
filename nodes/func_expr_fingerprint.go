// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FuncExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FuncExpr")
	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.Itoa(int(node.Funccollid)))
	ctx.WriteString(strconv.Itoa(int(node.Funcformat)))
	ctx.WriteString(strconv.Itoa(int(node.Funcid)))
	ctx.WriteString(strconv.Itoa(int(node.Funcresulttype)))
	ctx.WriteString(strconv.FormatBool(node.Funcretset))
	ctx.WriteString(strconv.FormatBool(node.Funcvariadic))
	ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
