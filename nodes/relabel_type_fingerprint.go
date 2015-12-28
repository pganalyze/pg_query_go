// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RelabelType) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RelabelType")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if int(node.Relabelformat) != 0 {
		ctx.WriteString("relabelformat")
		ctx.WriteString(strconv.Itoa(int(node.Relabelformat)))
	}

	if node.Resultcollid != 0 {
		ctx.WriteString("resultcollid")
		ctx.WriteString(strconv.Itoa(int(node.Resultcollid)))
	}

	if node.Resulttype != 0 {
		ctx.WriteString("resulttype")
		ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	}

	if node.Resulttypmod != 0 {
		ctx.WriteString("resulttypmod")
		ctx.WriteString(strconv.Itoa(int(node.Resulttypmod)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
