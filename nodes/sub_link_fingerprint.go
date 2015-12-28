// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SubLink) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SubLink")
	// Intentionally ignoring node.Location for fingerprinting

	if len(node.OperName.Items) > 0 {
		ctx.WriteString("operName")
		node.OperName.Fingerprint(ctx, "OperName")
	}

	if int(node.SubLinkType) != 0 {
		ctx.WriteString("subLinkType")
		ctx.WriteString(strconv.Itoa(int(node.SubLinkType)))
	}

	if node.Subselect != nil {
		ctx.WriteString("subselect")
		node.Subselect.Fingerprint(ctx, "Subselect")
	}

	if node.Testexpr != nil {
		ctx.WriteString("testexpr")
		node.Testexpr.Fingerprint(ctx, "Testexpr")
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
