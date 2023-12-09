package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/lamoda-tech/loikx/internal/config"
	"github.com/lamoda-tech/loikx/internal/product/domain"
	"github.com/lamoda-tech/loikx/internal/product/handlers"
	"github.com/lamoda-tech/loikx/internal/product/repository"
	"github.com/lamoda-tech/loikx/internal/product/usecases"
	domain2 "github.com/lamoda-tech/loikx/internal/warehouse/domain"
	handlers2 "github.com/lamoda-tech/loikx/internal/warehouse/handlers"
	repository2 "github.com/lamoda-tech/loikx/internal/warehouse/repository"
	usecases2 "github.com/lamoda-tech/loikx/internal/warehouse/usecases"
	"github.com/lamoda-tech/loikx/pkg/server"
)

type (
	Server interface {
		Serve() error
	}

	ProductsRepository interface {
		Release(ctx context.Context, ids []uuid.UUID) error
		Reserve(ctx context.Context, ids []uuid.UUID) error
	}
)

type App struct {
	config *config.Config

	server Server

	router http.Handler

	connection *pgx.Conn

	releaseHandler         *handlers.ReleaseProductHandler
	reserveHandler         *handlers.ReserveProductHandler
	findByWarehouseHandler *handlers.FindByWarehouseHandler
	createProductHandler   *handlers.CreateProductHandler
	deleteProductHandler   *handlers.DeleteProductHandler
	createWarehouseHandler *handlers2.CreateWarehouseHandler
	deleteWarehouseHandler *handlers2.DeleteWarehouseHandler

	releaseUseCase         *usecases.ReleaseUseCase
	reserveUseCase         *usecases.ReserveUseCase
	findByWarehouseUseCase *usecases.FindByWarehouseUseCase
	createProductUseCase   *usecases.CreateProductUseCase
	deleteProductUseCase   *usecases.DeleteProductUseCase
	createWarehouseUseCase *usecases2.CreateWarehouseUseCase
	deleteWarehouseUseCase *usecases2.DeleteWarehouseUseCase

	productsRepository  domain.ProductRepository
	warehouseRepository domain2.WarehouseRepository
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() error {
	log.Printf("server starts on: %s:%d\n", a.config.Address, a.config.Port)

	return a.server.Serve()
}

func (a *App) Init(ctx context.Context) error {
	cfg, err := config.LoadFromEnvironment()
	if err != nil {
		return fmt.Errorf("load from environment: %w", err)
	}

	a.config = cfg

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	if err = a.initServer(ctx, logger); err != nil {
		return fmt.Errorf("init server: %w", err)
	}

	return nil
}

func (a *App) initServer(ctx context.Context, logger *log.Logger) error {
	var err error
	a.connection, err = pgx.Connect(ctx, a.config.DataBaseURL)
	if err != nil {
		return fmt.Errorf("init database: %w", err)
	}

	a.productsRepository = repository.NewProductRepository(a.connection)
	a.warehouseRepository = repository2.NewWarehouseRepository(a.connection)

	a.releaseUseCase = usecases.NewReleaseUseCase(a.productsRepository)
	a.reserveUseCase = usecases.NewReserveUseCase(a.productsRepository)
	a.findByWarehouseUseCase = usecases.NewFindByWarehouseUseCase(a.productsRepository)
	a.createProductUseCase = usecases.NewCreateProductUseCase(a.productsRepository)
	a.deleteProductUseCase = usecases.NewDeleteProductUseCase(a.productsRepository)
	a.createWarehouseUseCase = usecases2.NewCreateWarehouseUseCase(a.warehouseRepository)
	a.deleteWarehouseUseCase = usecases2.NewDeleteWarehouseUseCase(a.warehouseRepository)

	a.releaseHandler = handlers.NewReleaseProductHandler(a.releaseUseCase)
	a.reserveHandler = handlers.NewReserveProductHandler(a.reserveUseCase)
	a.findByWarehouseHandler = handlers.NewFindByWarehouseHandler(a.findByWarehouseUseCase)
	a.createProductHandler = handlers.NewCreateProductHandler(a.createProductUseCase)
	a.deleteProductHandler = handlers.NewDeleteProductHandler(a.deleteProductUseCase)
	a.createWarehouseHandler = handlers2.NewCreateWarehouseHandler(a.createWarehouseUseCase)
	a.deleteWarehouseHandler = handlers2.NewDeleteWarehouseHandler(a.deleteWarehouseUseCase)

	router := a.createRouter()
	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	a.server = server.NewServer(
		fmt.Sprintf("%s:%d", a.config.Address, a.config.Port),
		tracing(nextRequestID)(logging(logger)(router)),
	)

	return nil
}

func (a *App) createRouter() http.Handler {
	router := mux.NewRouter()

	router.Handle("/api/products/reserve", a.reserveHandler).Methods(http.MethodPatch)
	router.Handle("/api/products/release", a.releaseHandler).Methods(http.MethodPatch)
	router.Handle("/api/products/find-by-warehouse/{id}", a.findByWarehouseHandler).Methods(http.MethodGet)
	router.Handle("/api/products/create", a.createProductHandler).Methods(http.MethodPost)
	router.Handle("/api/products/delete/{id}", a.deleteProductHandler).Methods(http.MethodDelete)

	router.Handle("/api/warehouse/create", a.createWarehouseHandler).Methods(http.MethodPost)
	router.Handle("/api/warehouse/delete/{id}", a.deleteWarehouseHandler).Methods(http.MethodDelete)

	return router
}

const requestIDKey = 0

func tracing(nextRequestID func() string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}
