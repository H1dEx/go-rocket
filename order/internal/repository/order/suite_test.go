package order

import (
	"context"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/H1dEx/go-rocket/order/internal/migrator"
	"github.com/H1dEx/go-rocket/order/internal/model"
	repoModel "github.com/H1dEx/go-rocket/order/internal/repository/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	repo *repository
}

func (s *ServiceSuite) SetupTest() {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../../../..")
	err := godotenv.Load(path.Join(root, ".env"))
	if err != nil {
		log.Printf("Failed to read env file %v\n", err)
		return
	}

	s.ctx = context.Background()
	ctx := context.Background()
	pgContainer, err := postgres.Run(ctx, "postgres:15",
		postgres.WithDatabase("test_db"),
		postgres.WithUsername("test_user"),
		postgres.WithPassword("test_password"),
		postgres.BasicWaitStrategies(),
	)

	defer func() {
		if err := testcontainers.TerminateContainer(pgContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	connString, err := pgContainer.ConnectionString(ctx)
	if err != nil {
		log.Printf("failed to get connection string: %s", err)
		return
	}
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return
	}

	defer pool.Close()

	migrationsDir := path.Join(root, os.Getenv("MIGRATIONS_DIR"))
	log.Printf("check : %v\n", migrationsDir)
	migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*pool.Config().ConnConfig.Copy()), migrationsDir)

	if err := migratorRunner.Up(); err != nil {
		log.Printf("Ошибка миграции базы данных: %v\n", err)
		return
	}

	s.repo = NewRepository(pool)
}

func (s *ServiceSuite) TearDownTest() {
}

func (s *ServiceSuite) InsertOrder(order model.Order) (pgconn.CommandTag, error) {
	ctx := context.Background()
	res, err := s.repo.pool.Exec(ctx, "INSERT INTO orders (orderUuid, userUuid, partUuids, totalPrice, transactionUuid, paymentMethod, status) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		order.OrderUUID, order.UserUUID, order.PartUuids, order.TotalPrice, order.TransactionUUID, order.PaymentMethod, order.Status,
	)
	return res, err
}

func (s *ServiceSuite) GenFakeOrder() model.Order {
	return model.Order{
		OrderUUID:       gofakeit.UUID(),
		UserUUID:        gofakeit.UUID(),
		PartUuids:       []string{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()},
		TotalPrice:      gofakeit.Float32(),
		TransactionUUID: gofakeit.UUID(),
		PaymentMethod:   model.PaymentMethodUnknown,
		Status:          model.OrderStatusPendingPayment,
	}
}

func (s *ServiceSuite) GenFakeRepoOrder() repoModel.Order {
	return repoModel.Order{
		OrderUUID:       gofakeit.UUID(),
		UserUUID:        gofakeit.UUID(),
		PartUuids:       []string{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()},
		TotalPrice:      gofakeit.Float32(),
		TransactionUUID: gofakeit.UUID(),
		PaymentMethod:   repoModel.PaymentMethodUnknown,
		Status:          repoModel.OrderStatusPendingPayment,
	}
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
