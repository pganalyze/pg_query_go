// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DefineStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DefineStmt")
	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if len(node.Definition.Items) > 0 {
		ctx.WriteString("definition")
		node.Definition.Fingerprint(ctx, "Definition")
	}

	if len(node.Defnames.Items) > 0 {
		ctx.WriteString("defnames")
		node.Defnames.Fingerprint(ctx, "Defnames")
	}

	if int(node.Kind) != 0 {
		ctx.WriteString("kind")
		ctx.WriteString(strconv.Itoa(int(node.Kind)))
	}

	if node.Oldstyle {
		ctx.WriteString("oldstyle")
		ctx.WriteString(strconv.FormatBool(node.Oldstyle))
	}
}
