package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/go-redis/redis/v8"
)

type redisAuthRepo struct {
	Cl *redis.Client
}

func NewRedisAuthRepo(cl *redis.Client) redisAuthRepo {
	return redisAuthRepo{
		Cl: cl,
	}
}

func (r *redisAuthRepo) CreateAuth(ctx context.Context, userid int64, td *domain.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := r.Cl.Set(ctx, td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := r.Cl.Set(ctx, td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
