// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node SubLink) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SubLink")

	for _, subNode := range node.OperName {
		subNode.Fingerprint(ctx)
	}

	if node.Subselect != nil {
		node.Subselect.Fingerprint(ctx)
	}

	if node.Testexpr != nil {
		node.Testexpr.Fingerprint(ctx)
	}
}
