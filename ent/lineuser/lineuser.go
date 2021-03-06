// Code generated by entc, DO NOT EDIT.

package lineuser

import (
	"time"
)

const (
	// Label holds the string label denoting the lineuser type in the database.
	Label = "line_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// Table holds the table name of the lineuser in the database.
	Table = "line_users"
)

// Columns holds all SQL columns for lineuser fields.
var Columns = []string{
	FieldID,
	FieldIsActive,
	FieldCreateAt,
	FieldUpdateAt,
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
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
