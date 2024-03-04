package cli

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

func GetJfrogCliPipelinesApp() components.App {
	app := components.CreateEmbeddedApp(
		"pipelines",
		getCommands(),
	)
	return app
}
