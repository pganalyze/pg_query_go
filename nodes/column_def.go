// Auto-generated - DO NOT EDIT

package pg_query

type ColumnDef struct {
	Colname       *string        `json:"colname"`        /* name of column */
	TypeName      *TypeName      `json:"typeName"`       /* type of column */
	Inhcount      int            `json:"inhcount"`       /* number of times column is inherited */
	IsLocal       bool           `json:"is_local"`       /* column has local (non-inherited) def'n */
	IsNotNull     bool           `json:"is_not_null"`    /* NOT NULL constraint specified? */
	IsFromType    bool           `json:"is_from_type"`   /* column definition came from table type */
	Storage       byte           `json:"storage"`        /* attstorage setting, or 0 for default */
	RawDefault    Node           `json:"raw_default"`    /* default value (untransformed parse tree) */
	CookedDefault Node           `json:"cooked_default"` /* default value (transformed expr tree) */
	CollClause    *CollateClause `json:"collClause"`     /* untransformed COLLATE spec, if any */
	CollOid       Oid            `json:"collOid"`        /* collation OID (InvalidOid if not set) */
	Constraints   []Node         `json:"constraints"`    /* other constraints on column */
	Fdwoptions    []Node         `json:"fdwoptions"`     /* per-column FDW options */
	Location      int            `json:"location"`       /* parse location, or -1 if none/unknown */
}
