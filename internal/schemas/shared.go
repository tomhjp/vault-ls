package schemas

import (
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var (
	// TODO: Agent supports an additional option and also unix type
	listenerBlockSchema = &schema.BlockSchema{
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
					{Index: 0, Value: "tcp"},
				},
			}): {
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
		},
	}
)
