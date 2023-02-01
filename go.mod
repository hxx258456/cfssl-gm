module github.com/hxx258456/cfssl-gm

go 1.17

require (
	gitee.com/zhaochuninhefei/zcgolog v0.0.16
	github.com/GeertJohan/go.rice v1.0.2
	github.com/cloudflare/backoff v0.0.0-20161212185259-647f3cdfc87a
	github.com/cloudflare/go-metrics v0.0.0-20151117154305-6a9aea36fb41
	github.com/cloudflare/redoctober v0.0.0-20211013234631-6a74ccc611f6
	github.com/google/certificate-transparency-go v1.1.4
	github.com/hxx258456/ccgo v0.0.4
	github.com/jmhodges/clock v1.2.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/kisielk/sqlstruct v0.0.0-20210630145711-dae28ed37023
	github.com/kisom/goutils v1.4.3
	github.com/lib/pq v1.10.7
	github.com/mattn/go-sqlite3 v1.14.15
	github.com/zmap/zcrypto v0.0.0-20190729165852-9051775e6a2e
	github.com/zmap/zlint v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.1.0
)

require (
	github.com/daaku/go.zipexe v1.0.0 // indirect
	github.com/getsentry/sentry-go v0.11.0 // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/weppos/publicsuffix-go v0.15.1-0.20220329081811-9a40b608a236 // indirect
	golang.org/x/net v0.1.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
)

replace (
	// zlint与zcrypto版本必须匹配，否则zlint编译出错
	github.com/zmap/zcrypto => github.com/zmap/zcrypto v0.0.0-20190729165852-9051775e6a2e
	github.com/zmap/zlint => github.com/zmap/zlint v0.0.0-20190806154020-fd021b4cfbeb
// github.com/hxx258456/ccgo => /home/zhaochun/work/sources/gm-zc/gmgo
)
