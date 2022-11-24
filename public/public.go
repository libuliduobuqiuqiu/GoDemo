package public

type DescriptionStruct struct {
	Description string `json:"description" db:"description"` //描述
}

type RemarkStruct struct {
	Remark string `json:"remark" db:"remark"` //备注
}
