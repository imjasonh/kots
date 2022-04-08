package types

import (
	"io"
	"time"

	kotsv1beta1 "github.com/replicatedhq/kots/kotskinds/apis/kots/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

type DeployOptions struct {
	Namespace              string
	Context                string
	SharedPassword         string
	SharedPasswordBcrypt   string
	S3AccessKey            string
	S3SecretKey            string
	JWT                    string
	PostgresPassword       string
	APIEncryptionKey       string
	AutoCreateClusterToken string
	ServiceType            string
	NodePort               int32
	ApplicationMetadata    []byte
	LimitRange             *corev1.LimitRange
	IsOpenShift            bool
	License                *kotsv1beta1.License
	ConfigValues           *kotsv1beta1.ConfigValues
	AppVersionLabel        string
	Airgap                 bool
	AirgapRootDir          string
	AirgapBundle           string
	AppImagesPushed        bool
	ProgressWriter         io.Writer
	IncludeMinio           bool
	IncludeMinioSnapshots  bool
	Timeout                time.Duration
	HTTPProxyEnvValue      string
	HTTPSProxyEnvValue     string
	NoProxyEnvValue        string
	ExcludeAdminConsole    bool
	EnsureKotsadmConfig    bool
	SkipPreflights         bool
	SkipCompatibilityCheck bool
	EnsureRBAC             bool
	SkipRBACCheck          bool
	UseMinimalRBAC         bool
	StrictSecurityContext  bool
	InstallID              string
	SimultaneousUploads    int
	DisableImagePush       bool
	UpstreamURI            string
	IsMinimalRBAC          bool
	AdditionalNamespaces   []string

	IdentityConfig kotsv1beta1.IdentityConfig
	IngressConfig  kotsv1beta1.IngressConfig

	KotsadmOptions KotsadmOptions
}
