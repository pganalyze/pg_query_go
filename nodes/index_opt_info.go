// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type IndexOptInfo struct {
	Indexoid      Oid         `json:"indexoid"`      /* OID of the index relation */
	Reltablespace Oid         `json:"reltablespace"` /* tablespace of index (not table) */
	Rel           *RelOptInfo `json:"rel"`           /* back-link to index's table */

	/* index-size statistics (from pg_class and elsewhere) */
	Pages      BlockNumber `json:"pages"`       /* number of disk pages in index */
	Tuples     float64     `json:"tuples"`      /* number of index tuples in index */
	TreeHeight int         `json:"tree_height"` /* index tree height, or -1 if unknown */

	/* index descriptor information */
	Ncolumns        int   `json:"ncolumns"`        /* number of columns in index */
	Indexkeys       *int  `json:"indexkeys"`       /* column numbers of index's keys, or 0 */
	Indexcollations *Oid  `json:"indexcollations"` /* OIDs of collations of index columns */
	Opfamily        *Oid  `json:"opfamily"`        /* OIDs of operator families for columns */
	Opcintype       *Oid  `json:"opcintype"`       /* OIDs of opclass declared input data types */
	Sortopfamily    *Oid  `json:"sortopfamily"`    /* OIDs of btree opfamilies, if orderable */
	ReverseSort     *bool `json:"reverse_sort"`    /* is sort order descending? */
	NullsFirst      *bool `json:"nulls_first"`     /* do NULLs come first in the sort order? */
	Relam           Oid   `json:"relam"`           /* OID of the access method (in pg_am) */

	Amcostestimate RegProcedure `json:"amcostestimate"` /* OID of the access method's cost fcn */

	Indexprs []Node `json:"indexprs"` /* expressions for non-simple index columns */
	Indpred  []Node `json:"indpred"`  /* predicate if a partial index, else NIL */

	Indextlist []Node `json:"indextlist"` /* targetlist representing index columns */

	PredOk         bool `json:"predOK"`         /* true if predicate matches query */
	Unique         bool `json:"unique"`         /* true if a unique index */
	Immediate      bool `json:"immediate"`      /* is uniqueness enforced immediately? */
	Hypothetical   bool `json:"hypothetical"`   /* true if index doesn't really exist */
	Canreturn      bool `json:"canreturn"`      /* can index return IndexTuples? */
	Amcanorderbyop bool `json:"amcanorderbyop"` /* does AM support order by operator result? */
	Amoptionalkey  bool `json:"amoptionalkey"`  /* can query omit key for the first column? */
	Amsearcharray  bool `json:"amsearcharray"`  /* can AM handle ScalarArrayOpExpr quals? */
	Amsearchnulls  bool `json:"amsearchnulls"`  /* can AM search for NULL/NOT NULL entries? */
	Amhasgettuple  bool `json:"amhasgettuple"`  /* does AM have amgettuple interface? */
	Amhasgetbitmap bool `json:"amhasgetbitmap"` /* does AM have amgetbitmap interface? */
}

func (node IndexOptInfo) MarshalJSON() ([]byte, error) {
	type IndexOptInfoMarshalAlias IndexOptInfo
	return json.Marshal(map[string]interface{}{
		"INDEXOPTINFO": (*IndexOptInfoMarshalAlias)(&node),
	})
}

func (node *IndexOptInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
