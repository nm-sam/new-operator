
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


// BuilderPlaygroundDeployment represents the top-level K8s CRD
type BuilderPlaygroundDeployment struct {
    // APIVersion is the Kubernetes API version for this resource
    APIVersion string `yaml:"apiVersion"`
    // Kind identifies this as a BuilderPlaygroundDeployment
    Kind string `yaml:"kind"`
    // Metadata contains the resource metadata
    Metadata BuilderPlaygroundMetadata `yaml:"metadata"`
    // Spec defines the desired state of the deployment
    Spec BuilderPlaygroundSpec `yaml:"spec"`
}

// BuilderPlaygroundMetadata contains the resource metadata
type BuilderPlaygroundMetadata struct {
    // Name is the name of the deployment
    Name string `yaml:"name"`
}

// BuilderPlaygroundSpec defines the desired state of the deployment
type BuilderPlaygroundSpec struct {
    // Recipe is the builder-playground recipe used (l1, opstack, etc)
    Recipe string `yaml:"recipe"`
    // Storage defines how persistent data should be stored
    Storage BuilderPlaygroundStorage `yaml:"storage"`
    // Network defines networking configuration (optional)
    Network *BuilderPlaygroundNetwork `yaml:"network,omitempty"`
    // Services is the list of services in this deployment
    Services []BuilderPlaygroundService `yaml:"services"`
}

// BuilderPlaygroundStorage defines storage configuration
type BuilderPlaygroundStorage struct {
    // Type is the storage type, either "local-path" or "pvc"
    Type string `yaml:"type"`
    // Path is the host path for local-path storage (used when type is "local-path")
    Path string `yaml:"path,omitempty"`
    // StorageClass is the K8s storage class (used when type is "pvc")
    StorageClass string `yaml:"storageClass,omitempty"`
    // Size is the storage size (used when type is "pvc")
    Size string `yaml:"size,omitempty"`
}

// BuilderPlaygroundNetwork defines network configuration
type BuilderPlaygroundNetwork struct {
    // Name is the name of the network
    Name string `yaml:"name"`
}

// BuilderPlaygroundService represents a single service in the deployment
type BuilderPlaygroundService struct {
    // Name is the service name
    Name string `yaml:"name"`
    // Image is the container image
    Image string `yaml:"image"`
    // Tag is the container image tag
    Tag string `yaml:"tag"`
    // Entrypoint overrides the container entrypoint
    Entrypoint []string `yaml:"entrypoint,omitempty"`
    // Args are the container command arguments
    Args []string `yaml:"args,omitempty"`
    // Env defines environment variables
    Env map[string]string `yaml:"env,omitempty"`
    // Ports are the container ports to expose
    Ports []BuilderPlaygroundPort `yaml:"ports,omitempty"`
    // Dependencies defines services this service depends on
    Dependencies []BuilderPlaygroundDependency `yaml:"dependencies,omitempty"`
    // ReadyCheck defines how to determine service readiness
    ReadyCheck *BuilderPlaygroundReadyCheck `yaml:"readyCheck,omitempty"`
    // Labels are the service labels
    Labels map[string]string `yaml:"labels,omitempty"`
    // UseHostExecution indicates whether to run on host instead of in container
    UseHostExecution bool `yaml:"useHostExecution,omitempty"`
    // Volumes are the volume mounts for the service
    Volumes []BuilderPlaygroundVolume `yaml:"volumes,omitempty"`
}

// BuilderPlaygroundPort represents a port configuration
type BuilderPlaygroundPort struct {
    // Name is a unique identifier for this port
    Name string `yaml:"name"`
    // Port is the container port number
    Port int `yaml:"port"`
    // Protocol is either "tcp" or "udp"
    Protocol string `yaml:"protocol,omitempty"`
    // HostPort is the port to expose on the host (if applicable)
    HostPort int `yaml:"hostPort,omitempty"`
}

// BuilderPlaygroundDependency represents a service dependency
type BuilderPlaygroundDependency struct {
    // Name is the name of the dependent service
    Name string `yaml:"name"`
    // Condition is either "running" or "healthy"
    Condition string `yaml:"condition"`
}

// BuilderPlaygroundReadyCheck defines readiness checking
type BuilderPlaygroundReadyCheck struct {
    // QueryURL is the URL to query for readiness
    QueryURL string `yaml:"queryURL,omitempty"`
    // Test is the command to run for readiness check
    Test []string `yaml:"test,omitempty"`
    // Interval is the time between checks
    Interval string `yaml:"interval,omitempty"`
    // Timeout is the maximum time for a check
    Timeout string `yaml:"timeout,omitempty"`
    // Retries is the number of retry attempts
    Retries int `yaml:"retries,omitempty"`
    // StartPeriod is the initial delay before checks begin
    StartPeriod string `yaml:"startPeriod,omitempty"`
}

// BuilderPlaygroundVolume represents a volume mount
type BuilderPlaygroundVolume struct {
    // Name is the volume name
    Name string `yaml:"name"`
    // MountPath is the path in the container
    MountPath string `yaml:"mountPath"`
    // SubPath is the path within the volume (optional)
    SubPath string `yaml:"subPath,omitempty"`
}