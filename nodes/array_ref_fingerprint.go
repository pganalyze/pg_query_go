// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ArrayRef) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ArrayRef")

	if node.Refarraytype != 0 {
		ctx.WriteString("refarraytype")
		ctx.WriteString(strconv.Itoa(int(node.Refarraytype)))
	}

	if node.Refassgnexpr != nil {
		ctx.WriteString("refassgnexpr")
		node.Refassgnexpr.Fingerprint(ctx, "Refassgnexpr")
	}

	if node.Refcollid != 0 {
		ctx.WriteString("refcollid")
		ctx.WriteString(strconv.Itoa(int(node.Refcollid)))
	}

	if node.Refelemtype != 0 {
		ctx.WriteString("refelemtype")
		ctx.WriteString(strconv.Itoa(int(node.Refelemtype)))
	}

	if node.Refexpr != nil {
		ctx.WriteString("refexpr")
		node.Refexpr.Fingerprint(ctx, "Refexpr")
	}

	if len(node.Reflowerindexpr.Items) > 0 {
		ctx.WriteString("reflowerindexpr")
		node.Reflowerindexpr.Fingerprint(ctx, "Reflowerindexpr")
	}

	if node.Reftypmod != 0 {
		ctx.WriteString("reftypmod")
		ctx.WriteString(strconv.Itoa(int(node.Reftypmod)))
	}

	if len(node.Refupperindexpr.Items) > 0 {
		ctx.WriteString("refupperindexpr")
		node.Refupperindexpr.Fingerprint(ctx, "Refupperindexpr")
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
