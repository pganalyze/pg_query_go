// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TypeName) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TypeName")

	for _, subNode := range node.ArrayBounds {
		subNode.Fingerprint(ctx, "ArrayBounds")
	}

	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.Names {
		subNode.Fingerprint(ctx, "Names")
	}

	ctx.WriteString(strconv.FormatBool(node.PctType))
	ctx.WriteString(strconv.FormatBool(node.Setof))
	ctx.WriteString(strconv.Itoa(int(node.TypeOid)))
	ctx.WriteString(strconv.Itoa(int(node.Typemod)))

	for _, subNode := range node.Typmods {
		subNode.Fingerprint(ctx, "Typmods")
	}
}
