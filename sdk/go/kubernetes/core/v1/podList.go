// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PodList is a list of Pods.
type PodList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// List of pods. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items PodTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewPodList registers a new resource with the given unique name, arguments, and options.
func NewPodList(ctx *pulumi.Context,
	name string, args *PodListArgs, opts ...pulumi.ResourceOption) (*PodList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("PodList")
	var resource PodList
	err := ctx.RegisterResource("kubernetes:core/v1:PodList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodList gets an existing PodList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodListState, opts ...pulumi.ResourceOption) (*PodList, error) {
	var resource PodList
	err := ctx.ReadResource("kubernetes:core/v1:PodList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodList resources.
type podListState struct {
}

type PodListState struct {
}

func (PodListState) ElementType() reflect.Type {
	return reflect.TypeOf((*podListState)(nil)).Elem()
}

type podListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of pods. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []PodType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a PodList resource.
type PodListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of pods. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items PodTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (PodListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podListArgs)(nil)).Elem()
}

type PodListInput interface {
	pulumi.Input

	ToPodListOutput() PodListOutput
	ToPodListOutputWithContext(ctx context.Context) PodListOutput
}

func (*PodList) ElementType() reflect.Type {
	return reflect.TypeOf((*PodList)(nil))
}

func (i *PodList) ToPodListOutput() PodListOutput {
	return i.ToPodListOutputWithContext(context.Background())
}

func (i *PodList) ToPodListOutputWithContext(ctx context.Context) PodListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodListOutput)
}

func (i *PodList) ToPodListPtrOutput() PodListPtrOutput {
	return i.ToPodListPtrOutputWithContext(context.Background())
}

func (i *PodList) ToPodListPtrOutputWithContext(ctx context.Context) PodListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodListPtrOutput)
}

type PodListPtrInput interface {
	pulumi.Input

	ToPodListPtrOutput() PodListPtrOutput
	ToPodListPtrOutputWithContext(ctx context.Context) PodListPtrOutput
}

type podListPtrType PodListArgs

func (*podListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PodList)(nil))
}

func (i *podListPtrType) ToPodListPtrOutput() PodListPtrOutput {
	return i.ToPodListPtrOutputWithContext(context.Background())
}

func (i *podListPtrType) ToPodListPtrOutputWithContext(ctx context.Context) PodListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodListPtrOutput)
}

// PodListArrayInput is an input type that accepts PodListArray and PodListArrayOutput values.
// You can construct a concrete instance of `PodListArrayInput` via:
//
//          PodListArray{ PodListArgs{...} }
type PodListArrayInput interface {
	pulumi.Input

	ToPodListArrayOutput() PodListArrayOutput
	ToPodListArrayOutputWithContext(context.Context) PodListArrayOutput
}

type PodListArray []PodListInput

func (PodListArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PodList)(nil))
}

func (i PodListArray) ToPodListArrayOutput() PodListArrayOutput {
	return i.ToPodListArrayOutputWithContext(context.Background())
}

func (i PodListArray) ToPodListArrayOutputWithContext(ctx context.Context) PodListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodListArrayOutput)
}

// PodListMapInput is an input type that accepts PodListMap and PodListMapOutput values.
// You can construct a concrete instance of `PodListMapInput` via:
//
//          PodListMap{ "key": PodListArgs{...} }
type PodListMapInput interface {
	pulumi.Input

	ToPodListMapOutput() PodListMapOutput
	ToPodListMapOutputWithContext(context.Context) PodListMapOutput
}

type PodListMap map[string]PodListInput

func (PodListMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PodList)(nil))
}

func (i PodListMap) ToPodListMapOutput() PodListMapOutput {
	return i.ToPodListMapOutputWithContext(context.Background())
}

func (i PodListMap) ToPodListMapOutputWithContext(ctx context.Context) PodListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodListMapOutput)
}

type PodListOutput struct {
	*pulumi.OutputState
}

func (PodListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PodList)(nil))
}

func (o PodListOutput) ToPodListOutput() PodListOutput {
	return o
}

func (o PodListOutput) ToPodListOutputWithContext(ctx context.Context) PodListOutput {
	return o
}

func (o PodListOutput) ToPodListPtrOutput() PodListPtrOutput {
	return o.ToPodListPtrOutputWithContext(context.Background())
}

func (o PodListOutput) ToPodListPtrOutputWithContext(ctx context.Context) PodListPtrOutput {
	return o.ApplyT(func(v PodList) *PodList {
		return &v
	}).(PodListPtrOutput)
}

type PodListPtrOutput struct {
	*pulumi.OutputState
}

func (PodListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodList)(nil))
}

func (o PodListPtrOutput) ToPodListPtrOutput() PodListPtrOutput {
	return o
}

func (o PodListPtrOutput) ToPodListPtrOutputWithContext(ctx context.Context) PodListPtrOutput {
	return o
}

type PodListArrayOutput struct{ *pulumi.OutputState }

func (PodListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PodList)(nil))
}

func (o PodListArrayOutput) ToPodListArrayOutput() PodListArrayOutput {
	return o
}

func (o PodListArrayOutput) ToPodListArrayOutputWithContext(ctx context.Context) PodListArrayOutput {
	return o
}

func (o PodListArrayOutput) Index(i pulumi.IntInput) PodListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PodList {
		return vs[0].([]PodList)[vs[1].(int)]
	}).(PodListOutput)
}

type PodListMapOutput struct{ *pulumi.OutputState }

func (PodListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PodList)(nil))
}

func (o PodListMapOutput) ToPodListMapOutput() PodListMapOutput {
	return o
}

func (o PodListMapOutput) ToPodListMapOutputWithContext(ctx context.Context) PodListMapOutput {
	return o
}

func (o PodListMapOutput) MapIndex(k pulumi.StringInput) PodListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PodList {
		return vs[0].(map[string]PodList)[vs[1].(string)]
	}).(PodListOutput)
}

func init() {
	pulumi.RegisterOutputType(PodListOutput{})
	pulumi.RegisterOutputType(PodListPtrOutput{})
	pulumi.RegisterOutputType(PodListArrayOutput{})
	pulumi.RegisterOutputType(PodListMapOutput{})
}
