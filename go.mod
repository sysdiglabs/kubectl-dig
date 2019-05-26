module github.com/sysdiglabs/kubectl-dig

go 1.12

replace (
	github.com/docker/docker => github.com/docker/docker v0.7.3-0.20190327010347-be7ac8be2ae0
	github.com/docker/go-connections => github.com/docker/go-connections v0.3.0
	github.com/docker/go-units => github.com/docker/go-units v0.3.3
)

require (
	cloud.google.com/go v0.39.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Azure/go-autorest v12.1.0+incompatible // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/coreos/etcd v3.3.13+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/dgryski/go-sip13 v0.0.0-20190329191031-25c5027a8c7b // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/elazarl/goproxy v0.0.0-20190421051319-9d40249d3c2f // indirect
	github.com/elazarl/goproxy/ext v0.0.0-20190421051319-9d40249d3c2f // indirect
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/evanphx/json-patch v4.2.0+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/go-openapi/spec v0.19.0 // indirect
	github.com/go-openapi/swag v0.19.0 // indirect
	github.com/golang/mock v1.3.1 // indirect
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/google/pprof v0.0.0-20190515194954-54271f7e092f // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gophercloud/gophercloud v0.0.0-20190525091829-c2d73b246b48 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pty v1.1.4 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20190414153302-2ae31c8b6b30 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/prometheus/common v0.4.1 // indirect
	github.com/prometheus/procfs v0.0.0-20190523193104-a7aeb8df3389 // indirect
	github.com/prometheus/tsdb v0.8.0 // indirect
	github.com/rogpeppe/fastuuid v1.1.0 // indirect
	github.com/russross/blackfriday v2.0.0+incompatible // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/cobra v0.0.4
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f // indirect
	golang.org/x/exp v0.0.0-20190510132918-efd6b22b2522 // indirect
	golang.org/x/image v0.0.0-20190523035834-f03afa92d3ff // indirect
	golang.org/x/lint v0.0.0-20190409202823-959b441ac422 // indirect
	golang.org/x/mobile v0.0.0-20190509164839-32b2708ab171 // indirect
	golang.org/x/oauth2 v0.0.0-20190523182746-aaccbc9213b0 // indirect
	golang.org/x/sys v0.0.0-20190526052359-791d8a0f4d09 // indirect
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/tools v0.0.0-20190525145741-7be61e1b0e51 // indirect
	google.golang.org/appengine v1.6.0 // indirect
	google.golang.org/genproto v0.0.0-20190522204451-c2c4e71fbf69 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gotest.tools v2.2.0+incompatible // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
	k8s.io/api v0.0.0-20190515023547-db5a9d1c40eb
	k8s.io/apimachinery v0.0.0-20190515023456-b74e4c97951f
	k8s.io/cli-runtime v0.0.0-20190515024640-178667528169
	k8s.io/client-go v0.0.0-20190515023709-78e94f51a042
	k8s.io/gengo v0.0.0-20190327210449-e17681d19d3a // indirect
	k8s.io/klog v0.3.1 // indirect
	k8s.io/kube-openapi v0.0.0-20190510232812-a01b7d5d6c22 // indirect
	k8s.io/kubernetes v1.14.2
	k8s.io/utils v0.0.0-20190520173318-324c5df7d3f0
	sigs.k8s.io/structured-merge-diff v0.0.0-20190525122527-15d366b2352e // indirect
)
