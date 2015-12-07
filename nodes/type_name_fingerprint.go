// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TypeName) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("TYPENAME")

	for _, subNode := range node.ArrayBounds {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.Names {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.PctType))
	ctx.WriteString(strconv.FormatBool(node.Setof))

	for _, subNode := range node.Typmods {
		subNode.Fingerprint(ctx)
	}
}
