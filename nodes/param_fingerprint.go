// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Param) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Param")
	// Intentionally ignoring node.Location for fingerprinting

	if node.Paramcollid != 0 {
		ctx.WriteString("paramcollid")
		ctx.WriteString(strconv.Itoa(int(node.Paramcollid)))
	}

	if node.Paramid != 0 {
		ctx.WriteString("paramid")
		ctx.WriteString(strconv.Itoa(int(node.Paramid)))
	}

	if int(node.Paramkind) != 0 {
		ctx.WriteString("paramkind")
		ctx.WriteString(strconv.Itoa(int(node.Paramkind)))
	}

	if node.Paramtype != 0 {
		ctx.WriteString("paramtype")
		ctx.WriteString(strconv.Itoa(int(node.Paramtype)))
	}

	if node.Paramtypmod != 0 {
		ctx.WriteString("paramtypmod")
		ctx.WriteString(strconv.Itoa(int(node.Paramtypmod)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
