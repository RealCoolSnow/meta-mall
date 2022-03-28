// Code generated by entc, DO NOT EDIT.

package userlog

import (
	"time"
)

const (
	// Label holds the string label denoting the userlog type in the database.
	Label = "user_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// Table holds the table name of the userlog in the database.
	Table = "user_logs"
)

// Columns holds all SQL columns for userlog fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldIP,
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
	// DefaultIP holds the default value on creation for the "ip" field.
	DefaultIP string
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// UpdateDefaultCreateTime holds the default value on update for the "create_time" field.
	UpdateDefaultCreateTime func() time.Time
)
