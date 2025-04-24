module github.com/tenderly/net-nitro

go 1.22

replace github.com/VictoriaMetrics/fastcache => ./fastcache

replace github.com/tenderly/net-nitro/go-ethereum => ./go-ethereum

replace github.com/offchainlabs/bold => ./bold

require (
	cloud.google.com/go/storage v1.43.0
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/Shopify/toxiproxy v2.1.4+incompatible
	github.com/alicebob/miniredis/v2 v2.32.1
	github.com/andybalholm/brotli v1.0.5
	github.com/aws/aws-sdk-go-v2 v1.27.2
	github.com/aws/aws-sdk-go-v2/config v1.27.11
	github.com/aws/aws-sdk-go-v2/credentials v1.17.11
	github.com/aws/aws-sdk-go-v2/service/lambda v1.47.0
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.16.15
	github.com/cavaliergopher/grab/v3 v3.0.1
	github.com/ccoveille/go-safecast v1.1.0
	github.com/cockroachdb/pebble v1.1.2
	github.com/codeclysm/extract/v3 v3.0.2
	github.com/dgraph-io/badger/v4 v4.2.0
	github.com/enescakir/emoji v1.0.0
	github.com/fatih/structtag v1.2.0
	github.com/gdamore/tcell/v2 v2.7.1
	github.com/gobwas/httphead v0.1.0
	github.com/gobwas/ws v1.2.1
	github.com/gobwas/ws-examples v0.0.0-20190625122829-a9e8908d9484
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/google/btree v1.1.2
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/holiman/uint256 v1.3.2
	github.com/jmoiron/sqlx v1.4.0
	github.com/knadh/koanf v1.4.0
	github.com/mailru/easygo v0.0.0-20190618140210-3c14a0dc985f
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/mitchellh/mapstructure v1.4.1
	github.com/offchainlabs/bold v0.0.3-0.20250313062923-4b76649f2abc
	github.com/pkg/errors v0.9.1
	github.com/r3labs/diff/v3 v3.0.1
	github.com/redis/go-redis/v9 v9.6.3
	github.com/rivo/tview v0.0.0-20240307173318-e804876934a1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.10.0
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	github.com/wealdtech/go-merkletree v1.0.0
	golang.org/x/crypto v0.32.0
	golang.org/x/sync v0.10.0
	golang.org/x/sys v0.29.0
	golang.org/x/term v0.28.0
	golang.org/x/tools v0.29.0
	google.golang.org/api v0.187.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)
