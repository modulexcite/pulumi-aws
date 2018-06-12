// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * `aws_vpc` provides details about a specific VPC.
 * 
 * This resource can prove useful when a module accepts a vpc id as
 * an input variable and needs to, for example, determine the CIDR block of that
 * VPC.
 */
export function getVpc(args?: GetVpcArgs): Promise<GetVpcResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:ec2/getVpc:getVpc", {
        "cidrBlock": args.cidrBlock,
        "default": args.default,
        "dhcpOptionsId": args.dhcpOptionsId,
        "filters": args.filters,
        "id": args.id,
        "state": args.state,
        "tags": args.tags,
    });
}

/**
 * A collection of arguments for invoking getVpc.
 */
export interface GetVpcArgs {
    /**
     * The cidr block of the desired VPC.
     */
    readonly cidrBlock?: pulumi.Input<string>;
    /**
     * Boolean constraint on whether the desired VPC is
     * the default VPC for the region.
     */
    readonly default?: pulumi.Input<boolean>;
    /**
     * The DHCP options id of the desired VPC.
     */
    readonly dhcpOptionsId?: pulumi.Input<string>;
    /**
     * Custom filter block as described below.
     */
    readonly filters?: pulumi.Input<{ name: pulumi.Input<string>, values: pulumi.Input<pulumi.Input<string>[]> }[]>;
    /**
     * The id of the specific VPC to retrieve.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The current state of the desired VPC.
     * Can be either `"pending"` or `"available"`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired VPC.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * A collection of values returned by getVpc.
 */
export interface GetVpcResult {
    readonly cidrBlock: string;
    readonly default: boolean;
    readonly dhcpOptionsId: string;
    /**
     * Whether or not the VPC has DNS hostname support
     */
    readonly enableDnsHostnames: boolean;
    /**
     * Whether or not the VPC has DNS support
     */
    readonly enableDnsSupport: boolean;
    readonly id: string;
    /**
     * The allowed tenancy of instances launched into the
     * selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
     */
    readonly instanceTenancy: string;
    /**
     * The association ID for the IPv6 CIDR block.
     */
    readonly ipv6AssociationId: string;
    /**
     * The IPv6 CIDR block.
     */
    readonly ipv6CidrBlock: string;
    readonly state: string;
    readonly tags: {[key: string]: any};
}