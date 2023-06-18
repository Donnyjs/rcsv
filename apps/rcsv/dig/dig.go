package dig

import (
	"go.uber.org/dig"
	"rcsv/apps/rcsv/internal/config"
	"rcsv/apps/rcsv/internal/service/svc_inscription"
	"rcsv/domain/cache"
	"rcsv/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(svc_inscription.NewInscriptionService)
	container.Provide(repo.NewInscriptionRepository)
	container.Provide(cache.NewInscriptionCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
