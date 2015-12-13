// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SetToDefault) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SetToDefault")
	ctx.WriteString(strconv.Itoa(int(node.Collation)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.TypeId)))
	ctx.WriteString(strconv.Itoa(int(node.TypeMod)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
