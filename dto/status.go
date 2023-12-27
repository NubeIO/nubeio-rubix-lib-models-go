package dto

type CreateStatus string

const (
	CreateNotAvailable CreateStatus = "N/A"
	Creating           CreateStatus = "Creating"
	Created            CreateStatus = "Created"
	CreateFailed       CreateStatus = "Failed"
)

type RestoreStatus string

const (
	RestoreNotAvailable RestoreStatus = "N/A"
	Restoring           RestoreStatus = "Restoring"
	Restored            RestoreStatus = "Restored"
	RestoreFailed       RestoreStatus = "Failed"
)

type EmailStatus string

const (
	EmailStatusSending EmailStatus = "Sending"
	EmailStatusSent    EmailStatus = "Sent"
	EmailStatusFailed  EmailStatus = "Failed"
)
