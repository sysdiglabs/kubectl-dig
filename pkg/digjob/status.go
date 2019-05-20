package digjob

import (
	batchv1 "k8s.io/api/batch/v1"
)

// Status is a label for the running status of a dig job at the current time.
type Status string

// These are the valid status of dig runs.
const (
	// Running means the digjob has active pods.
	Running Status = "Running"
	// Completed means the digjob does not have any active pod and has success pods.
	Completed Status = "Completed"
	// Failed means the digjob does not have any active or success pod and has fpods that failed.
	Failed Status = "Failed"
	// Unknown means that for some reason we do not have the information to determine the status.
	Unknown Status = "Unknown"
)

func jobStatus(j batchv1.Job) Status {
	if j.Status.Active > 0 {
		return Running
	}
	if j.Status.Succeeded > 0 {
		return Completed
	}
	if j.Status.Failed > 0 {
		return Failed
	}
	return Unknown
}
