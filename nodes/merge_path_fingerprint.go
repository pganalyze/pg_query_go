// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node MergePath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("MERGEPATH")

	for _, subNode := range node.Innersortkeys {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.MaterializeInner))

	for _, subNode := range node.Outersortkeys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.PathMergeclauses {
		subNode.Fingerprint(ctx)
	}
}
