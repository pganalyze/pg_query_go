// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RelOptInfo) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RELOPTINFO")

	for _, subNode := range node.Baserestrictinfo {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.CheapestParameterizedPaths {
		subNode.Fingerprint(ctx)
	}

	if node.CheapestStartupPath != nil {
		node.CheapestStartupPath.Fingerprint(ctx)
	}

	if node.CheapestTotalPath != nil {
		node.CheapestTotalPath.Fingerprint(ctx)
	}

	if node.CheapestUniquePath != nil {
		node.CheapestUniquePath.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.ConsiderParamStartup))
	ctx.WriteString(strconv.FormatBool(node.ConsiderStartup))
	ctx.WriteString(strconv.FormatBool(node.HasEclassJoins))

	for _, subNode := range node.Indexlist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Joininfo {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.LateralVars {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Pathlist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ppilist {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Reloptkind)))

	for _, subNode := range node.Reltargetlist {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Rtekind)))

	if node.Subplan != nil {
		node.Subplan.Fingerprint(ctx)
	}

	for _, subNode := range node.SubplanParams {
		subNode.Fingerprint(ctx)
	}

	if node.Subroot != nil {
		node.Subroot.Fingerprint(ctx)
	}
}
