package model

type ViewTemplate struct {
	CommonUUID
	CommonNameUnique
	ViewTemplateWidgets []*ViewTemplateWidget `json:"view_template_widgets,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
