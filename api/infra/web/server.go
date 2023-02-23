package web

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/martinjirku/zasobar/adapters/handler"
	"github.com/martinjirku/zasobar/adapters/repository"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/infra/db"
	"github.com/martinjirku/zasobar/pkg/web"
	"github.com/martinjirku/zasobar/usecase"
)

func InitMiddlewares(r *chi.Mux) {
	r.Use(middleware.AllowContentType("application/json", "application/x-www-form-urlencoded"))
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	// r.Use(middleware.Timeout(30 * time.Second))
}

func InitServer() *chi.Mux {
	r := chi.NewRouter()
	InitMiddlewares(r)
	user := handler.CreateUserHandler(db.SqlDb, config.GetConfiguration())
	units := handler.CreateUnitHandler()
	categories := handler.CreateCategoryHandler(CategoryUsecaseProvider)
	storagePlaceHandler := handler.CreateStoragePlaceHandler(db.SqlDb)
	storageItemHandler := handler.CreateStorageItemHandler(StorageItemUsecaseProvider)

	r.Route("/api", func(r chi.Router) {
		// Unauthorized routes
		r.Post("/user/login", user.Login)
		r.Post("/user/auth/google", user.GoogleLogin)

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
	c := config.GetConfiguration()
	log.Printf("Application start, listening on %s:%s", c.Domain, c.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", c.Domain, c.Port), r)
	return r
}

func CategoryUsecaseProvider(ctx context.Context) *usecase.CategoryUsecase {
	repository := repository.NewCategoryRepository(ctx, db.SqlDb)
	usecase := usecase.CreateCategoryUsecase(repository)
	return usecase
}
func StorageItemUsecaseProvider(ctx context.Context) *usecase.StorageItemUsecase {
	repository := repository.NewStorageItemRepository(ctx, db.SqlDb)
	usecase := usecase.NewStorageItemUsecase(repository)
	return usecase
}
