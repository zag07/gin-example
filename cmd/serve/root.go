package serve

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	example "github.com/zag07/gin-example"
	"github.com/zag07/gin-example/pkg/transport/http"
)

var cfg string

func NewServeCmd() (serveCmd *cobra.Command) {
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "start app",
		// Run:   start,
	}

	serveCmd.PersistentFlags().StringVar(&cfg, "cfg", "./configs/config.yaml", "config path, eg: -conf config.yaml")

	return serveCmd
}

// startRunE start the web server and initializes the daemon
/*func start(cmd *cobra.Command, args []string) {
	cfg, err := config.Load(cfg)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	ctx := cmd.Context()
}*/

func newApp(logger *zap.Logger, hs *http.Server) *example.App {
	return example.New(
		example.Logger(logger),
		example.Server(hs),
	)
}

func RegisterCommandRecursive(parent *cobra.Command) {
	parent.AddCommand(NewServeCmd())
}
