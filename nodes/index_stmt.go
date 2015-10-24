// Auto-generated - DO NOT EDIT

package pg_query

type IndexStmt struct {
	Idxname        *string   `json:"idxname"`        /* name of new index, or NULL for default */
	Relation       *RangeVar `json:"relation"`       /* relation to build index on */
	AccessMethod   *string   `json:"accessMethod"`   /* name of access method (eg. btree) */
	TableSpace     *string   `json:"tableSpace"`     /* tablespace, or NULL for default */
	IndexParams    []Node    `json:"indexParams"`    /* columns to index: a list of IndexElem */
	Options        []Node    `json:"options"`        /* WITH clause options: a list of DefElem */
	WhereClause    Node      `json:"whereClause"`    /* qualification (partial-index predicate) */
	ExcludeOpNames []Node    `json:"excludeOpNames"` /* exclusion operator names, or NIL if none */
	Idxcomment     *string   `json:"idxcomment"`     /* comment to apply to index, or NULL */
	IndexOid       Oid       `json:"indexOid"`       /* OID of an existing index, if any */
	OldNode        Oid       `json:"oldNode"`        /* relfilenode of existing storage, if any */
	Unique         bool      `json:"unique"`         /* is index unique? */
	Primary        bool      `json:"primary"`        /* is index a primary key? */
	Isconstraint   bool      `json:"isconstraint"`   /* is it for a pkey/unique constraint? */
	Deferrable     bool      `json:"deferrable"`     /* is the constraint DEFERRABLE? */
	Initdeferred   bool      `json:"initdeferred"`   /* is the constraint INITIALLY DEFERRED? */
	Concurrent     bool      `json:"concurrent"`     /* should this be a concurrent index build? */
}
