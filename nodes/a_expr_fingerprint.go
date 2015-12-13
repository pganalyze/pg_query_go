// Auto-generated - DO NOT EDIT

package pg_query

import (
	"sort"
	"strconv"
)

func (node A_Expr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_Expr")
	ctx.WriteString(strconv.Itoa(int(node.Kind)))

	for _, subNode := range node.Lexpr {
		subNode.Fingerprint(ctx, "Lexpr")
	}

	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.Name {
		subNode.Fingerprint(ctx, "Name")
	}

	var rexprFingerprints FingerprintSubContextSlice

	for _, subNode := range node.Rexpr {
		subCtx := FingerprintSubContext{}
		subNode.Fingerprint(&subCtx, "Rexpr")
		rexprFingerprints.AddIfUnique(subCtx)
	}

	sort.Sort(rexprFingerprints)

	for _, fingerprint := range rexprFingerprints {
		for _, part := range fingerprint.parts {
			ctx.WriteString(part)
		}
	}
}
