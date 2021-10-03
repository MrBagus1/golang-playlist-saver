package servstatus

import (
	"context"
)

type StatusService interface {
	GetAllStatus(ctx context.Context) ([]Status, error)
	CronPremiumChecker(ctx context.Context) error
}
