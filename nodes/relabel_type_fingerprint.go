// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RelabelType) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RelabelType")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Relabelformat)))
	ctx.WriteString(strconv.Itoa(int(node.Resultcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	ctx.WriteString(strconv.Itoa(int(node.Resulttypmod)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
