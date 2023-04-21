package repository

import (
	"context"
)

func (h Handler) Del(key string) {
	if err := h.redis.Client.Del(context.TODO(), key).Err(); err != nil {
		h.logger.Error(err)
	}
}
