// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRole, and will no longer be served in v1.20.
type ClusterRolePatch struct {
	pulumi.CustomResourceState

	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule AggregationRulePatchPtrOutput `pulumi:"aggregationRule"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRulePatchArrayOutput `pulumi:"rules"`
}

// NewClusterRolePatch registers a new resource with the given unique name, arguments, and options.
func NewClusterRolePatch(ctx *pulumi.Context,
	name string, args *ClusterRolePatchArgs, opts ...pulumi.ResourceOption) (*ClusterRolePatch, error) {
	if args == nil {
		args = &ClusterRolePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ClusterRole")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1:ClusterRolePatch"),
		},
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRolePatch"),
		},
	})
	opts = append(opts, aliases)
	var resource ClusterRolePatch
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRolePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRolePatch gets an existing ClusterRolePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterRolePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterRolePatchState, opts ...pulumi.ResourceOption) (*ClusterRolePatch, error) {
	var resource ClusterRolePatch
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRolePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterRolePatch resources.
type clusterRolePatchState struct {
}

type ClusterRolePatchState struct {
}

func (ClusterRolePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRolePatchState)(nil)).Elem()
}

type clusterRolePatchArgs struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule *AggregationRulePatch `pulumi:"aggregationRule"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Rules holds all the PolicyRules for this ClusterRole
	Rules []PolicyRulePatch `pulumi:"rules"`
}

// The set of arguments for constructing a ClusterRolePatch resource.
type ClusterRolePatchArgs struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule AggregationRulePatchPtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchPtrInput
	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRulePatchArrayInput
}

func (ClusterRolePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRolePatchArgs)(nil)).Elem()
}

type ClusterRolePatchInput interface {
	pulumi.Input

	ToClusterRolePatchOutput() ClusterRolePatchOutput
	ToClusterRolePatchOutputWithContext(ctx context.Context) ClusterRolePatchOutput
}

func (*ClusterRolePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRolePatch)(nil)).Elem()
}

func (i *ClusterRolePatch) ToClusterRolePatchOutput() ClusterRolePatchOutput {
	return i.ToClusterRolePatchOutputWithContext(context.Background())
}

func (i *ClusterRolePatch) ToClusterRolePatchOutputWithContext(ctx context.Context) ClusterRolePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRolePatchOutput)
}

// ClusterRolePatchArrayInput is an input type that accepts ClusterRolePatchArray and ClusterRolePatchArrayOutput values.
// You can construct a concrete instance of `ClusterRolePatchArrayInput` via:
//
//	ClusterRolePatchArray{ ClusterRolePatchArgs{...} }
type ClusterRolePatchArrayInput interface {
	pulumi.Input

	ToClusterRolePatchArrayOutput() ClusterRolePatchArrayOutput
	ToClusterRolePatchArrayOutputWithContext(context.Context) ClusterRolePatchArrayOutput
}

type ClusterRolePatchArray []ClusterRolePatchInput

func (ClusterRolePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterRolePatch)(nil)).Elem()
}

func (i ClusterRolePatchArray) ToClusterRolePatchArrayOutput() ClusterRolePatchArrayOutput {
	return i.ToClusterRolePatchArrayOutputWithContext(context.Background())
}

func (i ClusterRolePatchArray) ToClusterRolePatchArrayOutputWithContext(ctx context.Context) ClusterRolePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRolePatchArrayOutput)
}

// ClusterRolePatchMapInput is an input type that accepts ClusterRolePatchMap and ClusterRolePatchMapOutput values.
// You can construct a concrete instance of `ClusterRolePatchMapInput` via:
//
//	ClusterRolePatchMap{ "key": ClusterRolePatchArgs{...} }
type ClusterRolePatchMapInput interface {
	pulumi.Input

	ToClusterRolePatchMapOutput() ClusterRolePatchMapOutput
	ToClusterRolePatchMapOutputWithContext(context.Context) ClusterRolePatchMapOutput
}

type ClusterRolePatchMap map[string]ClusterRolePatchInput

func (ClusterRolePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterRolePatch)(nil)).Elem()
}

func (i ClusterRolePatchMap) ToClusterRolePatchMapOutput() ClusterRolePatchMapOutput {
	return i.ToClusterRolePatchMapOutputWithContext(context.Background())
}

func (i ClusterRolePatchMap) ToClusterRolePatchMapOutputWithContext(ctx context.Context) ClusterRolePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRolePatchMapOutput)
}

type ClusterRolePatchOutput struct{ *pulumi.OutputState }

func (ClusterRolePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRolePatch)(nil)).Elem()
}

func (o ClusterRolePatchOutput) ToClusterRolePatchOutput() ClusterRolePatchOutput {
	return o
}

func (o ClusterRolePatchOutput) ToClusterRolePatchOutputWithContext(ctx context.Context) ClusterRolePatchOutput {
	return o
}

// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
func (o ClusterRolePatchOutput) AggregationRule() AggregationRulePatchPtrOutput {
	return o.ApplyT(func(v *ClusterRolePatch) AggregationRulePatchPtrOutput { return v.AggregationRule }).(AggregationRulePatchPtrOutput)
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ClusterRolePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterRolePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterRolePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterRolePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata.
func (o ClusterRolePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ClusterRolePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Rules holds all the PolicyRules for this ClusterRole
func (o ClusterRolePatchOutput) Rules() PolicyRulePatchArrayOutput {
	return o.ApplyT(func(v *ClusterRolePatch) PolicyRulePatchArrayOutput { return v.Rules }).(PolicyRulePatchArrayOutput)
}

type ClusterRolePatchArrayOutput struct{ *pulumi.OutputState }

func (ClusterRolePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterRolePatch)(nil)).Elem()
}

func (o ClusterRolePatchArrayOutput) ToClusterRolePatchArrayOutput() ClusterRolePatchArrayOutput {
	return o
}

func (o ClusterRolePatchArrayOutput) ToClusterRolePatchArrayOutputWithContext(ctx context.Context) ClusterRolePatchArrayOutput {
	return o
}

func (o ClusterRolePatchArrayOutput) Index(i pulumi.IntInput) ClusterRolePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterRolePatch {
		return vs[0].([]*ClusterRolePatch)[vs[1].(int)]
	}).(ClusterRolePatchOutput)
}

type ClusterRolePatchMapOutput struct{ *pulumi.OutputState }

func (ClusterRolePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterRolePatch)(nil)).Elem()
}

func (o ClusterRolePatchMapOutput) ToClusterRolePatchMapOutput() ClusterRolePatchMapOutput {
	return o
}

func (o ClusterRolePatchMapOutput) ToClusterRolePatchMapOutputWithContext(ctx context.Context) ClusterRolePatchMapOutput {
	return o
}

func (o ClusterRolePatchMapOutput) MapIndex(k pulumi.StringInput) ClusterRolePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterRolePatch {
		return vs[0].(map[string]*ClusterRolePatch)[vs[1].(string)]
	}).(ClusterRolePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRolePatchInput)(nil)).Elem(), &ClusterRolePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRolePatchArrayInput)(nil)).Elem(), ClusterRolePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRolePatchMapInput)(nil)).Elem(), ClusterRolePatchMap{})
	pulumi.RegisterOutputType(ClusterRolePatchOutput{})
	pulumi.RegisterOutputType(ClusterRolePatchArrayOutput{})
	pulumi.RegisterOutputType(ClusterRolePatchMapOutput{})
}
