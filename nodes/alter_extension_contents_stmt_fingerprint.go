// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterExtensionContentsStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterExtensionContentsStmt")

	if node.Action != 0 {
		ctx.WriteString("action")
		ctx.WriteString(strconv.Itoa(int(node.Action)))
	}

	if node.Extname != nil {
		ctx.WriteString("extname")
		ctx.WriteString(*node.Extname)
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
}
