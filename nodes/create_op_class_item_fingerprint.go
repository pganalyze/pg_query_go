// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateOpClassItem) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateOpClassItem")
	node.Args.Fingerprint(ctx, "Args")
	node.ClassArgs.Fingerprint(ctx, "ClassArgs")
	ctx.WriteString(strconv.Itoa(int(node.Itemtype)))
	node.Name.Fingerprint(ctx, "Name")
	ctx.WriteString(strconv.Itoa(int(node.Number)))
	node.OrderFamily.Fingerprint(ctx, "OrderFamily")

	if node.Storedtype != nil {
		node.Storedtype.Fingerprint(ctx, "Storedtype")
	}
}
