// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WithCheckOption) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("WithCheckOption")

	if node.Cascaded {
		ctx.WriteString("cascaded")
		ctx.WriteString(strconv.FormatBool(node.Cascaded))
	}

	if node.Qual != nil {
		subCtx := FingerprintSubContext{}
		node.Qual.Fingerprint(&subCtx, "Qual")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("qual")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Viewname != nil {
		ctx.WriteString("viewname")
		ctx.WriteString(*node.Viewname)
	}
}
