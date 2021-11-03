// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DEPRECATED 1.9 - This group version of NetworkPolicy is deprecated by networking/v1/NetworkPolicy. NetworkPolicy describes what network traffic is allowed for a set of Pods
type NetworkPolicy struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Specification of the desired behavior for this NetworkPolicy.
	Spec NetworkPolicySpecPtrOutput `pulumi:"spec"`
}

// NewNetworkPolicy registers a new resource with the given unique name, arguments, and options.
func NewNetworkPolicy(ctx *pulumi.Context,
	name string, args *NetworkPolicyArgs, opts ...pulumi.ResourceOption) (*NetworkPolicy, error) {
	if args == nil {
		args = &NetworkPolicyArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("extensions/v1beta1")
	args.Kind = pulumi.StringPtr("NetworkPolicy")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:networking.k8s.io/v1:NetworkPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkPolicy
	err := ctx.RegisterResource("kubernetes:extensions/v1beta1:NetworkPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPolicy gets an existing NetworkPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPolicyState, opts ...pulumi.ResourceOption) (*NetworkPolicy, error) {
	var resource NetworkPolicy
	err := ctx.ReadResource("kubernetes:extensions/v1beta1:NetworkPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPolicy resources.
type networkPolicyState struct {
}

type NetworkPolicyState struct {
}

func (NetworkPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPolicyState)(nil)).Elem()
}

type networkPolicyArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired behavior for this NetworkPolicy.
	Spec *NetworkPolicySpec `pulumi:"spec"`
}

// The set of arguments for constructing a NetworkPolicy resource.
type NetworkPolicyArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the desired behavior for this NetworkPolicy.
	Spec NetworkPolicySpecPtrInput
}

func (NetworkPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPolicyArgs)(nil)).Elem()
}

type NetworkPolicyInput interface {
	pulumi.Input

	ToNetworkPolicyOutput() NetworkPolicyOutput
	ToNetworkPolicyOutputWithContext(ctx context.Context) NetworkPolicyOutput
}

func (*NetworkPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPolicy)(nil))
}

func (i *NetworkPolicy) ToNetworkPolicyOutput() NetworkPolicyOutput {
	return i.ToNetworkPolicyOutputWithContext(context.Background())
}

func (i *NetworkPolicy) ToNetworkPolicyOutputWithContext(ctx context.Context) NetworkPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyOutput)
}

func (i *NetworkPolicy) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return i.ToNetworkPolicyPtrOutputWithContext(context.Background())
}

func (i *NetworkPolicy) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyPtrOutput)
}

type NetworkPolicyPtrInput interface {
	pulumi.Input

	ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput
	ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput
}

type networkPolicyPtrType NetworkPolicyArgs

func (*networkPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPolicy)(nil))
}

func (i *networkPolicyPtrType) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return i.ToNetworkPolicyPtrOutputWithContext(context.Background())
}

func (i *networkPolicyPtrType) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyPtrOutput)
}

// NetworkPolicyArrayInput is an input type that accepts NetworkPolicyArray and NetworkPolicyArrayOutput values.
// You can construct a concrete instance of `NetworkPolicyArrayInput` via:
//
//          NetworkPolicyArray{ NetworkPolicyArgs{...} }
type NetworkPolicyArrayInput interface {
	pulumi.Input

	ToNetworkPolicyArrayOutput() NetworkPolicyArrayOutput
	ToNetworkPolicyArrayOutputWithContext(context.Context) NetworkPolicyArrayOutput
}

type NetworkPolicyArray []NetworkPolicyInput

func (NetworkPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPolicy)(nil)).Elem()
}

func (i NetworkPolicyArray) ToNetworkPolicyArrayOutput() NetworkPolicyArrayOutput {
	return i.ToNetworkPolicyArrayOutputWithContext(context.Background())
}

func (i NetworkPolicyArray) ToNetworkPolicyArrayOutputWithContext(ctx context.Context) NetworkPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyArrayOutput)
}

// NetworkPolicyMapInput is an input type that accepts NetworkPolicyMap and NetworkPolicyMapOutput values.
// You can construct a concrete instance of `NetworkPolicyMapInput` via:
//
//          NetworkPolicyMap{ "key": NetworkPolicyArgs{...} }
type NetworkPolicyMapInput interface {
	pulumi.Input

	ToNetworkPolicyMapOutput() NetworkPolicyMapOutput
	ToNetworkPolicyMapOutputWithContext(context.Context) NetworkPolicyMapOutput
}

type NetworkPolicyMap map[string]NetworkPolicyInput

func (NetworkPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPolicy)(nil)).Elem()
}

func (i NetworkPolicyMap) ToNetworkPolicyMapOutput() NetworkPolicyMapOutput {
	return i.ToNetworkPolicyMapOutputWithContext(context.Background())
}

func (i NetworkPolicyMap) ToNetworkPolicyMapOutputWithContext(ctx context.Context) NetworkPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPolicyMapOutput)
}

type NetworkPolicyOutput struct{ *pulumi.OutputState }

func (NetworkPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPolicy)(nil))
}

func (o NetworkPolicyOutput) ToNetworkPolicyOutput() NetworkPolicyOutput {
	return o
}

func (o NetworkPolicyOutput) ToNetworkPolicyOutputWithContext(ctx context.Context) NetworkPolicyOutput {
	return o
}

func (o NetworkPolicyOutput) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return o.ToNetworkPolicyPtrOutputWithContext(context.Background())
}

func (o NetworkPolicyOutput) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkPolicy) *NetworkPolicy {
		return &v
	}).(NetworkPolicyPtrOutput)
}

type NetworkPolicyPtrOutput struct{ *pulumi.OutputState }

func (NetworkPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPolicy)(nil))
}

func (o NetworkPolicyPtrOutput) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return o
}

func (o NetworkPolicyPtrOutput) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return o
}

func (o NetworkPolicyPtrOutput) Elem() NetworkPolicyOutput {
	return o.ApplyT(func(v *NetworkPolicy) NetworkPolicy {
		if v != nil {
			return *v
		}
		var ret NetworkPolicy
		return ret
	}).(NetworkPolicyOutput)
}

type NetworkPolicyArrayOutput struct{ *pulumi.OutputState }

func (NetworkPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkPolicy)(nil))
}

func (o NetworkPolicyArrayOutput) ToNetworkPolicyArrayOutput() NetworkPolicyArrayOutput {
	return o
}

func (o NetworkPolicyArrayOutput) ToNetworkPolicyArrayOutputWithContext(ctx context.Context) NetworkPolicyArrayOutput {
	return o
}

func (o NetworkPolicyArrayOutput) Index(i pulumi.IntInput) NetworkPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkPolicy {
		return vs[0].([]NetworkPolicy)[vs[1].(int)]
	}).(NetworkPolicyOutput)
}

type NetworkPolicyMapOutput struct{ *pulumi.OutputState }

func (NetworkPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkPolicy)(nil))
}

func (o NetworkPolicyMapOutput) ToNetworkPolicyMapOutput() NetworkPolicyMapOutput {
	return o
}

func (o NetworkPolicyMapOutput) ToNetworkPolicyMapOutputWithContext(ctx context.Context) NetworkPolicyMapOutput {
	return o
}

func (o NetworkPolicyMapOutput) MapIndex(k pulumi.StringInput) NetworkPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkPolicy {
		return vs[0].(map[string]NetworkPolicy)[vs[1].(string)]
	}).(NetworkPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyInput)(nil)).Elem(), &NetworkPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyPtrInput)(nil)).Elem(), &NetworkPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyArrayInput)(nil)).Elem(), NetworkPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyMapInput)(nil)).Elem(), NetworkPolicyMap{})
	pulumi.RegisterOutputType(NetworkPolicyOutput{})
	pulumi.RegisterOutputType(NetworkPolicyPtrOutput{})
	pulumi.RegisterOutputType(NetworkPolicyArrayOutput{})
	pulumi.RegisterOutputType(NetworkPolicyMapOutput{})
}
