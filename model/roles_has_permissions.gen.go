// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRolesHasPermission = "roles_has_permissions"

// RolesHasPermission mapped from table <roles_has_permissions>
type RolesHasPermission struct {
	ID            uint32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RolesID       uint32 `gorm:"column:roles_id;not null" json:"roles_id"`
	PermissionsID uint32 `gorm:"column:permissions_id;not null" json:"permissions_id"`
}

// TableName RolesHasPermission's table name
func (*RolesHasPermission) TableName() string {
	return TableNameRolesHasPermission
}