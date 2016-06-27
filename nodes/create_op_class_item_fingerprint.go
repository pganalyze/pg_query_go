// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateOpClassItem) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("CreateOpClassItem")

	if len(node.Args.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Args.Fingerprint(&subCtx, node, "Args")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("args")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.ClassArgs.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.ClassArgs.Fingerprint(&subCtx, node, "ClassArgs")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("class_args")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Itemtype != 0 {
		ctx.WriteString("itemtype")
		ctx.WriteString(strconv.Itoa(int(node.Itemtype)))
	}

	if len(node.Name.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Name.Fingerprint(&subCtx, node, "Name")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("name")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Number != 0 {
		ctx.WriteString("number")
		ctx.WriteString(strconv.Itoa(int(node.Number)))
	}

	if len(node.OrderFamily.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.OrderFamily.Fingerprint(&subCtx, node, "OrderFamily")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("order_family")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Storedtype != nil {
		subCtx := FingerprintSubContext{}
		node.Storedtype.Fingerprint(&subCtx, node, "Storedtype")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("storedtype")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
