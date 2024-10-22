// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHistory = "history"

// History mapped from table <history>
type History struct {
	ID             int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Operation      string `gorm:"column:operation;not null" json:"operation"`
	ModelID        string `gorm:"column:model_id;not null" json:"model_id"`
	InstanceID     int64  `gorm:"column:instance_id;not null" json:"instance_id"`
	User           string `gorm:"column:user;not null" json:"user"`
	CreateTime     string `gorm:"column:create_time;not null" json:"create_time"`
	ModelInstance  string `gorm:"column:model_instance;not null" json:"model_instance"`
	Description    string `gorm:"column:description;not null" json:"description"`
	FieldData      string `gorm:"column:field_data;not null" json:"field_data"`
	RollbackStatus int32  `gorm:"column:rollback_status;not null" json:"rollback_status"`
	FailReason     string `gorm:"column:fail_reason;not null" json:"fail_reason"`
	ModifiedData   string `gorm:"column:modified_data;not null;comment:save modified data" json:"modified_data"` // save modified data
}

// TableName History's table name
func (*History) TableName() string {
	return TableNameHistory
}