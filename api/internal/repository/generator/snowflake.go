package generator

import (
	"github.com/kytruong0712/getgo/api/internal/pkg/snowflake"
)

var (
	ProductSNF snowflake.SnowflakeGenerator
)

func InitSnowflakeGenerators() error {
	ProductSNF = snowflake.New()

	return nil
}
