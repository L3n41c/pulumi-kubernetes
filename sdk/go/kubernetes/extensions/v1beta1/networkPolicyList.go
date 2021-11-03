// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DEPRECATED 1.9 - This group version of NetworkPolicyList is deprecated by networking/v1/NetworkPolicyList. Network Policy List is a list of NetworkPolicy objects.
type NetworkPolicyList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items NetworkPolicyTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewNetworkPolicyList registers a new resource with the given unique name, arguments, and options.
func NewNetworkPolicyList(ctx *pulumi.Context,
	name string, args *NetworkPolicyListArgs, opts ...pulumi.ResourceOption) (*NetworkPolicyList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("extensions/v1beta1")
	args.Kind = pulumi.StringPtr("NetworkPolicyList")
	var resource NetworkPolicyList
	err := ctx.RegisterResource("kubernetes:extensions/v1beta1:NetworkPolicyList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPolicyList gets an existing NetworkPolicyList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPolicyList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPolicyListState, opts ...pulumi.ResourceOption) (*NetworkPolicyList, error) {
	var resource NetworkPolicyList
	err := ctx.ReadResource("kubernetes:extensions/v1beta1:NetworkPolicyList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPolicyList resources.
type networkPolicyListState struct {
}

type NetworkPolicyListState struct {
}

func (NetworkPolicyListState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPolicyListState)(nil)).Elem()
}

type networkPolicyListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items []NetworkPolicyType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a NetworkPolicyList resource.
type NetworkPolicyListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of schema objects.
	Items NetworkPolicyTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (NetworkPolicyListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPolicyListArgs)(nil)).Elem()
}

type NetworkPolicyListInput interface {
	pulumi.Input

	ToNetworkPolicyListOutput() NetworkPolicyListOutput
	ToNetworkPolicyListOutputWithContext(ctx context.Context) NetworkPolicyListOutput
}

func (*NetworkPolicyList) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPolicyList)(nil))
}

func (i *NetworkPolicyList) ToNetworkPolicyListOutput() NetworkPolicyListOutput {
	return i.ToNetworkPolicyListOutputWithContext(context.Background())
}

func (i *NetworkPolicyList) ToNetworkPolicyListOutputWithContext(ctx context.Context) NetworkPolicyListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyListOutput)
}

func (i *NetworkPolicyList) ToNetworkPolicyListPtrOutput() NetworkPolicyListPtrOutput {
	return i.ToNetworkPolicyListPtrOutputWithContext(context.Background())
}

func (i *NetworkPolicyList) ToNetworkPolicyListPtrOutputWithContext(ctx context.Context) NetworkPolicyListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyListPtrOutput)
}

type NetworkPolicyListPtrInput interface {
	pulumi.Input

	ToNetworkPolicyListPtrOutput() NetworkPolicyListPtrOutput
	ToNetworkPolicyListPtrOutputWithContext(ctx context.Context) NetworkPolicyListPtrOutput
}

type networkPolicyListPtrType NetworkPolicyListArgs

func (*networkPolicyListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPolicyList)(nil))
}

func (i *networkPolicyListPtrType) ToNetworkPolicyListPtrOutput() NetworkPolicyListPtrOutput {
	return i.ToNetworkPolicyListPtrOutputWithContext(context.Background())
}

func (i *networkPolicyListPtrType) ToNetworkPolicyListPtrOutputWithContext(ctx context.Context) NetworkPolicyListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyListPtrOutput)
}

// NetworkPolicyListArrayInput is an input type that accepts NetworkPolicyListArray and NetworkPolicyListArrayOutput values.
// You can construct a concrete instance of `NetworkPolicyListArrayInput` via:
//
//          NetworkPolicyListArray{ NetworkPolicyListArgs{...} }
type NetworkPolicyListArrayInput interface {
	pulumi.Input

	ToNetworkPolicyListArrayOutput() NetworkPolicyListArrayOutput
	ToNetworkPolicyListArrayOutputWithContext(context.Context) NetworkPolicyListArrayOutput
}

type NetworkPolicyListArray []NetworkPolicyListInput

func (NetworkPolicyListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPolicyList)(nil)).Elem()
}

func (i NetworkPolicyListArray) ToNetworkPolicyListArrayOutput() NetworkPolicyListArrayOutput {
	return i.ToNetworkPolicyListArrayOutputWithContext(context.Background())
}

func (i NetworkPolicyListArray) ToNetworkPolicyListArrayOutputWithContext(ctx context.Context) NetworkPolicyListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyListArrayOutput)
}

// NetworkPolicyListMapInput is an input type that accepts NetworkPolicyListMap and NetworkPolicyListMapOutput values.
// You can construct a concrete instance of `NetworkPolicyListMapInput` via:
//
//          NetworkPolicyListMap{ "key": NetworkPolicyListArgs{...} }
type NetworkPolicyListMapInput interface {
	pulumi.Input

	ToNetworkPolicyListMapOutput() NetworkPolicyListMapOutput
	ToNetworkPolicyListMapOutputWithContext(context.Context) NetworkPolicyListMapOutput
}

type NetworkPolicyListMap map[string]NetworkPolicyListInput

func (NetworkPolicyListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPolicyList)(nil)).Elem()
}

func (i NetworkPolicyListMap) ToNetworkPolicyListMapOutput() NetworkPolicyListMapOutput {
	return i.ToNetworkPolicyListMapOutputWithContext(context.Background())
}

func (i NetworkPolicyListMap) ToNetworkPolicyListMapOutputWithContext(ctx context.Context) NetworkPolicyListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyListMapOutput)
}

type NetworkPolicyListOutput struct{ *pulumi.OutputState }

func (NetworkPolicyListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPolicyList)(nil))
}

func (o NetworkPolicyListOutput) ToNetworkPolicyListOutput() NetworkPolicyListOutput {
	return o
}

func (o NetworkPolicyListOutput) ToNetworkPolicyListOutputWithContext(ctx context.Context) NetworkPolicyListOutput {
	return o
}

func (o NetworkPolicyListOutput) ToNetworkPolicyListPtrOutput() NetworkPolicyListPtrOutput {
	return o.ToNetworkPolicyListPtrOutputWithContext(context.Background())
}

func (o NetworkPolicyListOutput) ToNetworkPolicyListPtrOutputWithContext(ctx context.Context) NetworkPolicyListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkPolicyList) *NetworkPolicyList {
		return &v
	}).(NetworkPolicyListPtrOutput)
}

type NetworkPolicyListPtrOutput struct{ *pulumi.OutputState }

func (NetworkPolicyListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPolicyList)(nil))
}

func (o NetworkPolicyListPtrOutput) ToNetworkPolicyListPtrOutput() NetworkPolicyListPtrOutput {
	return o
}

func (o NetworkPolicyListPtrOutput) ToNetworkPolicyListPtrOutputWithContext(ctx context.Context) NetworkPolicyListPtrOutput {
	return o
}

func (o NetworkPolicyListPtrOutput) Elem() NetworkPolicyListOutput {
	return o.ApplyT(func(v *NetworkPolicyList) NetworkPolicyList {
		if v != nil {
			return *v
		}
		var ret NetworkPolicyList
		return ret
	}).(NetworkPolicyListOutput)
}

type NetworkPolicyListArrayOutput struct{ *pulumi.OutputState }

func (NetworkPolicyListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkPolicyList)(nil))
}

func (o NetworkPolicyListArrayOutput) ToNetworkPolicyListArrayOutput() NetworkPolicyListArrayOutput {
	return o
}

func (o NetworkPolicyListArrayOutput) ToNetworkPolicyListArrayOutputWithContext(ctx context.Context) NetworkPolicyListArrayOutput {
	return o
}

func (o NetworkPolicyListArrayOutput) Index(i pulumi.IntInput) NetworkPolicyListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkPolicyList {
		return vs[0].([]NetworkPolicyList)[vs[1].(int)]
	}).(NetworkPolicyListOutput)
}

type NetworkPolicyListMapOutput struct{ *pulumi.OutputState }

func (NetworkPolicyListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkPolicyList)(nil))
}

func (o NetworkPolicyListMapOutput) ToNetworkPolicyListMapOutput() NetworkPolicyListMapOutput {
	return o
}

func (o NetworkPolicyListMapOutput) ToNetworkPolicyListMapOutputWithContext(ctx context.Context) NetworkPolicyListMapOutput {
	return o
}

func (o NetworkPolicyListMapOutput) MapIndex(k pulumi.StringInput) NetworkPolicyListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkPolicyList {
		return vs[0].(map[string]NetworkPolicyList)[vs[1].(string)]
	}).(NetworkPolicyListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyListInput)(nil)).Elem(), &NetworkPolicyList{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyListPtrInput)(nil)).Elem(), &NetworkPolicyList{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyListArrayInput)(nil)).Elem(), NetworkPolicyListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyListMapInput)(nil)).Elem(), NetworkPolicyListMap{})
	pulumi.RegisterOutputType(NetworkPolicyListOutput{})
	pulumi.RegisterOutputType(NetworkPolicyListPtrOutput{})
	pulumi.RegisterOutputType(NetworkPolicyListArrayOutput{})
	pulumi.RegisterOutputType(NetworkPolicyListMapOutput{})
}
