package task

import "todo-planner/infrastructure"

type IRepository interface {}

type Repository struct {
	db infrastructure.IDatabase
	logger infrastructure.ILogger
}

func NewRepository(db infrastructure.IDatabase, logger infrastructure.ILogger) IRepository {
	return &Repository{
		db: db,
		logger: logger,
	}
}
