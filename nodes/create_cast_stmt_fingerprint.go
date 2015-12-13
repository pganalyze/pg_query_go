// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateCastStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateCastStmt")
	ctx.WriteString(strconv.Itoa(int(node.Context)))

	if node.Func != nil {
		node.Func.Fingerprint(ctx, "Func")
	}

	ctx.WriteString(strconv.FormatBool(node.Inout))

	if node.Sourcetype != nil {
		node.Sourcetype.Fingerprint(ctx, "Sourcetype")
	}

	if node.Targettype != nil {
		node.Targettype.Fingerprint(ctx, "Targettype")
	}
}
