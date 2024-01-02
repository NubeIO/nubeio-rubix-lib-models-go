package model

type PointMetaTag struct {
	PointUUID string `json:"point_uuid,omitempty" gorm:"type:varchar(255) references points;not null;default:null;primaryKey"`
	Key       string `json:"key,omitempty" gorm:"primaryKey"`
	Value     string `json:"value,omitempty"`
}
