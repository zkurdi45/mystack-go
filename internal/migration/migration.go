// internal/migrations/migrations.go
package migrations

import "embed"

// The go:embed directive is now relative to this file. It looks for a
// directory named "sql" within this same package.
//
//go:embed sql/*.sql
var FS embed.FS
