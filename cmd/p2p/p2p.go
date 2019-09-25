package main

import (
	"fmt"
	"github.com/spacemeshos/go-spacemesh/activation"
	"github.com/spacemeshos/go-spacemesh/api"
	cmdp "github.com/spacemeshos/go-spacemesh/cmd"
	"github.com/spacemeshos/go-spacemesh/database"
	"github.com/spacemeshos/go-spacemesh/log"
	"github.com/spacemeshos/go-spacemesh/metrics"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/service"
	"github.com/spacemeshos/post/config"
	"github.com/spf13/cobra"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// Cmd is the p2p cmd
var Cmd = &cobra.Command{
	Use:   "p2p",
	Short: "start p2p",
	Run: func(cmd *cobra.Command, args []string) {
		p2pApp := NewP2PApp()
		defer p2pApp.Cleanup()

		p2pApp.Initialize(cmd)
		p2pApp.Start(cmd, args)
	},
}

func init() {
	cmdp.AddCommands(Cmd)
}

type P2PApp struct {
	*cmdp.BaseApp
	p2p p2p.Service
}

func NewP2PApp() *P2PApp {
	return &P2PApp{BaseApp: cmdp.NewBaseApp()}
}

func (app *P2PApp) Cleanup() {
	// TODO: move to array of cleanup functions and execute all here
}

func (app *P2PApp) Start(cmd *cobra.Command, args []string) {
	// init p2p services
	log.JSONLog(true)
	//log.DebugMode(true)

	log.Event().Info("Starting Spacemesh")
	if app.Config.MemProfile != "" {
		log.Info("Starting mem profiling")
		f, err := os.Create(app.Config.MemProfile)
		if err != nil {
			log.Error("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Error("could not write memory profile: ", err)
		}
	}

	if app.Config.CpuProfile != "" {
		log.Info("Starting cpu profile")
		f, err := os.Create(app.Config.CpuProfile)
		if err != nil {
			log.Error("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Error("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if app.Config.PprofHttpServer {
		log.Info("Starting pprof server")
		go func() {
			err := http.ListenAndServe(":6060", nil)
			if err != nil {
				log.Error("cannot start http server", err)
			}
		}()
	}

	log.Info("Initializing P2P services")

	swarm, err := p2p.New(cmdp.Ctx, app.Config.P2P)
	if err != nil {
		log.Panic("Error init p2p services, err: %v", err)
	}
	app.p2p = swarm

	// Testing stuff
	api.ApproveAPIGossipMessages(cmdp.Ctx, app.p2p)
	metrics.StartCollectingMetrics(app.Config.MetricsPort)

	somelogger := log.NewDefault(fmt.Sprintf("p2p.test.%v", swarm.LocalNode().PublicKey().String()))

	dbStorepath := fmt.Sprintf("./data/%v", config.DefaultDataDirName)

	poetDbStore, err := database.NewLDBDatabase(dbStorepath+"poet", 0, 0, somelogger.WithName("poetDbStore"))
	if err != nil {
		panic(err)
	}

	poetDb := activation.NewPoetDb(poetDbStore, somelogger.WithName("poetDb"))
	poetListener := activation.NewPoetListener(swarm, poetDb, somelogger.WithName("poetListener"))
	poetListener.Start()

	for a := 'a'; a < 10; a++ {
		c := app.p2p.RegisterGossipProtocol(string(a))
		go func(a string, c chan service.GossipMessage) {
			log.Info("Registered protocol %v, starting to approve messages", a)
			for {
				select {
				case msg := <-c:
					msg.ReportValidation(a)
				case <-cmdp.Ctx.Done():
					return
				}
			}
		}(string(a), c)
	}

	// start the node
	err = app.p2p.Start()
	defer app.p2p.Shutdown()

	if err != nil {
		log.Panic("Error starting p2p services, err: %v", err)
	}

	// start api servers
	if app.Config.API.StartGrpcServer || app.Config.API.StartJSONServer {
		// start grpc if specified or if json rpc specified
		log.Info("Started the GRPC Service")
		grpc := api.NewGrpcService(app.p2p, nil, nil, nil, nil, nil)
		grpc.StartService()
	}

	if app.Config.API.StartJSONServer {
		log.Info("Started the JSON Service")
		json := api.NewJSONHTTPServer()
		json.StartService()
	}

	go func() {
		time.Sleep(10 * time.Second)
		for {
			select {
			case <-cmdp.Ctx.Done():
				return
			default:
				break
			}
			log.Info("Propagating messages")
			for a := 'a'; a < 10; a++ {
				err := app.p2p.Broadcast(string(a), []byte(api.RandString(100000)))
				log.Error("Err broadcasting random message err:%v", err)
			}
			time.Sleep(10 * time.Second)
		}
	}()

	<-cmdp.Ctx.Done()
}

func main() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
