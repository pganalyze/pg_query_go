// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DefElem) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DefElem")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	ctx.WriteString(strconv.Itoa(int(node.Defaction)))

	if node.Defname != nil {
		ctx.WriteString(*node.Defname)
	}

	if node.Defnamespace != nil {
		ctx.WriteString(*node.Defnamespace)
	}
}
