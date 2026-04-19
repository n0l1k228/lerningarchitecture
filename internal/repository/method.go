package repository

import (
	"context"
	"fmt"
	"myavito/internal/domain"
)

func (r *PostgresRepo) SaveUser(ctx context.Context, name, phone_number string) (int, error) {
	var id int
	sqlQuery := `
	INSERT INTO avitoApp.users (name,phone_number)
	VALUES($1,$2)
	RETURNING id
	`
	err := r.conn.QueryRow(ctx, sqlQuery, name, phone_number).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("failed saved user in db:%w", err)
	}
	return id, nil
}

func (r *PostgresRepo) SaveProduct(ctx context.Context, user_id, price int, title, description, category string) error {
	sqlQuery := `
	INSERT INTO avitoApp.product(author_product_id, price, title, description, category)
	VALUES($1, $2, $3, $4, $5);
	`

	if _, err := r.conn.Exec(ctx, sqlQuery, user_id, price, title, description, category); err != nil {
		return fmt.Errorf("failed saved product in db:%w", err)
	}
	return nil
}

func (r *PostgresRepo) SearchProduct(ctx context.Context, title string) ([]domain.Product, error) {
	sqlQuery := `
	SELECT id, title, description, price, category FROM product WHERE title = $1
	`
	rows, err := r.conn.Query(ctx, sqlQuery, title)
	if err != nil {
		return nil, fmt.Errorf("failed get rows from db: %w", err)
	}
	defer rows.Close()
	var prod []domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.Category); err != nil {
			return nil, fmt.Errorf("failed scan product: %w", err)
		}

		prod = append(prod, p)
	}

	return prod, nil
}
