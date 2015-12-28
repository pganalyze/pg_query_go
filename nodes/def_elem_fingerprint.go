// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DefElem) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DefElem")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if int(node.Defaction) != 0 {
		ctx.WriteString("defaction")
		ctx.WriteString(strconv.Itoa(int(node.Defaction)))
	}

	if node.Defname != nil {
		ctx.WriteString("defname")
		ctx.WriteString(*node.Defname)
	}

	if node.Defnamespace != nil {
		ctx.WriteString("defnamespace")
		ctx.WriteString(*node.Defnamespace)
	}
}
