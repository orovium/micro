package main

import (
	"github.com/micro/go-micro"

	"github.com/orovium/micro/microserver"
)

var log = microserver.GetLogger()

func main() {

	/* //Attaching one option
	serviceOptions := microserver.NewOptions()
	serviceOptions.Logger(microserver.DefaultLoggerOptions())
	*/

	/* //Get default options
	serviceOptions := microserver.NewOptions().WithDefaultOptions()
	*/

	/* // Initialize a service with custom options
	microserver.Init(serviceOptions)
	*/

	// Initialize a service with default options attached
	microserver.InitDefault()

	service := microserver.StartDefaultService(
		micro.Name("ping"),
		micro.Version("1.0.0"),
		micro.Address(":8091"),
	)

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")

	if err := service.Run(); err != nil {
		log.Error(err)
	}
}
