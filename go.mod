module github.com/shantanubansal/mongo-tools

go 1.14

replace gopkg.in/mgo.v2 8c133fd1d0fc316302df84fc12769e17fb473bb1 => github.com/10gen/mgo v0.0.0-20181212170345-8c133fd1d0fc

require (
	github.com/10gen/mgo v0.0.0-20181212170345-8c133fd1d0fc
	github.com/3rf/mongo-lint v0.0.0-20140604191638-3550fdcf1f43
	github.com/aws/aws-sdk-go v1.29.11
	github.com/go-stack/stack v1.8.0
	github.com/golang/snappy v0.0.0-20160529050041-d9eb7a3d35ec
	github.com/google/go-cmp v0.3.0
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c
	github.com/jessevdk/go-flags v1.4.0
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/jtolds/gls v4.20.0+incompatible
	github.com/klauspost/compress v1.10.1
	github.com/mattn/go-runewidth v0.0.4
	github.com/mongodb/mongo-tools-common v4.0.2+incompatible
	github.com/nsf/termbox-go v0.0.0-20160718140619-0723e7c3d0a3
	github.com/pkg/errors v0.9.1
	github.com/smartystreets/assertions v0.0.0-20160205033931-287b4346dc4e
	github.com/smartystreets/goconvey v1.6.1-0.20160205033552-bf58a9a12912
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c
	github.com/xdg/stringprep v1.0.1-0.20180714160509-73f8eece6fdc
	go.mongodb.org/mongo-driver v1.4.0-rc0
	golang.org/x/crypto v0.0.0-20190530122614-20be4c3c3ed5
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527
	golang.org/x/text v0.3.2
	gopkg.in/tomb.v2 v2.0.0-20140626144623-14b3d72120e8
)
