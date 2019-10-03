// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabase":       schema_pkg_apis_api_v1alpha1_ExternalDatabase(ref),
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseRef":    schema_pkg_apis_api_v1alpha1_ExternalDatabaseRef(ref),
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseSpec":   schema_pkg_apis_api_v1alpha1_ExternalDatabaseSpec(ref),
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseStatus": schema_pkg_apis_api_v1alpha1_ExternalDatabaseStatus(ref),
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLDatabase":          schema_pkg_apis_api_v1alpha1_MySQLDatabase(ref),
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLDatabaseSpec":      schema_pkg_apis_api_v1alpha1_MySQLDatabaseSpec(ref),
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLDatabaseStatus":    schema_pkg_apis_api_v1alpha1_MySQLDatabaseStatus(ref),
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLUser":              schema_pkg_apis_api_v1alpha1_MySQLUser(ref),
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLUserSpec":          schema_pkg_apis_api_v1alpha1_MySQLUserSpec(ref),
		"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLUserStatus":        schema_pkg_apis_api_v1alpha1_MySQLUserStatus(ref),
	}
}

func schema_pkg_apis_api_v1alpha1_ExternalDatabase(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExternalDatabase is the Schema for the externaldatabases API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseSpec", "github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_api_v1alpha1_ExternalDatabaseRef(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExternalDatabaseRef talks about an external database.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_api_v1alpha1_ExternalDatabaseSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExternalDatabaseSpec defines the desired state of ExternalDatabase",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"host": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"adminUser": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"adminPasswordRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"namespaceSelector": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
				},
				Required: []string{"host", "adminUser", "adminPasswordRef", "type"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SecretKeySelector", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_api_v1alpha1_ExternalDatabaseStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExternalDatabaseStatus defines the observed state of ExternalDatabase",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"reachable": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"error": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_api_v1alpha1_MySQLDatabase(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MySQLDatabase is the Schema for the mysqldatabases API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLDatabaseSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLDatabaseStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLDatabaseSpec", "github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLDatabaseStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_api_v1alpha1_MySQLDatabaseSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MySQLDatabaseSpec defines the desired state of MySqlDatabase",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"owner": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"externalDatabaseRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseRef"),
						},
					},
				},
				Required: []string{"name", "externalDatabaseRef"},
			},
		},
		Dependencies: []string{
			"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseRef"},
	}
}

func schema_pkg_apis_api_v1alpha1_MySQLDatabaseStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MySQLDatabaseStatus defines the observed state of MySQLDatabase",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"error": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"created": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"secretCreated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"existingSecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"secretCreated"},
			},
		},
	}
}

func schema_pkg_apis_api_v1alpha1_MySQLUser(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MySQLUser is the Schema for the mysqlusers API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLUserSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLUserStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLUserSpec", "github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.MySQLUserStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_api_v1alpha1_MySQLUserSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MySQLUserSpec defines the desired state of MySqlUser",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"host": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"externalDatabaseRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseRef"),
						},
					},
					"passwordSecretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name", "externalDatabaseRef", "passwordSecretName"},
			},
		},
		Dependencies: []string{
			"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1.ExternalDatabaseRef"},
	}
}

func schema_pkg_apis_api_v1alpha1_MySQLUserStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MySQLUserStatus defines the observed state of MySqlUser",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"created": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"error": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"created", "error"},
			},
		},
	}
}
