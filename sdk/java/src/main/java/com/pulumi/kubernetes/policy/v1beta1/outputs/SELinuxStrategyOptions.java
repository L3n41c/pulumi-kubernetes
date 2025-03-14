// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.policy.v1beta1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.core.v1.outputs.SELinuxOptions;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SELinuxStrategyOptions {
    /**
     * @return rule is the strategy that will dictate the allowable labels that may be set.
     * 
     */
    private String rule;
    /**
     * @return seLinuxOptions required to run as; required for MustRunAs More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
     * 
     */
    private @Nullable SELinuxOptions seLinuxOptions;

    private SELinuxStrategyOptions() {}
    /**
     * @return rule is the strategy that will dictate the allowable labels that may be set.
     * 
     */
    public String rule() {
        return this.rule;
    }
    /**
     * @return seLinuxOptions required to run as; required for MustRunAs More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
     * 
     */
    public Optional<SELinuxOptions> seLinuxOptions() {
        return Optional.ofNullable(this.seLinuxOptions);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SELinuxStrategyOptions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String rule;
        private @Nullable SELinuxOptions seLinuxOptions;
        public Builder() {}
        public Builder(SELinuxStrategyOptions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.rule = defaults.rule;
    	      this.seLinuxOptions = defaults.seLinuxOptions;
        }

        @CustomType.Setter
        public Builder rule(String rule) {
            this.rule = Objects.requireNonNull(rule);
            return this;
        }
        @CustomType.Setter
        public Builder seLinuxOptions(@Nullable SELinuxOptions seLinuxOptions) {
            this.seLinuxOptions = seLinuxOptions;
            return this;
        }
        public SELinuxStrategyOptions build() {
            final var o = new SELinuxStrategyOptions();
            o.rule = rule;
            o.seLinuxOptions = seLinuxOptions;
            return o;
        }
    }
}
