// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Authorization.V1Beta1
{

    public class SelfSubjectRulesReviewSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Namespace to evaluate rules for. Required.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public SelfSubjectRulesReviewSpecArgs()
        {
        }
        public static new SelfSubjectRulesReviewSpecArgs Empty => new SelfSubjectRulesReviewSpecArgs();
    }
}
