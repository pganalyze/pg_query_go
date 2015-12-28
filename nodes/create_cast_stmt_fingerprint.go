// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateCastStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateCastStmt")

	if int(node.Context) != 0 {
		ctx.WriteString("context")
		ctx.WriteString(strconv.Itoa(int(node.Context)))
	}

	if node.Func != nil {
		ctx.WriteString("func")
		node.Func.Fingerprint(ctx, "Func")
	}

	if node.Inout {
		ctx.WriteString("inout")
		ctx.WriteString(strconv.FormatBool(node.Inout))
	}

	if node.Sourcetype != nil {
		ctx.WriteString("sourcetype")
		node.Sourcetype.Fingerprint(ctx, "Sourcetype")
	}

	if node.Targettype != nil {
		ctx.WriteString("targettype")
		node.Targettype.Fingerprint(ctx, "Targettype")
	}
}
