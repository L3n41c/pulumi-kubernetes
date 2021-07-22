// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lease defines a lease concept.
type Lease struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec LeaseSpecPtrOutput `pulumi:"spec"`
}

// NewLease registers a new resource with the given unique name, arguments, and options.
func NewLease(ctx *pulumi.Context,
	name string, args *LeaseArgs, opts ...pulumi.ResourceOption) (*Lease, error) {
	if args == nil {
		args = &LeaseArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("coordination.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("Lease")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:coordination.k8s.io/v1:Lease"),
		},
	})
	opts = append(opts, aliases)
	var resource Lease
	err := ctx.RegisterResource("kubernetes:coordination.k8s.io/v1beta1:Lease", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLease gets an existing Lease resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LeaseState, opts ...pulumi.ResourceOption) (*Lease, error) {
	var resource Lease
	err := ctx.ReadResource("kubernetes:coordination.k8s.io/v1beta1:Lease", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Lease resources.
type leaseState struct {
}

type LeaseState struct {
}

func (LeaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*leaseState)(nil)).Elem()
}

type leaseArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *LeaseSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Lease resource.
type LeaseArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec LeaseSpecPtrInput
}

func (LeaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*leaseArgs)(nil)).Elem()
}

type LeaseInput interface {
	pulumi.Input

	ToLeaseOutput() LeaseOutput
	ToLeaseOutputWithContext(ctx context.Context) LeaseOutput
}

func (*Lease) ElementType() reflect.Type {
	return reflect.TypeOf((*Lease)(nil))
}

func (i *Lease) ToLeaseOutput() LeaseOutput {
	return i.ToLeaseOutputWithContext(context.Background())
}

func (i *Lease) ToLeaseOutputWithContext(ctx context.Context) LeaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseOutput)
}

func (i *Lease) ToLeasePtrOutput() LeasePtrOutput {
	return i.ToLeasePtrOutputWithContext(context.Background())
}

func (i *Lease) ToLeasePtrOutputWithContext(ctx context.Context) LeasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeasePtrOutput)
}

type LeasePtrInput interface {
	pulumi.Input

	ToLeasePtrOutput() LeasePtrOutput
	ToLeasePtrOutputWithContext(ctx context.Context) LeasePtrOutput
}

type leasePtrType LeaseArgs

func (*leasePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Lease)(nil))
}

func (i *leasePtrType) ToLeasePtrOutput() LeasePtrOutput {
	return i.ToLeasePtrOutputWithContext(context.Background())
}

func (i *leasePtrType) ToLeasePtrOutputWithContext(ctx context.Context) LeasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeasePtrOutput)
}

// LeaseArrayInput is an input type that accepts LeaseArray and LeaseArrayOutput values.
// You can construct a concrete instance of `LeaseArrayInput` via:
//
//          LeaseArray{ LeaseArgs{...} }
type LeaseArrayInput interface {
	pulumi.Input

	ToLeaseArrayOutput() LeaseArrayOutput
	ToLeaseArrayOutputWithContext(context.Context) LeaseArrayOutput
}

type LeaseArray []LeaseInput

func (LeaseArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Lease)(nil))
}

func (i LeaseArray) ToLeaseArrayOutput() LeaseArrayOutput {
	return i.ToLeaseArrayOutputWithContext(context.Background())
}

func (i LeaseArray) ToLeaseArrayOutputWithContext(ctx context.Context) LeaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseArrayOutput)
}

// LeaseMapInput is an input type that accepts LeaseMap and LeaseMapOutput values.
// You can construct a concrete instance of `LeaseMapInput` via:
//
//          LeaseMap{ "key": LeaseArgs{...} }
type LeaseMapInput interface {
	pulumi.Input

	ToLeaseMapOutput() LeaseMapOutput
	ToLeaseMapOutputWithContext(context.Context) LeaseMapOutput
}

type LeaseMap map[string]LeaseInput

func (LeaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Lease)(nil))
}

func (i LeaseMap) ToLeaseMapOutput() LeaseMapOutput {
	return i.ToLeaseMapOutputWithContext(context.Background())
}

func (i LeaseMap) ToLeaseMapOutputWithContext(ctx context.Context) LeaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseMapOutput)
}

type LeaseOutput struct {
	*pulumi.OutputState
}

func (LeaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Lease)(nil))
}

func (o LeaseOutput) ToLeaseOutput() LeaseOutput {
	return o
}

func (o LeaseOutput) ToLeaseOutputWithContext(ctx context.Context) LeaseOutput {
	return o
}

func (o LeaseOutput) ToLeasePtrOutput() LeasePtrOutput {
	return o.ToLeasePtrOutputWithContext(context.Background())
}

func (o LeaseOutput) ToLeasePtrOutputWithContext(ctx context.Context) LeasePtrOutput {
	return o.ApplyT(func(v Lease) *Lease {
		return &v
	}).(LeasePtrOutput)
}

type LeasePtrOutput struct {
	*pulumi.OutputState
}

func (LeasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lease)(nil))
}

func (o LeasePtrOutput) ToLeasePtrOutput() LeasePtrOutput {
	return o
}

func (o LeasePtrOutput) ToLeasePtrOutputWithContext(ctx context.Context) LeasePtrOutput {
	return o
}

type LeaseArrayOutput struct{ *pulumi.OutputState }

func (LeaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Lease)(nil))
}

func (o LeaseArrayOutput) ToLeaseArrayOutput() LeaseArrayOutput {
	return o
}

func (o LeaseArrayOutput) ToLeaseArrayOutputWithContext(ctx context.Context) LeaseArrayOutput {
	return o
}

func (o LeaseArrayOutput) Index(i pulumi.IntInput) LeaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Lease {
		return vs[0].([]Lease)[vs[1].(int)]
	}).(LeaseOutput)
}

type LeaseMapOutput struct{ *pulumi.OutputState }

func (LeaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Lease)(nil))
}

func (o LeaseMapOutput) ToLeaseMapOutput() LeaseMapOutput {
	return o
}

func (o LeaseMapOutput) ToLeaseMapOutputWithContext(ctx context.Context) LeaseMapOutput {
	return o
}

func (o LeaseMapOutput) MapIndex(k pulumi.StringInput) LeaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Lease {
		return vs[0].(map[string]Lease)[vs[1].(string)]
	}).(LeaseOutput)
}

func init() {
	pulumi.RegisterOutputType(LeaseOutput{})
	pulumi.RegisterOutputType(LeasePtrOutput{})
	pulumi.RegisterOutputType(LeaseArrayOutput{})
	pulumi.RegisterOutputType(LeaseMapOutput{})
}
