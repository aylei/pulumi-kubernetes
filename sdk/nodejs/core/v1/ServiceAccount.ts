// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { core } from "../..";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import { getVersion } from "../../version";

    /**
     * ServiceAccount binds together: * a name, understood by users, and perhaps by peripheral
     * systems, for an identity * a principal that can be authenticated and authorized * a set of
     * secrets
     */
    export class ServiceAccount extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"v1">;

      /**
       * AutomountServiceAccountToken indicates whether pods running as this service account should
       * have an API token automatically mounted. Can be overridden at the pod level.
       */
      public readonly automountServiceAccountToken: pulumi.Output<boolean>;

      /**
       * ImagePullSecrets is a list of references to secrets in the same namespace to use for
       * pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are
       * distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are
       * only accessed by the kubelet. More info:
       * https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
       */
      public readonly imagePullSecrets: pulumi.Output<outputs.core.v1.LocalObjectReference[]>;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"ServiceAccount">;

      /**
       * Standard object's metadata. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputs.meta.v1.ObjectMeta>;

      /**
       * Secrets is the list of secrets allowed to be used by pods running using this
       * ServiceAccount. More info: https://kubernetes.io/docs/concepts/configuration/secret
       */
      public readonly secrets: pulumi.Output<outputs.core.v1.ObjectReference[]>;

      /**
       * Get the state of an existing `ServiceAccount` resource, as identified by `id`.
       * The ID is of the form `[namespace]/<name>`; if `namespace` is omitted, then (per
       * Kubernetes convention) the ID becomes `default/<name>`.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form `[namespace]/<name>`.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceAccount {
          return new ServiceAccount(name, undefined, { ...opts, id: id });
      }

      /** @internal */
      private static readonly __pulumiType = "kubernetes:core/v1:ServiceAccount";

      /**
       * Returns true if the given object is an instance of ServiceAccount.  This is designed to work even
       * when multiple copies of the Pulumi SDK have been loaded into the same process.
       */
      public static isInstance(obj: any): obj is ServiceAccount {
          if (obj === undefined || obj === null) {
              return false;
          }

          return obj["__pulumiType"] === ServiceAccount.__pulumiType;
      }

      /**
       * Create a core.v1.ServiceAccount resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputs.core.v1.ServiceAccount, opts?: pulumi.CustomResourceOptions) {
          const props: pulumi.Inputs = {};

          props["apiVersion"] = "v1";
          props["automountServiceAccountToken"] = args && args.automountServiceAccountToken || undefined;
          props["imagePullSecrets"] = args && args.imagePullSecrets || undefined;
          props["kind"] = "ServiceAccount";
          props["metadata"] = args && args.metadata || undefined;
          props["secrets"] = args && args.secrets || undefined;

          props["status"] = undefined;

          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }

          const _opts = pulumi.mergeOptions(opts, {
              additionalSecretOutputs: [
              ],
              aliases: [
              ]
          });

          super(ServiceAccount.__pulumiType, name, props, _opts);
      }
    }
