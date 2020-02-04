# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class EventDestination(pulumi.CustomResource):
    cloudwatch_destinations: pulumi.Output[list]
    """
    CloudWatch destination for the events

      * `default_value` (`str`) - The default value for the event
      * `dimensionName` (`str`) - The name for the dimension
      * `valueSource` (`str`) - The source for the value. It can be either `"messageTag"` or `"emailHeader"`
    """
    configuration_set_name: pulumi.Output[str]
    """
    The name of the configuration set
    """
    enabled: pulumi.Output[bool]
    """
    If true, the event destination will be enabled
    """
    kinesis_destination: pulumi.Output[dict]
    """
    Send the events to a kinesis firehose destination

      * `role_arn` (`str`) - The ARN of the role that has permissions to access the Kinesis Stream
      * `streamArn` (`str`) - The ARN of the Kinesis Stream
    """
    matching_types: pulumi.Output[list]
    """
    A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
    """
    name: pulumi.Output[str]
    """
    The name of the event destination
    """
    sns_destination: pulumi.Output[dict]
    """
    Send the events to an SNS Topic destination

      * `topic_arn` (`str`) - The ARN of the SNS topic
    """
    def __init__(__self__, resource_name, opts=None, cloudwatch_destinations=None, configuration_set_name=None, enabled=None, kinesis_destination=None, matching_types=None, name=None, sns_destination=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an SES event destination

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_event_destination.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] cloudwatch_destinations: CloudWatch destination for the events
        :param pulumi.Input[str] configuration_set_name: The name of the configuration set
        :param pulumi.Input[bool] enabled: If true, the event destination will be enabled
        :param pulumi.Input[dict] kinesis_destination: Send the events to a kinesis firehose destination
        :param pulumi.Input[list] matching_types: A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
        :param pulumi.Input[str] name: The name of the event destination
        :param pulumi.Input[dict] sns_destination: Send the events to an SNS Topic destination

        The **cloudwatch_destinations** object supports the following:

          * `default_value` (`pulumi.Input[str]`) - The default value for the event
          * `dimensionName` (`pulumi.Input[str]`) - The name for the dimension
          * `valueSource` (`pulumi.Input[str]`) - The source for the value. It can be either `"messageTag"` or `"emailHeader"`

        The **kinesis_destination** object supports the following:

          * `role_arn` (`pulumi.Input[str]`) - The ARN of the role that has permissions to access the Kinesis Stream
          * `streamArn` (`pulumi.Input[str]`) - The ARN of the Kinesis Stream

        The **sns_destination** object supports the following:

          * `topic_arn` (`pulumi.Input[str]`) - The ARN of the SNS topic
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['cloudwatch_destinations'] = cloudwatch_destinations
            if configuration_set_name is None:
                raise TypeError("Missing required property 'configuration_set_name'")
            __props__['configuration_set_name'] = configuration_set_name
            __props__['enabled'] = enabled
            __props__['kinesis_destination'] = kinesis_destination
            if matching_types is None:
                raise TypeError("Missing required property 'matching_types'")
            __props__['matching_types'] = matching_types
            __props__['name'] = name
            __props__['sns_destination'] = sns_destination
        super(EventDestination, __self__).__init__(
            'aws:ses/eventDestination:EventDestination',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cloudwatch_destinations=None, configuration_set_name=None, enabled=None, kinesis_destination=None, matching_types=None, name=None, sns_destination=None):
        """
        Get an existing EventDestination resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] cloudwatch_destinations: CloudWatch destination for the events
        :param pulumi.Input[str] configuration_set_name: The name of the configuration set
        :param pulumi.Input[bool] enabled: If true, the event destination will be enabled
        :param pulumi.Input[dict] kinesis_destination: Send the events to a kinesis firehose destination
        :param pulumi.Input[list] matching_types: A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
        :param pulumi.Input[str] name: The name of the event destination
        :param pulumi.Input[dict] sns_destination: Send the events to an SNS Topic destination

        The **cloudwatch_destinations** object supports the following:

          * `default_value` (`pulumi.Input[str]`) - The default value for the event
          * `dimensionName` (`pulumi.Input[str]`) - The name for the dimension
          * `valueSource` (`pulumi.Input[str]`) - The source for the value. It can be either `"messageTag"` or `"emailHeader"`

        The **kinesis_destination** object supports the following:

          * `role_arn` (`pulumi.Input[str]`) - The ARN of the role that has permissions to access the Kinesis Stream
          * `streamArn` (`pulumi.Input[str]`) - The ARN of the Kinesis Stream

        The **sns_destination** object supports the following:

          * `topic_arn` (`pulumi.Input[str]`) - The ARN of the SNS topic
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cloudwatch_destinations"] = cloudwatch_destinations
        __props__["configuration_set_name"] = configuration_set_name
        __props__["enabled"] = enabled
        __props__["kinesis_destination"] = kinesis_destination
        __props__["matching_types"] = matching_types
        __props__["name"] = name
        __props__["sns_destination"] = sns_destination
        return EventDestination(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

