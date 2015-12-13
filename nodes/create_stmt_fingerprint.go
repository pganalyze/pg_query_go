// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateStmt")

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx, "Constraints")
	}

	ctx.WriteString(strconv.FormatBool(node.IfNotExists))

	for _, subNode := range node.InhRelations {
		subNode.Fingerprint(ctx, "InhRelations")
	}

	if node.OfTypename != nil {
		node.OfTypename.Fingerprint(ctx, "OfTypename")
	}

	ctx.WriteString(strconv.Itoa(int(node.Oncommit)))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	for _, subNode := range node.TableElts {
		subNode.Fingerprint(ctx, "TableElts")
	}

	if node.Tablespacename != nil {
		ctx.WriteString(*node.Tablespacename)
	}
}
