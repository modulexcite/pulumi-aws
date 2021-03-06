// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage a VPC peering connection.
//
// > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
// both a standalone VPC Peering Connection Options and a VPC Peering Connection
// resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
// connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
// Doing so will cause a conflict of options and will overwrite the options.
// Using a VPC Peering Connection Options resource decouples management of the connection options from
// management of the VPC Peering Connection and allows options to be set correctly in cross-account scenarios.
//
// > **Note:** For cross-account (requester's AWS account differs from the accepter's AWS account) or inter-region
// VPC Peering Connections use the `ec2.VpcPeeringConnection` resource to manage the requester's side of the
// connection and use the `ec2.VpcPeeringConnectionAccepter` resource to manage the accepter's side of the connection.
//
//
// ## Notes
//
// If both VPCs are not in the same AWS account do not enable the `autoAccept` attribute.
// The accepter can manage its side of the connection using the `ec2.VpcPeeringConnectionAccepter` resource
// or accept the connection manually using the AWS Management Console, AWS CLI, through SDKs, etc.
type VpcPeeringConnection struct {
	pulumi.CustomResourceState

	// The status of the VPC Peering Connection request.
	AcceptStatus pulumi.StringOutput `pulumi:"acceptStatus"`
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter VpcPeeringConnectionAccepterTypeOutput `pulumi:"accepter"`
	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept pulumi.BoolPtrOutput `pulumi:"autoAccept"`
	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId pulumi.StringOutput `pulumi:"peerOwnerId"`
	// The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
	// and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
	PeerRegion pulumi.StringOutput `pulumi:"peerRegion"`
	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId pulumi.StringOutput `pulumi:"peerVpcId"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester VpcPeeringConnectionRequesterOutput `pulumi:"requester"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the requester VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcPeeringConnection registers a new resource with the given unique name, arguments, and options.
func NewVpcPeeringConnection(ctx *pulumi.Context,
	name string, args *VpcPeeringConnectionArgs, opts ...pulumi.ResourceOption) (*VpcPeeringConnection, error) {
	if args == nil || args.PeerVpcId == nil {
		return nil, errors.New("missing required argument 'PeerVpcId'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &VpcPeeringConnectionArgs{}
	}
	var resource VpcPeeringConnection
	err := ctx.RegisterResource("aws:ec2/vpcPeeringConnection:VpcPeeringConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPeeringConnection gets an existing VpcPeeringConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPeeringConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPeeringConnectionState, opts ...pulumi.ResourceOption) (*VpcPeeringConnection, error) {
	var resource VpcPeeringConnection
	err := ctx.ReadResource("aws:ec2/vpcPeeringConnection:VpcPeeringConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPeeringConnection resources.
type vpcPeeringConnectionState struct {
	// The status of the VPC Peering Connection request.
	AcceptStatus *string `pulumi:"acceptStatus"`
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter *VpcPeeringConnectionAccepterType `pulumi:"accepter"`
	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept *bool `pulumi:"autoAccept"`
	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId *string `pulumi:"peerOwnerId"`
	// The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
	// and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
	PeerRegion *string `pulumi:"peerRegion"`
	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId *string `pulumi:"peerVpcId"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester *VpcPeeringConnectionRequester `pulumi:"requester"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the requester VPC.
	VpcId *string `pulumi:"vpcId"`
}

type VpcPeeringConnectionState struct {
	// The status of the VPC Peering Connection request.
	AcceptStatus pulumi.StringPtrInput
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter VpcPeeringConnectionAccepterTypePtrInput
	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept pulumi.BoolPtrInput
	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId pulumi.StringPtrInput
	// The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
	// and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
	PeerRegion pulumi.StringPtrInput
	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId pulumi.StringPtrInput
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester VpcPeeringConnectionRequesterPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the requester VPC.
	VpcId pulumi.StringPtrInput
}

func (VpcPeeringConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPeeringConnectionState)(nil)).Elem()
}

type vpcPeeringConnectionArgs struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter *VpcPeeringConnectionAccepterType `pulumi:"accepter"`
	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept *bool `pulumi:"autoAccept"`
	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId *string `pulumi:"peerOwnerId"`
	// The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
	// and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
	PeerRegion *string `pulumi:"peerRegion"`
	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId string `pulumi:"peerVpcId"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester *VpcPeeringConnectionRequester `pulumi:"requester"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the requester VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcPeeringConnection resource.
type VpcPeeringConnectionArgs struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter VpcPeeringConnectionAccepterTypePtrInput
	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept pulumi.BoolPtrInput
	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId pulumi.StringPtrInput
	// The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
	// and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
	PeerRegion pulumi.StringPtrInput
	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId pulumi.StringInput
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester VpcPeeringConnectionRequesterPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the requester VPC.
	VpcId pulumi.StringInput
}

func (VpcPeeringConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPeeringConnectionArgs)(nil)).Elem()
}
