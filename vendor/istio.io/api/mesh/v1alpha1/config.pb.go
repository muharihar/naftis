// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/config.proto

package v1alpha1 // import "istio.io/api/mesh/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AuthenticationPolicy defines authentication policy. It can be set for
// different scopes (mesh, service …), and the most narrow scope with
// non-INHERIT value will be used.
// Mesh policy cannot be INHERIT.
type AuthenticationPolicy int32

const (
	// Do not encrypt Envoy to Envoy traffic.
	AuthenticationPolicy_NONE AuthenticationPolicy = 0
	// Envoy to Envoy traffic is wrapped into mutual TLS connections.
	AuthenticationPolicy_MUTUAL_TLS AuthenticationPolicy = 1
	// Use the policy defined by the parent scope. Should not be used for mesh
	// policy.
	AuthenticationPolicy_INHERIT AuthenticationPolicy = 1000
)

var AuthenticationPolicy_name = map[int32]string{
	0:    "NONE",
	1:    "MUTUAL_TLS",
	1000: "INHERIT",
}
var AuthenticationPolicy_value = map[string]int32{
	"NONE":       0,
	"MUTUAL_TLS": 1,
	"INHERIT":    1000,
}

func (x AuthenticationPolicy) String() string {
	return proto.EnumName(AuthenticationPolicy_name, int32(x))
}
func (AuthenticationPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_cc82c7a62bb884d0, []int{0}
}

// The mode used to redirect inbound traffic to Envoy.
// This setting has no effect on outbound traffic: iptables REDIRECT is always used for
// outbound connections.
type ProxyConfig_InboundInterceptionMode int32

const (
	// The REDIRECT mode uses iptables REDIRECT to NAT and redirect to Envoy. This mode loses
	// source IP addresses during redirection.
	ProxyConfig_REDIRECT ProxyConfig_InboundInterceptionMode = 0
	// The TPROXY mode uses iptables TPROXY to redirect to Envoy. This mode preserves both the
	// source and destination IP addresses and ports, so that they can be used for advanced
	// filtering and manipulation. This mode also configures the sidecar to run with the
	// CAP_NET_ADMIN capability, which is required to use TPROXY.
	ProxyConfig_TPROXY ProxyConfig_InboundInterceptionMode = 1
)

var ProxyConfig_InboundInterceptionMode_name = map[int32]string{
	0: "REDIRECT",
	1: "TPROXY",
}
var ProxyConfig_InboundInterceptionMode_value = map[string]int32{
	"REDIRECT": 0,
	"TPROXY":   1,
}

func (x ProxyConfig_InboundInterceptionMode) String() string {
	return proto.EnumName(ProxyConfig_InboundInterceptionMode_name, int32(x))
}
func (ProxyConfig_InboundInterceptionMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_cc82c7a62bb884d0, []int{0, 0}
}

type MeshConfig_IngressControllerMode int32

const (
	// Disables Istio ingress controller.
	MeshConfig_OFF MeshConfig_IngressControllerMode = 0
	// Istio ingress controller will act on ingress resources that do not
	// contain any annotation or whose annotations match the value
	// specified in the ingress_class parameter described earlier. Use this
	// mode if Istio ingress controller will be the default ingress
	// controller for the entire kubernetes cluster.
	MeshConfig_DEFAULT MeshConfig_IngressControllerMode = 1
	// Istio ingress controller will only act on ingress resources whose
	// annotations match the value specified in the ingress_class parameter
	// described earlier. Use this mode if Istio ingress controller will be
	// a secondary ingress controller (e.g., in addition to a
	// cloud-provided ingress controller).
	MeshConfig_STRICT MeshConfig_IngressControllerMode = 2
)

var MeshConfig_IngressControllerMode_name = map[int32]string{
	0: "OFF",
	1: "DEFAULT",
	2: "STRICT",
}
var MeshConfig_IngressControllerMode_value = map[string]int32{
	"OFF":     0,
	"DEFAULT": 1,
	"STRICT":  2,
}

func (x MeshConfig_IngressControllerMode) String() string {
	return proto.EnumName(MeshConfig_IngressControllerMode_name, int32(x))
}
func (MeshConfig_IngressControllerMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_cc82c7a62bb884d0, []int{1, 0}
}

// TODO AuthPolicy needs to be removed and merged with AuthPolicy defined above
type MeshConfig_AuthPolicy int32

const (
	// Do not encrypt Envoy to Envoy traffic.
	MeshConfig_NONE MeshConfig_AuthPolicy = 0
	// Envoy to Envoy traffic is wrapped into mutual TLS connections.
	MeshConfig_MUTUAL_TLS MeshConfig_AuthPolicy = 1
)

var MeshConfig_AuthPolicy_name = map[int32]string{
	0: "NONE",
	1: "MUTUAL_TLS",
}
var MeshConfig_AuthPolicy_value = map[string]int32{
	"NONE":       0,
	"MUTUAL_TLS": 1,
}

func (x MeshConfig_AuthPolicy) String() string {
	return proto.EnumName(MeshConfig_AuthPolicy_name, int32(x))
}
func (MeshConfig_AuthPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_cc82c7a62bb884d0, []int{1, 1}
}

type MeshConfig_OutboundTrafficPolicy_Mode int32

const (
	// outbound traffic will be restricted to services defined in the service registry as well as those defined
	// through ServiceEntries
	MeshConfig_OutboundTrafficPolicy_REGISTRY_ONLY MeshConfig_OutboundTrafficPolicy_Mode = 0
	// outbound traffic to unknown destinations will be allowed
	MeshConfig_OutboundTrafficPolicy_ALLOW_ANY MeshConfig_OutboundTrafficPolicy_Mode = 1
	// not implemented. outbound traffic will be restricted to destinations defined in VirtualServices only
	MeshConfig_OutboundTrafficPolicy_VIRTUAL_SERVICE_ONLY MeshConfig_OutboundTrafficPolicy_Mode = 2
)

var MeshConfig_OutboundTrafficPolicy_Mode_name = map[int32]string{
	0: "REGISTRY_ONLY",
	1: "ALLOW_ANY",
	2: "VIRTUAL_SERVICE_ONLY",
}
var MeshConfig_OutboundTrafficPolicy_Mode_value = map[string]int32{
	"REGISTRY_ONLY":        0,
	"ALLOW_ANY":            1,
	"VIRTUAL_SERVICE_ONLY": 2,
}

func (x MeshConfig_OutboundTrafficPolicy_Mode) String() string {
	return proto.EnumName(MeshConfig_OutboundTrafficPolicy_Mode_name, int32(x))
}
func (MeshConfig_OutboundTrafficPolicy_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_cc82c7a62bb884d0, []int{1, 0, 0}
}

// ProxyConfig defines variables for individual Envoy instances.
type ProxyConfig struct {
	// Path to the generated configuration file directory.
	// Proxy agent generates the actual configuration and stores it in this directory.
	ConfigPath string `protobuf:"bytes,1,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	// Path to the proxy binary
	BinaryPath string `protobuf:"bytes,2,opt,name=binary_path,json=binaryPath,proto3" json:"binary_path,omitempty"`
	// Service cluster defines the name for the service_cluster that is
	// shared by all Envoy instances. This setting corresponds to
	// _--service-cluster_ flag in Envoy.  In a typical Envoy deployment, the
	// _service-cluster_ flag is used to identify the caller, for
	// source-based routing scenarios.
	//
	// Since Istio does not assign a local service/service version to each
	// Envoy instance, the name is same for all of them.  However, the
	// source/caller's identity (e.g., IP address) is encoded in the
	// _--service-node_ flag when launching Envoy.  When the RDS service
	// receives API calls from Envoy, it uses the value of the _service-node_
	// flag to compute routes that are relative to the service instances
	// located at that IP address.
	ServiceCluster string `protobuf:"bytes,3,opt,name=service_cluster,json=serviceCluster,proto3" json:"service_cluster,omitempty"`
	// The time in seconds that Envoy will drain connections during a hot
	// restart. MUST be >=1s (e.g., _1s/1m/1h_)
	DrainDuration *duration.Duration `protobuf:"bytes,4,opt,name=drain_duration,json=drainDuration,proto3" json:"drain_duration,omitempty"`
	// The time in seconds that Envoy will wait before shutting down the
	// parent process during a hot restart. MUST be >=1s (e.g., _1s/1m/1h_).
	// MUST BE greater than _drain_duration_ parameter.
	ParentShutdownDuration *duration.Duration `protobuf:"bytes,5,opt,name=parent_shutdown_duration,json=parentShutdownDuration,proto3" json:"parent_shutdown_duration,omitempty"`
	// Address of the discovery service exposing xDS with mTLS connection.
	DiscoveryAddress string `protobuf:"bytes,6,opt,name=discovery_address,json=discoveryAddress,proto3" json:"discovery_address,omitempty"`
	// Polling interval for service discovery (used by EDS, CDS, LDS, but not RDS). (MUST BE >=1ms)
	DiscoveryRefreshDelay *duration.Duration `protobuf:"bytes,7,opt,name=discovery_refresh_delay,json=discoveryRefreshDelay,proto3" json:"discovery_refresh_delay,omitempty"`
	// Address of the Zipkin service (e.g. _zipkin:9411_).
	ZipkinAddress string `protobuf:"bytes,8,opt,name=zipkin_address,json=zipkinAddress,proto3" json:"zipkin_address,omitempty"`
	// Connection timeout used by Envoy for supporting services. (MUST BE >=1ms)
	ConnectTimeout *duration.Duration `protobuf:"bytes,9,opt,name=connect_timeout,json=connectTimeout,proto3" json:"connect_timeout,omitempty"`
	// IP Address and Port of a statsd UDP listener (e.g. _10.75.241.127:9125_).
	StatsdUdpAddress string `protobuf:"bytes,10,opt,name=statsd_udp_address,json=statsdUdpAddress,proto3" json:"statsd_udp_address,omitempty"`
	// Port on which Envoy should listen for administrative commands.
	ProxyAdminPort int32 `protobuf:"varint,11,opt,name=proxy_admin_port,json=proxyAdminPort,proto3" json:"proxy_admin_port,omitempty"`
	// The availability zone where this Envoy instance is running. When running
	// Envoy as a sidecar in Kubernetes, this flag must be one of the availability
	// zones assigned to a node using failure-domain.beta.kubernetes.io/zone annotation.
	AvailabilityZone string `protobuf:"bytes,12,opt,name=availability_zone,json=availabilityZone,proto3" json:"availability_zone,omitempty"`
	// Authentication policy defines the global switch to control authentication
	// for Envoy-to-Envoy communication for istio components Mixer and Pilot.
	ControlPlaneAuthPolicy AuthenticationPolicy `protobuf:"varint,13,opt,name=control_plane_auth_policy,json=controlPlaneAuthPolicy,proto3,enum=istio.mesh.v1alpha1.AuthenticationPolicy" json:"control_plane_auth_policy,omitempty"`
	// File path of custom proxy configuration, currently used by proxies
	// in front of Mixer and Pilot.
	CustomConfigFile string `protobuf:"bytes,14,opt,name=custom_config_file,json=customConfigFile,proto3" json:"custom_config_file,omitempty"`
	// Maximum length of name field in Envoy's metrics. The length of the name field
	// is determined by the length of a name field in a service and the set of labels that
	// comprise a particular version of the service. The default value is set to 189 characters.
	// Envoy's internal metrics take up 67 characters, for a total of 256 character name per metric.
	// Increase the value of this field if you find that the metrics from Envoys are truncated.
	StatNameLength int32 `protobuf:"varint,15,opt,name=stat_name_length,json=statNameLength,proto3" json:"stat_name_length,omitempty"`
	// The number of worker threads to run. Default value is number of cores on the machine.
	Concurrency int32 `protobuf:"varint,16,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// Path to the proxy bootstrap template file
	ProxyBootstrapTemplatePath string `protobuf:"bytes,17,opt,name=proxy_bootstrap_template_path,json=proxyBootstrapTemplatePath,proto3" json:"proxy_bootstrap_template_path,omitempty"`
	// The mode used to redirect inbound traffic to Envoy.
	InterceptionMode     ProxyConfig_InboundInterceptionMode `protobuf:"varint,18,opt,name=interception_mode,json=interceptionMode,proto3,enum=istio.mesh.v1alpha1.ProxyConfig_InboundInterceptionMode" json:"interception_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ProxyConfig) Reset()         { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()    {}
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_cc82c7a62bb884d0, []int{0}
}
func (m *ProxyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyConfig.Unmarshal(m, b)
}
func (m *ProxyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyConfig.Marshal(b, m, deterministic)
}
func (dst *ProxyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyConfig.Merge(dst, src)
}
func (m *ProxyConfig) XXX_Size() int {
	return xxx_messageInfo_ProxyConfig.Size(m)
}
func (m *ProxyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyConfig proto.InternalMessageInfo

func (m *ProxyConfig) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *ProxyConfig) GetBinaryPath() string {
	if m != nil {
		return m.BinaryPath
	}
	return ""
}

func (m *ProxyConfig) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

func (m *ProxyConfig) GetDrainDuration() *duration.Duration {
	if m != nil {
		return m.DrainDuration
	}
	return nil
}

func (m *ProxyConfig) GetParentShutdownDuration() *duration.Duration {
	if m != nil {
		return m.ParentShutdownDuration
	}
	return nil
}

func (m *ProxyConfig) GetDiscoveryAddress() string {
	if m != nil {
		return m.DiscoveryAddress
	}
	return ""
}

func (m *ProxyConfig) GetDiscoveryRefreshDelay() *duration.Duration {
	if m != nil {
		return m.DiscoveryRefreshDelay
	}
	return nil
}

func (m *ProxyConfig) GetZipkinAddress() string {
	if m != nil {
		return m.ZipkinAddress
	}
	return ""
}

func (m *ProxyConfig) GetConnectTimeout() *duration.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

func (m *ProxyConfig) GetStatsdUdpAddress() string {
	if m != nil {
		return m.StatsdUdpAddress
	}
	return ""
}

func (m *ProxyConfig) GetProxyAdminPort() int32 {
	if m != nil {
		return m.ProxyAdminPort
	}
	return 0
}

func (m *ProxyConfig) GetAvailabilityZone() string {
	if m != nil {
		return m.AvailabilityZone
	}
	return ""
}

func (m *ProxyConfig) GetControlPlaneAuthPolicy() AuthenticationPolicy {
	if m != nil {
		return m.ControlPlaneAuthPolicy
	}
	return AuthenticationPolicy_NONE
}

func (m *ProxyConfig) GetCustomConfigFile() string {
	if m != nil {
		return m.CustomConfigFile
	}
	return ""
}

func (m *ProxyConfig) GetStatNameLength() int32 {
	if m != nil {
		return m.StatNameLength
	}
	return 0
}

func (m *ProxyConfig) GetConcurrency() int32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *ProxyConfig) GetProxyBootstrapTemplatePath() string {
	if m != nil {
		return m.ProxyBootstrapTemplatePath
	}
	return ""
}

func (m *ProxyConfig) GetInterceptionMode() ProxyConfig_InboundInterceptionMode {
	if m != nil {
		return m.InterceptionMode
	}
	return ProxyConfig_REDIRECT
}

// MeshConfig defines mesh-wide variables shared by all Envoy instances in the
// Istio service mesh.
//
// NOTE: This configuration type should be used for the low-level global
// configuration, such as component addresses and port numbers. It should not
// be used for the features of the mesh that can be scoped by service or by
// namespace. Some of the fields in the mesh config are going to be deprecated
// and replaced with several individual configuration types (for example,
// tracing configuration).
type MeshConfig struct {
	// Address of the server that will be used by the proxies for policy
	// check calls. By using different names for mixerCheckServer and
	// mixerReportServer, it is possible to have one set of mixer servers handle
	// policy check calls while another set of mixer servers handle telemetry
	// calls.
	//
	// NOTE: Omitting mixerCheckServer while specifying mixerReportServer is
	// equivalent to setting disablePolicyChecks to true.
	MixerCheckServer string `protobuf:"bytes,1,opt,name=mixer_check_server,json=mixerCheckServer,proto3" json:"mixer_check_server,omitempty"`
	// Address of the server that will be used by the proxies for policy report
	// calls.
	MixerReportServer string `protobuf:"bytes,2,opt,name=mixer_report_server,json=mixerReportServer,proto3" json:"mixer_report_server,omitempty"`
	// Disable policy checks by the mixer service. Default
	// is false, i.e. mixer policy check is enabled by default.
	DisablePolicyChecks bool `protobuf:"varint,3,opt,name=disable_policy_checks,json=disablePolicyChecks,proto3" json:"disable_policy_checks,omitempty"`
	// Port on which Envoy should listen for incoming connections from
	// other services.
	ProxyListenPort int32 `protobuf:"varint,4,opt,name=proxy_listen_port,json=proxyListenPort,proto3" json:"proxy_listen_port,omitempty"`
	// Port on which Envoy should listen for HTTP PROXY requests if set.
	ProxyHttpPort int32 `protobuf:"varint,5,opt,name=proxy_http_port,json=proxyHttpPort,proto3" json:"proxy_http_port,omitempty"`
	// Connection timeout used by Envoy. (MUST BE >=1ms)
	ConnectTimeout *duration.Duration `protobuf:"bytes,6,opt,name=connect_timeout,json=connectTimeout,proto3" json:"connect_timeout,omitempty"`
	// Class of ingress resources to be processed by Istio ingress
	// controller.  This corresponds to the value of
	// "kubernetes.io/ingress.class" annotation.
	IngressClass string `protobuf:"bytes,7,opt,name=ingress_class,json=ingressClass,proto3" json:"ingress_class,omitempty"`
	// Name of the kubernetes service used for the istio ingress controller.
	IngressService string `protobuf:"bytes,8,opt,name=ingress_service,json=ingressService,proto3" json:"ingress_service,omitempty"`
	// Defines whether to use Istio ingress controller for annotated or all ingress resources.
	IngressControllerMode MeshConfig_IngressControllerMode `protobuf:"varint,9,opt,name=ingress_controller_mode,json=ingressControllerMode,proto3,enum=istio.mesh.v1alpha1.MeshConfig_IngressControllerMode" json:"ingress_controller_mode,omitempty"`
	// Authentication policy defines the global switch to control authentication
	// for Envoy-to-Envoy communication.
	// Use authentication_policy instead.
	AuthPolicy MeshConfig_AuthPolicy `protobuf:"varint,10,opt,name=auth_policy,json=authPolicy,proto3,enum=istio.mesh.v1alpha1.MeshConfig_AuthPolicy" json:"auth_policy,omitempty"` // Deprecated: Do not use.
	// Polling interval for RDS (MUST BE >=1ms)
	RdsRefreshDelay *duration.Duration `protobuf:"bytes,11,opt,name=rds_refresh_delay,json=rdsRefreshDelay,proto3" json:"rds_refresh_delay,omitempty"`
	// Flag to control generation of trace spans and request IDs.
	// Requires a trace span collector defined in the proxy configuration.
	EnableTracing bool `protobuf:"varint,12,opt,name=enable_tracing,json=enableTracing,proto3" json:"enable_tracing,omitempty"`
	// File address for the proxy access log (e.g. /dev/stdout).
	// Empty value disables access logging.
	AccessLogFile string `protobuf:"bytes,13,opt,name=access_log_file,json=accessLogFile,proto3" json:"access_log_file,omitempty"`
	// Default proxy config used by the proxy injection mechanism operating in the mesh
	// (e.g. Kubernetes admission controller)
	// In case of Kubernetes, the proxy config is applied once during the injection process,
	// and remain constant for the duration of the pod. The rest of the mesh config can be changed
	// at runtime and config gets distributed dynamically.
	DefaultConfig *ProxyConfig `protobuf:"bytes,14,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	// DEPRECATED. Mixer address. This option will be removed soon. Please
	// use mixer_check and mixer_report.
	MixerAddress string `protobuf:"bytes,16,opt,name=mixer_address,json=mixerAddress,proto3" json:"mixer_address,omitempty"`
	// Set the default behavior of the sidecar for handling outbound traffic from the application.
	// While the default mode should work out of the box, if your application uses one or more external services that
	// are not known apriori, setting the policy to ALLOW_ANY will cause the sidecars to route traffic to the any
	// requested destination.
	// Users are strongly encouraged to use ServiceEntries to explicitly declare any external dependencies,
	// instead of using allow_any.
	OutboundTrafficPolicy *MeshConfig_OutboundTrafficPolicy `protobuf:"bytes,17,opt,name=outbound_traffic_policy,json=outboundTrafficPolicy,proto3" json:"outbound_traffic_policy,omitempty"`
	// Enables clide side policy checks.
	EnableClientSidePolicyCheck bool `protobuf:"varint,19,opt,name=enable_client_side_policy_check,json=enableClientSidePolicyCheck,proto3" json:"enable_client_side_policy_check,omitempty"`
	// Unix Domain Socket through which envoy communicates with NodeAgent SDS to get key/cert for mTLS.
	// Use secret-mount files instead of SDS if set to empty.
	SdsUdsPath string `protobuf:"bytes,20,opt,name=sds_uds_path,json=sdsUdsPath,proto3" json:"sds_uds_path,omitempty"`
	// Polling interval for SDS (MUST BE >=1ms)
	SdsRefreshDelay      *duration.Duration `protobuf:"bytes,21,opt,name=sds_refresh_delay,json=sdsRefreshDelay,proto3" json:"sds_refresh_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MeshConfig) Reset()         { *m = MeshConfig{} }
func (m *MeshConfig) String() string { return proto.CompactTextString(m) }
func (*MeshConfig) ProtoMessage()    {}
func (*MeshConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_cc82c7a62bb884d0, []int{1}
}
func (m *MeshConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshConfig.Unmarshal(m, b)
}
func (m *MeshConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshConfig.Marshal(b, m, deterministic)
}
func (dst *MeshConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshConfig.Merge(dst, src)
}
func (m *MeshConfig) XXX_Size() int {
	return xxx_messageInfo_MeshConfig.Size(m)
}
func (m *MeshConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MeshConfig proto.InternalMessageInfo

func (m *MeshConfig) GetMixerCheckServer() string {
	if m != nil {
		return m.MixerCheckServer
	}
	return ""
}

func (m *MeshConfig) GetMixerReportServer() string {
	if m != nil {
		return m.MixerReportServer
	}
	return ""
}

func (m *MeshConfig) GetDisablePolicyChecks() bool {
	if m != nil {
		return m.DisablePolicyChecks
	}
	return false
}

func (m *MeshConfig) GetProxyListenPort() int32 {
	if m != nil {
		return m.ProxyListenPort
	}
	return 0
}

func (m *MeshConfig) GetProxyHttpPort() int32 {
	if m != nil {
		return m.ProxyHttpPort
	}
	return 0
}

func (m *MeshConfig) GetConnectTimeout() *duration.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

func (m *MeshConfig) GetIngressClass() string {
	if m != nil {
		return m.IngressClass
	}
	return ""
}

func (m *MeshConfig) GetIngressService() string {
	if m != nil {
		return m.IngressService
	}
	return ""
}

func (m *MeshConfig) GetIngressControllerMode() MeshConfig_IngressControllerMode {
	if m != nil {
		return m.IngressControllerMode
	}
	return MeshConfig_OFF
}

// Deprecated: Do not use.
func (m *MeshConfig) GetAuthPolicy() MeshConfig_AuthPolicy {
	if m != nil {
		return m.AuthPolicy
	}
	return MeshConfig_NONE
}

func (m *MeshConfig) GetRdsRefreshDelay() *duration.Duration {
	if m != nil {
		return m.RdsRefreshDelay
	}
	return nil
}

func (m *MeshConfig) GetEnableTracing() bool {
	if m != nil {
		return m.EnableTracing
	}
	return false
}

func (m *MeshConfig) GetAccessLogFile() string {
	if m != nil {
		return m.AccessLogFile
	}
	return ""
}

func (m *MeshConfig) GetDefaultConfig() *ProxyConfig {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func (m *MeshConfig) GetMixerAddress() string {
	if m != nil {
		return m.MixerAddress
	}
	return ""
}

func (m *MeshConfig) GetOutboundTrafficPolicy() *MeshConfig_OutboundTrafficPolicy {
	if m != nil {
		return m.OutboundTrafficPolicy
	}
	return nil
}

func (m *MeshConfig) GetEnableClientSidePolicyCheck() bool {
	if m != nil {
		return m.EnableClientSidePolicyCheck
	}
	return false
}

func (m *MeshConfig) GetSdsUdsPath() string {
	if m != nil {
		return m.SdsUdsPath
	}
	return ""
}

func (m *MeshConfig) GetSdsRefreshDelay() *duration.Duration {
	if m != nil {
		return m.SdsRefreshDelay
	}
	return nil
}

type MeshConfig_OutboundTrafficPolicy struct {
	Mode                 MeshConfig_OutboundTrafficPolicy_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=istio.mesh.v1alpha1.MeshConfig_OutboundTrafficPolicy_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *MeshConfig_OutboundTrafficPolicy) Reset()         { *m = MeshConfig_OutboundTrafficPolicy{} }
func (m *MeshConfig_OutboundTrafficPolicy) String() string { return proto.CompactTextString(m) }
func (*MeshConfig_OutboundTrafficPolicy) ProtoMessage()    {}
func (*MeshConfig_OutboundTrafficPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_cc82c7a62bb884d0, []int{1, 0}
}
func (m *MeshConfig_OutboundTrafficPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshConfig_OutboundTrafficPolicy.Unmarshal(m, b)
}
func (m *MeshConfig_OutboundTrafficPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshConfig_OutboundTrafficPolicy.Marshal(b, m, deterministic)
}
func (dst *MeshConfig_OutboundTrafficPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshConfig_OutboundTrafficPolicy.Merge(dst, src)
}
func (m *MeshConfig_OutboundTrafficPolicy) XXX_Size() int {
	return xxx_messageInfo_MeshConfig_OutboundTrafficPolicy.Size(m)
}
func (m *MeshConfig_OutboundTrafficPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshConfig_OutboundTrafficPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_MeshConfig_OutboundTrafficPolicy proto.InternalMessageInfo

func (m *MeshConfig_OutboundTrafficPolicy) GetMode() MeshConfig_OutboundTrafficPolicy_Mode {
	if m != nil {
		return m.Mode
	}
	return MeshConfig_OutboundTrafficPolicy_REGISTRY_ONLY
}

func init() {
	proto.RegisterType((*ProxyConfig)(nil), "istio.mesh.v1alpha1.ProxyConfig")
	proto.RegisterType((*MeshConfig)(nil), "istio.mesh.v1alpha1.MeshConfig")
	proto.RegisterType((*MeshConfig_OutboundTrafficPolicy)(nil), "istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy")
	proto.RegisterEnum("istio.mesh.v1alpha1.AuthenticationPolicy", AuthenticationPolicy_name, AuthenticationPolicy_value)
	proto.RegisterEnum("istio.mesh.v1alpha1.ProxyConfig_InboundInterceptionMode", ProxyConfig_InboundInterceptionMode_name, ProxyConfig_InboundInterceptionMode_value)
	proto.RegisterEnum("istio.mesh.v1alpha1.MeshConfig_IngressControllerMode", MeshConfig_IngressControllerMode_name, MeshConfig_IngressControllerMode_value)
	proto.RegisterEnum("istio.mesh.v1alpha1.MeshConfig_AuthPolicy", MeshConfig_AuthPolicy_name, MeshConfig_AuthPolicy_value)
	proto.RegisterEnum("istio.mesh.v1alpha1.MeshConfig_OutboundTrafficPolicy_Mode", MeshConfig_OutboundTrafficPolicy_Mode_name, MeshConfig_OutboundTrafficPolicy_Mode_value)
}

func init() { proto.RegisterFile("mesh/v1alpha1/config.proto", fileDescriptor_config_cc82c7a62bb884d0) }

var fileDescriptor_config_cc82c7a62bb884d0 = []byte{
	// 1209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x6f, 0x6f, 0x1a, 0xc7,
	0x13, 0x36, 0x0e, 0xb1, 0xf1, 0x60, 0xf0, 0xb1, 0x8e, 0xe3, 0x8b, 0x7f, 0x7f, 0x82, 0x1c, 0x35,
	0x75, 0xd3, 0x08, 0x2b, 0x8e, 0x2a, 0xb5, 0x7d, 0x51, 0x15, 0x63, 0x9c, 0x10, 0x11, 0x70, 0x8f,
	0x73, 0x5a, 0xe7, 0xcd, 0x6a, 0xb9, 0x5b, 0xcc, 0x2a, 0xc7, 0xed, 0x69, 0x77, 0xcf, 0x0d, 0x79,
	0xdb, 0x8f, 0xd3, 0xef, 0xd0, 0xcf, 0xd3, 0x8f, 0x51, 0xdd, 0xec, 0x81, 0x49, 0x4a, 0x84, 0xda,
	0x97, 0x3c, 0xfb, 0xcc, 0xcc, 0xee, 0xcc, 0x3c, 0xcf, 0x01, 0x07, 0x13, 0xae, 0xc7, 0xc7, 0x37,
	0xcf, 0x58, 0x94, 0x8c, 0xd9, 0xb3, 0xe3, 0x40, 0xc6, 0x23, 0x71, 0xdd, 0x48, 0x94, 0x34, 0x92,
	0xec, 0x0a, 0x6d, 0x84, 0x6c, 0x64, 0x8c, 0xc6, 0x8c, 0x71, 0xf0, 0xff, 0x6b, 0x29, 0xaf, 0x23,
	0x7e, 0x8c, 0x94, 0x61, 0x3a, 0x3a, 0x0e, 0x53, 0xc5, 0x8c, 0x90, 0xb1, 0x0d, 0x3a, 0xfc, 0xa3,
	0x04, 0xe5, 0x0b, 0x25, 0xdf, 0x4f, 0x5b, 0x98, 0x8a, 0x3c, 0x84, 0xb2, 0x4d, 0x4a, 0x13, 0x66,
	0xc6, 0x6e, 0xa1, 0x5e, 0x38, 0xda, 0xf2, 0xc0, 0x42, 0x17, 0xcc, 0x8c, 0x33, 0xc2, 0x50, 0xc4,
	0x4c, 0x4d, 0x2d, 0x61, 0xdd, 0x12, 0x2c, 0x84, 0x84, 0x2f, 0x61, 0x47, 0x73, 0x75, 0x23, 0x02,
	0x4e, 0x83, 0x28, 0xd5, 0x86, 0x2b, 0xf7, 0x0e, 0x92, 0xaa, 0x39, 0xdc, 0xb2, 0x28, 0xf9, 0x11,
	0xaa, 0xa1, 0x62, 0x22, 0xa6, 0xb3, 0x2b, 0xb9, 0xc5, 0x7a, 0xe1, 0xa8, 0x7c, 0xf2, 0xa0, 0x61,
	0xef, 0xdc, 0x98, 0xdd, 0xb9, 0x71, 0x96, 0x13, 0xbc, 0x0a, 0x06, 0xcc, 0x7e, 0x92, 0x01, 0xb8,
	0x09, 0x53, 0x3c, 0x36, 0x54, 0x8f, 0x53, 0x13, 0xca, 0x5f, 0x17, 0x72, 0xdd, 0x5d, 0x95, 0xeb,
	0xbe, 0x0d, 0x1d, 0xe4, 0x91, 0xf3, 0xa4, 0x5f, 0x43, 0x2d, 0x14, 0x3a, 0x90, 0x37, 0x5c, 0x4d,
	0x29, 0x0b, 0x43, 0xc5, 0xb5, 0x76, 0x37, 0xf0, 0x05, 0xce, 0xfc, 0xa0, 0x69, 0x71, 0xf2, 0x13,
	0xec, 0xdf, 0x92, 0x15, 0x1f, 0x29, 0xae, 0xc7, 0x34, 0xe4, 0x11, 0x9b, 0xba, 0x9b, 0xab, 0x2e,
	0xb0, 0x37, 0x8f, 0xf4, 0x6c, 0xe0, 0x59, 0x16, 0x47, 0xbe, 0x80, 0xea, 0x07, 0x91, 0xbc, 0x13,
	0xf1, 0xbc, 0x78, 0x09, 0x8b, 0x57, 0x2c, 0x3a, 0xab, 0x7c, 0x0a, 0x3b, 0x81, 0x8c, 0x63, 0x1e,
	0x18, 0x6a, 0xc4, 0x84, 0xcb, 0xd4, 0xb8, 0x5b, 0xab, 0x2a, 0x56, 0xf3, 0x08, 0xdf, 0x06, 0x90,
	0xa7, 0x40, 0xb4, 0x61, 0x46, 0x87, 0x34, 0x0d, 0x93, 0x79, 0x39, 0xb0, 0x6f, 0xb5, 0x27, 0x97,
	0x61, 0x32, 0xab, 0x78, 0x04, 0x4e, 0x92, 0x6d, 0x0a, 0x65, 0xe1, 0x44, 0xc4, 0x34, 0x91, 0xca,
	0xb8, 0xe5, 0x7a, 0xe1, 0xe8, 0xae, 0x57, 0x45, 0xbc, 0x99, 0xc1, 0x17, 0x52, 0x99, 0xac, 0x85,
	0xec, 0x86, 0x89, 0x88, 0x0d, 0x45, 0x24, 0xcc, 0x94, 0x7e, 0x90, 0x31, 0x77, 0xb7, 0x6d, 0xda,
	0xc5, 0x83, 0xb7, 0x32, 0xe6, 0x24, 0x84, 0x07, 0x81, 0x8c, 0x8d, 0x92, 0x11, 0x4d, 0x22, 0x16,
	0x73, 0xca, 0x52, 0x33, 0xa6, 0x89, 0x8c, 0x44, 0x30, 0x75, 0x2b, 0xf5, 0xc2, 0x51, 0xf5, 0xe4,
	0xab, 0xc6, 0x92, 0xd5, 0x6e, 0x34, 0x53, 0x33, 0xe6, 0xb1, 0x11, 0x01, 0x3e, 0xee, 0x02, 0x03,
	0xbc, 0xfb, 0x79, 0xae, 0x8b, 0x2c, 0x55, 0xc6, 0xb0, 0x78, 0xf6, 0xd4, 0x20, 0xd5, 0x46, 0x4e,
	0x68, 0xbe, 0xde, 0x23, 0x11, 0x71, 0xb7, 0x6a, 0xef, 0x64, 0x4f, 0xac, 0x02, 0xce, 0x45, 0xc4,
	0xb3, 0xa7, 0x66, 0xcf, 0xa7, 0x31, 0x9b, 0x70, 0x1a, 0xf1, 0xf8, 0xda, 0x8c, 0xdd, 0x1d, 0xfb,
	0xd4, 0x0c, 0xef, 0xb1, 0x09, 0xef, 0x22, 0x4a, 0xea, 0xa8, 0x97, 0x20, 0x55, 0x8a, 0xc7, 0xc1,
	0xd4, 0x75, 0x90, 0xb4, 0x08, 0x91, 0x26, 0xfc, 0xcf, 0xb6, 0x6d, 0x28, 0xa5, 0xd1, 0x46, 0xb1,
	0x84, 0x1a, 0x3e, 0x49, 0x22, 0x66, 0xb8, 0x95, 0x50, 0x0d, 0x2f, 0x71, 0x80, 0xa4, 0xd3, 0x19,
	0xc7, 0xcf, 0x29, 0x28, 0x29, 0x0e, 0x35, 0x11, 0x1b, 0xae, 0x02, 0x9e, 0x64, 0x4f, 0xa5, 0x13,
	0x19, 0x72, 0x97, 0x60, 0x6b, 0xbe, 0x5d, 0xda, 0x9a, 0x05, 0x45, 0x37, 0x3a, 0xf1, 0x50, 0xa6,
	0x71, 0xd8, 0x59, 0x48, 0xf0, 0x5a, 0x86, 0xdc, 0x73, 0xc4, 0x27, 0xc8, 0xe1, 0x73, 0xd8, 0xff,
	0x0c, 0x99, 0x6c, 0x43, 0xc9, 0x6b, 0x9f, 0x75, 0xbc, 0x76, 0xcb, 0x77, 0xd6, 0x08, 0xc0, 0x86,
	0x7f, 0xe1, 0xf5, 0x7f, 0xb9, 0x72, 0x0a, 0x87, 0xbf, 0x95, 0x01, 0x5e, 0x73, 0x3d, 0xce, 0xfd,
	0xe3, 0x29, 0x90, 0x89, 0x78, 0xcf, 0x15, 0x0d, 0xc6, 0x3c, 0x78, 0x47, 0x33, 0xc9, 0x73, 0x95,
	0xdb, 0x88, 0x83, 0x27, 0xad, 0xec, 0x60, 0x80, 0x38, 0x69, 0xc0, 0xae, 0x65, 0x2b, 0x9e, 0xad,
	0xd3, 0x8c, 0x6e, 0x4d, 0xa5, 0x86, 0x47, 0x1e, 0x9e, 0xe4, 0xfc, 0x13, 0xc8, 0x44, 0xc3, 0x86,
	0x11, 0xcf, 0x17, 0xc4, 0x96, 0xd1, 0xe8, 0x30, 0x25, 0x6f, 0x37, 0x3f, 0xb4, 0x33, 0xc7, 0x42,
	0x9a, 0x3c, 0x81, 0x9a, 0xed, 0x7f, 0x24, 0xb4, 0xe1, 0xf9, 0xde, 0x16, 0x71, 0x4e, 0x3b, 0x78,
	0xd0, 0x45, 0x1c, 0x17, 0xf7, 0x31, 0x58, 0x88, 0x8e, 0x8d, 0x49, 0x2c, 0xf3, 0x2e, 0x32, 0x2b,
	0x08, 0xbf, 0x34, 0x26, 0x41, 0xde, 0x12, 0xf1, 0x6d, 0xfc, 0x53, 0xf1, 0x3d, 0x82, 0x8a, 0x88,
	0xaf, 0x33, 0x65, 0xd1, 0x20, 0x62, 0x5a, 0xa3, 0x61, 0x6c, 0x79, 0xdb, 0x39, 0xd8, 0xca, 0xb0,
	0xcc, 0x4c, 0x67, 0xa4, 0xdc, 0x3d, 0x73, 0x37, 0xa8, 0xe6, 0xf0, 0xc0, 0xa2, 0x64, 0x02, 0xfb,
	0xf3, 0x6c, 0x56, 0x01, 0x11, 0x57, 0x76, 0x51, 0xb6, 0x70, 0x51, 0xbe, 0x59, 0xba, 0x28, 0xb7,
	0x93, 0x6b, 0x74, 0xf2, 0xba, 0xf3, 0x68, 0xdc, 0x92, 0x3d, 0xb1, 0x0c, 0x26, 0x7d, 0x28, 0x2f,
	0xca, 0x14, 0xb0, 0xc4, 0x93, 0x55, 0x25, 0x6e, 0xf5, 0x78, 0xba, 0xee, 0x16, 0x3c, 0x60, 0xb7,
	0xfa, 0x6c, 0x43, 0x4d, 0x85, 0xfa, 0x13, 0x0b, 0x2d, 0xaf, 0xea, 0xe9, 0x8e, 0x0a, 0xf5, 0xa7,
	0xe6, 0xc9, 0x63, 0xdc, 0x0f, 0xa3, 0x58, 0x20, 0xe2, 0x6b, 0xb4, 0x9d, 0x92, 0x57, 0xb1, 0xa8,
	0x6f, 0xc1, 0x6c, 0xce, 0x2c, 0x08, 0xb2, 0x66, 0x45, 0x32, 0xb7, 0x82, 0x8a, 0x35, 0x59, 0x0b,
	0x77, 0xa5, 0xf5, 0x81, 0x17, 0x50, 0x0d, 0xf9, 0x88, 0xa5, 0x91, 0xc9, 0x6d, 0x03, 0x1d, 0xa3,
	0x7c, 0x52, 0x5f, 0xa5, 0x3a, 0xaf, 0x92, 0xc7, 0xe5, 0xb2, 0x78, 0x04, 0x15, 0xbb, 0xe8, 0x33,
	0x93, 0x75, 0xec, 0xb0, 0x11, 0x9c, 0x19, 0xec, 0x04, 0xf6, 0x65, 0x6a, 0x50, 0x80, 0xd9, 0xf5,
	0x47, 0x23, 0x11, 0xcc, 0x1a, 0x5c, 0xc3, 0xb2, 0x2b, 0x67, 0xd8, 0xcf, 0xc3, 0x7d, 0x1b, 0x9d,
	0x7b, 0xe2, 0x9e, 0x5c, 0x06, 0x93, 0x33, 0x78, 0x98, 0xf7, 0x2a, 0x88, 0x04, 0x7e, 0x44, 0x45,
	0xf8, 0xb1, 0xae, 0xdc, 0x5d, 0x6c, 0xde, 0x7f, 0x2c, 0xad, 0x85, 0xac, 0x81, 0x08, 0x17, 0xf5,
	0x45, 0xea, 0xb0, 0xad, 0x43, 0x4d, 0xd3, 0x50, 0x5b, 0x37, 0xbb, 0x67, 0xff, 0x10, 0xe8, 0x50,
	0x5f, 0x86, 0x1a, 0xdd, 0xab, 0x0d, 0x35, 0xfd, 0xb7, 0xd1, 0xee, 0xad, 0x1c, 0xad, 0xfe, 0x78,
	0xb4, 0x07, 0xbf, 0x17, 0x60, 0x6f, 0xe9, 0xfb, 0x48, 0x0f, 0x8a, 0xb8, 0xe8, 0x05, 0xdc, 0xc2,
	0xef, 0xff, 0x55, 0x93, 0x1a, 0xb8, 0xed, 0x98, 0xe7, 0xf0, 0x14, 0x8a, 0xb8, 0xe4, 0x35, 0xa8,
	0x78, 0xed, 0x17, 0x9d, 0x81, 0xef, 0x5d, 0xd1, 0x7e, 0xaf, 0x7b, 0xe5, 0xac, 0x91, 0x0a, 0x6c,
	0x35, 0xbb, 0xdd, 0xfe, 0xcf, 0xb4, 0xd9, 0xbb, 0x72, 0x0a, 0xc4, 0x85, 0x7b, 0x6f, 0x3a, 0x9e,
	0x7f, 0xd9, 0xec, 0xd2, 0x41, 0xdb, 0x7b, 0xd3, 0x69, 0xb5, 0x2d, 0x71, 0xfd, 0xf0, 0x3b, 0xd8,
	0x5b, 0x2a, 0x28, 0xb2, 0x09, 0x77, 0xfa, 0xe7, 0xe7, 0xce, 0x1a, 0x29, 0xc3, 0xe6, 0x59, 0xfb,
	0xbc, 0x79, 0xd9, 0xf5, 0x9d, 0x42, 0xe6, 0xa8, 0x03, 0xdf, 0xeb, 0xb4, 0x7c, 0x67, 0xfd, 0xf0,
	0x31, 0xc0, 0xc2, 0x87, 0xab, 0x04, 0xc5, 0x5e, 0xbf, 0xd7, 0x76, 0xd6, 0x48, 0x15, 0xe0, 0xf5,
	0x25, 0xd6, 0xf2, 0xbb, 0x03, 0xa7, 0xf0, 0xaa, 0x58, 0xda, 0x71, 0x9c, 0x57, 0xc5, 0x12, 0x71,
	0x76, 0x9f, 0xfc, 0x00, 0xf7, 0x96, 0x7d, 0x0e, 0x3f, 0x1f, 0x4d, 0xb6, 0x61, 0xb3, 0xd3, 0x7b,
	0xd9, 0xf6, 0x3a, 0xbe, 0xf3, 0xe7, 0xe6, 0xe9, 0x7f, 0xdf, 0x1e, 0xd8, 0xae, 0x09, 0x79, 0xcc,
	0x12, 0x71, 0xfc, 0xd1, 0xdf, 0xcc, 0xe1, 0x06, 0x8e, 0xe7, 0xf9, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xb5, 0xd2, 0x0b, 0xd7, 0x7e, 0x0a, 0x00, 0x00,
}
