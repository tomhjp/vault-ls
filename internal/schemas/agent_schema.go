package schemas

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var (
	AgentSchema = &schema.BodySchema{
		Blocks: map[string]*schema.BlockSchema{
			"auto_auth":       autoAuthBlockSchema,
			"cache":           cacheBlockSchema,
			"vault":           vaultBlockSchema,
			"template_config": templateConfigBlockSchema,
			"template":        templateBlockSchema,
			"listener":        listenerBlockSchema,
		},

		Attributes: map[string]*schema.AttributeSchema{
			"pid_file": {
				Expr:        schema.LiteralTypeOnly(cty.String),
				IsOptional:  true,
				Description: lang.Markdown("File to store Vault Agent's process ID in"),
			},
			"exit_after_auth": {
				Expr:        schema.LiteralTypeOnly(cty.Bool),
				IsOptional:  true,
				Description: lang.Markdown("If true, Vault Agent will exit after authenticating and writing to any configured auto_auth sinks"),
			},
		},
	}

	autoAuthBlockSchema = &schema.BlockSchema{
		// Address: &schema.BlockAddrSchema{
		// 	Steps: []schema.AddrStep{
		// 		schema.LabelStep{Index: 0},
		// 		schema.AttrValueStep{Name: "alias", IsOptional: true},
		// 	},
		// 	FriendlyName: "provider",
		// 	ScopeId:      refscope.ProviderScope,
		// 	AsReference:  true,
		// },
		Description: lang.PlainText("An auto_auth block is used to specify the auth method Agent should use to authenticate with Vault"),
		MaxItems:    1,
		Body: &schema.BodySchema{
			Blocks: map[string]*schema.BlockSchema{
				"method": {
					Labels: []*schema.LabelSchema{
						{
							Name:        "type",
							IsDepKey:    true,
							Completable: true,
						},
					},
					MinItems: 1,
					MaxItems: 1,
					Body: &schema.BodySchema{
						Attributes: map[string]*schema.AttributeSchema{
							"mount_path": {
								Expr: schema.LiteralTypeOnly(cty.String),
							},
							"namespace": {
								Expr: schema.LiteralTypeOnly(cty.String),
							},
							"wrap_ttl": {
								Expr: schema.LiteralTypeOnly(cty.String),
							},
							"max_backoff": {
								Expr: schema.LiteralTypeOnly(cty.String),
							},
						},
					},
					DependentBody: map[schema.SchemaKey]*schema.BodySchema{
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "alicloud"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"role": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"region": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"credential_poll_interval": {
												Expr: schema.LiteralTypeOnly(cty.Number),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "approle"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"role_id_file_path": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"secret_id_file_path": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"remove_secret_id_file_after_reading": {
												Expr: schema.LiteralTypeOnly(cty.Bool),
											},
											"secret_id_response_wrapping_path": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "aws"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"type": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"role": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"credential_poll_interval": {
												Expr: schema.LiteralTypeOnly(cty.Number),
											},
											"access_key": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"secret_key": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"region": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"session_token": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"header_value": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"nonce": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "azure"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"role": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"resource": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "cert"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"name": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"ca_cert": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"client_cert": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"client_key": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "cf"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"role": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "gcp"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"type": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"role": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"credentials": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"service_account": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"jwt_exp": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "jwt"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"path": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"role": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "kerberos"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"krb5conf_path": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"keytab_path": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"username": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"service": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"realm": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"name": {
												Expr: schema.LiteralTypeOnly(cty.Bool),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
						schema.NewSchemaKey(schema.DependencyKeys{
							Labels: []schema.LabelDependent{
								{Index: 0, Value: "kubernetes"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"role": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"token_path": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
					},
				},
				"sink": {
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
								{Index: 0, Value: "file"},
							},
						}): {
							Blocks: map[string]*schema.BlockSchema{
								"config": {
									Body: &schema.BodySchema{
										Attributes: map[string]*schema.AttributeSchema{
											"path": {
												Expr: schema.LiteralTypeOnly(cty.String),
											},
											"mode": {
												Expr: schema.LiteralTypeOnly(cty.Number),
											},
										},
									},
									MinItems: 1,
									MaxItems: 1,
								},
							},
						},
					},
					Body: &schema.BodySchema{
						Attributes: map[string]*schema.AttributeSchema{
							"wrap_ttl": {
								Expr: schema.LiteralTypeOnly(cty.String),
								Description: lang.MarkupContent{
									Value: "If specified, the written token" +
										"will be response-wrapped by the agent. This is more secure than wrapping by " +
										"sinks, but does not allow the agent to keep the token renewed or " +
										"automatically reauthenticate when it expires. Rather than a simple string, " +
										"the written value will be a JSON-encoded " +
										"[SecretWrapInfo](https://godoc.org/github.com/hashicorp/vault/api#SecretWrapInfo) " +
										"structure. Values can be an integer number of seconds or a stringish value " +
										"like `5m`.",
									Kind: lang.MarkdownKind,
								},
							},
							"dh_type": {
								Expr: schema.LiteralTypeOnly(cty.String),
								Description: lang.MarkupContent{
									Value: "If specified, the type of Diffie-Hellman exchange to " +
										"perform, meaning, which ciphers and/or curves. Currently only `curve25519` is " +
										"supported.",
									Kind: lang.MarkdownKind,
								},
							},
							"dh_path": {
								Expr: schema.LiteralTypeOnly(cty.String),
								Description: lang.MarkupContent{
									Value: "The path from which the" +
										"agent should read the client's initial parameters (e.g. curve25519 public " +
										"key).",
									Kind: lang.MarkdownKind,
								},
							},
							"derive_key": {
								Expr: schema.LiteralTypeOnly(cty.Bool),
								Description: lang.MarkupContent{
									Value: "If specified, the final encryption key is" +
										"calculated by using HKDF-SHA256 to derive a key from the calculated shared " +
										"secret and the two public keys for enhanced security. This is recommended " +
										"if backward compatibility isn't a concern.",
									Kind: lang.MarkdownKind,
								},
							},
							"aad": {
								Expr: schema.LiteralTypeOnly(cty.String),
								Description: lang.MarkupContent{
									Value: "If specified, additional authenticated data to " +
										"use with the AES-GCM encryption of the token. Can be any string, including " +
										"serialized data.",
									Kind: lang.MarkdownKind,
								},
							},
							"aad_env_var": {
								Expr: schema.LiteralTypeOnly(cty.String),
								Description: lang.MarkupContent{
									Value: "If specified, AAD will be read from the " +
										"given environment variable rather than a value in the configuration file.",
									Kind: lang.MarkdownKind,
								},
							},
						},
					},
				},
			},
		},
	}
	cacheBlockSchema = &schema.BlockSchema{
		MaxItems: 1,
		Body: &schema.BodySchema{
			Attributes: map[string]*schema.AttributeSchema{
				"use_auto_auth_token": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"enforce_consistency": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"when_inconsistent": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
			},
			Blocks: map[string]*schema.BlockSchema{
				"persist": {
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
								{Index: 0, Value: "kubernetes"},
							},
						}): {
							Attributes: map[string]*schema.AttributeSchema{
								"path": {
									Expr: schema.LiteralTypeOnly(cty.String),
								},
								"keep_after_import": {
									Expr: schema.LiteralTypeOnly(cty.Bool),
								},
								"exit_on_err": {
									Expr: schema.LiteralTypeOnly(cty.Bool),
								},
							},
						},
					},
				},
			},
		},
	}
	vaultBlockSchema = &schema.BlockSchema{
		MaxItems: 1,
		Body: &schema.BodySchema{
			Attributes: map[string]*schema.AttributeSchema{
				"address": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"ca_cert": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"ca_path": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"client_cert": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"client_key": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"tls_skip_verify": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"tls_server_name": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
			},
			Blocks: map[string]*schema.BlockSchema{
				"retry": {
					Body: &schema.BodySchema{
						Attributes: map[string]*schema.AttributeSchema{
							"num_retries": {
								Expr: schema.LiteralTypeOnly(cty.Number),
							},
						},
					},
				},
			},
		},
	}
	templateConfigBlockSchema = &schema.BlockSchema{
		MaxItems: 1,
		Body: &schema.BodySchema{
			Attributes: map[string]*schema.AttributeSchema{
				"exit_on_retry_failure": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"static_secret_render_interval": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
			},
		},
	}
	templateBlockSchema = &schema.BlockSchema{
		Body: &schema.BodySchema{
			Attributes: map[string]*schema.AttributeSchema{
				"source": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"destination": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"create_dest_dirs": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"contents": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"commmand": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"command_timeout": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"error_on_missing_key": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"perms": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"backup": {
					Expr: schema.LiteralTypeOnly(cty.Bool),
				},
				"left_delimiter": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"right_delimiter": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"sandbox_path": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"wait": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
			},
		},
	}
	listenerBlockSchema = &schema.BlockSchema{
		Labels: []*schema.LabelSchema{
			{
				Name:        "type",
				IsDepKey:    true,
				Completable: true,
			},
		},
		Body: &schema.BodySchema{
			Attributes: map[string]*schema.AttributeSchema{
				"address": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"cluster_address": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"http_idle_timeout": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"http_read_header_timeout": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"http_read_timeout": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"http_write_timeout": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"max_request_size": {
					Expr: schema.LiteralTypeOnly(cty.Number),
				},
				"max_request_duration": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"proxy_protocol_behavior": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"proxy_protocol_authorized_addrs": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"tls_disable": {
					Expr: schema.LiteralTypeOnly(cty.String), // TODO: docs say string??
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
				"tls_cipher_suites": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"tls_prefer_server_cipher_suites": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"tls_require_and_verify_client_cert": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"tls_client_ca_file": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"tls_disable_client_certs": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"x_forwarded_for_authorized_addrs": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"x_forwarded_for_hop_skips": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"x_forwarded_for_reject_not_authorized": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
				"x_forwarded_for_reject_not_present": {
					Expr: schema.LiteralTypeOnly(cty.String),
				},
			},
			Blocks: map[string]*schema.BlockSchema{
				"telemetry": {
					Body: &schema.BodySchema{
						Attributes: map[string]*schema.AttributeSchema{
							"unauthenticated_metrics_access": {
								Expr: schema.LiteralTypeOnly(cty.Bool),
							},
						},
					},
				},
				"profiling": {
					Body: &schema.BodySchema{
						Attributes: map[string]*schema.AttributeSchema{
							"unauthenticated_pprof_access": {
								Expr: schema.LiteralTypeOnly(cty.Bool),
							},
						},
					},
				},
				"custom_response_headers": {
					Body: &schema.BodySchema{
						Attributes: map[string]*schema.AttributeSchema{
							"default": {
								// Expr: schema.MapExpr(), //TODO: type is a map
							},
							// TODO: Aribtrary status code keys
						},
					},
				},
			},
		},
	}
)
