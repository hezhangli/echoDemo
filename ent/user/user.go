// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldQiniuFilePath holds the string denoting the qiniu_file_path field in the database.
	FieldQiniuFilePath = "qiniu_file_path"
	// FieldQiniuTxPath holds the string denoting the qiniu_tx_path field in the database.
	FieldQiniuTxPath = "qiniu_tx_path"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldWxID holds the string denoting the wx_id field in the database.
	FieldWxID = "wx_id"
	// FieldCorpID holds the string denoting the corp_id field in the database.
	FieldCorpID = "corp_id"
	// FieldDeptID holds the string denoting the dept_id field in the database.
	FieldDeptID = "dept_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldName,
	FieldQiniuFilePath,
	FieldQiniuTxPath,
	FieldTags,
	FieldUserID,
	FieldWxID,
	FieldCorpID,
	FieldDeptID,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByQiniuFilePath orders the results by the qiniu_file_path field.
func ByQiniuFilePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQiniuFilePath, opts...).ToFunc()
}

// ByQiniuTxPath orders the results by the qiniu_tx_path field.
func ByQiniuTxPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQiniuTxPath, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByWxID orders the results by the wx_id field.
func ByWxID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWxID, opts...).ToFunc()
}

// ByCorpID orders the results by the corp_id field.
func ByCorpID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCorpID, opts...).ToFunc()
}

// ByDeptID orders the results by the dept_id field.
func ByDeptID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeptID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}