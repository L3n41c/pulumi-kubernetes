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

// TokenReview attempts to authenticate a token to a known user. Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.
type TokenReview struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec TokenReviewSpecOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request can be authenticated.
	Status TokenReviewStatusPtrOutput `pulumi:"status"`
}

// NewTokenReview registers a new resource with the given unique name, arguments, and options.
func NewTokenReview(ctx *pulumi.Context,
	name string, args *TokenReviewArgs, opts ...pulumi.ResourceOption) (*TokenReview, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("authentication.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("TokenReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authentication.k8s.io/v1:TokenReview"),
		},
	})
	opts = append(opts, aliases)
	var resource TokenReview
	err := ctx.RegisterResource("kubernetes:authentication.k8s.io/v1beta1:TokenReview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTokenReview gets an existing TokenReview resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTokenReview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenReviewState, opts ...pulumi.ResourceOption) (*TokenReview, error) {
	var resource TokenReview
	err := ctx.ReadResource("kubernetes:authentication.k8s.io/v1beta1:TokenReview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TokenReview resources.
type tokenReviewState struct {
}

type TokenReviewState struct {
}

func (TokenReviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenReviewState)(nil)).Elem()
}

type tokenReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec TokenReviewSpec `pulumi:"spec"`
}

// The set of arguments for constructing a TokenReview resource.
type TokenReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated
	Spec TokenReviewSpecInput
}

func (TokenReviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenReviewArgs)(nil)).Elem()
}

type TokenReviewInput interface {
	pulumi.Input

	ToTokenReviewOutput() TokenReviewOutput
	ToTokenReviewOutputWithContext(ctx context.Context) TokenReviewOutput
}

func (*TokenReview) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenReview)(nil))
}

func (i *TokenReview) ToTokenReviewOutput() TokenReviewOutput {
	return i.ToTokenReviewOutputWithContext(context.Background())
}

func (i *TokenReview) ToTokenReviewOutputWithContext(ctx context.Context) TokenReviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenReviewOutput)
}

func (i *TokenReview) ToTokenReviewPtrOutput() TokenReviewPtrOutput {
	return i.ToTokenReviewPtrOutputWithContext(context.Background())
}

func (i *TokenReview) ToTokenReviewPtrOutputWithContext(ctx context.Context) TokenReviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenReviewPtrOutput)
}

type TokenReviewPtrInput interface {
	pulumi.Input

	ToTokenReviewPtrOutput() TokenReviewPtrOutput
	ToTokenReviewPtrOutputWithContext(ctx context.Context) TokenReviewPtrOutput
}

type tokenReviewPtrType TokenReviewArgs

func (*tokenReviewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenReview)(nil))
}

func (i *tokenReviewPtrType) ToTokenReviewPtrOutput() TokenReviewPtrOutput {
	return i.ToTokenReviewPtrOutputWithContext(context.Background())
}

func (i *tokenReviewPtrType) ToTokenReviewPtrOutputWithContext(ctx context.Context) TokenReviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenReviewPtrOutput)
}

// TokenReviewArrayInput is an input type that accepts TokenReviewArray and TokenReviewArrayOutput values.
// You can construct a concrete instance of `TokenReviewArrayInput` via:
//
//          TokenReviewArray{ TokenReviewArgs{...} }
type TokenReviewArrayInput interface {
	pulumi.Input

	ToTokenReviewArrayOutput() TokenReviewArrayOutput
	ToTokenReviewArrayOutputWithContext(context.Context) TokenReviewArrayOutput
}

type TokenReviewArray []TokenReviewInput

func (TokenReviewArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TokenReview)(nil))
}

func (i TokenReviewArray) ToTokenReviewArrayOutput() TokenReviewArrayOutput {
	return i.ToTokenReviewArrayOutputWithContext(context.Background())
}

func (i TokenReviewArray) ToTokenReviewArrayOutputWithContext(ctx context.Context) TokenReviewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenReviewArrayOutput)
}

// TokenReviewMapInput is an input type that accepts TokenReviewMap and TokenReviewMapOutput values.
// You can construct a concrete instance of `TokenReviewMapInput` via:
//
//          TokenReviewMap{ "key": TokenReviewArgs{...} }
type TokenReviewMapInput interface {
	pulumi.Input

	ToTokenReviewMapOutput() TokenReviewMapOutput
	ToTokenReviewMapOutputWithContext(context.Context) TokenReviewMapOutput
}

type TokenReviewMap map[string]TokenReviewInput

func (TokenReviewMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TokenReview)(nil))
}

func (i TokenReviewMap) ToTokenReviewMapOutput() TokenReviewMapOutput {
	return i.ToTokenReviewMapOutputWithContext(context.Background())
}

func (i TokenReviewMap) ToTokenReviewMapOutputWithContext(ctx context.Context) TokenReviewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenReviewMapOutput)
}

type TokenReviewOutput struct {
	*pulumi.OutputState
}

func (TokenReviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenReview)(nil))
}

func (o TokenReviewOutput) ToTokenReviewOutput() TokenReviewOutput {
	return o
}

func (o TokenReviewOutput) ToTokenReviewOutputWithContext(ctx context.Context) TokenReviewOutput {
	return o
}

func (o TokenReviewOutput) ToTokenReviewPtrOutput() TokenReviewPtrOutput {
	return o.ToTokenReviewPtrOutputWithContext(context.Background())
}

func (o TokenReviewOutput) ToTokenReviewPtrOutputWithContext(ctx context.Context) TokenReviewPtrOutput {
	return o.ApplyT(func(v TokenReview) *TokenReview {
		return &v
	}).(TokenReviewPtrOutput)
}

type TokenReviewPtrOutput struct {
	*pulumi.OutputState
}

func (TokenReviewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenReview)(nil))
}

func (o TokenReviewPtrOutput) ToTokenReviewPtrOutput() TokenReviewPtrOutput {
	return o
}

func (o TokenReviewPtrOutput) ToTokenReviewPtrOutputWithContext(ctx context.Context) TokenReviewPtrOutput {
	return o
}

type TokenReviewArrayOutput struct{ *pulumi.OutputState }

func (TokenReviewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenReview)(nil))
}

func (o TokenReviewArrayOutput) ToTokenReviewArrayOutput() TokenReviewArrayOutput {
	return o
}

func (o TokenReviewArrayOutput) ToTokenReviewArrayOutputWithContext(ctx context.Context) TokenReviewArrayOutput {
	return o
}

func (o TokenReviewArrayOutput) Index(i pulumi.IntInput) TokenReviewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenReview {
		return vs[0].([]TokenReview)[vs[1].(int)]
	}).(TokenReviewOutput)
}

type TokenReviewMapOutput struct{ *pulumi.OutputState }

func (TokenReviewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TokenReview)(nil))
}

func (o TokenReviewMapOutput) ToTokenReviewMapOutput() TokenReviewMapOutput {
	return o
}

func (o TokenReviewMapOutput) ToTokenReviewMapOutputWithContext(ctx context.Context) TokenReviewMapOutput {
	return o
}

func (o TokenReviewMapOutput) MapIndex(k pulumi.StringInput) TokenReviewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TokenReview {
		return vs[0].(map[string]TokenReview)[vs[1].(string)]
	}).(TokenReviewOutput)
}

func init() {
	pulumi.RegisterOutputType(TokenReviewOutput{})
	pulumi.RegisterOutputType(TokenReviewPtrOutput{})
	pulumi.RegisterOutputType(TokenReviewArrayOutput{})
	pulumi.RegisterOutputType(TokenReviewMapOutput{})
}
