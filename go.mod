module github.com/skip-mev/connect-mmu

go 1.23.2

replace (
	cosmossdk.io/core => cosmossdk.io/core v0.11.0

	// required replace for initia
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	// compatibility with older version of datadog used in dydxprotocol
	github.com/google/pprof => github.com/google/pprof v0.0.0-20210423192551-a2663126120b

	github.com/skip-mev/connect/v2 => github.com/initia-labs/connect/v2 v2.3.2
)

require (
	cosmossdk.io/api v0.7.6
	cosmossdk.io/math v1.4.0
	cosmossdk.io/x/tx v0.13.7
	github.com/avast/retry-go/v4 v4.6.0
	github.com/cometbft/cometbft v0.38.15
	github.com/cosmos/cosmos-sdk v0.50.11
	github.com/cosmos/gogoproto v1.7.0
	github.com/gagliardetto/binary v0.8.0
	github.com/gagliardetto/solana-go v1.12.0
	github.com/golangci/golangci-lint v1.62.2
	github.com/josephburnett/jd/v2 v2.0.0-20240818191833-6125a15c637a
	github.com/mitchellh/mapstructure v1.5.0
	github.com/skip-mev/connect/v2 v2.2.1
	github.com/skip-mev/slinky v1.0.8
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.10.0
	github.com/vektra/mockery/v2 v2.50.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20240909161429-701f63a606c0
	golang.org/x/sync v0.10.0
	golang.org/x/vuln v1.1.3
	gonum.org/v1/gonum v0.15.1
	google.golang.org/grpc v1.68.1
	google.golang.org/protobuf v1.35.2
	gopkg.in/typ.v4 v4.3.1
	mvdan.cc/gofumpt v0.7.0
)

require (
	4d63.com/gocheckcompilerdirectives v1.2.1 // indirect
	4d63.com/gochecknoglobals v0.2.1 // indirect
	cosmossdk.io/collections v0.4.0 // indirect
	cosmossdk.io/core v0.11.1 // indirect
	cosmossdk.io/depinject v1.1.0 // indirect
	cosmossdk.io/errors v1.0.1 // indirect
	cosmossdk.io/log v1.4.1 // indirect
	cosmossdk.io/store v1.1.1 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/4meepo/tagalign v1.3.4 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.1 // indirect
	github.com/Abirdcfly/dupword v0.1.3 // indirect
	github.com/Antonboom/errname v1.0.0 // indirect
	github.com/Antonboom/nilnil v1.0.0 // indirect
	github.com/Antonboom/testifylint v1.5.2 // indirect
	github.com/BurntSushi/toml v1.4.1-0.20240526193622-a339e1f7089c // indirect
	github.com/Crocmagnon/fatcontext v0.5.3 // indirect
	github.com/DataDog/datadog-go v3.2.0+incompatible // indirect
	github.com/DataDog/zstd v1.5.5 // indirect
	github.com/Djarvur/go-err113 v0.0.0-20210108212216-aea10b59be24 // indirect
	github.com/GaijinEntertainment/go-exhaustruct/v3 v3.3.0 // indirect
	github.com/GeertJohan/go.rice v1.0.3 // indirect
	github.com/Masterminds/semver/v3 v3.3.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/OpenPeeDeeP/depguard/v2 v2.2.0 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/alecthomas/go-check-sumtype v0.2.0 // indirect
	github.com/alexkohler/nakedret/v2 v2.0.5 // indirect
	github.com/alexkohler/prealloc v1.0.0 // indirect
	github.com/alingse/asasalint v0.0.11 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/ashanbrown/forbidigo v1.6.0 // indirect
	github.com/ashanbrown/makezero v1.1.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/speakeasy v0.1.1-0.20220910012023-760eaf8b6816 // indirect
	github.com/bits-and-blooms/bitset v1.14.2 // indirect
	github.com/bkielbasa/cyclop v1.2.3 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/blizzy78/varnamelen v0.8.0 // indirect
	github.com/bombsimon/wsl/v4 v4.4.1 // indirect
	github.com/breml/bidichk v0.3.2 // indirect
	github.com/breml/errchkjson v0.4.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/butuzov/ireturn v0.3.0 // indirect
	github.com/butuzov/mirror v1.2.0 // indirect
	github.com/catenacyber/perfsprint v0.7.1 // indirect
	github.com/ccojocar/zxcvbn-go v1.0.2 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/charithe/durationcheck v0.0.10 // indirect
	github.com/chavacava/garif v0.1.0 // indirect
	github.com/chigopher/pathlib v0.19.1 // indirect
	github.com/ckaznocha/intrange v0.2.1 // indirect
	github.com/cockroachdb/errors v1.11.3 // indirect
	github.com/cockroachdb/fifo v0.0.0-20240606204812-0bbfbd93a7ce // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v1.1.2 // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/cometbft/cometbft-db v0.15.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/cosmos/btcutil v1.0.5 // indirect
	github.com/cosmos/cosmos-db v1.1.0 // indirect
	github.com/cosmos/cosmos-proto v1.0.0-beta.5 // indirect
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/cosmos/gogogateway v1.2.0 // indirect
	github.com/cosmos/iavl v1.2.2 // indirect
	github.com/cosmos/ics23/go v0.11.0 // indirect
	github.com/cosmos/ledger-cosmos-go v0.13.3 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20240223125850-b1e8a79f509c // indirect
	github.com/crate-crypto/go-kzg-4844 v1.0.0 // indirect
	github.com/curioswitch/go-reassign v0.2.0 // indirect
	github.com/daaku/go.zipexe v1.0.2 // indirect
	github.com/daixiang0/gci v0.13.5 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set/v2 v2.6.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/denis-tingaikin/go-header v0.5.0 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgraph-io/badger/v4 v4.3.0 // indirect
	github.com/dgraph-io/ristretto v0.1.2-0.20240116140435-c67e07994f91 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/dvsekhvalnov/jose2go v1.6.0 // indirect
	github.com/emicklei/dot v1.6.2 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.14.11 // indirect
	github.com/ethereum/go-verkle v0.1.1-0.20240829091221-dffa7562dbe9 // indirect
	github.com/ettle/strcase v0.2.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/firefart/nonamedreturns v1.0.5 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/fzipp/gocyclo v0.6.0 // indirect
	github.com/gagliardetto/treeout v0.1.4 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/ghostiam/protogetter v0.3.8 // indirect
	github.com/go-critic/go-critic v0.11.5 // indirect
	github.com/go-kit/kit v0.13.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-toolsmith/astcast v1.1.0 // indirect
	github.com/go-toolsmith/astcopy v1.1.0 // indirect
	github.com/go-toolsmith/astequal v1.2.0 // indirect
	github.com/go-toolsmith/astfmt v1.1.0 // indirect
	github.com/go-toolsmith/astp v1.1.0 // indirect
	github.com/go-toolsmith/strparse v1.1.0 // indirect
	github.com/go-toolsmith/typep v1.1.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/go-xmlfmt/xmlfmt v1.1.2 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gofrs/flock v0.12.1 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/golangci/dupl v0.0.0-20180902072040-3e9179ac440a // indirect
	github.com/golangci/go-printf-func-name v0.1.0 // indirect
	github.com/golangci/gofmt v0.0.0-20240816233607-d8596aa466a9 // indirect
	github.com/golangci/misspell v0.6.0 // indirect
	github.com/golangci/modinfo v0.3.4 // indirect
	github.com/golangci/plugin-module-register v0.1.1 // indirect
	github.com/golangci/revgrep v0.5.3 // indirect
	github.com/golangci/unconvert v0.0.0-20240309020433-c5143eacb3ed // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/google/flatbuffers v1.12.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gordonklaus/ineffassign v0.1.0 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/rpc v1.2.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/gostaticanalysis/analysisutil v0.7.1 // indirect
	github.com/gostaticanalysis/comment v1.4.2 // indirect
	github.com/gostaticanalysis/forcetypeassert v0.1.0 // indirect
	github.com/gostaticanalysis/nilerr v0.1.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-metrics v0.5.3 // indirect
	github.com/hashicorp/go-plugin v1.5.2 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/hdevalence/ed25519consensus v0.1.0 // indirect
	github.com/hexops/gotextdiff v1.0.3 // indirect
	github.com/holiman/uint256 v1.3.1 // indirect
	github.com/huandu/skiplist v1.2.0 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/improbable-eng/grpc-web v0.15.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jgautheron/goconst v1.7.1 // indirect
	github.com/jingyugao/rowserrcheck v1.1.1 // indirect
	github.com/jinzhu/copier v0.3.5 // indirect
	github.com/jjti/go-spancheck v0.6.2 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/julz/importas v0.1.0 // indirect
	github.com/karamaru-alpha/copyloopvar v1.1.0 // indirect
	github.com/kisielk/errcheck v1.8.0 // indirect
	github.com/kkHAIKE/contextcheck v1.1.5 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/kulti/thelper v0.6.3 // indirect
	github.com/kunwardeep/paralleltest v1.0.10 // indirect
	github.com/kyoh86/exportloopref v0.1.11 // indirect
	github.com/lasiar/canonicalheader v1.1.2 // indirect
	github.com/ldez/gomoddirectives v0.2.4 // indirect
	github.com/ldez/tagliatelle v0.5.0 // indirect
	github.com/leonklingele/grouper v1.1.2 // indirect
	github.com/linxGnu/grocksdb v1.9.3 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/macabu/inamedparam v0.1.3 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/maratori/testableexamples v1.0.0 // indirect
	github.com/maratori/testpackage v1.1.1 // indirect
	github.com/matoous/godox v0.0.0-20230222163458-006bad1f9d26 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mgechev/revive v1.5.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/moricho/tparallel v0.3.2 // indirect
	github.com/mostynb/zstdpool-freelist v0.0.0-20201229113212-927304c0c3b1 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nakabonne/nestif v0.3.1 // indirect
	github.com/nishanths/exhaustive v0.12.0 // indirect
	github.com/nishanths/predeclared v0.2.2 // indirect
	github.com/nunnatsa/ginkgolinter v0.18.3 // indirect
	github.com/oasisprotocol/curve25519-voi v0.0.0-20230904125328-1f23a7beb09a // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/petermattis/goid v0.0.0-20240813172612-4fcff4a6cae7 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/polyfloyd/go-errorlint v1.7.0 // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.60.1 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/quasilyte/go-ruleguard v0.4.3-0.20240823090925-0fe6f58b47b1 // indirect
	github.com/quasilyte/go-ruleguard/dsl v0.3.22 // indirect
	github.com/quasilyte/gogrep v0.5.0 // indirect
	github.com/quasilyte/regex/syntax v0.0.0-20210819130434-b3f0c404a727 // indirect
	github.com/quasilyte/stdinfo v0.0.0-20220114132959-f7386bf02567 // indirect
	github.com/raeperd/recvcheck v0.1.2 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/ryancurrah/gomodguard v1.3.5 // indirect
	github.com/ryanrolds/sqlclosecheck v0.5.1 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sanposhiho/wastedassign/v2 v2.0.7 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1 // indirect
	github.com/sasha-s/go-deadlock v0.3.5 // indirect
	github.com/sashamelentyev/interfacebloat v1.1.0 // indirect
	github.com/sashamelentyev/usestdlibvars v1.27.0 // indirect
	github.com/securego/gosec/v2 v2.21.4 // indirect
	github.com/shazow/go-diff v0.0.0-20160112020656-b6b7b6733b8c // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sivchari/containedctx v1.0.3 // indirect
	github.com/sivchari/tenv v1.12.1 // indirect
	github.com/sonatard/noctx v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/sourcegraph/go-diff v0.7.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.7.0 // indirect
	github.com/spf13/viper v1.19.0 // indirect
	github.com/ssgreg/nlreturn/v2 v2.2.1 // indirect
	github.com/stbenjam/no-sprintf-host-port v0.1.1 // indirect
	github.com/streamingfast/logging v0.0.0-20230608130331-f22c91403091 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/supranational/blst v0.3.13 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	github.com/tdakkota/asciicheck v0.2.0 // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tetafro/godot v1.4.18 // indirect
	github.com/tidwall/btree v1.7.0 // indirect
	github.com/timakin/bodyclose v0.0.0-20230421092635-574207250966 // indirect
	github.com/timonwong/loggercheck v0.10.1 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/tomarrell/wrapcheck/v2 v2.9.0 // indirect
	github.com/tommy-muehle/go-mnd/v2 v2.5.1 // indirect
	github.com/ultraware/funlen v0.1.0 // indirect
	github.com/ultraware/whitespace v0.1.1 // indirect
	github.com/uudashr/gocognit v1.1.3 // indirect
	github.com/uudashr/iface v1.2.1 // indirect
	github.com/xen0n/gosmopolitan v1.2.2 // indirect
	github.com/yagipy/maintidx v1.0.0 // indirect
	github.com/yeya24/promlinter v0.3.0 // indirect
	github.com/ykadowak/zerologlint v0.1.5 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/zondax/hid v0.9.2 // indirect
	github.com/zondax/ledger-go v0.14.3 // indirect
	gitlab.com/bosi/decorder v0.4.2 // indirect
	go-simpler.org/musttag v0.13.0 // indirect
	go-simpler.org/sloglint v0.7.2 // indirect
	go.etcd.io/bbolt v1.4.0-alpha.0.0.20240404170359-43604f3112c5 // indirect
	go.mongodb.org/mongo-driver v1.12.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/automaxprocs v1.6.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20241108190413-2d47ceb2692f // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/telemetry v0.0.0-20240522233618-39ace7a40ae7 // indirect
	golang.org/x/term v0.27.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	golang.org/x/tools v0.27.0 // indirect
	google.golang.org/genproto v0.0.0-20240722135656-d784300faade // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240930140551-af27646dc61f // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.1 // indirect
	honnef.co/go/tools v0.5.1 // indirect
	mvdan.cc/unparam v0.0.0-20240528143540-8a5130ca722f // indirect
	nhooyr.io/websocket v1.8.6 // indirect
	pgregory.net/rapid v1.1.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
