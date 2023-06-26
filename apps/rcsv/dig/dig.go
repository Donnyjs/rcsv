package dig

import (
	"go.uber.org/dig"
	"rcsv/apps/rcsv/config"
	"rcsv/apps/rcsv/internal/service/svc_inscription"
	"rcsv/apps/rcsv/internal/service/svc_inscription_monitor"
	"rcsv/domain/cache"
	"rcsv/domain/oss"
	"rcsv/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(svc_inscription.NewInscriptionService)
	container.Provide(repo.NewInscriptionRepository)
	container.Provide(cache.NewInscriptionCache)
	container.Provide(oss.NewInscriptionOss)
	container.Provide(svc_inscription_monitor.NewInscriptionMonitor)
	container.Invoke(func(im *svc_inscription_monitor.InscriptionMonitor) {
		go im.Run()
	})
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
