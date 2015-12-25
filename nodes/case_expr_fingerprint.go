// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CaseExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CaseExpr")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.Itoa(int(node.Casecollid)))
	ctx.WriteString(strconv.Itoa(int(node.Casetype)))

	if node.Defresult != nil {
		node.Defresult.Fingerprint(ctx, "Defresult")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
