# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class IdentityNotificationTopic(pulumi.CustomResource):
    identity: pulumi.Output[str]
    """
    The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
    """
    include_original_headers: pulumi.Output[bool]
    """
    Whether SES should include original email headers in SNS notifications of this type. *false* by default.
    """
    notification_type: pulumi.Output[str]
    """
    The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
    """
    topic_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
    """
    def __init__(__self__, resource_name, opts=None, identity=None, include_original_headers=None, notification_type=None, topic_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Resource for managing SES Identity Notification Topics



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity: The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
        :param pulumi.Input[bool] include_original_headers: Whether SES should include original email headers in SNS notifications of this type. *false* by default.
        :param pulumi.Input[str] notification_type: The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
        :param pulumi.Input[str] topic_arn: The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
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

            if identity is None:
                raise TypeError("Missing required property 'identity'")
            __props__['identity'] = identity
            __props__['include_original_headers'] = include_original_headers
            if notification_type is None:
                raise TypeError("Missing required property 'notification_type'")
            __props__['notification_type'] = notification_type
            __props__['topic_arn'] = topic_arn
        super(IdentityNotificationTopic, __self__).__init__(
            'aws:ses/identityNotificationTopic:IdentityNotificationTopic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, identity=None, include_original_headers=None, notification_type=None, topic_arn=None):
        """
        Get an existing IdentityNotificationTopic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity: The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
        :param pulumi.Input[bool] include_original_headers: Whether SES should include original email headers in SNS notifications of this type. *false* by default.
        :param pulumi.Input[str] notification_type: The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
        :param pulumi.Input[str] topic_arn: The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["identity"] = identity
        __props__["include_original_headers"] = include_original_headers
        __props__["notification_type"] = notification_type
        __props__["topic_arn"] = topic_arn
        return IdentityNotificationTopic(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

