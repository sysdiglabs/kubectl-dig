package digjob

import (
	"io"
	"io/ioutil"

	"github.com/leodido/kubectl-dig/pkg/meta"
	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	batchv1typed "k8s.io/client-go/kubernetes/typed/batch/v1"

	pointerutils "k8s.io/utils/pointer"
)

// Client ...
type Client struct {
	JobClient batchv1typed.JobInterface

	outStream io.Writer
}

// Job contains info needed to create the job responsible for digging stuff.
type Job struct {
	Name           string
	ID             types.UID
	Namespace      string
	ServiceAccount string
	Hostname       string
	ImageNameTag   string
	StartTime      metav1.Time
	Status         Status
}

// WithOutStream setup a file stream to output digjob operation information
func (c *Client) WithOutStream(o io.Writer) {
	if o == nil {
		c.outStream = ioutil.Discard
	}
	c.outStream = o
}

// CreateJob ...
func (c *Client) CreateJob(j Job) (*batchv1.Job, error) {
	command := []string{
		"csysdig",
	}

	if len(j.ServiceAccount) > 0 {
		// TODO: this actually works only if the connection is made in https,
		// we should check whether this connection can be plain http too.
		command = []string{
			"bash",
			"-c",
			"csysdig -K /var/run/secrets/kubernetes.io/serviceaccount/token -k https://${KUBERNETES_SERVICE_HOST}",
		}
	}

	commonMeta := metav1.ObjectMeta{
		Name:      j.Name,
		Namespace: j.Namespace,
		Labels: map[string]string{
			meta.DigLabelKey:   j.Name,
			meta.DigIDLabelKey: string(j.ID),
		},
		Annotations: map[string]string{
			meta.DigLabelKey:   j.Name,
			meta.DigIDLabelKey: string(j.ID),
		},
	}

	job := &batchv1.Job{
		ObjectMeta: commonMeta,
		Spec: batchv1.JobSpec{
			TTLSecondsAfterFinished: pointerutils.Int32Ptr(5),
			Parallelism:             pointerutils.Int32Ptr(1),
			Completions:             pointerutils.Int32Ptr(1),
			BackoffLimit:            pointerutils.Int32Ptr(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: commonMeta,
				Spec: apiv1.PodSpec{
					HostPID:            true,
					ServiceAccountName: j.ServiceAccount,
					Volumes: []apiv1.Volume{
						{
							Name: "hostroot",
							VolumeSource: apiv1.VolumeSource{
								HostPath: &apiv1.HostPathVolumeSource{
									Path: "/",
								},
							},
						},
					},
					Containers: []apiv1.Container{
						{
							Name:    j.Name,
							Image:   j.ImageNameTag,
							Command: command,
							TTY:     true,
							Stdin:   true,
							Resources: apiv1.ResourceRequirements{
								Requests: apiv1.ResourceList{
									apiv1.ResourceCPU:    resource.MustParse("100m"),
									apiv1.ResourceMemory: resource.MustParse("100Mi"),
								},
								Limits: apiv1.ResourceList{
									apiv1.ResourceCPU:    resource.MustParse("1"),
									apiv1.ResourceMemory: resource.MustParse("1G"),
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "hostroot",
									MountPath: "/host",
									ReadOnly:  true,
								},
							},
							SecurityContext: &apiv1.SecurityContext{
								Privileged: pointerutils.BoolPtr(true),
							},
						},
					},
					RestartPolicy: "Never",
					Affinity: &apiv1.Affinity{
						NodeAffinity: &apiv1.NodeAffinity{
							RequiredDuringSchedulingIgnoredDuringExecution: &apiv1.NodeSelector{
								NodeSelectorTerms: []apiv1.NodeSelectorTerm{
									{
										MatchExpressions: []apiv1.NodeSelectorRequirement{
											{
												Key:      "kubernetes.io/hostname",
												Operator: apiv1.NodeSelectorOpIn,
												Values: []string{
													j.Hostname,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return c.JobClient.Create(job)
}
