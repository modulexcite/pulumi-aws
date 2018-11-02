// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Route53 Hosted Zone.
type Zone struct {
	s *pulumi.ResourceState
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOpt) (*Zone, error) {
	inputs := make(map[string]interface{})
	inputs["comment"] = "Managed by Pulumi"
	if args == nil {
		inputs["delegationSetId"] = nil
		inputs["forceDestroy"] = nil
		inputs["name"] = nil
		inputs["tags"] = nil
		inputs["vpcs"] = nil
		inputs["vpcId"] = nil
		inputs["vpcRegion"] = nil
	} else {
		inputs["comment"] = args.Comment
		inputs["delegationSetId"] = args.DelegationSetId
		inputs["forceDestroy"] = args.ForceDestroy
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
		inputs["vpcs"] = args.Vpcs
		inputs["vpcId"] = args.VpcId
		inputs["vpcRegion"] = args.VpcRegion
	}
	inputs["nameServers"] = nil
	inputs["zoneId"] = nil
	s, err := ctx.RegisterResource("aws:route53/zone:Zone", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Zone{s: s}, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ZoneState, opts ...pulumi.ResourceOpt) (*Zone, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["comment"] = state.Comment
		inputs["delegationSetId"] = state.DelegationSetId
		inputs["forceDestroy"] = state.ForceDestroy
		inputs["name"] = state.Name
		inputs["nameServers"] = state.NameServers
		inputs["tags"] = state.Tags
		inputs["vpcs"] = state.Vpcs
		inputs["vpcId"] = state.VpcId
		inputs["vpcRegion"] = state.VpcRegion
		inputs["zoneId"] = state.ZoneId
	}
	s, err := ctx.ReadResource("aws:route53/zone:Zone", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Zone{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Zone) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Zone) ID() *pulumi.IDOutput {
	return r.s.ID
}

// A comment for the hosted zone. Defaults to 'Managed by Terraform'.
func (r *Zone) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` and `vpc_id` as delegation sets can only be used for public zones.
func (r *Zone) DelegationSetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["delegationSetId"])
}

// Whether to destroy all records (possibly managed outside of Terraform) in the zone when destroying the zone.
func (r *Zone) ForceDestroy() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["forceDestroy"])
}

// This is the name of the hosted zone.
func (r *Zone) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A list of name servers in associated (or default) delegation set.
// Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
func (r *Zone) NameServers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["nameServers"])
}

// A mapping of tags to assign to the zone.
func (r *Zone) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with `delegation_set_id`, `vpc_id`, and `vpc_region` in this resource and any [`aws_route53_zone_association` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone_association.html) specifying the same zone ID. Detailed below.
func (r *Zone) Vpcs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["vpcs"])
}

// ID of the VPC to associate.
func (r *Zone) VpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcId"])
}

// Region of the VPC to associate. Defaults to AWS provider region.
func (r *Zone) VpcRegion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcRegion"])
}

// The Hosted Zone ID. This can be referenced by zone records.
func (r *Zone) ZoneId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zoneId"])
}

// Input properties used for looking up and filtering Zone resources.
type ZoneState struct {
	// A comment for the hosted zone. Defaults to 'Managed by Terraform'.
	Comment interface{}
	// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` and `vpc_id` as delegation sets can only be used for public zones.
	DelegationSetId interface{}
	// Whether to destroy all records (possibly managed outside of Terraform) in the zone when destroying the zone.
	ForceDestroy interface{}
	// This is the name of the hosted zone.
	Name interface{}
	// A list of name servers in associated (or default) delegation set.
	// Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
	NameServers interface{}
	// A mapping of tags to assign to the zone.
	Tags interface{}
	// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with `delegation_set_id`, `vpc_id`, and `vpc_region` in this resource and any [`aws_route53_zone_association` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone_association.html) specifying the same zone ID. Detailed below.
	Vpcs interface{}
	// ID of the VPC to associate.
	VpcId interface{}
	// Region of the VPC to associate. Defaults to AWS provider region.
	VpcRegion interface{}
	// The Hosted Zone ID. This can be referenced by zone records.
	ZoneId interface{}
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// A comment for the hosted zone. Defaults to 'Managed by Terraform'.
	Comment interface{}
	// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` and `vpc_id` as delegation sets can only be used for public zones.
	DelegationSetId interface{}
	// Whether to destroy all records (possibly managed outside of Terraform) in the zone when destroying the zone.
	ForceDestroy interface{}
	// This is the name of the hosted zone.
	Name interface{}
	// A mapping of tags to assign to the zone.
	Tags interface{}
	// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with `delegation_set_id`, `vpc_id`, and `vpc_region` in this resource and any [`aws_route53_zone_association` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone_association.html) specifying the same zone ID. Detailed below.
	Vpcs interface{}
	// ID of the VPC to associate.
	VpcId interface{}
	// Region of the VPC to associate. Defaults to AWS provider region.
	VpcRegion interface{}
}
