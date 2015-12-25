// Auto-generated - DO NOT EDIT

package pg_query

import "sort"

func (node List) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	if parentFieldName == "TargetList" || parentFieldName == "Cols" || parentFieldName == "Rexpr" {
		var itemsFingerprints FingerprintSubContextSlice

		for _, subNode := range node.Items {
			subCtx := FingerprintSubContext{}
			subNode.Fingerprint(&subCtx, parentFieldName)
			itemsFingerprints.AddIfUnique(subCtx)
		}

		sort.Sort(itemsFingerprints)

		for _, fingerprint := range itemsFingerprints {
			for _, part := range fingerprint.parts {
				ctx.WriteString(part)
			}
		}
	} else {
		for _, subNode := range node.Items {
			subNode.Fingerprint(ctx, parentFieldName)
		}
	}
}
