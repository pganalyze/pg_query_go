// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FieldSelect) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FieldSelect")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	ctx.WriteString(strconv.Itoa(int(node.Fieldnum)))
	ctx.WriteString(strconv.Itoa(int(node.Resultcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	ctx.WriteString(strconv.Itoa(int(node.Resulttypmod)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
