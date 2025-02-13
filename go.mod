module github.com/kopia/kopia

go 1.16

require (
	cloud.google.com/go/storage v1.16.0
	contrib.go.opencensus.io/exporter/prometheus v0.3.0
	github.com/Azure/azure-pipeline-go v0.2.3
	github.com/Azure/azure-storage-blob-go v0.14.0
	github.com/GehirnInc/crypt v0.0.0-20200316065508-bb7000b8a962 // indirect
	github.com/alecthomas/kingpin v0.0.0-20200323085623-b6657d9477a6 // this is pulling master, which is newer than v2
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15
	github.com/aws/aws-sdk-go v1.40.17
	github.com/chmduquesne/rollinghash v4.0.0+incompatible
	github.com/dustinkirkland/golang-petname v0.0.0-20191129215211-8e5a1ed0cff0
	github.com/efarrer/iothrottler v0.0.1
	github.com/fatih/color v1.12.0
	github.com/foomo/htpasswd v0.0.0-20200116085101-e3a90e78da9c
	github.com/frankban/quicktest v1.10.2 // indirect
	github.com/gofrs/flock v0.8.1
	github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/golang/protobuf v1.5.2
	github.com/google/fswalker v0.2.1-0.20200214223026-f0e929ba4126
	github.com/google/go-cmp v0.5.6
	github.com/google/readahead v0.0.0-20161222183148-eaceba169032 // indirect
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/hanwen/go-fuse/v2 v2.0.4-0.20210104155004-09a3c381714c
	github.com/klauspost/compress v1.13.3
	github.com/klauspost/cpuid/v2 v2.0.5 // indirect
	github.com/klauspost/pgzip v1.2.5
	github.com/kr/text v0.2.0 // indirect
	github.com/kylelemons/godebug v1.1.0
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/minio-go/v7 v7.0.11-0.20210302210017-6ae69c73ce78
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/natefinch/atomic v0.0.0-20200526193002-18c0533a5b09
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pierrec/lz4 v2.6.1+incompatible
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.6.0
	github.com/pkg/sftp v1.13.2
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/sanity-io/litter v1.5.1
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/studio-b12/gowebdav v0.0.0-20210630100626-7ff61aa87be8
	github.com/tg123/go-htpasswd v1.0.0
	github.com/zalando/go-keyring v0.1.1
	github.com/zeebo/blake3 v0.2.0
	go.opencensus.io v0.23.0
	gocloud.dev v0.23.0
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf
	golang.org/x/exp v0.0.0-20210220032938-85be41e4509f
	golang.org/x/mod v0.4.2
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/oauth2 v0.0.0-20210628180205-a41e5a781914
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
	golang.org/x/term v0.0.0-20210317153231-de623e64d2a6
	google.golang.org/api v0.52.0
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/kothar/go-backblaze.v0 v0.0.0-20210124194846-35409b867216
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
