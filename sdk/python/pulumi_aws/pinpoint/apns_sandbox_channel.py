# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class ApnsSandboxChannel(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, application_id=None, bundle_id=None, certificate=None, default_authentication_method=None, enabled=None, private_key=None, team_id=None, token_key=None, token_key_id=None):
        """Create a ApnsSandboxChannel resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not application_id:
            raise TypeError('Missing required property application_id')
        elif not isinstance(application_id, basestring):
            raise TypeError('Expected property application_id to be a basestring')
        __self__.application_id = application_id
        __props__['applicationId'] = application_id

        if bundle_id and not isinstance(bundle_id, basestring):
            raise TypeError('Expected property bundle_id to be a basestring')
        __self__.bundle_id = bundle_id
        __props__['bundleId'] = bundle_id

        if certificate and not isinstance(certificate, basestring):
            raise TypeError('Expected property certificate to be a basestring')
        __self__.certificate = certificate
        __props__['certificate'] = certificate

        if default_authentication_method and not isinstance(default_authentication_method, basestring):
            raise TypeError('Expected property default_authentication_method to be a basestring')
        __self__.default_authentication_method = default_authentication_method
        __props__['defaultAuthenticationMethod'] = default_authentication_method

        if enabled and not isinstance(enabled, bool):
            raise TypeError('Expected property enabled to be a bool')
        __self__.enabled = enabled
        __props__['enabled'] = enabled

        if private_key and not isinstance(private_key, basestring):
            raise TypeError('Expected property private_key to be a basestring')
        __self__.private_key = private_key
        __props__['privateKey'] = private_key

        if team_id and not isinstance(team_id, basestring):
            raise TypeError('Expected property team_id to be a basestring')
        __self__.team_id = team_id
        __props__['teamId'] = team_id

        if token_key and not isinstance(token_key, basestring):
            raise TypeError('Expected property token_key to be a basestring')
        __self__.token_key = token_key
        __props__['tokenKey'] = token_key

        if token_key_id and not isinstance(token_key_id, basestring):
            raise TypeError('Expected property token_key_id to be a basestring')
        __self__.token_key_id = token_key_id
        __props__['tokenKeyId'] = token_key_id

        super(ApnsSandboxChannel, __self__).__init__(
            'aws:pinpoint/apnsSandboxChannel:ApnsSandboxChannel',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'applicationId' in outs:
            self.application_id = outs['applicationId']
        if 'bundleId' in outs:
            self.bundle_id = outs['bundleId']
        if 'certificate' in outs:
            self.certificate = outs['certificate']
        if 'defaultAuthenticationMethod' in outs:
            self.default_authentication_method = outs['defaultAuthenticationMethod']
        if 'enabled' in outs:
            self.enabled = outs['enabled']
        if 'privateKey' in outs:
            self.private_key = outs['privateKey']
        if 'teamId' in outs:
            self.team_id = outs['teamId']
        if 'tokenKey' in outs:
            self.token_key = outs['tokenKey']
        if 'tokenKeyId' in outs:
            self.token_key_id = outs['tokenKeyId']
