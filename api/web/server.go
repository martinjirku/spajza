package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/martinjirku/zasobar/config"
	spajzaMiddleware "github.com/martinjirku/zasobar/web/middleware"
)

func InitMiddlewares(r *chi.Mux) {
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	// r.Use(middleware.Timeout(30 * time.Second))
	r.Use(spajzaMiddleware.JwtMiddleware)
}

func InitServer() *chi.Mux {
	r := chi.NewRouter()
	InitMiddlewares(r)
	user := createUserHandler()
	units := createUnitHandler()
	categories := createCategoryHandler()
	storagePlaceHandler := createStoragePlaceHandler()
	storageItemHandler := createStorageItemHandler()

	r.Route("/api", func(r chi.Router) {
		// Unauthorized routes
		r.Post("/user/login", user.login)

		// Authorized routes
		r.Group(func(r chi.Router) {
			r.Use(spajzaMiddleware.JwtMiddleware)

			// user
			r.Post("/user/logout", user.logout)
			r.Post("/user/register", user.register)
			r.Get("/user/me", user.aboutMe)

			// units
			r.Get("/units", units.list)
			r.Get("/units/{quantity}", units.listUnitsByQuantity)

			// categories
			r.Get("/categories", categories.listCategories)
			r.Delete("/categories/{id}", categories.deleteCategory)
			r.Post("/categories", categories.saveCategory)
			r.Post("/categories/{id}", categories.saveCategory)

			// storage place
			r.Post("/storage/places", storagePlaceHandler.createStoragePlace)
			r.Get("/storage/places", storagePlaceHandler.listStoragePlace)
			r.Post("/storage/places/{id}", storagePlaceHandler.updateStoragePlace)
			r.Get("/storage/places/{id}", storagePlaceHandler.getStoragePlace)
			r.Delete("/storage/places/{id}", storagePlaceHandler.deleteStoragePlace)

			// storage item
			r.Get("/storage/items", storageItemHandler.list)
			r.Post("/storage/items", storageItemHandler.createStorageItem)
			r.Post("/storage/items/{id}/title", storageItemHandler.updateTitle)
			r.Post("/storage/items/{id}/consumpt", storageItemHandler.consumpt)
		})
	})
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.DefaultConfiguration.Domain, config.DefaultConfiguration.Port), r)
	return r
}
