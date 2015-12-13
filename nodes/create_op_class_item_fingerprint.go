// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateOpClassItem) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateOpClassItem")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx, "Args")
	}

	for _, subNode := range node.ClassArgs {
		subNode.Fingerprint(ctx, "ClassArgs")
	}

	ctx.WriteString(strconv.Itoa(int(node.Itemtype)))

	for _, subNode := range node.Name {
		subNode.Fingerprint(ctx, "Name")
	}

	ctx.WriteString(strconv.Itoa(int(node.Number)))

	for _, subNode := range node.OrderFamily {
		subNode.Fingerprint(ctx, "OrderFamily")
	}

	if node.Storedtype != nil {
		node.Storedtype.Fingerprint(ctx, "Storedtype")
	}
}
