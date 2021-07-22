# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ... import meta as _meta

__all__ = [
    'AuditSinkSpecArgs',
    'AuditSinkArgs',
    'PolicyArgs',
    'ServiceReferenceArgs',
    'WebhookClientConfigArgs',
    'WebhookThrottleConfigArgs',
    'WebhookArgs',
]

@pulumi.input_type
class AuditSinkSpecArgs:
    def __init__(__self__, *,
                 policy: pulumi.Input['PolicyArgs'],
                 webhook: pulumi.Input['WebhookArgs']):
        """
        AuditSinkSpec holds the spec for the audit sink
        :param pulumi.Input['PolicyArgs'] policy: Policy defines the policy for selecting which events should be sent to the webhook required
        :param pulumi.Input['WebhookArgs'] webhook: Webhook to send events required
        """
        pulumi.set(__self__, "policy", policy)
        pulumi.set(__self__, "webhook", webhook)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input['PolicyArgs']:
        """
        Policy defines the policy for selecting which events should be sent to the webhook required
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input['PolicyArgs']):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def webhook(self) -> pulumi.Input['WebhookArgs']:
        """
        Webhook to send events required
        """
        return pulumi.get(self, "webhook")

    @webhook.setter
    def webhook(self, value: pulumi.Input['WebhookArgs']):
        pulumi.set(self, "webhook", value)


@pulumi.input_type
class AuditSinkArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['AuditSinkSpecArgs']] = None):
        """
        AuditSink represents a cluster level audit sink
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['AuditSinkSpecArgs'] spec: Spec defines the audit configuration spec
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'auditregistration.k8s.io/v1alpha1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'AuditSink')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['AuditSinkSpecArgs']]:
        """
        Spec defines the audit configuration spec
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['AuditSinkSpecArgs']]):
        pulumi.set(self, "spec", value)


@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 level: pulumi.Input[str],
                 stages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Policy defines the configuration of how audit events are logged
        :param pulumi.Input[str] level: The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required
        :param pulumi.Input[Sequence[pulumi.Input[str]]] stages: Stages is a list of stages for which events are created.
        """
        pulumi.set(__self__, "level", level)
        if stages is not None:
            pulumi.set(__self__, "stages", stages)

    @property
    @pulumi.getter
    def level(self) -> pulumi.Input[str]:
        """
        The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: pulumi.Input[str]):
        pulumi.set(self, "level", value)

    @property
    @pulumi.getter
    def stages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Stages is a list of stages for which events are created.
        """
        return pulumi.get(self, "stages")

    @stages.setter
    def stages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "stages", value)


@pulumi.input_type
class ServiceReferenceArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 namespace: pulumi.Input[str],
                 path: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        ServiceReference holds a reference to Service.legacy.k8s.io
        :param pulumi.Input[str] name: `name` is the name of the service. Required
        :param pulumi.Input[str] namespace: `namespace` is the namespace of the service. Required
        :param pulumi.Input[str] path: `path` is an optional URL path which will be sent in any request to this service.
        :param pulumi.Input[int] port: If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        `name` is the name of the service. Required
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        `namespace` is the namespace of the service. Required
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        `path` is an optional URL path which will be sent in any request to this service.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class WebhookClientConfigArgs:
    def __init__(__self__, *,
                 ca_bundle: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input['ServiceReferenceArgs']] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        WebhookClientConfig contains the information to make a connection with the webhook
        :param pulumi.Input[str] ca_bundle: `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
        :param pulumi.Input['ServiceReferenceArgs'] service: `service` is a reference to the service for this webhook. Either `service` or `url` must be specified.
               
               If the webhook is running within the cluster, then you should use `service`.
        :param pulumi.Input[str] url: `url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.
               
               The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.
               
               Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.
               
               The scheme must be "https"; the URL must begin with "https://".
               
               A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.
               
               Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
        """
        if ca_bundle is not None:
            pulumi.set(__self__, "ca_bundle", ca_bundle)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="caBundle")
    def ca_bundle(self) -> Optional[pulumi.Input[str]]:
        """
        `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
        """
        return pulumi.get(self, "ca_bundle")

    @ca_bundle.setter
    def ca_bundle(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_bundle", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input['ServiceReferenceArgs']]:
        """
        `service` is a reference to the service for this webhook. Either `service` or `url` must be specified.

        If the webhook is running within the cluster, then you should use `service`.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input['ServiceReferenceArgs']]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        `url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.

        The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.

        Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.

        The scheme must be "https"; the URL must begin with "https://".

        A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.

        Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class WebhookThrottleConfigArgs:
    def __init__(__self__, *,
                 burst: Optional[pulumi.Input[int]] = None,
                 qps: Optional[pulumi.Input[int]] = None):
        """
        WebhookThrottleConfig holds the configuration for throttling events
        :param pulumi.Input[int] burst: ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS
        :param pulumi.Input[int] qps: ThrottleQPS maximum number of batches per second default 10 QPS
        """
        if burst is not None:
            pulumi.set(__self__, "burst", burst)
        if qps is not None:
            pulumi.set(__self__, "qps", qps)

    @property
    @pulumi.getter
    def burst(self) -> Optional[pulumi.Input[int]]:
        """
        ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS
        """
        return pulumi.get(self, "burst")

    @burst.setter
    def burst(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "burst", value)

    @property
    @pulumi.getter
    def qps(self) -> Optional[pulumi.Input[int]]:
        """
        ThrottleQPS maximum number of batches per second default 10 QPS
        """
        return pulumi.get(self, "qps")

    @qps.setter
    def qps(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "qps", value)


@pulumi.input_type
class WebhookArgs:
    def __init__(__self__, *,
                 client_config: pulumi.Input['WebhookClientConfigArgs'],
                 throttle: Optional[pulumi.Input['WebhookThrottleConfigArgs']] = None):
        """
        Webhook holds the configuration of the webhook
        :param pulumi.Input['WebhookClientConfigArgs'] client_config: ClientConfig holds the connection parameters for the webhook required
        :param pulumi.Input['WebhookThrottleConfigArgs'] throttle: Throttle holds the options for throttling the webhook
        """
        pulumi.set(__self__, "client_config", client_config)
        if throttle is not None:
            pulumi.set(__self__, "throttle", throttle)

    @property
    @pulumi.getter(name="clientConfig")
    def client_config(self) -> pulumi.Input['WebhookClientConfigArgs']:
        """
        ClientConfig holds the connection parameters for the webhook required
        """
        return pulumi.get(self, "client_config")

    @client_config.setter
    def client_config(self, value: pulumi.Input['WebhookClientConfigArgs']):
        pulumi.set(self, "client_config", value)

    @property
    @pulumi.getter
    def throttle(self) -> Optional[pulumi.Input['WebhookThrottleConfigArgs']]:
        """
        Throttle holds the options for throttling the webhook
        """
        return pulumi.get(self, "throttle")

    @throttle.setter
    def throttle(self, value: Optional[pulumi.Input['WebhookThrottleConfigArgs']]):
        pulumi.set(self, "throttle", value)


