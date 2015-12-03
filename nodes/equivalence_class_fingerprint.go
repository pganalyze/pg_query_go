// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node EquivalenceClass) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "EquivalenceClass")

	for _, subNode := range node.EcDerives {
		subNode.Fingerprint(ctx)
	}

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
