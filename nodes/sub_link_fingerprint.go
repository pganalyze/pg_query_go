// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SubLink) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SubLink")
	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.OperName {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.SubLinkType)))

	if node.Subselect != nil {
		node.Subselect.Fingerprint(ctx)
	}

	if node.Testexpr != nil {
		node.Testexpr.Fingerprint(ctx)
	}

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
