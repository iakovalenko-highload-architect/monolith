package storage

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"

	"monolith/internal/models"
	dto "monolith/internal/storage/models/friendship"
)

func (p *Postgres) InsertFriendship(ctx context.Context, userID string, friendID string) error {
	query := `
		insert into friendships (user_id, friend_id)
		values($1, $2)
	`

	_, err := p.conn.ExecContext(ctx, query, userID, friendID)
	if err != nil {
		return errors.Wrap(err, "insert friendship error")
	}

	return nil
}

func (p *Postgres) DeleteFriendship(ctx context.Context, userID string, friendID string) error {
	query := `
		delete from friendships
		where user_id = $1 and friend_id = $2
	`

	_, err := p.conn.ExecContext(ctx, query, userID, friendID)
	if err != nil {
		return errors.Wrap(err, "delete friendship error")
	}

	return nil
}

func (p *Postgres) FindFriendshipByUserID(ctx context.Context, userID string) ([]models.Friendship, error) {
	query := `
		select user_id, friend_id
		from friendships
		where friend_id = $1
	`

	query, args, err := sqlx.In(query, userID)
	if err != nil {
		return nil, errors.Wrap(err, "create friendships select error")
	}
	dbQuery := p.conn.Rebind(query)

	var friendships []dto.Friendship
	err = p.conn.SelectContext(ctx, &friendships, dbQuery, args...)
	if err != nil {
		return nil, errors.Wrap(err, "friendships select error")
	}
	if len(friendships) == 0 {
		return nil, nil
	}

	res := make([]models.Friendship, 0, len(friendships))
	for _, friend := range friendships {
		res = append(res, dto.Exported(friend))
	}

	return res, nil
}

func (p *Postgres) GetAllFriends(ctx context.Context) ([]models.Friendship, error) {
	query := `
		select user_id, friend_id
		from friendships
	`

	dbQuery := p.conn.Rebind(query)

	var friendships []dto.Friendship
	err := p.conn.SelectContext(ctx, &friendships, dbQuery)
	if err != nil {
		return nil, errors.Wrap(err, "friendships select error")
	}
	if len(friendships) == 0 {
		return nil, nil
	}

	res := make([]models.Friendship, 0, len(friendships))
	for _, friend := range friendships {
		res = append(res, dto.Exported(friend))
	}

	return res, nil
}
