// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node XmlExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("XMLEXPR")

	for _, subNode := range node.ArgNames {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	for _, subNode := range node.NamedArgs {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Op)))
	ctx.WriteString(strconv.Itoa(int(node.Xmloption)))
}
