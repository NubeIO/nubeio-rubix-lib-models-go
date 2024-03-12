package model

type AlertCondition struct {
	CommonUUID
	CommonName
	AlertTypes      []string                   `json:"alert_types"`
	AlertStatuses   []string                   `json:"alert_statuses"`
	AlertSeverities []string                   `json:"alert_severities"`
	Tags            []*Tag                   `json:"tags,omitempty" gorm:"many2many:alert_conditions_tags;constraint:OnDelete:CASCADE"`
	MetaTags        []*AlertConditionMetaTag `json:"meta_tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams           []*AlertConditionTeam    `json:"teams,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
