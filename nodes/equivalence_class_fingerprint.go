// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node EquivalenceClass) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("EQUIVALENCECLASS")
	ctx.WriteString(strconv.FormatBool(node.EcBelowOuterJoin))
	ctx.WriteString(strconv.FormatBool(node.EcBroken))

	for _, subNode := range node.EcDerives {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.EcHasConst))
	ctx.WriteString(strconv.FormatBool(node.EcHasVolatile))

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
