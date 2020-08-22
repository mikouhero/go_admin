package model

import "github.com/jinzhu/gorm"

type SysWorkflow struct {
	gorm.Model
	WorkflowNickName    string `json:"workflowNickName" gorm:"comment:'工作流中文名称'"`
	WorkflowName        string `json:"workflowName" gorm:"comment:''"`
	WorkflowDescription string `json:"workflowDescription" gorm:"comment:'工作流描述'"` // 工作流描述
	WorkflowStepInfo    string `json:"workflowStep" gorm:"comment:'工作流步骤'"`        // 工作流步骤
}

type SysWorkflowStepInfo struct {
	gorm.Model
	SysWorkflowID   uint    `json:"workflowID" gorm:"comment:'所属工作流ID'"`      // 所属工作流ID
	IsStrat         bool    `json:"isStrat" gorm:"comment:'是否是开始流节点'"`        // 是否是开始流节点
	StepName        string  `json:"stepName" gorm:"comment:'工作流节点名称'"`        // 工作流名称
	StepNo          float64 `json:"stepNo" gorm:"comment:'步骤id （第几步）'"`       // 步骤id （第几步）
	StepAuthorityID string  `json:"stepAuthorityID" gorm:"comment:'操作者级别id'"` // 操作者级别id
	IsEnd           bool    `json:"isEnd" gorm:"comment:'是否是完结流节点'"`          // 是否是完结流节点
}
