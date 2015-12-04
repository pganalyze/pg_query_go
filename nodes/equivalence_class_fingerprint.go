// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node EquivalenceClass) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "EQUIVALENCECLASS")
	io.WriteString(ctx.hash, strconv.FormatBool(node.EcBelowOuterJoin))
	io.WriteString(ctx.hash, strconv.FormatBool(node.EcBroken))

	for _, subNode := range node.EcDerives {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.EcHasConst))
	io.WriteString(ctx.hash, strconv.FormatBool(node.EcHasVolatile))

	for _, subNode := range node.EcMembers {
		subNode.Fingerprint(ctx)
	}

	if node.EcMerged != nil {
		node.EcMerged.Fingerprint(ctx)
	}

	for _, subNode := range node.EcOpfamilies {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.EcSources {
		subNode.Fingerprint(ctx)
	}
}
