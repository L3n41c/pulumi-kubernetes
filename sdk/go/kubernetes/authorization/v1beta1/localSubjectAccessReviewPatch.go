// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
// LocalSubjectAccessReview checks whether or not a user or group can perform an action in a given namespace. Having a namespace scoped resource makes it much easier to grant namespace scoped policy that includes permissions checking.
type LocalSubjectAccessReviewPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput          `pulumi:"kind"`
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
	Spec SubjectAccessReviewSpecPatchPtrOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusPatchPtrOutput `pulumi:"status"`
}

// NewLocalSubjectAccessReviewPatch registers a new resource with the given unique name, arguments, and options.
func NewLocalSubjectAccessReviewPatch(ctx *pulumi.Context,
	name string, args *LocalSubjectAccessReviewPatchArgs, opts ...pulumi.ResourceOption) (*LocalSubjectAccessReviewPatch, error) {
	if args == nil {
		args = &LocalSubjectAccessReviewPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("authorization.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("LocalSubjectAccessReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authorization.k8s.io/v1:LocalSubjectAccessReviewPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource LocalSubjectAccessReviewPatch
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1beta1:LocalSubjectAccessReviewPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalSubjectAccessReviewPatch gets an existing LocalSubjectAccessReviewPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalSubjectAccessReviewPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalSubjectAccessReviewPatchState, opts ...pulumi.ResourceOption) (*LocalSubjectAccessReviewPatch, error) {
	var resource LocalSubjectAccessReviewPatch
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1beta1:LocalSubjectAccessReviewPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalSubjectAccessReviewPatch resources.
type localSubjectAccessReviewPatchState struct {
}

type LocalSubjectAccessReviewPatchState struct {
}

func (LocalSubjectAccessReviewPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*localSubjectAccessReviewPatchState)(nil)).Elem()
}

type localSubjectAccessReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string                 `pulumi:"kind"`
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
	Spec *SubjectAccessReviewSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a LocalSubjectAccessReviewPatch resource.
type LocalSubjectAccessReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPatchPtrInput
	// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
	Spec SubjectAccessReviewSpecPatchPtrInput
}

func (LocalSubjectAccessReviewPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localSubjectAccessReviewPatchArgs)(nil)).Elem()
}

type LocalSubjectAccessReviewPatchInput interface {
	pulumi.Input

	ToLocalSubjectAccessReviewPatchOutput() LocalSubjectAccessReviewPatchOutput
	ToLocalSubjectAccessReviewPatchOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPatchOutput
}

func (*LocalSubjectAccessReviewPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalSubjectAccessReviewPatch)(nil)).Elem()
}

func (i *LocalSubjectAccessReviewPatch) ToLocalSubjectAccessReviewPatchOutput() LocalSubjectAccessReviewPatchOutput {
	return i.ToLocalSubjectAccessReviewPatchOutputWithContext(context.Background())
}

func (i *LocalSubjectAccessReviewPatch) ToLocalSubjectAccessReviewPatchOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSubjectAccessReviewPatchOutput)
}

// LocalSubjectAccessReviewPatchArrayInput is an input type that accepts LocalSubjectAccessReviewPatchArray and LocalSubjectAccessReviewPatchArrayOutput values.
// You can construct a concrete instance of `LocalSubjectAccessReviewPatchArrayInput` via:
//
//	LocalSubjectAccessReviewPatchArray{ LocalSubjectAccessReviewPatchArgs{...} }
type LocalSubjectAccessReviewPatchArrayInput interface {
	pulumi.Input

	ToLocalSubjectAccessReviewPatchArrayOutput() LocalSubjectAccessReviewPatchArrayOutput
	ToLocalSubjectAccessReviewPatchArrayOutputWithContext(context.Context) LocalSubjectAccessReviewPatchArrayOutput
}

type LocalSubjectAccessReviewPatchArray []LocalSubjectAccessReviewPatchInput

func (LocalSubjectAccessReviewPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalSubjectAccessReviewPatch)(nil)).Elem()
}

func (i LocalSubjectAccessReviewPatchArray) ToLocalSubjectAccessReviewPatchArrayOutput() LocalSubjectAccessReviewPatchArrayOutput {
	return i.ToLocalSubjectAccessReviewPatchArrayOutputWithContext(context.Background())
}

func (i LocalSubjectAccessReviewPatchArray) ToLocalSubjectAccessReviewPatchArrayOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSubjectAccessReviewPatchArrayOutput)
}

// LocalSubjectAccessReviewPatchMapInput is an input type that accepts LocalSubjectAccessReviewPatchMap and LocalSubjectAccessReviewPatchMapOutput values.
// You can construct a concrete instance of `LocalSubjectAccessReviewPatchMapInput` via:
//
//	LocalSubjectAccessReviewPatchMap{ "key": LocalSubjectAccessReviewPatchArgs{...} }
type LocalSubjectAccessReviewPatchMapInput interface {
	pulumi.Input

	ToLocalSubjectAccessReviewPatchMapOutput() LocalSubjectAccessReviewPatchMapOutput
	ToLocalSubjectAccessReviewPatchMapOutputWithContext(context.Context) LocalSubjectAccessReviewPatchMapOutput
}

type LocalSubjectAccessReviewPatchMap map[string]LocalSubjectAccessReviewPatchInput

func (LocalSubjectAccessReviewPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalSubjectAccessReviewPatch)(nil)).Elem()
}

func (i LocalSubjectAccessReviewPatchMap) ToLocalSubjectAccessReviewPatchMapOutput() LocalSubjectAccessReviewPatchMapOutput {
	return i.ToLocalSubjectAccessReviewPatchMapOutputWithContext(context.Background())
}

func (i LocalSubjectAccessReviewPatchMap) ToLocalSubjectAccessReviewPatchMapOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSubjectAccessReviewPatchMapOutput)
}

type LocalSubjectAccessReviewPatchOutput struct{ *pulumi.OutputState }

func (LocalSubjectAccessReviewPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalSubjectAccessReviewPatch)(nil)).Elem()
}

func (o LocalSubjectAccessReviewPatchOutput) ToLocalSubjectAccessReviewPatchOutput() LocalSubjectAccessReviewPatchOutput {
	return o
}

func (o LocalSubjectAccessReviewPatchOutput) ToLocalSubjectAccessReviewPatchOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o LocalSubjectAccessReviewPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalSubjectAccessReviewPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o LocalSubjectAccessReviewPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalSubjectAccessReviewPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LocalSubjectAccessReviewPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *LocalSubjectAccessReviewPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
func (o LocalSubjectAccessReviewPatchOutput) Spec() SubjectAccessReviewSpecPatchPtrOutput {
	return o.ApplyT(func(v *LocalSubjectAccessReviewPatch) SubjectAccessReviewSpecPatchPtrOutput { return v.Spec }).(SubjectAccessReviewSpecPatchPtrOutput)
}

// Status is filled in by the server and indicates whether the request is allowed or not
func (o LocalSubjectAccessReviewPatchOutput) Status() SubjectAccessReviewStatusPatchPtrOutput {
	return o.ApplyT(func(v *LocalSubjectAccessReviewPatch) SubjectAccessReviewStatusPatchPtrOutput { return v.Status }).(SubjectAccessReviewStatusPatchPtrOutput)
}

type LocalSubjectAccessReviewPatchArrayOutput struct{ *pulumi.OutputState }

func (LocalSubjectAccessReviewPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalSubjectAccessReviewPatch)(nil)).Elem()
}

func (o LocalSubjectAccessReviewPatchArrayOutput) ToLocalSubjectAccessReviewPatchArrayOutput() LocalSubjectAccessReviewPatchArrayOutput {
	return o
}

func (o LocalSubjectAccessReviewPatchArrayOutput) ToLocalSubjectAccessReviewPatchArrayOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPatchArrayOutput {
	return o
}

func (o LocalSubjectAccessReviewPatchArrayOutput) Index(i pulumi.IntInput) LocalSubjectAccessReviewPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalSubjectAccessReviewPatch {
		return vs[0].([]*LocalSubjectAccessReviewPatch)[vs[1].(int)]
	}).(LocalSubjectAccessReviewPatchOutput)
}

type LocalSubjectAccessReviewPatchMapOutput struct{ *pulumi.OutputState }

func (LocalSubjectAccessReviewPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalSubjectAccessReviewPatch)(nil)).Elem()
}

func (o LocalSubjectAccessReviewPatchMapOutput) ToLocalSubjectAccessReviewPatchMapOutput() LocalSubjectAccessReviewPatchMapOutput {
	return o
}

func (o LocalSubjectAccessReviewPatchMapOutput) ToLocalSubjectAccessReviewPatchMapOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPatchMapOutput {
	return o
}

func (o LocalSubjectAccessReviewPatchMapOutput) MapIndex(k pulumi.StringInput) LocalSubjectAccessReviewPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalSubjectAccessReviewPatch {
		return vs[0].(map[string]*LocalSubjectAccessReviewPatch)[vs[1].(string)]
	}).(LocalSubjectAccessReviewPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalSubjectAccessReviewPatchInput)(nil)).Elem(), &LocalSubjectAccessReviewPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalSubjectAccessReviewPatchArrayInput)(nil)).Elem(), LocalSubjectAccessReviewPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalSubjectAccessReviewPatchMapInput)(nil)).Elem(), LocalSubjectAccessReviewPatchMap{})
	pulumi.RegisterOutputType(LocalSubjectAccessReviewPatchOutput{})
	pulumi.RegisterOutputType(LocalSubjectAccessReviewPatchArrayOutput{})
	pulumi.RegisterOutputType(LocalSubjectAccessReviewPatchMapOutput{})
}
