// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CaseExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CaseExpr")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Casecollid)))
	ctx.WriteString(strconv.Itoa(int(node.Casetype)))

	if node.Defresult != nil {
		node.Defresult.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
