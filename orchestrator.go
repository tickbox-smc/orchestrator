package orchestrator

import (
	"time"
)

// //go:generate stringer -type=MoodState
// type MoodState int

// // Here we define all the possible mood states using an
// // iota enumerator.
// const (
// 	MoodStateNeutral MoodState = iota
// 	MoodStateHappy
// 	MoodStateSad
// 	MoodStateAngry
// 	MoodStateHopeful
// 	MoodStateThrilled
// 	MoodStateBored
// 	MoodStateShy
// 	MoodStateComical
// 	MoodStateOnCloudNine
// )

// This is a type we embed into types we want to keep a
// check on for auditing purposes
type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

// This is the type that represents a VPC
type Vpc struct {
	AuditableContent // Embedded type
	Name string
	CIDR string
	AZs string
	PrivateSubnets string //This is a list of subnets
	PublicSubnets string //This is a list of subnets
	DatabaseSubnets	string //This is a list of subnets
	ElasticacheSubnets string //This is a list of subnets
	RedshiftSubnets	string //This is a list of subnets
	IntraSubnets string //This is a list of subnets
	CreateDatabaseSubnetGroup	bool
	EnableDnsHostnames	bool
	EnableDnsSupport bool
	EnableNatGateway bool
	SingleNatGateway bool
	EnableVpnGateway bool
	EnableDhcpOptions bool
	DhcpOptionsDomainName string
	DhcpOptionsDomainNameServers string //This is a list of IPs
	//VPC endpoint for S3
	EnableS3Endpoint bool
	//VPC endpoint for DynamoDB
	EnableDynamodbEndpoint bool
	//VPC endpoint for SSM
	EnableSsmEndpoint bool
	SsmEndpointPrivateDnsEnabled bool
	SsmEndpointSecurityGroupIDs string // = ["${data.aws_security_group.default.id}"] # ssm_endpoint_subnet_ids = ["..."]
	//VPC endpoint for SSMMESSAGES
	EnableSsmMessagesEndpoint bool
	SsmMessagesEndpointPrivateDnsEnabled bool
	SsmMessagesEndpointSecurityGroupIDs string // = ["${data.aws_security_group.default.id}"]
	// VPC Endpoint for EC2
	EnableEC2Endpoint bool
	EC2EndpointPrivateDnsEnabled bool
	EC2EndpointSecurityGroupIDs string // = ["${data.aws_security_group.default.id}"]
	//VPC Endpoint for EC2MESSAGES
	EnableEC2MessagesEndpoint bool
	EC2MessagesEndpointPrivateDnsEnabled bool
	EC2MessagesEndpointSecurityGroupIDs string // = ["${data.aws_security_group.default.id}"]
	//VPC Endpoint for ECR API
	EnableEcrApiEndpoint bool
	EcrApiEndpointPrivateDnsEnabled bool
	EcrApiEndpointSecurityGroupIDs string //  = ["${data.aws_security_group.default.id}"]
	//VPC Endpoint for ECR DKR
	EnableEcrDkrEndpoint bool
	EcrDkrEndpointPrivateDnsEnabled bool
	EcrDkrEndpointSecurityGroupIDs string // = ["${data.aws_security_group.default.id}"]


}

// // Map that holds the various mood states with keys to serve as
// // aliases to their respective mood state
// var Moods map[string]MoodState

// The init() function is responsible for initializing the mood state
// func init() {
// 	Moods = map[string]MoodState{"netural": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry, "hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored, "shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateOnCloudNine}
// }

// This is the function responsible for creating a new VPC.
func NewVPC(
	username string,
	name string,
	cidr string,
	azs string,
	privateSubnets []string,
	publicSubnets []string,
	databaseSubnets []string,
	elasticacheSubnets []string,
	redshiftSubnets []string,
	intraSubnets []string,
	createDatabaseSubnetGroup bool,
	enableDnsHostnames bool,
	enableDnsSupport bool,
	enableNatGateway  bool,
	singleNatGateway bool,
	enableVpnGateway bool,
	enableDhcpOptions bool,
	dhcpOptionsDomainName string,
	dhcpOptionsDomainNameServers []string,
	enableS3Endpoint bool,
	enableDynamodbEndpoint bool,
	enableSsmEndpoint bool,
	ssmEndpointPrivateDnsEnabled bool,
	ssmEndpointSecurityGroupIDs []string,
	enableSsmMessagesEndpoint bool,
	ssmMessagesEndpointPrivateDnsEnabled bool,
	ssmMessagesEndpointSecurityGroupIDs []string,
	enableEC2Endpoint bool,
	ec2EndpointPrivateDnsEnabled bool,
	ec2EndpointSecurityGroupIDs []string,
	enableEC2MessagesEndpoint bool,
	ec2MessagesEndpointPrivateDnsEnabled bool,
	ec2MessagesEndpointSecurityGroupIDs []string,
	enableEcrApiEndpoint bool,
	ecrApiEndpointPrivateDnsEnabled bool,
	ecrApiEndpointSecurityGroupIDs []string
	enableEcrDkrEndpoint bool,
	ecrDkrEndpointPrivateDnsEnabled bool,
	ecrDkrEndpointSecurityGroupIDs []string
) *Vpc {
	
	//username string, mood MoodState, caption string, messageBody string, url string, imageURI string, thumbnailURI string, keywords []string) *Post {

	auditableContent := AuditableContent{CreatedBy: username, TimeCreated: time.Now()}
	return &Vpc{
		Name: name,
		CIDR: cidr,
		AZs: azs,
		PrivateSubnets: privateSubnets,
		PublicSubnets: publicSubnets,
		DatabaseSubnets: databaseSubnets,
		ElasticacheSubnets: elasticacheSubnets,
		RedshiftSubnets: redshiftSubnets,
		IntraSubnets: intraSubnets,
		CreateDatabaseSubnetGroup: createDatabaseSubnetGroup,
		EnableDnsHostnames: enableDnsHostnames,
		EnableDnsSupport: enableDnsSupport,
		EnableNatGateway: enableNatGateway,
		SingleNatGateway: singleNatGateway,
		EnableVpnGateway: enableVpnGateway,
		EnableDhcpOptions: enableDhcpOptions,
		DhcpOptionsDomainName: dhcpOptionsDomainName,
		DhcpOptionsDomainNameServers: dhcpOptionsDomainNameServers,
		EnableS3Endpoint: enableS3Endpoint,
		EnableDynamodbEndpoint: enableDynamodbEndpoint,
		EnableSsmEndpoint: enableSsmEndpoint,
		SsmEndpointPrivateDnsEnabled: ssmEndpointPrivateDnsEnabled,
		SsmEndpointSecurityGroupIDs: ssmEndpointSecurityGroupIDs,
		EnableSsmMessagesEndpoint: enableSsmMessagesEndpoint,
		SsmMessagesEndpointPrivateDnsEnabled: ssmMessagesEndpointPrivateDnsEnabled,
		SsmMessagesEndpointSecurityGroupIDs: ssmMessagesEndpointSecurityGroupIDs,
		EnableEC2Endpoint: enableEC2Endpoint,
		EC2EndpointPrivateDnsEnabled: ec2EndpointPrivateDnsEnabled,
		EC2EndpointSecurityGroupIDs: ec2EndpointSecurityGroupIDs,
		EnableEC2MessagesEndpoint: enableEC2MessagesEndpoint,
		EC2MessagesEndpointPrivateDnsEnabled: ec2MessagesEndpointPrivateDnsEnabled,
		EC2MessagesEndpointSecurityGroupIDs: ec2MessagesEndpointSecurityGroupIDs,
		EnableEcrApiEndpoint: enableEcrApiEndpoint,
		EcrApiEndpointPrivateDnsEnabled: ecrApiEndpointPrivateDnsEnabled,
		EcrApiEndpointSecurityGroupIDs: ecrApiEndpointSecurityGroupIDs,
		EnableEcrDkrEndpoint: enableEcrDkrEndpoint,
		EcrDkrEndpointPrivateDnsEnabled: ecrDkrEndpointPrivateDnsEnabled,
		EcrDkrEndpointSecurityGroupIDs: ecrDkrEndpointSecurityGroupIDs,
		AuditableContent: auditableContent
	}
}
