// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node VacuumStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("VacuumStmt")
	ctx.WriteString(strconv.Itoa(int(node.FreezeMinAge)))
	ctx.WriteString(strconv.Itoa(int(node.FreezeTableAge)))
	ctx.WriteString(strconv.Itoa(int(node.MultixactFreezeMinAge)))
	ctx.WriteString(strconv.Itoa(int(node.MultixactFreezeTableAge)))
	ctx.WriteString(strconv.Itoa(int(node.Options)))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	for _, subNode := range node.VaCols {
		subNode.Fingerprint(ctx)
	}
}
