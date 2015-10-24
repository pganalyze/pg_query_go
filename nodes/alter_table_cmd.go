// Auto-generated - DO NOT EDIT

package pg_query

type AlterTableCmd struct {
	Subtype AlterTableType `json:"subtype"` /* Type of table alteration to apply */
	Name    *string        `json:"name"`    /* column, constraint, or trigger to act on,
	 * or new owner or tablespace */
	Def Node `json:"def"` /* definition of new column, index,
	 * constraint, or parent table */
	Behavior  DropBehavior `json:"behavior"`   /* RESTRICT or CASCADE for DROP cases */
	MissingOk bool         `json:"missing_ok"` /* skip error if missing? */
}
