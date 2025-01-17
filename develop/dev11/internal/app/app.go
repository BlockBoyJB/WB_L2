package app

import (
	"event-server/config"
	v1 "event-server/internal/api/v1"
	"event-server/internal/repo"
	"event-server/internal/service"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("init config error:", err)
	}

	repos := repo.NewRepositories()

	d := &service.ServicesDependencies{
		Repos: repos,
	}
	services := service.NewServices(d)

	v1.NewRouter(services)

	go func() {
		if err := http.ListenAndServe(net.JoinHostPort("", cfg.HTTP.Port), nil); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("app serving http connections on port", cfg.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
}
