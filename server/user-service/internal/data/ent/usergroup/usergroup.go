// Code generated by entc, DO NOT EDIT.

package usergroup

import (
	"time"
)

const (
	// Label holds the string label denoting the usergroup type in the database.
	Label = "user_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// Table holds the table name of the usergroup in the database.
	Table = "user_groups"
)

// Columns holds all SQL columns for usergroup fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// UpdateDefaultCreateTime holds the default value on update for the "create_time" field.
	UpdateDefaultCreateTime func() time.Time
)