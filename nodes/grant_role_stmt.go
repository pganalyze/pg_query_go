// Auto-generated - DO NOT EDIT

package pg_query

type GrantRoleStmt struct {
	GrantedRoles []Node       `json:"granted_roles"` /* list of roles to be granted/revoked */
	GranteeRoles []Node       `json:"grantee_roles"` /* list of member roles to add/delete */
	IsGrant      bool         `json:"is_grant"`      /* true = GRANT, false = REVOKE */
	AdminOpt     bool         `json:"admin_opt"`     /* with admin option */
	Grantor      *string      `json:"grantor"`       /* set grantor to other than current role */
	Behavior     DropBehavior `json:"behavior"`      /* drop behavior (for REVOKE) */
}
