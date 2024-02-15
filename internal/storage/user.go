package storage

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"

	"monolith/internal/models"
	dto "monolith/internal/storage/models/user"
)

func (p *Postgres) Insert(ctx context.Context, user models.User) (string, error) {
	query := `
		insert into users (password, first_name, second_name, birthday, city, biography)
		values(:password, :first_name, :second_name, :birthday, :city, :biography)
		returning id
	`

	rows, err := p.conn.NamedQueryContext(ctx, query, dto.Imported(user))
	if err != nil {
		return "", errors.Wrap(err, "insert user error")
	}

	var userID string
	for rows.Next() {
		if err := rows.Scan(&userID); err != nil {
			return "", errors.Wrap(err, "scan user insert result error")
		}
	}

	return userID, nil
}

func (p *Postgres) FindByUserID(ctx context.Context, userID string) (*models.User, error) {
	query := `
		select id, password, first_name, second_name, birthday, city, biography
		from users
		where id = $1
	`

	query, args, err := sqlx.In(query, userID)
	if err != nil {
		return nil, errors.Wrap(err, "create user select error")
	}
	dbQuery := p.conn.Rebind(query)

	var users []dto.User
	err = p.conn.SelectContext(ctx, &users, dbQuery, args...)
	if err != nil {
		return nil, errors.Wrap(err, "user select error")
	}
	if len(users) == 0 {
		return nil, nil
	}

	return pointer.To(dto.Exported(users[0])), nil
}
