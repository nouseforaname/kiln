package main

import (
	"log"
	"os"

	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/kiln/builder"
	"github.com/pivotal-cf/kiln/commands"
	"github.com/pivotal-cf/kiln/fetcher"
	"github.com/pivotal-cf/kiln/helper"
	"github.com/pivotal-cf/kiln/internal/baking"
	"github.com/pivotal-cf/kiln/internal/providers"
)

var version = "unknown"
var assetsFile = "unknown"
var releasesDir = "unknown"

func main() {
	errLogger := log.New(os.Stderr, "", 0)
	outLogger := log.New(os.Stdout, "", 0)

	var global struct {
		Help    bool `short:"h" long:"help"    description:"prints this usage information"   default:"false"`
		Version bool `short:"v" long:"version" description:"prints the kiln release version" default:"false"`
	}

	args, err := jhanda.Parse(&global, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	globalFlagsUsage, err := jhanda.PrintUsage(global)
	if err != nil {
		log.Fatal(err)
	}

	var command string
	if len(args) > 0 {
		command, args = args[0], args[1:]
	}

	if global.Version {
		command = "version"
	}

	if global.Help {
		command = "help"
	}

	if command == "" {
		command = "help"
	}

	filesystem := helper.NewFilesystem()
	zipper := builder.NewZipper()
	interpolator := builder.NewInterpolator()
	tileWriter := builder.NewTileWriter(filesystem, &zipper, errLogger)

	releaseManifestReader := builder.NewReleaseManifestReader()
	releasesService := baking.NewReleasesService(errLogger, releaseManifestReader)

	stemcellManifestReader := builder.NewStemcellManifestReader(filesystem)
	stemcellService := baking.NewStemcellService(errLogger, stemcellManifestReader)

	templateVariablesService := baking.NewTemplateVariablesService()

	boshVariableDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	boshVariablesService := baking.NewBOSHVariablesService(errLogger, boshVariableDirectoryReader)

	formDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	formsService := baking.NewFormsService(errLogger, formDirectoryReader)

	instanceGroupDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	instanceGroupsService := baking.NewInstanceGroupsService(errLogger, instanceGroupDirectoryReader)

	jobsDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	jobsService := baking.NewJobsService(errLogger, jobsDirectoryReader)

	propertiesDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	propertiesService := baking.NewPropertiesService(errLogger, propertiesDirectoryReader)

	runtimeConfigsDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	runtimeConfigsService := baking.NewRuntimeConfigsService(errLogger, runtimeConfigsDirectoryReader)

	iconService := baking.NewIconService(errLogger)

	metadataService := baking.NewMetadataService()
	checksummer := baking.NewChecksummer(errLogger)

	s3Provider := providers.NewS3Provider()

	downloader := fetcher.NewDownloader(outLogger, s3Provider)
	releaseMatcher := fetcher.NewReleaseMatcher(s3Provider)
	localReleaseDirectory := fetcher.NewLocalReleaseDirectory(releasesService)

	commandSet := jhanda.CommandSet{}
	commandSet["help"] = commands.NewHelp(os.Stdout, globalFlagsUsage, commandSet)
	commandSet["version"] = commands.NewVersion(outLogger, version)
	commandSet["fetch"] = commands.NewFetch(outLogger, downloader, releaseMatcher, localReleaseDirectory)
	commandSet["bake"] = commands.NewBake(
		interpolator,
		tileWriter,
		outLogger,
		templateVariablesService,
		boshVariablesService,
		releasesService,
		stemcellService,
		formsService,
		instanceGroupsService,
		jobsService,
		propertiesService,
		runtimeConfigsService,
		iconService,
		metadataService,
		checksummer,
	)

	err = commandSet.Execute(command, args)
	if err != nil {
		log.Fatal(err)
	}
}
