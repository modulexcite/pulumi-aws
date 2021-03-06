// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **Note:** There is only a single subscription allowed per account.
//
// To help you understand the charges for your Spot instances, Amazon EC2 provides a data feed that describes your Spot instance usage and pricing.
// This data feed is sent to an Amazon S3 bucket that you specify when you subscribe to the data feed.
type SpotDatafeedSubscription struct {
	pulumi.CustomResourceState

	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
}

// NewSpotDatafeedSubscription registers a new resource with the given unique name, arguments, and options.
func NewSpotDatafeedSubscription(ctx *pulumi.Context,
	name string, args *SpotDatafeedSubscriptionArgs, opts ...pulumi.ResourceOption) (*SpotDatafeedSubscription, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil {
		args = &SpotDatafeedSubscriptionArgs{}
	}
	var resource SpotDatafeedSubscription
	err := ctx.RegisterResource("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpotDatafeedSubscription gets an existing SpotDatafeedSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpotDatafeedSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpotDatafeedSubscriptionState, opts ...pulumi.ResourceOption) (*SpotDatafeedSubscription, error) {
	var resource SpotDatafeedSubscription
	err := ctx.ReadResource("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpotDatafeedSubscription resources.
type spotDatafeedSubscriptionState struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket *string `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix *string `pulumi:"prefix"`
}

type SpotDatafeedSubscriptionState struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringPtrInput
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringPtrInput
}

func (SpotDatafeedSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*spotDatafeedSubscriptionState)(nil)).Elem()
}

type spotDatafeedSubscriptionArgs struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket string `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix *string `pulumi:"prefix"`
}

// The set of arguments for constructing a SpotDatafeedSubscription resource.
type SpotDatafeedSubscriptionArgs struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringInput
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringPtrInput
}

func (SpotDatafeedSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spotDatafeedSubscriptionArgs)(nil)).Elem()
}
