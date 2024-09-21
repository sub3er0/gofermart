package main

import (
	"github.com/go-chi/chi/v5"
	"gofermart/db"
	"gofermart/internal/config"
	"gofermart/internal/handlers"
	"gofermart/internal/middleware"
	"gofermart/internal/repository"
	"gofermart/internal/service"
	"log"
	"net/http"
)

func main() {
	PgsStorage := &db.PgStorage{}
	cfg, err := config.InitConfig()

	if err != nil {
		log.Fatalf("Error while initializing config: %v", err)
	}

	err = PgsStorage.Init(cfg.DatabaseDsn)

	if err != nil {
		log.Fatalf("Error while initializing db connection: %v", err)
	}

	defer PgsStorage.Close()

	userRepository := repository.UserRepository{
		DBStorage: PgsStorage,
	}
	userService := service.UserService{
		UserRepository: &userRepository,
	}
	orderRepository := repository.OrderRepository{
		DBStorage: PgsStorage,
	}
	orderService := service.OrderService{
		OrderRepository: &orderRepository,
	}
	withdrawRepository := repository.WithdrawRepository{
		DBStorage: PgsStorage,
	}
	withdrawService := service.WithdrawService{
		WithdrawRepository: &withdrawRepository,
	}
	userHandler := handlers.UserHandler{
		UserService:     userService,
		OrderService:    orderService,
		WithdrawService: withdrawService,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestDecompressor)
	r.Post("/api/user/register", userHandler.Register)
	r.Post("/api/user/login", userHandler.Login)

	r.With(middleware.TokenAuthMiddleware).Route("/", func(r chi.Router) {
		r.Post("/api/user/orders", userHandler.SaveOrder)
		r.Get("/api/user/orders", userHandler.GetOrders)
		r.Get("/api/user/balance", userHandler.GetBalance)
		r.Post("/api/user/withdraw", userHandler.Withdraw)
		r.Get("/api/user/withdrawals", userHandler.Withdrawals)
	})

	err = http.ListenAndServe(cfg.ServerAddress, r)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
