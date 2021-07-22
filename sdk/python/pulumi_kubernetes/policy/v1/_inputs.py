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
    'PodDisruptionBudgetSpecArgs',
    'PodDisruptionBudgetStatusArgs',
    'PodDisruptionBudgetArgs',
]

@pulumi.input_type
class PodDisruptionBudgetSpecArgs:
    def __init__(__self__, *,
                 max_unavailable: Optional[pulumi.Input[Union[int, str]]] = None,
                 min_available: Optional[pulumi.Input[Union[int, str]]] = None,
                 selector: Optional[pulumi.Input['_meta.v1.LabelSelectorArgs']] = None):
        """
        PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
        :param pulumi.Input[Union[int, str]] max_unavailable: An eviction is allowed if at most "maxUnavailable" pods selected by "selector" are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with "minAvailable".
        :param pulumi.Input[Union[int, str]] min_available: An eviction is allowed if at least "minAvailable" pods selected by "selector" will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying "100%".
        :param pulumi.Input['_meta.v1.LabelSelectorArgs'] selector: Label query over pods whose evictions are managed by the disruption budget. A null selector will match no pods, while an empty ({}) selector will select all pods within the namespace.
        """
        if max_unavailable is not None:
            pulumi.set(__self__, "max_unavailable", max_unavailable)
        if min_available is not None:
            pulumi.set(__self__, "min_available", min_available)
        if selector is not None:
            pulumi.set(__self__, "selector", selector)

    @property
    @pulumi.getter(name="maxUnavailable")
    def max_unavailable(self) -> Optional[pulumi.Input[Union[int, str]]]:
        """
        An eviction is allowed if at most "maxUnavailable" pods selected by "selector" are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with "minAvailable".
        """
        return pulumi.get(self, "max_unavailable")

    @max_unavailable.setter
    def max_unavailable(self, value: Optional[pulumi.Input[Union[int, str]]]):
        pulumi.set(self, "max_unavailable", value)

    @property
    @pulumi.getter(name="minAvailable")
    def min_available(self) -> Optional[pulumi.Input[Union[int, str]]]:
        """
        An eviction is allowed if at least "minAvailable" pods selected by "selector" will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying "100%".
        """
        return pulumi.get(self, "min_available")

    @min_available.setter
    def min_available(self, value: Optional[pulumi.Input[Union[int, str]]]):
        pulumi.set(self, "min_available", value)

    @property
    @pulumi.getter
    def selector(self) -> Optional[pulumi.Input['_meta.v1.LabelSelectorArgs']]:
        """
        Label query over pods whose evictions are managed by the disruption budget. A null selector will match no pods, while an empty ({}) selector will select all pods within the namespace.
        """
        return pulumi.get(self, "selector")

    @selector.setter
    def selector(self, value: Optional[pulumi.Input['_meta.v1.LabelSelectorArgs']]):
        pulumi.set(self, "selector", value)


@pulumi.input_type
class PodDisruptionBudgetStatusArgs:
    def __init__(__self__, *,
                 current_healthy: pulumi.Input[int],
                 desired_healthy: pulumi.Input[int],
                 disruptions_allowed: pulumi.Input[int],
                 expected_pods: pulumi.Input[int],
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['_meta.v1.ConditionArgs']]]] = None,
                 disrupted_pods: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 observed_generation: Optional[pulumi.Input[int]] = None):
        """
        PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.
        :param pulumi.Input[int] current_healthy: current number of healthy pods
        :param pulumi.Input[int] desired_healthy: minimum desired number of healthy pods
        :param pulumi.Input[int] disruptions_allowed: Number of pod disruptions that are currently allowed.
        :param pulumi.Input[int] expected_pods: total number of pods counted by this disruption budget
        :param pulumi.Input[Sequence[pulumi.Input['_meta.v1.ConditionArgs']]] conditions: Conditions contain conditions for PDB. The disruption controller sets the DisruptionAllowed condition. The following are known values for the reason field (additional reasons could be added in the future): - SyncFailed: The controller encountered an error and wasn't able to compute
                             the number of allowed disruptions. Therefore no disruptions are
                             allowed and the status of the condition will be False.
               - InsufficientPods: The number of pods are either at or below the number
                                   required by the PodDisruptionBudget. No disruptions are
                                   allowed and the status of the condition will be False.
               - SufficientPods: There are more pods than required by the PodDisruptionBudget.
                                 The condition will be True, and the number of allowed
                                 disruptions are provided by the disruptionsAllowed property.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] disrupted_pods: DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
        :param pulumi.Input[int] observed_generation: Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid only if observedGeneration equals to PDB's object generation.
        """
        pulumi.set(__self__, "current_healthy", current_healthy)
        pulumi.set(__self__, "desired_healthy", desired_healthy)
        pulumi.set(__self__, "disruptions_allowed", disruptions_allowed)
        pulumi.set(__self__, "expected_pods", expected_pods)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if disrupted_pods is not None:
            pulumi.set(__self__, "disrupted_pods", disrupted_pods)
        if observed_generation is not None:
            pulumi.set(__self__, "observed_generation", observed_generation)

    @property
    @pulumi.getter(name="currentHealthy")
    def current_healthy(self) -> pulumi.Input[int]:
        """
        current number of healthy pods
        """
        return pulumi.get(self, "current_healthy")

    @current_healthy.setter
    def current_healthy(self, value: pulumi.Input[int]):
        pulumi.set(self, "current_healthy", value)

    @property
    @pulumi.getter(name="desiredHealthy")
    def desired_healthy(self) -> pulumi.Input[int]:
        """
        minimum desired number of healthy pods
        """
        return pulumi.get(self, "desired_healthy")

    @desired_healthy.setter
    def desired_healthy(self, value: pulumi.Input[int]):
        pulumi.set(self, "desired_healthy", value)

    @property
    @pulumi.getter(name="disruptionsAllowed")
    def disruptions_allowed(self) -> pulumi.Input[int]:
        """
        Number of pod disruptions that are currently allowed.
        """
        return pulumi.get(self, "disruptions_allowed")

    @disruptions_allowed.setter
    def disruptions_allowed(self, value: pulumi.Input[int]):
        pulumi.set(self, "disruptions_allowed", value)

    @property
    @pulumi.getter(name="expectedPods")
    def expected_pods(self) -> pulumi.Input[int]:
        """
        total number of pods counted by this disruption budget
        """
        return pulumi.get(self, "expected_pods")

    @expected_pods.setter
    def expected_pods(self, value: pulumi.Input[int]):
        pulumi.set(self, "expected_pods", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_meta.v1.ConditionArgs']]]]:
        """
        Conditions contain conditions for PDB. The disruption controller sets the DisruptionAllowed condition. The following are known values for the reason field (additional reasons could be added in the future): - SyncFailed: The controller encountered an error and wasn't able to compute
                      the number of allowed disruptions. Therefore no disruptions are
                      allowed and the status of the condition will be False.
        - InsufficientPods: The number of pods are either at or below the number
                            required by the PodDisruptionBudget. No disruptions are
                            allowed and the status of the condition will be False.
        - SufficientPods: There are more pods than required by the PodDisruptionBudget.
                          The condition will be True, and the number of allowed
                          disruptions are provided by the disruptionsAllowed property.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_meta.v1.ConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="disruptedPods")
    def disrupted_pods(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
        """
        return pulumi.get(self, "disrupted_pods")

    @disrupted_pods.setter
    def disrupted_pods(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "disrupted_pods", value)

    @property
    @pulumi.getter(name="observedGeneration")
    def observed_generation(self) -> Optional[pulumi.Input[int]]:
        """
        Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid only if observedGeneration equals to PDB's object generation.
        """
        return pulumi.get(self, "observed_generation")

    @observed_generation.setter
    def observed_generation(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "observed_generation", value)


@pulumi.input_type
class PodDisruptionBudgetArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['PodDisruptionBudgetSpecArgs']] = None,
                 status: Optional[pulumi.Input['PodDisruptionBudgetStatusArgs']] = None):
        """
        PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['PodDisruptionBudgetSpecArgs'] spec: Specification of the desired behavior of the PodDisruptionBudget.
        :param pulumi.Input['PodDisruptionBudgetStatusArgs'] status: Most recently observed status of the PodDisruptionBudget.
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'policy/v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'PodDisruptionBudget')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)

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
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['PodDisruptionBudgetSpecArgs']]:
        """
        Specification of the desired behavior of the PodDisruptionBudget.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['PodDisruptionBudgetSpecArgs']]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['PodDisruptionBudgetStatusArgs']]:
        """
        Most recently observed status of the PodDisruptionBudget.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['PodDisruptionBudgetStatusArgs']]):
        pulumi.set(self, "status", value)


