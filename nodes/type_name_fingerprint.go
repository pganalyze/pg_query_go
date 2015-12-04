// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node TypeName) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TYPENAME")

	for _, subNode := range node.ArrayBounds {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.Names {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.PctType))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Setof))

	for _, subNode := range node.Typmods {
		subNode.Fingerprint(ctx)
	}
}
