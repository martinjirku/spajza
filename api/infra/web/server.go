package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/martinjirku/zasobar/adapters/handler"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/infra/db"
	"github.com/martinjirku/zasobar/pkg/web"
)

func InitMiddlewares(r *chi.Mux) {
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	// r.Use(middleware.Timeout(30 * time.Second))
}

func InitServer() *chi.Mux {
	r := chi.NewRouter()
	InitMiddlewares(r)
	user := handler.CreateUserHandler(db.SqlDb, config.DefaultConfiguration)
	units := handler.CreateUnitHandler()
	categories := handler.CreateCategoryHandler(db.SqlDb)
	storagePlaceHandler := handler.CreateStoragePlaceHandler(db.SqlDb)
	storageItemHandler := handler.CreateStorageItemHandler(db.SqlDb)

	r.Route("/api", func(r chi.Router) {
		// Unauthorized routes
		r.Post("/user/login", user.Login)

		// Authorized routes
		r.Group(func(r chi.Router) {
			r.Use(web.JwtMiddlware(config.GetJwtSecret))

			// user
			r.Post("/user/logout", user.Logout)
			r.Post("/user/register", user.Register)
			r.Get("/user/me", user.AboutMe)

			// units
			r.Get("/units", units.List)
			r.Get("/units/{quantity}", units.ListUnitsByQuantity)

			// categories
			r.Get("/categories", categories.ListCategories)
			r.Delete("/categories/{id}", categories.DeleteCategory)
			r.Post("/categories", categories.SaveCategory)
			r.Post("/categories/{id}", categories.SaveCategory)

			// storage place
			r.Post("/storage/places", storagePlaceHandler.CreateStoragePlace)
			r.Get("/storage/places", storagePlaceHandler.ListStoragePlace)
			r.Post("/storage/places/{id}", storagePlaceHandler.UpdateStoragePlace)
			r.Get("/storage/places/{id}", storagePlaceHandler.GetStoragePlace)
			r.Delete("/storage/places/{id}", storagePlaceHandler.DeleteStoragePlace)

			// storage item
			r.Get("/storage/items", storageItemHandler.List)
			r.Post("/storage/items", storageItemHandler.CreateStorageItem)
			r.Post("/storage/items/{id}/consumpt", storageItemHandler.Consumpt)
			r.Post("/storage/items/{id}/{fieldName}", storageItemHandler.UpdateField)
		})
	})
	log.Printf("Application start, listening on %s:%s", config.DefaultConfiguration.Domain, config.DefaultConfiguration.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.DefaultConfiguration.Domain, config.DefaultConfiguration.Port), r)
	return r
}
