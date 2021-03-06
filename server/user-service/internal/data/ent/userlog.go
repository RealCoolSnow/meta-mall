// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"user-service/internal/data/ent/userlog"

	"entgo.io/ent/dialect/sql"
)

// UserLog is the model entity for the UserLog schema.
type UserLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// Extra holds the value of the "extra" field.
	Extra string `json:"extra,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserLog) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userlog.FieldID, userlog.FieldUserID:
			values[i] = new(sql.NullInt64)
		case userlog.FieldIP, userlog.FieldExtra:
			values[i] = new(sql.NullString)
		case userlog.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserLog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserLog fields.
func (ul *UserLog) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ul.ID = int64(value.Int64)
		case userlog.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ul.UserID = value.Int64
			}
		case userlog.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				ul.IP = value.String
			}
		case userlog.FieldExtra:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extra", values[i])
			} else if value.Valid {
				ul.Extra = value.String
			}
		case userlog.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ul.CreateTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserLog.
// Note that you need to call UserLog.Unwrap() before calling this method if this UserLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (ul *UserLog) Update() *UserLogUpdateOne {
	return (&UserLogClient{config: ul.config}).UpdateOne(ul)
}

// Unwrap unwraps the UserLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ul *UserLog) Unwrap() *UserLog {
	tx, ok := ul.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserLog is not a transactional entity")
	}
	ul.config.driver = tx.drv
	return ul
}

// String implements the fmt.Stringer.
func (ul *UserLog) String() string {
	var builder strings.Builder
	builder.WriteString("UserLog(")
	builder.WriteString(fmt.Sprintf("id=%v", ul.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", ul.UserID))
	builder.WriteString(", ip=")
	builder.WriteString(ul.IP)
	builder.WriteString(", extra=")
	builder.WriteString(ul.Extra)
	builder.WriteString(", create_time=")
	builder.WriteString(ul.CreateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserLogs is a parsable slice of UserLog.
type UserLogs []*UserLog

func (ul UserLogs) config(cfg config) {
	for _i := range ul {
		ul[_i].config = cfg
	}
}
