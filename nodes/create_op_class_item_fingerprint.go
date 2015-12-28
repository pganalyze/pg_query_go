// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateOpClassItem) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateOpClassItem")
	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if len(node.ClassArgs.Items) > 0 {
		ctx.WriteString("class_args")
		node.ClassArgs.Fingerprint(ctx, "ClassArgs")
	}

	if node.Itemtype != 0 {
		ctx.WriteString("itemtype")
		ctx.WriteString(strconv.Itoa(int(node.Itemtype)))
	}

	if len(node.Name.Items) > 0 {
		ctx.WriteString("name")
		node.Name.Fingerprint(ctx, "Name")
	}

	if node.Number != 0 {
		ctx.WriteString("number")
		ctx.WriteString(strconv.Itoa(int(node.Number)))
	}

	if len(node.OrderFamily.Items) > 0 {
		ctx.WriteString("order_family")
		node.OrderFamily.Fingerprint(ctx, "OrderFamily")
	}

	if node.Storedtype != nil {
		ctx.WriteString("storedtype")
		node.Storedtype.Fingerprint(ctx, "Storedtype")
	}
}
