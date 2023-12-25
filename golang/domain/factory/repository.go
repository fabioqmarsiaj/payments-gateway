package factory

import "github.com/fabioqmarsiaj/payments-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
