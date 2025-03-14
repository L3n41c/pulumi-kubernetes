// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Authorization.V1
{

    /// <summary>
    /// SelfSubjectRulesReviewSpec defines the specification for SelfSubjectRulesReview.
    /// </summary>
    [OutputType]
    public sealed class SelfSubjectRulesReviewSpecPatch
    {
        /// <summary>
        /// Namespace to evaluate rules for. Required.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private SelfSubjectRulesReviewSpecPatch(string @namespace)
        {
            Namespace = @namespace;
        }
    }
}
