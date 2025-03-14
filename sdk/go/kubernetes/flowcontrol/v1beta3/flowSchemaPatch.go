// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta3

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
// FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a "flow distinguisher".
type FlowSchemaPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec FlowSchemaSpecPatchPtrOutput `pulumi:"spec"`
	// `status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status FlowSchemaStatusPatchPtrOutput `pulumi:"status"`
}

// NewFlowSchemaPatch registers a new resource with the given unique name, arguments, and options.
func NewFlowSchemaPatch(ctx *pulumi.Context,
	name string, args *FlowSchemaPatchArgs, opts ...pulumi.ResourceOption) (*FlowSchemaPatch, error) {
	if args == nil {
		args = &FlowSchemaPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("flowcontrol.apiserver.k8s.io/v1beta3")
	args.Kind = pulumi.StringPtr("FlowSchema")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchemaPatch"),
		},
		{
			Type: pulumi.String("kubernetes:flowcontrol.apiserver.k8s.io/v1beta1:FlowSchemaPatch"),
		},
		{
			Type: pulumi.String("kubernetes:flowcontrol.apiserver.k8s.io/v1beta2:FlowSchemaPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource FlowSchemaPatch
	err := ctx.RegisterResource("kubernetes:flowcontrol.apiserver.k8s.io/v1beta3:FlowSchemaPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowSchemaPatch gets an existing FlowSchemaPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowSchemaPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowSchemaPatchState, opts ...pulumi.ResourceOption) (*FlowSchemaPatch, error) {
	var resource FlowSchemaPatch
	err := ctx.ReadResource("kubernetes:flowcontrol.apiserver.k8s.io/v1beta3:FlowSchemaPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowSchemaPatch resources.
type flowSchemaPatchState struct {
}

type FlowSchemaPatchState struct {
}

func (FlowSchemaPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSchemaPatchState)(nil)).Elem()
}

type flowSchemaPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *FlowSchemaSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a FlowSchemaPatch resource.
type FlowSchemaPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec FlowSchemaSpecPatchPtrInput
}

func (FlowSchemaPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSchemaPatchArgs)(nil)).Elem()
}

type FlowSchemaPatchInput interface {
	pulumi.Input

	ToFlowSchemaPatchOutput() FlowSchemaPatchOutput
	ToFlowSchemaPatchOutputWithContext(ctx context.Context) FlowSchemaPatchOutput
}

func (*FlowSchemaPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowSchemaPatch)(nil)).Elem()
}

func (i *FlowSchemaPatch) ToFlowSchemaPatchOutput() FlowSchemaPatchOutput {
	return i.ToFlowSchemaPatchOutputWithContext(context.Background())
}

func (i *FlowSchemaPatch) ToFlowSchemaPatchOutputWithContext(ctx context.Context) FlowSchemaPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaPatchOutput)
}

// FlowSchemaPatchArrayInput is an input type that accepts FlowSchemaPatchArray and FlowSchemaPatchArrayOutput values.
// You can construct a concrete instance of `FlowSchemaPatchArrayInput` via:
//
//	FlowSchemaPatchArray{ FlowSchemaPatchArgs{...} }
type FlowSchemaPatchArrayInput interface {
	pulumi.Input

	ToFlowSchemaPatchArrayOutput() FlowSchemaPatchArrayOutput
	ToFlowSchemaPatchArrayOutputWithContext(context.Context) FlowSchemaPatchArrayOutput
}

type FlowSchemaPatchArray []FlowSchemaPatchInput

func (FlowSchemaPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowSchemaPatch)(nil)).Elem()
}

func (i FlowSchemaPatchArray) ToFlowSchemaPatchArrayOutput() FlowSchemaPatchArrayOutput {
	return i.ToFlowSchemaPatchArrayOutputWithContext(context.Background())
}

func (i FlowSchemaPatchArray) ToFlowSchemaPatchArrayOutputWithContext(ctx context.Context) FlowSchemaPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaPatchArrayOutput)
}

// FlowSchemaPatchMapInput is an input type that accepts FlowSchemaPatchMap and FlowSchemaPatchMapOutput values.
// You can construct a concrete instance of `FlowSchemaPatchMapInput` via:
//
//	FlowSchemaPatchMap{ "key": FlowSchemaPatchArgs{...} }
type FlowSchemaPatchMapInput interface {
	pulumi.Input

	ToFlowSchemaPatchMapOutput() FlowSchemaPatchMapOutput
	ToFlowSchemaPatchMapOutputWithContext(context.Context) FlowSchemaPatchMapOutput
}

type FlowSchemaPatchMap map[string]FlowSchemaPatchInput

func (FlowSchemaPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowSchemaPatch)(nil)).Elem()
}

func (i FlowSchemaPatchMap) ToFlowSchemaPatchMapOutput() FlowSchemaPatchMapOutput {
	return i.ToFlowSchemaPatchMapOutputWithContext(context.Background())
}

func (i FlowSchemaPatchMap) ToFlowSchemaPatchMapOutputWithContext(ctx context.Context) FlowSchemaPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaPatchMapOutput)
}

type FlowSchemaPatchOutput struct{ *pulumi.OutputState }

func (FlowSchemaPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowSchemaPatch)(nil)).Elem()
}

func (o FlowSchemaPatchOutput) ToFlowSchemaPatchOutput() FlowSchemaPatchOutput {
	return o
}

func (o FlowSchemaPatchOutput) ToFlowSchemaPatchOutputWithContext(ctx context.Context) FlowSchemaPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o FlowSchemaPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSchemaPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o FlowSchemaPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSchemaPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o FlowSchemaPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *FlowSchemaPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o FlowSchemaPatchOutput) Spec() FlowSchemaSpecPatchPtrOutput {
	return o.ApplyT(func(v *FlowSchemaPatch) FlowSchemaSpecPatchPtrOutput { return v.Spec }).(FlowSchemaSpecPatchPtrOutput)
}

// `status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o FlowSchemaPatchOutput) Status() FlowSchemaStatusPatchPtrOutput {
	return o.ApplyT(func(v *FlowSchemaPatch) FlowSchemaStatusPatchPtrOutput { return v.Status }).(FlowSchemaStatusPatchPtrOutput)
}

type FlowSchemaPatchArrayOutput struct{ *pulumi.OutputState }

func (FlowSchemaPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowSchemaPatch)(nil)).Elem()
}

func (o FlowSchemaPatchArrayOutput) ToFlowSchemaPatchArrayOutput() FlowSchemaPatchArrayOutput {
	return o
}

func (o FlowSchemaPatchArrayOutput) ToFlowSchemaPatchArrayOutputWithContext(ctx context.Context) FlowSchemaPatchArrayOutput {
	return o
}

func (o FlowSchemaPatchArrayOutput) Index(i pulumi.IntInput) FlowSchemaPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlowSchemaPatch {
		return vs[0].([]*FlowSchemaPatch)[vs[1].(int)]
	}).(FlowSchemaPatchOutput)
}

type FlowSchemaPatchMapOutput struct{ *pulumi.OutputState }

func (FlowSchemaPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowSchemaPatch)(nil)).Elem()
}

func (o FlowSchemaPatchMapOutput) ToFlowSchemaPatchMapOutput() FlowSchemaPatchMapOutput {
	return o
}

func (o FlowSchemaPatchMapOutput) ToFlowSchemaPatchMapOutputWithContext(ctx context.Context) FlowSchemaPatchMapOutput {
	return o
}

func (o FlowSchemaPatchMapOutput) MapIndex(k pulumi.StringInput) FlowSchemaPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlowSchemaPatch {
		return vs[0].(map[string]*FlowSchemaPatch)[vs[1].(string)]
	}).(FlowSchemaPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSchemaPatchInput)(nil)).Elem(), &FlowSchemaPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSchemaPatchArrayInput)(nil)).Elem(), FlowSchemaPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSchemaPatchMapInput)(nil)).Elem(), FlowSchemaPatchMap{})
	pulumi.RegisterOutputType(FlowSchemaPatchOutput{})
	pulumi.RegisterOutputType(FlowSchemaPatchArrayOutput{})
	pulumi.RegisterOutputType(FlowSchemaPatchMapOutput{})
}
