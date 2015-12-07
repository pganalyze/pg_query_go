// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATESTMT")

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.IfNotExists))

	for _, subNode := range node.InhRelations {
		subNode.Fingerprint(ctx)
	}

	if node.OfTypename != nil {
		node.OfTypename.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Oncommit)))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	for _, subNode := range node.TableElts {
		subNode.Fingerprint(ctx)
	}

	if node.Tablespacename != nil {
		ctx.WriteString(*node.Tablespacename)
	}
}
