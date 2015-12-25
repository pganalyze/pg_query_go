// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TypeName) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TypeName")
	node.ArrayBounds.Fingerprint(ctx, "ArrayBounds")
	// Intentionally ignoring node.Location for fingerprinting

	node.Names.Fingerprint(ctx, "Names")
	ctx.WriteString(strconv.FormatBool(node.PctType))
	ctx.WriteString(strconv.FormatBool(node.Setof))
	ctx.WriteString(strconv.Itoa(int(node.TypeOid)))
	ctx.WriteString(strconv.Itoa(int(node.Typemod)))
	node.Typmods.Fingerprint(ctx, "Typmods")
}
