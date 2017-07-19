// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Activity extends lumi.NamedResource implements ActivityArgs {
    public /*out*/ readonly creationDate: string;
    public readonly activityName?: string;

    constructor(name: string, args: ActivityArgs) {
        super(name);
        this.activityName = args.activityName;
    }
}

export interface ActivityArgs {
    readonly activityName?: string;
}
