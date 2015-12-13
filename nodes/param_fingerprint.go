// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Param) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("Param")
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Paramcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Paramid)))
	ctx.WriteString(strconv.Itoa(int(node.Paramkind)))
	ctx.WriteString(strconv.Itoa(int(node.Paramtype)))
	ctx.WriteString(strconv.Itoa(int(node.Paramtypmod)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
