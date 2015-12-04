// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node SubLink) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SUBLINK")
	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.OperName {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.SubLinkType)))

	if node.Subselect != nil {
		node.Subselect.Fingerprint(ctx)
	}

	if node.Testexpr != nil {
		node.Testexpr.Fingerprint(ctx)
	}
}
