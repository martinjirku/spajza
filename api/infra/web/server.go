package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/martinjirku/zasobar/adapters/handler"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/infra/db"
	spajzaMiddleware "github.com/martinjirku/zasobar/infra/web/middleware"
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
	units := createUnitHandler()
	categories := createCategoryHandler()
	storagePlaceHandler := handler.CreateStoragePlaceHandler(db.SqlDb)
	storageItemHandler := createStorageItemHandler()

	r.Route("/api", func(r chi.Router) {
		// Unauthorized routes
		r.Post("/user/login", user.Login)

		// Authorized routes
		r.Group(func(r chi.Router) {
			r.Use(spajzaMiddleware.JwtMiddleware)

			// user
			r.Post("/user/logout", user.Logout)
			r.Post("/user/register", user.Register)
			r.Get("/user/me", user.AboutMe)

			// units
			r.Get("/units", units.list)
			r.Get("/units/{quantity}", units.listUnitsByQuantity)

			// categories
			r.Get("/categories", categories.listCategories)
			r.Delete("/categories/{id}", categories.deleteCategory)
			r.Post("/categories", categories.saveCategory)
			r.Post("/categories/{id}", categories.saveCategory)

			// storage place
			r.Post("/storage/places", storagePlaceHandler.CreateStoragePlace)
			r.Get("/storage/places", storagePlaceHandler.ListStoragePlace)
			r.Post("/storage/places/{id}", storagePlaceHandler.UpdateStoragePlace)
			r.Get("/storage/places/{id}", storagePlaceHandler.GetStoragePlace)
			r.Delete("/storage/places/{id}", storagePlaceHandler.DeleteStoragePlace)

			// storage item
			r.Get("/storage/items", storageItemHandler.list)
			r.Post("/storage/items", storageItemHandler.createStorageItem)
			r.Post("/storage/items/{id}/consumpt", storageItemHandler.consumpt)
			r.Post("/storage/items/{id}/{fieldName}", storageItemHandler.updateField)
		})
	})
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.DefaultConfiguration.Domain, config.DefaultConfiguration.Port), r)
	return r
}
