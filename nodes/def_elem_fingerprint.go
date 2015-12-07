// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DefElem) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DEFELEM")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Defaction)))

	if node.Defname != nil {
		ctx.WriteString(*node.Defname)
	}

	if node.Defnamespace != nil {
		ctx.WriteString(*node.Defnamespace)
	}
}
