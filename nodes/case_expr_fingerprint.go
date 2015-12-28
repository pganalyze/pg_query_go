// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CaseExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CaseExpr")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if node.Casecollid != 0 {
		ctx.WriteString("casecollid")
		ctx.WriteString(strconv.Itoa(int(node.Casecollid)))
	}

	if node.Casetype != 0 {
		ctx.WriteString("casetype")
		ctx.WriteString(strconv.Itoa(int(node.Casetype)))
	}

	if node.Defresult != nil {
		ctx.WriteString("defresult")
		node.Defresult.Fingerprint(ctx, "Defresult")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
