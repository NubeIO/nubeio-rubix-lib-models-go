package dto

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

type AlertCondition struct {
	UUID            string                         `json:"uuid"`
	Name            string                         `json:"name"`
	AlertTypes      []string                       `json:"alert_types"`
	AlertStatuses   []string                       `json:"alert_statuses"`
	AlertSeverities []string                       `json:"alert_severities"`
	Tags            []*model.Tag                   `json:"tags,omitempty"`
	MetaTags        []*model.AlertConditionMetaTag `json:"meta_tags,omitempty"`
	Teams           []*model.AlertConditionTeam    `json:"teams,omitempty"`
}
