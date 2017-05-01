package main

import (
	"k8s.io/gengo/args"
	"k8s.io/kubernetes/cmd/libs/go2idl/informer-gen/generators"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

func main() {
	customArgs := &generators.CustomArgs{}
	arguments := &args.GeneratorArgs{
		OutputBase:        args.DefaultSourceTree(),
		GoHeaderFilePath:  "hack/boilerplate.txt",
		GeneratedBuildTag: "ignore_autogenerated_openshift",
		CustomArgs:        customArgs,
	}
	arguments.AddFlags(pflag.CommandLine)
	customArgs.AddFlags(pflag.CommandLine)

	// Run it.
	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
	glog.V(2).Info("Completed successfully.")
}
