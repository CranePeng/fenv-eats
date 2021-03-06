package models

type PipelineNodePivot struct {
	Id         string    `json:"id" gorm:"not null; primary_key; comment:'ID';type: CHAR(36)"`
	PipelineId string    `json:"pipeline_id" validate:"required,uuid4" gorm:"not null; index; comment:'流水线ID';type: CHAR(36)"`
	NodeId     string    `json:"node_id" validate:"required,uuid4" gorm:"not null; index; comment:'节点ID'; type:CHAR(36)"`
	Pipeline   *Pipeline `json:"pipeline" gorm:"-"`
	CommonColumn
}

// 定义模型的数据表名称
func (pivot *PipelineNodePivot) TableName() string {
	return "pipeline_node_pivot"
}

// 创建关联
func (pivot *PipelineNodePivot) Create() error {
	err := Engine.Create(pivot).Error
	return err
}

// 更新
func (pivot *PipelineNodePivot) Update() int {
	i := Engine.Model(pivot).Updates(pivot).RowsAffected
	return int(i)
}

// 解除关联
func (pivot *PipelineNodePivot) Delete() error {
	err := Engine.Delete(pivot).Error
	return err
}
