# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class S3Location(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the DataSync Location.
    """
    s3_bucket_arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the S3 Bucket.
    """
    s3_config: pulumi.Output[dict]
    """
    Configuration block containing information for connecting to S3.

      * `bucketAccessRoleArn` (`str`) - Amazon Resource Names (ARN) of the IAM Role used to connect to the S3 Bucket.
    """
    subdirectory: pulumi.Output[str]
    """
    Prefix to perform actions as source or destination.
    """
    tags: pulumi.Output[dict]
    """
    Key-value pairs of resource tags to assign to the DataSync Location.
    """
    uri: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, s3_bucket_arn=None, s3_config=None, subdirectory=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an S3 Location within AWS DataSync.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/datasync_location_s3.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] s3_bucket_arn: Amazon Resource Name (ARN) of the S3 Bucket.
        :param pulumi.Input[dict] s3_config: Configuration block containing information for connecting to S3.
        :param pulumi.Input[str] subdirectory: Prefix to perform actions as source or destination.
        :param pulumi.Input[dict] tags: Key-value pairs of resource tags to assign to the DataSync Location.

        The **s3_config** object supports the following:

          * `bucketAccessRoleArn` (`pulumi.Input[str]`) - Amazon Resource Names (ARN) of the IAM Role used to connect to the S3 Bucket.
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

            if s3_bucket_arn is None:
                raise TypeError("Missing required property 's3_bucket_arn'")
            __props__['s3_bucket_arn'] = s3_bucket_arn
            if s3_config is None:
                raise TypeError("Missing required property 's3_config'")
            __props__['s3_config'] = s3_config
            if subdirectory is None:
                raise TypeError("Missing required property 'subdirectory'")
            __props__['subdirectory'] = subdirectory
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['uri'] = None
        super(S3Location, __self__).__init__(
            'aws:datasync/s3Location:S3Location',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, s3_bucket_arn=None, s3_config=None, subdirectory=None, tags=None, uri=None):
        """
        Get an existing S3Location resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the DataSync Location.
        :param pulumi.Input[str] s3_bucket_arn: Amazon Resource Name (ARN) of the S3 Bucket.
        :param pulumi.Input[dict] s3_config: Configuration block containing information for connecting to S3.
        :param pulumi.Input[str] subdirectory: Prefix to perform actions as source or destination.
        :param pulumi.Input[dict] tags: Key-value pairs of resource tags to assign to the DataSync Location.

        The **s3_config** object supports the following:

          * `bucketAccessRoleArn` (`pulumi.Input[str]`) - Amazon Resource Names (ARN) of the IAM Role used to connect to the S3 Bucket.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["s3_bucket_arn"] = s3_bucket_arn
        __props__["s3_config"] = s3_config
        __props__["subdirectory"] = subdirectory
        __props__["tags"] = tags
        __props__["uri"] = uri
        return S3Location(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

