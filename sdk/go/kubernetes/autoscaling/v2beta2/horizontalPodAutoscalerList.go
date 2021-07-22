// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// HorizontalPodAutoscalerList is a list of horizontal pod autoscaler objects.
type HorizontalPodAutoscalerList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items is the list of horizontal pod autoscaler objects.
	Items HorizontalPodAutoscalerTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// metadata is the standard list metadata.
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewHorizontalPodAutoscalerList registers a new resource with the given unique name, arguments, and options.
func NewHorizontalPodAutoscalerList(ctx *pulumi.Context,
	name string, args *HorizontalPodAutoscalerListArgs, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscalerList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("autoscaling/v2beta2")
	args.Kind = pulumi.StringPtr("HorizontalPodAutoscalerList")
	var resource HorizontalPodAutoscalerList
	err := ctx.RegisterResource("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscalerList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHorizontalPodAutoscalerList gets an existing HorizontalPodAutoscalerList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHorizontalPodAutoscalerList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HorizontalPodAutoscalerListState, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscalerList, error) {
	var resource HorizontalPodAutoscalerList
	err := ctx.ReadResource("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscalerList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HorizontalPodAutoscalerList resources.
type horizontalPodAutoscalerListState struct {
}

type HorizontalPodAutoscalerListState struct {
}

func (HorizontalPodAutoscalerListState) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerListState)(nil)).Elem()
}

type horizontalPodAutoscalerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of horizontal pod autoscaler objects.
	Items []HorizontalPodAutoscalerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata is the standard list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a HorizontalPodAutoscalerList resource.
type HorizontalPodAutoscalerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of horizontal pod autoscaler objects.
	Items HorizontalPodAutoscalerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata is the standard list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (HorizontalPodAutoscalerListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerListArgs)(nil)).Elem()
}

type HorizontalPodAutoscalerListInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerListOutput() HorizontalPodAutoscalerListOutput
	ToHorizontalPodAutoscalerListOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListOutput
}

func (*HorizontalPodAutoscalerList) ElementType() reflect.Type {
	return reflect.TypeOf((*HorizontalPodAutoscalerList)(nil))
}

func (i *HorizontalPodAutoscalerList) ToHorizontalPodAutoscalerListOutput() HorizontalPodAutoscalerListOutput {
	return i.ToHorizontalPodAutoscalerListOutputWithContext(context.Background())
}

func (i *HorizontalPodAutoscalerList) ToHorizontalPodAutoscalerListOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerListOutput)
}

func (i *HorizontalPodAutoscalerList) ToHorizontalPodAutoscalerListPtrOutput() HorizontalPodAutoscalerListPtrOutput {
	return i.ToHorizontalPodAutoscalerListPtrOutputWithContext(context.Background())
}

func (i *HorizontalPodAutoscalerList) ToHorizontalPodAutoscalerListPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerListPtrOutput)
}

type HorizontalPodAutoscalerListPtrInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerListPtrOutput() HorizontalPodAutoscalerListPtrOutput
	ToHorizontalPodAutoscalerListPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListPtrOutput
}

type horizontalPodAutoscalerListPtrType HorizontalPodAutoscalerListArgs

func (*horizontalPodAutoscalerListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscalerList)(nil))
}

func (i *horizontalPodAutoscalerListPtrType) ToHorizontalPodAutoscalerListPtrOutput() HorizontalPodAutoscalerListPtrOutput {
	return i.ToHorizontalPodAutoscalerListPtrOutputWithContext(context.Background())
}

func (i *horizontalPodAutoscalerListPtrType) ToHorizontalPodAutoscalerListPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerListPtrOutput)
}

// HorizontalPodAutoscalerListArrayInput is an input type that accepts HorizontalPodAutoscalerListArray and HorizontalPodAutoscalerListArrayOutput values.
// You can construct a concrete instance of `HorizontalPodAutoscalerListArrayInput` via:
//
//          HorizontalPodAutoscalerListArray{ HorizontalPodAutoscalerListArgs{...} }
type HorizontalPodAutoscalerListArrayInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerListArrayOutput() HorizontalPodAutoscalerListArrayOutput
	ToHorizontalPodAutoscalerListArrayOutputWithContext(context.Context) HorizontalPodAutoscalerListArrayOutput
}

type HorizontalPodAutoscalerListArray []HorizontalPodAutoscalerListInput

func (HorizontalPodAutoscalerListArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*HorizontalPodAutoscalerList)(nil))
}

func (i HorizontalPodAutoscalerListArray) ToHorizontalPodAutoscalerListArrayOutput() HorizontalPodAutoscalerListArrayOutput {
	return i.ToHorizontalPodAutoscalerListArrayOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerListArray) ToHorizontalPodAutoscalerListArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerListArrayOutput)
}

// HorizontalPodAutoscalerListMapInput is an input type that accepts HorizontalPodAutoscalerListMap and HorizontalPodAutoscalerListMapOutput values.
// You can construct a concrete instance of `HorizontalPodAutoscalerListMapInput` via:
//
//          HorizontalPodAutoscalerListMap{ "key": HorizontalPodAutoscalerListArgs{...} }
type HorizontalPodAutoscalerListMapInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerListMapOutput() HorizontalPodAutoscalerListMapOutput
	ToHorizontalPodAutoscalerListMapOutputWithContext(context.Context) HorizontalPodAutoscalerListMapOutput
}

type HorizontalPodAutoscalerListMap map[string]HorizontalPodAutoscalerListInput

func (HorizontalPodAutoscalerListMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*HorizontalPodAutoscalerList)(nil))
}

func (i HorizontalPodAutoscalerListMap) ToHorizontalPodAutoscalerListMapOutput() HorizontalPodAutoscalerListMapOutput {
	return i.ToHorizontalPodAutoscalerListMapOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerListMap) ToHorizontalPodAutoscalerListMapOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerListMapOutput)
}

type HorizontalPodAutoscalerListOutput struct {
	*pulumi.OutputState
}

func (HorizontalPodAutoscalerListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HorizontalPodAutoscalerList)(nil))
}

func (o HorizontalPodAutoscalerListOutput) ToHorizontalPodAutoscalerListOutput() HorizontalPodAutoscalerListOutput {
	return o
}

func (o HorizontalPodAutoscalerListOutput) ToHorizontalPodAutoscalerListOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListOutput {
	return o
}

func (o HorizontalPodAutoscalerListOutput) ToHorizontalPodAutoscalerListPtrOutput() HorizontalPodAutoscalerListPtrOutput {
	return o.ToHorizontalPodAutoscalerListPtrOutputWithContext(context.Background())
}

func (o HorizontalPodAutoscalerListOutput) ToHorizontalPodAutoscalerListPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListPtrOutput {
	return o.ApplyT(func(v HorizontalPodAutoscalerList) *HorizontalPodAutoscalerList {
		return &v
	}).(HorizontalPodAutoscalerListPtrOutput)
}

type HorizontalPodAutoscalerListPtrOutput struct {
	*pulumi.OutputState
}

func (HorizontalPodAutoscalerListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscalerList)(nil))
}

func (o HorizontalPodAutoscalerListPtrOutput) ToHorizontalPodAutoscalerListPtrOutput() HorizontalPodAutoscalerListPtrOutput {
	return o
}

func (o HorizontalPodAutoscalerListPtrOutput) ToHorizontalPodAutoscalerListPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListPtrOutput {
	return o
}

type HorizontalPodAutoscalerListArrayOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HorizontalPodAutoscalerList)(nil))
}

func (o HorizontalPodAutoscalerListArrayOutput) ToHorizontalPodAutoscalerListArrayOutput() HorizontalPodAutoscalerListArrayOutput {
	return o
}

func (o HorizontalPodAutoscalerListArrayOutput) ToHorizontalPodAutoscalerListArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListArrayOutput {
	return o
}

func (o HorizontalPodAutoscalerListArrayOutput) Index(i pulumi.IntInput) HorizontalPodAutoscalerListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HorizontalPodAutoscalerList {
		return vs[0].([]HorizontalPodAutoscalerList)[vs[1].(int)]
	}).(HorizontalPodAutoscalerListOutput)
}

type HorizontalPodAutoscalerListMapOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HorizontalPodAutoscalerList)(nil))
}

func (o HorizontalPodAutoscalerListMapOutput) ToHorizontalPodAutoscalerListMapOutput() HorizontalPodAutoscalerListMapOutput {
	return o
}

func (o HorizontalPodAutoscalerListMapOutput) ToHorizontalPodAutoscalerListMapOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListMapOutput {
	return o
}

func (o HorizontalPodAutoscalerListMapOutput) MapIndex(k pulumi.StringInput) HorizontalPodAutoscalerListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HorizontalPodAutoscalerList {
		return vs[0].(map[string]HorizontalPodAutoscalerList)[vs[1].(string)]
	}).(HorizontalPodAutoscalerListOutput)
}

func init() {
	pulumi.RegisterOutputType(HorizontalPodAutoscalerListOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerListPtrOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerListArrayOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerListMapOutput{})
}
