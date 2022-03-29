// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "group_id", Type: field.TypeInt32, Default: 0},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "mobile", Type: field.TypeString, Default: ""},
		{Name: "email", Type: field.TypeString, Default: ""},
		{Name: "create_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "update_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0]},
	}
	// UserLogsColumns holds the columns for the "user_logs" table.
	UserLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "ip", Type: field.TypeString, Default: ""},
		{Name: "extra", Type: field.TypeString, Default: ""},
		{Name: "create_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UserLogsTable holds the schema information for the "user_logs" table.
	UserLogsTable = &schema.Table{
		Name:       "user_logs",
		Columns:    UserLogsColumns,
		PrimaryKey: []*schema.Column{UserLogsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
		UserGroupsTable,
		UserLogsTable,
	}
)

func init() {
}
