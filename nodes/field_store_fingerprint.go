// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FieldStore) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FieldStore")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if len(node.Fieldnums.Items) > 0 {
		ctx.WriteString("fieldnums")
		node.Fieldnums.Fingerprint(ctx, "Fieldnums")
	}

	if len(node.Newvals.Items) > 0 {
		ctx.WriteString("newvals")
		node.Newvals.Fingerprint(ctx, "Newvals")
	}

	if node.Resulttype != 0 {
		ctx.WriteString("resulttype")
		ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
