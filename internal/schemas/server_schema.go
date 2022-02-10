package schemas

import (
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var (
	ServerSchema = &schema.BodySchema{
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
			"storage":              {},
			"ha_storage":           {},
			"listener":             {},
			"telemetry":            {},
			"seal":                 {},
			"replication":          {},
			"sentinel":             {},
			"service_registration": {},
			"entropy":              {},
		},
	}
)
