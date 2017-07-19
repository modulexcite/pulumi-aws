// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Key extends lumi.NamedResource implements KeyArgs {
    public /*out*/ readonly arn: string;
    public readonly deletionWindowInDays?: number;
    public readonly description: string;
    public readonly enableKeyRotation?: boolean;
    public readonly isEnabled?: boolean;
    public /*out*/ readonly keyId: string;
    public readonly keyUsage: string;
    public readonly policy: string;
    public readonly tags?: {[key: string]: any};

    constructor(name: string, args: KeyArgs) {
        super(name);
        this.deletionWindowInDays = args.deletionWindowInDays;
        this.description = args.description;
        this.enableKeyRotation = args.enableKeyRotation;
        this.isEnabled = args.isEnabled;
        this.keyUsage = args.keyUsage;
        this.policy = args.policy;
        this.tags = args.tags;
    }
}

export interface KeyArgs {
    readonly deletionWindowInDays?: number;
    readonly description?: string;
    readonly enableKeyRotation?: boolean;
    readonly isEnabled?: boolean;
    readonly keyUsage?: string;
    readonly policy?: string;
    readonly tags?: {[key: string]: any};
}
