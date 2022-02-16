package schemas

import (
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var (
	ServerSchema = &schema.BodySchema{
		DocsLink: &schema.DocsLink{
			URL: "https://www.vaultproject.io/docs/configuration",
		},
		Attributes: map[string]*schema.AttributeSchema{
			"cluster_name": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"cache_size": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"disable_cache": {
				Expr: schema.LiteralTypeOnly(cty.Bool),
			},
			"disable_mlock": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"plugin_directory": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"log_level": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"log_format": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"default_lease_ttl": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"max_lease_ttl": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"default_max_request_duration": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"raw_storage_endpoint": {
				Expr: schema.LiteralTypeOnly(cty.Bool),
			},
			"ui": {
				Expr: schema.LiteralTypeOnly(cty.Bool),
			},
			"pid_file": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"enable_response_header_hostname": {
				Expr: schema.LiteralTypeOnly(cty.Bool),
			},
			"enable_response_header_raft_node_id": {
				Expr: schema.LiteralTypeOnly(cty.Bool),
			},
			"api_addr": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"cluster_addr": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
			"disable_clustering": {
				Expr: schema.LiteralTypeOnly(cty.Bool),
			},
			"disable_sealwrap": {
				Expr: schema.LiteralTypeOnly(cty.Bool),
			},
			"disable_performance_standby": {
				Expr: schema.LiteralTypeOnly(cty.Bool),
			},
			"license_path": {
				Expr: schema.LiteralTypeOnly(cty.String),
			},
		},
		Blocks: map[string]*schema.BlockSchema{
			"storage":              storageBlockSchema,
			"ha_storage":           storageBlockSchema, // TODO: Limit this to just backends supporting HA
			"listener":             listenerBlockSchema,
			"telemetry":            telemetryBlockSchema,
			"seal":                 sealBlockSchema,
			"replication":          replicationBlockSchema,
			"sentinel":             sentinelBlockSchema,
			"service_registration": serviceRegistrationBlockSchema,
			"entropy":              entropyBlockSchema,
		},
	}

	sealBlockSchema = &schema.BlockSchema{
		Labels: []*schema.LabelSchema{
			{
				Name:        "type",
				IsDepKey:    true,
				Completable: true,
			},
		},
		DependentBody: map[schema.SchemaKey]*schema.BodySchema{
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "alicloudkms"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/seal/alicloudkms",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"region": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"domain": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"access_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"secret_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"kms_key_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "awskms"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/seal/awskms",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"region": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"access_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"session_token": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"secret_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"kms_key_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"endpoint": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "azurekeyvault"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/seal/azurekeyvault",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"tenant_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"client_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"client_secret": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"environment": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"vault_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"key_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"resource": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "gcpckms"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/seal/gcpckms",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"credentials": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"project": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"region": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"key_ring": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"crypto_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "ocikms"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/seal/ocikms",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"key_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"crypto_endpoint": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"management_endpoint": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"auth_type_api_key": {
						Expr: schema.LiteralTypeOnly(cty.Bool),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "pkcs11"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/seal/pkcs11",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"lib": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"slot": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"token_label": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"pin": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"key_label": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"default_key_label": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"key_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"hmac_key_label": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"default_hmac_key_label": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"hmac_key_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"mechanism": {
						Expr: schema.LiteralTypeOnly(cty.String),
						// TODO: Can we make enums like this work?
						// Expr: schema.ExprConstraints{
						// 	schema.LiteralValue{
						// 		Val:         cty.StringVal("0x1085"),
						// 		Description: lang.Markdown("CKM_AES_CBC_PAD"),
						// 	},
						// 	schema.LiteralValue{
						// 		Val:         cty.StringVal("0x1082"),
						// 		Description: lang.Markdown("CKM_AES_CBC"),
						// 	},
						// 	schema.LiteralValue{
						// 		Val:         cty.StringVal("0x1087"),
						// 		Description: lang.Markdown("CKM_AES_GCM"),
						// 	},
						// 	schema.LiteralValue{
						// 		Val:         cty.StringVal("0x0009"),
						// 		Description: lang.Markdown("CKM_RSA_PKCS_OAEP"),
						// 	},
						// 	schema.LiteralValue{
						// 		Val:         cty.StringVal("0x0001"),
						// 		Description: lang.Markdown("CKM_RSA_PKCS"),
						// 	},
						// },
					},
					"hmac_mechanism": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"generate_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"force_rw_session": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					// TODO: Dependent on `mechanism` attribute
					"rsa_encrypt_local": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					// TODO: Dependent on `mechanism` attribute
					"rsa_oaep_hash": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "transit"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/seal/transit",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"address": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"token": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"key_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"mount_path": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"namespace": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"disable_renewal": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_ca_cert": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_client_cert": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_client_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_server_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_skip_verify": {
						Expr: schema.LiteralTypeOnly(cty.Bool),
					},
				},
			},
		},
	}

	serviceRegistrationBlockSchema = &schema.BlockSchema{
		Labels: []*schema.LabelSchema{
			{
				Name:        "type",
				IsDepKey:    true,
				Completable: true,
			},
		},
		DependentBody: map[schema.SchemaKey]*schema.BodySchema{
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "consul"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/service-registration/consul",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"address": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"check_timeout": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"disable_registration": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"scheme": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"service": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"service_tags": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"service_address": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"token": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_ca_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_cert_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_key_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_min_version": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_skip_verify": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "kubernetes"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/service-registration/kubernetes",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"namespace": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"pod_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
		},
	}

	replicationBlockSchema = &schema.BlockSchema{
		Body: &schema.BodySchema{
			DocsLink: &schema.DocsLink{
				URL: "https://www.vaultproject.io/docs/configuration/replication",
			},
			Attributes: map[string]*schema.AttributeSchema{
				"resolver_discover_servers": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"logshipper_buffer_length": {
					Expr: schema.LiteralTypeOnly(cty.Number),
				},
				"logshipper_buffer_size": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"allow_forwarding_via_header": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"best_effort_wal_wait_duration": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
			},
		},
	}

	sentinelBlockSchema = &schema.BlockSchema{
		Body: &schema.BodySchema{
			DocsLink: &schema.DocsLink{
				URL: "https://www.vaultproject.io/docs/configuration/sentinel",
			},
			Attributes: map[string]*schema.AttributeSchema{
				"additional_enabled_modules": {
					Expr: schema.ExprConstraints{
						schema.ListExpr{
							Elem: schema.LiteralTypeOnly(cty.String),
						},
					},
				},
			},
		},
	}

	entropyBlockSchema = &schema.BlockSchema{
		Labels: []*schema.LabelSchema{
			{
				Name:        "type",
				IsDepKey:    true,
				Completable: true,
			},
		},
		DependentBody: map[schema.SchemaKey]*schema.BodySchema{
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "seal"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/entropy-augmentation",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"mode": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
		},
	}

	telemetryBlockSchema = &schema.BlockSchema{
		Body: &schema.BodySchema{
			DocsLink: &schema.DocsLink{
				URL: "https://www.vaultproject.io/docs/configuration/telemetry",
			},
			Attributes: map[string]*schema.AttributeSchema{
				"usage_gauge_period": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"maximum_gauge_cardinality": {
					Expr: schema.LiteralTypeOnly(cty.Number),
				},
				"disable_hostname": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"enable_hostname_label": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"lease_metrics_epsilon": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"num_lease_metrics_buckets": {
					Expr: schema.LiteralTypeOnly(cty.Number),
				},
				"add_lease_metrics_namespace_labels": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"filter_default": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"prefix_filter": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"statsite_address": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"statsd_address": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_api_token": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_api_app": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_api_url": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_submission_interval": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_submission_url": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_check_id": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_check_force_metric_activation": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"circonus_check_instance_id": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_check_search_tag": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_check_display_name": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_check_tags": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_broker_id": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"circonus_broker_select_tag": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"dogstatsd_addr": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"prometheus_retention_time": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"stackdriver_project_id": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"stackdriver_location": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"stackdriver_namespace": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"stackdriver_debug_logs": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
			},
		},
	}

	storageBlockSchema = &schema.BlockSchema{
		Labels: []*schema.LabelSchema{
			{
				Name:        "type",
				IsDepKey:    true,
				Completable: true,
			},
		},
		DependentBody: map[schema.SchemaKey]*schema.BodySchema{
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "aerospike"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/aerospike",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"hostname": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"port": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"hostlist": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"namespace": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"set": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"username": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"password": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"cluster_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"auth_mode": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"timeout": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"idle_timeout": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "alicloudoss"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/alicloudoss",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"bucket": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"endpoint": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"access_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"secret_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "azure"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/azure",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"accountName": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"accountKey": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"container": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"environment": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"arm_endpoint": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "cassandra"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/cassandra",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"hosts": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"keyspace": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"consistency": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"protocol_version": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"username": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"password": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"connection_timeout": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"tls": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"pem_bundle_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"pem_json_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_skip_verify": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"tls_min_version": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "cockroachdb"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/cockroachdb",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"connection_url": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "consul"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/consul",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"address": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"check_timeout": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"consistency_mode": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"disable_registration": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"path": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"scheme": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"service": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"service_tags": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"service_address": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"token": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"session_ttl": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"lock_wait_time": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_ca_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_cert_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_key_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_min_version": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_skip_verify": {
						Expr: schema.LiteralTypeOnly(cty.Bool),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "couchdb"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/couchdb",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"endpoint": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"username": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"password": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "dynamodb"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/dynamodb",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"endpoint": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"ha_enabled": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"region": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"read_capacity": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"write_capacity": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"access_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"secret_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"session_token": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "etcd"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/etcd",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"address": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"discovery_srv": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"discovery_srv_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"etcd_api": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"ha_enabled": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"path": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"sync": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"username": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"password": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_ca_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_cert_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_key_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"request_timeout": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"lock_timeout": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "file"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/filesystem",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"path": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "foundationdb"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/foundationdb",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"api_version": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"cluster_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_verify_peers": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_ca_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_cert_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_key_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_password": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"path": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"ha_enabled": {
						Expr: schema.LiteralTypeOnly(cty.Bool),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "spanner"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/google-cloud-spanner",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"database": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"ha_enabled": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"ha_table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "gcs"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/google-cloud-storage",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"bucket": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"chunk_size": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"ha_enabled": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "inmem"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/in-memory",
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "manta"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/manta",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"directory": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"user": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"key_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"subuser": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"url": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "mssql"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/mssql",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"server": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"port": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"username": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"password": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"database": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"schema": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"appname": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"connectionTimeout": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"logLevel": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "mysql"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/mysql",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"address": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"database": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_ca_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"plaintext_credentials_transmission": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_idle_connections": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_connection_lifetime": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"username": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"password": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"ha_enabled": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"lock_table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "oci"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/oci-object-storage",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"region": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"namespace_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"bucket_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"ha_enabled": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"lock_bucket_name": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "postgresql"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/postgresql",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"connection_url": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_idle_connections": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"ha_enabled": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"ha_table": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "raft"},
				},
			}): {
				Blocks: map[string]*schema.BlockSchema{
					"retry_join": {
						Body: &schema.BodySchema{
							DocsLink: &schema.DocsLink{
								URL: "https://www.vaultproject.io/docs/configuration/storage/raft#retry_join-stanza",
							},
							Attributes: map[string]*schema.AttributeSchema{
								"leader_api_addr": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"auto_join": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"auto_join_scheme": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"auto_join_port": {
									Expr: schema.LiteralTypeOnly(cty.Number),
								},
								"leader_tls_servername": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"leader_ca_cert_file": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"leader_client_cert_file": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"leader_client_key_file": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"leader_ca_cert": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"leader_client_cert": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"leader_client_key": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
							},
						},
					},
				},
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/raft",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"path": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"node_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"performance_multiplier": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"trailing_logs": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"snapshot_threshold": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"max_entry_size": {
						Expr: schema.LiteralTypeOnly(cty.Number),
					},
					"autopilot_reconcile_interval": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "s3"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/s3",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"bucket": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"endpoint": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"region": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"access_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"secret_key": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"session_token": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"s3_force_path_style": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"disable_ssl": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"kms_key_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"path": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "swift"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/swift",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"auth_url": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"container": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"max_parallel": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"password": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tenant": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"username": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"region": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tenant_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"domain": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"project": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"trust_id": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"storage_url": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"auth_token": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
				},
			},
			schema.NewSchemaKey(schema.DependencyKeys{
				Labels: []schema.LabelDependent{
					{Index: 0, Value: "zookeeper"},
				},
			}): {
				DocsLink: &schema.DocsLink{
					URL: "https://www.vaultproject.io/docs/configuration/storage/zookeeper",
				},
				Attributes: map[string]*schema.AttributeSchema{
					"address": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"path": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"auth_info": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"znode_owner": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_enabled": {
						Expr: schema.LiteralTypeOnly(cty.Bool),
					},
					"tls_ca_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_cert_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_key_file": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_min_version": {
						Expr: schema.LiteralTypeOnly(cty.String),
					},
					"tls_skip_verify": {
						Expr: schema.LiteralTypeOnly(cty.Bool),
					},
					"tls_verify_ip": {
						Expr: schema.LiteralTypeOnly(cty.Bool),
					},
				},
			},
		},
	}
)
