package hook

import (
	"os"
	"os/signal"
)

type ShutdownHooker interface {
	BeforShutdown()
	Shutdown()
	AfterShutdown()
}

type App struct {
	hooks []ShutdownHooker
	stop  chan bool
}

func NewApp() *App {
	return &App{
		hooks: make([]ShutdownHooker, 5),
	}
}

func (s *App) InstallShutdownHook(hook ShutdownHooker) {
	s.hooks = append(s.hooks, hook)
}

func (s *App) Start() {
	c := make(chan os.Signal)
	//监听所有信号
	signal.Notify(c)

	for {
		select {
		case sig := <-c:
			switch sig {
			case os.Interrupt:
				os.Exit(0)
			}
		case <-s.stop:
			return
		}
	}
}

func (s *App) Stop() {
	s.stop <- true
}
