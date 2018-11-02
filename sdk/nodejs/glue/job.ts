// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Glue Job resource.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState): Job {
        return new Job(name, <any>state, { id });
    }

    /**
     * The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
     */
    public readonly allocatedCapacity: pulumi.Output<number | undefined>;
    /**
     * The command of the job. Defined below.
     */
    public readonly command: pulumi.Output<{ name?: string, scriptLocation: string }>;
    /**
     * The list of connections used for this job.
     */
    public readonly connections: pulumi.Output<string[] | undefined>;
    /**
     * The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
     */
    public readonly defaultArguments: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Description of the job.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * Execution property of the job. Defined below.
     */
    public readonly executionProperty: pulumi.Output<{ maxConcurrentRuns?: number }>;
    /**
     * The maximum number of times to retry this job if it fails.
     */
    public readonly maxRetries: pulumi.Output<number | undefined>;
    /**
     * The name of the job command. Defaults to `glueetl`
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The ARN of the IAM role associated with this job.
     */
    public readonly roleArn: pulumi.Output<string>;
    /**
     * The name of the Security Configuration to be associated with the job. 
     */
    public readonly securityConfiguration: pulumi.Output<string | undefined>;
    /**
     * The job timeout in minutes. The default is 2880 minutes (48 hours).
     */
    public readonly timeout: pulumi.Output<number | undefined>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: JobState = argsOrState as JobState | undefined;
            inputs["allocatedCapacity"] = state ? state.allocatedCapacity : undefined;
            inputs["command"] = state ? state.command : undefined;
            inputs["connections"] = state ? state.connections : undefined;
            inputs["defaultArguments"] = state ? state.defaultArguments : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["executionProperty"] = state ? state.executionProperty : undefined;
            inputs["maxRetries"] = state ? state.maxRetries : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["securityConfiguration"] = state ? state.securityConfiguration : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if (!args || args.command === undefined) {
                throw new Error("Missing required property 'command'");
            }
            if (!args || args.roleArn === undefined) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["allocatedCapacity"] = args ? args.allocatedCapacity : undefined;
            inputs["command"] = args ? args.command : undefined;
            inputs["connections"] = args ? args.connections : undefined;
            inputs["defaultArguments"] = args ? args.defaultArguments : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["executionProperty"] = args ? args.executionProperty : undefined;
            inputs["maxRetries"] = args ? args.maxRetries : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["securityConfiguration"] = args ? args.securityConfiguration : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
        }
        super("aws:glue/job:Job", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
     */
    readonly allocatedCapacity?: pulumi.Input<number>;
    /**
     * The command of the job. Defined below.
     */
    readonly command?: pulumi.Input<{ name?: pulumi.Input<string>, scriptLocation: pulumi.Input<string> }>;
    /**
     * The list of connections used for this job.
     */
    readonly connections?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
     */
    readonly defaultArguments?: pulumi.Input<{[key: string]: any}>;
    /**
     * Description of the job.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Execution property of the job. Defined below.
     */
    readonly executionProperty?: pulumi.Input<{ maxConcurrentRuns?: pulumi.Input<number> }>;
    /**
     * The maximum number of times to retry this job if it fails.
     */
    readonly maxRetries?: pulumi.Input<number>;
    /**
     * The name of the job command. Defaults to `glueetl`
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN of the IAM role associated with this job.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The name of the Security Configuration to be associated with the job. 
     */
    readonly securityConfiguration?: pulumi.Input<string>;
    /**
     * The job timeout in minutes. The default is 2880 minutes (48 hours).
     */
    readonly timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
     */
    readonly allocatedCapacity?: pulumi.Input<number>;
    /**
     * The command of the job. Defined below.
     */
    readonly command: pulumi.Input<{ name?: pulumi.Input<string>, scriptLocation: pulumi.Input<string> }>;
    /**
     * The list of connections used for this job.
     */
    readonly connections?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
     */
    readonly defaultArguments?: pulumi.Input<{[key: string]: any}>;
    /**
     * Description of the job.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Execution property of the job. Defined below.
     */
    readonly executionProperty?: pulumi.Input<{ maxConcurrentRuns?: pulumi.Input<number> }>;
    /**
     * The maximum number of times to retry this job if it fails.
     */
    readonly maxRetries?: pulumi.Input<number>;
    /**
     * The name of the job command. Defaults to `glueetl`
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN of the IAM role associated with this job.
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * The name of the Security Configuration to be associated with the job. 
     */
    readonly securityConfiguration?: pulumi.Input<string>;
    /**
     * The job timeout in minutes. The default is 2880 minutes (48 hours).
     */
    readonly timeout?: pulumi.Input<number>;
}
