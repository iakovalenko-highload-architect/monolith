package storage

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"

	"monolith/internal/models"
	dto "monolith/internal/storage/models/post"
)

func (p *Postgres) InsertPost(ctx context.Context, post models.Post) (string, error) {
	query := `
		insert into posts (user_id, text_)
		values(:user_id, :text_)
		returning id
	`

	rows, err := p.conn.NamedQueryContext(ctx, query, dto.Imported(post))
	if err != nil {
		return "", errors.Wrap(err, "insert post error")
	}

	var postID string
	for rows.Next() {
		if err := rows.Scan(&postID); err != nil {
			return "", errors.Wrap(err, "scan post insert result error")
		}
	}

	return postID, nil
}

func (p *Postgres) UpdatePost(ctx context.Context, post models.Post) error {
	query := `
		update posts
		set text_ = $1
		where id = $2
	`

	_, err := p.conn.ExecContext(ctx, query, post.Text, post.ID)
	if err != nil {
		return errors.Wrap(err, "update post error")
	}

	return nil
}

func (p *Postgres) FindByPostID(ctx context.Context, postID string) (*models.Post, error) {
	query := `
		select id, user_id, text_
		from posts
		where id = $1
	`

	query, args, err := sqlx.In(query, postID)
	if err != nil {
		return nil, errors.Wrap(err, "create post select error")
	}
	dbQuery := p.conn.Rebind(query)

	var posts []dto.Post
	err = p.conn.SelectContext(ctx, &posts, dbQuery, args...)
	if err != nil {
		return nil, errors.Wrap(err, "post select error")
	}
	if len(posts) == 0 {
		return nil, nil
	}

	return pointer.To(dto.Exported(posts[0])), nil
}

func (p *Postgres) FindPostsByUserID(ctx context.Context, userID string, limit int64) ([]models.Post, error) {
	query := `
		select id, user_id, text_
		from posts
		where user_id = $1
		order by created_at desc
		limit $2
	`

	query, args, err := sqlx.In(query, userID, limit)
	if err != nil {
		return nil, errors.Wrap(err, "create post select error")
	}
	dbQuery := p.conn.Rebind(query)

	var posts []dto.Post
	err = p.conn.SelectContext(ctx, &posts, dbQuery, args...)
	if err != nil {
		return nil, errors.Wrap(err, "post select error")
	}
	if len(posts) == 0 {
		return nil, nil
	}

	res := make([]models.Post, 0, len(posts))
	for _, post := range posts {
		res = append(res, dto.Exported(post))
	}

	return res, nil
}

func (p *Postgres) DeletePost(ctx context.Context, postID string) error {
	query := `
		delete from posts
		where id = $1
	`

	_, err := p.conn.ExecContext(ctx, query, postID)
	if err != nil {
		return errors.Wrap(err, "delete post error")
	}

	return nil
}
