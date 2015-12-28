// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SecLabelStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SecLabelStmt")

	if node.Label != nil {
		ctx.WriteString("label")
		ctx.WriteString(*node.Label)
	}

	if len(node.Objargs.Items) > 0 {
		ctx.WriteString("objargs")
		node.Objargs.Fingerprint(ctx, "Objargs")
	}

	if len(node.Objname.Items) > 0 {
		ctx.WriteString("objname")
		node.Objname.Fingerprint(ctx, "Objname")
	}

	if int(node.Objtype) != 0 {
		ctx.WriteString("objtype")
		ctx.WriteString(strconv.Itoa(int(node.Objtype)))
	}

	if node.Provider != nil {
		ctx.WriteString("provider")
		ctx.WriteString(*node.Provider)
	}
}
