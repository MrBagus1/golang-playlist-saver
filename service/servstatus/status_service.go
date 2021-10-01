package servstatus

import (
	"context"
)

type StatusService interface {
	GetStatusByUserId(ctx context.Context) Status
	GetAllStatus(ctx context.Context) ([]Status, error)
	CronPremiumChecker(ctx context.Context) error
}
