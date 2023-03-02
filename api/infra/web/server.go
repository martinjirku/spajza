package web

import (
	"context"
	"database/sql"
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

func InitServer(config config.Configuration) *chi.Mux {
	r := chi.NewRouter()
	InitMiddlewares(r)
	db := db.NewDB(config.DB)
	user := handler.CreateUserHandler(db, config)
	units := handler.CreateUnitHandler()
	categories := handler.CreateCategoryHandler(GetCategoryUsecaseProvider(db))
	storagePlaceHandler := handler.CreateStoragePlaceHandler(db)
	storageItemHandler := handler.CreateStorageItemHandler(GetStorageItemUsecaseProvider(db))

	r.Route("/api", func(r chi.Router) {
		// Unauthorized routes
		r.Post("/user/login", user.Login)
		r.Post("/user/auth/google", user.GoogleLogin)

		// Authorized routes
		r.Group(func(r chi.Router) {
			r.Use(web.JwtMiddlware(func() string { return config.Jwt.Secret }))

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
	log.Printf("Application start, listening on %s:%s", config.Domain, config.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.Domain, config.Port), r)
	return r
}

func GetCategoryUsecaseProvider(db *sql.DB) handler.CategoryUsecaseProvider {
	return func(ctx context.Context) *usecase.CategoryUsecase {
		repository := repository.NewCategoryRepository(ctx, db)
		usecase := usecase.CreateCategoryUsecase(repository)
		return usecase
	}

}
func GetStorageItemUsecaseProvider(db *sql.DB) handler.StorageItemUsecaseProvider {
	return func(ctx context.Context) *usecase.StorageItemUsecase {
		repository := repository.NewStorageItemRepository(ctx, db)
		usecase := usecase.NewStorageItemUsecase(repository)
		return usecase
	}
}
