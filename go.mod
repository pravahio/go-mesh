module github.com/pravahio/go-mesh

go 1.12

require (
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/apilayer/freegeoip v3.5.0+incompatible // indirect
	github.com/aristanetworks/goarista v0.0.0-20190712234253-ed1100a1c015 // indirect
	github.com/bluele/gcache v0.0.0-20190518031135-bc40bd653833
	github.com/cespare/cp v1.1.1 // indirect
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/elastic/gosigar v0.10.5 // indirect
	github.com/ethereum/go-ethereum v1.9.2
	github.com/fjl/memsize v0.0.0-20180929194037-2a09253e352a // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20190607065134-2772fd86a8ff // indirect
	github.com/golang/mock v1.3.1 // indirect
	github.com/golang/protobuf v1.3.3
	github.com/graph-gophers/graphql-go v0.0.0-20190902214650-641ae197eec7 // indirect
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/improbable-eng/grpc-web v0.11.0
	github.com/influxdata/influxdb v1.7.8 // indirect
	github.com/ipfs/go-ipfs-addr v0.0.1
	github.com/ipfs/go-log v0.0.1
	github.com/karalabe/usb v0.0.0-20190819132248-550797b1cad8 // indirect
	github.com/libp2p/go-libp2p v0.5.1
	github.com/libp2p/go-libp2p-autonat-svc v0.1.0
	github.com/libp2p/go-libp2p-circuit v0.1.4
	github.com/libp2p/go-libp2p-core v0.3.0
	github.com/libp2p/go-libp2p-discovery v0.2.0
	github.com/libp2p/go-libp2p-kad-dht v0.5.0
	github.com/libp2p/go-libp2p-peerstore v0.1.4
	github.com/libp2p/go-libp2p-pubsub v0.1.0
	github.com/libp2p/go-libp2p-quic-transport v0.2.3
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/olekukonko/tablewriter v0.0.1 // indirect
	github.com/oschwald/maxminddb-golang v1.4.0 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pravahio/go-auth-provider v0.0.0
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/status-im/keycard-go v0.0.0-20190424133014-d95853db0f48 // indirect
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570 // indirect
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/tyler-smith/go-bip39 v1.0.2 // indirect
	github.com/urfave/cli v1.22.0
	github.com/wsddn/go-ecdh v0.0.0-20161211032359-48726bab9208 // indirect
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
	golang.org/x/sys v0.0.0-20200202164722-d101bd2416d5 // indirect
	google.golang.org/genproto v0.0.0-20200204135345-fa8e72b47b90 // indirect
	google.golang.org/grpc v1.27.0
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20190709231704-1e4459ed25ff // indirect
	gopkg.in/urfave/cli.v1 v1.20.0 // indirect
)

replace github.com/libp2p/go-libp2p-pubsub v0.1.0 => github.com/upperwal/go-libp2p-pubsub v0.1.1-0.20190822125434-affd4e4c6c42

replace github.com/pravahio/go-auth-provider v0.0.0 => ../auth/go-auth-provider
