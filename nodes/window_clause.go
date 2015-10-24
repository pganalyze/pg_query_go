// Auto-generated - DO NOT EDIT

package pg_query

type WindowClause struct {
	Name            *string `json:"name"`            /* window name (NULL in an OVER clause) */
	Refname         *string `json:"refname"`         /* referenced window name, if any */
	PartitionClause []Node  `json:"partitionClause"` /* PARTITION BY list */
	OrderClause     []Node  `json:"orderClause"`     /* ORDER BY list */
	FrameOptions    int     `json:"frameOptions"`    /* frame_clause options, see WindowDef */
	StartOffset     Node    `json:"startOffset"`     /* expression for starting bound, if any */
	EndOffset       Node    `json:"endOffset"`       /* expression for ending bound, if any */
	Winref          Index   `json:"winref"`          /* ID referenced by window functions */
	CopiedOrder     bool    `json:"copiedOrder"`     /* did we copy orderClause from refname? */
}
