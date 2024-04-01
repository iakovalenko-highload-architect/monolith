package cache

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/redis/go-redis/v9"

	"monolith/internal/models"
)

type Cache struct {
	redisCli *redis.Client
}

func New(redisCli *redis.Client) *Cache {
	return &Cache{
		redisCli: redisCli,
	}
}

func (c *Cache) Append(ctx context.Context, userID string, post models.Post) error {
	encoded, err := encode(imported(post))
	if err != nil {
		return errors.Wrap(err, "encode post error")
	}

	if err := c.redisCli.LPush(ctx, userID, encoded).Err(); err != nil {
		return errors.Wrap(err, "push into cache by user id error")
	}

	if err := c.redisCli.LTrim(ctx, userID, 0, CacheLen-1).Err(); err != nil {
		return errors.Wrap(err, "trim cache by user id error")
	}
	return nil
}

func (c *Cache) Get(ctx context.Context, userID string, limit int64, offset int64) ([]models.Post, error) {
	encoded, err := c.redisCli.LRange(ctx, userID, offset, limit).Result()
	if err != nil {
		return nil, errors.Wrap(err, "get cache by user id error")
	}

	posts := make([]models.Post, 0, len(encoded))
	for _, e := range encoded {
		p, err := decode(e)
		if err != nil {
			return nil, errors.Wrap(err, "decode cache elem by user id error")
		}
		posts = append(posts, exported(p))
	}
	return posts, nil
}

func (c *Cache) Update(ctx context.Context, userID string, post models.Post) error {
	cachedPosts, err := c.Get(ctx, userID, CacheLen, 0)
	if err != nil {
		return errors.Wrap(err, "get cache by user id error")
	}

	index := -1
	for i := range cachedPosts {
		if cachedPosts[i].ID == post.ID {
			index = i
			break
		}
	}

	if index >= 0 {
		encoded, err := encode(imported(post))
		if err != nil {
			return errors.Wrap(err, "encode post error")
		}

		if c.redisCli.LSet(ctx, userID, int64(index), encoded).Err() != nil {
			return errors.Wrap(err, "set post error")
		}
	}

	return nil
}

func (c *Cache) Delete(ctx context.Context, userID string, post models.Post) error {
	encoded, err := encode(imported(post))
	if err != nil {
		return errors.Wrap(err, "encode post error")
	}

	if c.redisCli.LRem(ctx, userID, 1, encoded).Err() != nil {
		return errors.Wrap(err, "delete post error")
	}
	return nil
}

func (c *Cache) Clear(ctx context.Context, userID string) error {
	if err := c.redisCli.Del(ctx, userID).Err(); err != nil {
		return errors.Wrap(err, "delete cache by user id error")
	}
	return nil
}
