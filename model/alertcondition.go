package model

import "strings"

const separator = ";"

type AlertCondition struct {
	CommonUUID
	CommonName
	AlertTypes      string                   `json:"alert_types"`
	AlertStatuses   string                   `json:"alert_statuses"`
	AlertSeverities string                   `json:"alert_severities"`
	Tags            []*Tag                   `json:"tags,omitempty" gorm:"many2many:alert_conditions_tags;constraint:OnDelete:CASCADE"`
	MetaTags        []*AlertConditionMetaTag `json:"meta_tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams           []*AlertConditionTeam    `json:"teams,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}

type AlertConditionMetaTag struct {
	AlertConditionUUID string `json:"alert_condition_uuid,omitempty" gorm:"type:varchar(255) references alert_conditions;not null;default:null;primaryKey"`
	Key                string `json:"key,omitempty" gorm:"primaryKey"`
	Value              string `json:"value,omitempty"`
}

type AlertConditionDto struct {
	UUID            string                   `json:"uuid"`
	Name            string                   `json:"name"`
	AlertTypes      []string                 `json:"alert_types"`
	AlertStatuses   []string                 `json:"alert_statuses"`
	AlertSeverities []string                 `json:"alert_severities"`
	Tags            []*Tag                   `json:"tags,omitempty"`
	MetaTags        []*AlertConditionMetaTag `json:"meta_tags,omitempty"`
	Teams           []*AlertConditionTeam    `json:"teams,omitempty"`
}

func (a *AlertConditionDto) ToModel() *AlertCondition {
	if a == nil {
		return nil
	}
	return &AlertCondition{
		CommonUUID:      CommonUUID{UUID: a.UUID},
		CommonName:      CommonName{Name: a.Name},
		AlertTypes:      strings.Join(a.AlertTypes, separator),
		AlertStatuses:   strings.Join(a.AlertStatuses, separator),
		AlertSeverities: strings.Join(a.AlertSeverities, separator),
		Tags:            a.Tags,
		MetaTags:        a.MetaTags,
		Teams:           a.Teams,
	}
}

func (a *AlertCondition) ToDto() *AlertConditionDto {
	if a == nil {
		return nil
	}
	return &AlertConditionDto{
		UUID:            a.UUID,
		Name:            a.Name,
		AlertTypes:      strings.Split(a.AlertTypes, separator),
		AlertStatuses:   strings.Split(a.AlertStatuses, separator),
		AlertSeverities: strings.Split(a.AlertSeverities, separator),
		Tags:            a.Tags,
		MetaTags:        a.MetaTags,
		Teams:           a.Teams,
	}
}
