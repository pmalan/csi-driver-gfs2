package main

import (
	"flag"
	"os"

	"github.com/psmalan/csi-driver-gfs2/pkg/gfs2"
	"k8s.io/klog/v2"
)

var (
	endpoint              = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	nodeID                = flag.String("nodeid", "", "node id")
	mountPermissions      = flag.Uint64("mount-permissions", 0, "mounted folder permissions")
	driverName            = flag.String("drivername", "GFS2Driver", "name of the driver")
	workingMountDir       = flag.String("working-mount-dir", "/tmp", "working directory for provisioner to mount gfs2 shares temporarily")
	defaultOnDeletePolicy = flag.String("default-ondelete-policy", "", "default policy for deleting subdirectory when deleting a volume")
)

func init() {
	_ = flag.Set("logtostderr", "true")
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	if *nodeID == "" {
		klog.Warning("nodeid is empty")
	}

	handle()
	os.Exit(0)
}

func handle() {
	driverOptions := gfs2.DriverOptions{
		NodeID:                *nodeID,
		DriverName:            *driverName,
		Endpoint:              *endpoint,
		MountPermissions:      *mountPermissions,
		WorkingMountDir:       *workingMountDir,
		DefaultOnDeletePolicy: *defaultOnDeletePolicy,
	}
	d := gfs2.NewDriver(&driverOptions)
	d.Run(false)
}
