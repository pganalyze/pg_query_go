// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SubLink) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SubLink")
	// Intentionally ignoring node.Location for fingerprinting

	node.OperName.Fingerprint(ctx, "OperName")
	ctx.WriteString(strconv.Itoa(int(node.SubLinkType)))

	if node.Subselect != nil {
		node.Subselect.Fingerprint(ctx, "Subselect")
	}

	if node.Testexpr != nil {
		node.Testexpr.Fingerprint(ctx, "Testexpr")
	}

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
