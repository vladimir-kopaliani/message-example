package messagerepo

import (
	"context"
	"github.com/vladimir-kopaliani/message-example/internal/models"
	"time"
)

// SaveMessage save message in `messages` table
func (r *Repository) SaveMessage(ctx context.Context, message *models.Message) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO messages (
			id,
			text,
			created_at
			) VALUES ($1, $2, $3);`,
		message.ID,
		message.Text,
		message.CreatedAt,
	)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	return nil
}

// GetMessages returns message by given time range
func (r *Repository) GetMessages(ctx context.Context, input models.GetMessagesInput) ([]*models.Message, error) {
	rows, err := r.db.Query(ctx,
		`SELECT
			id,
			text,
			created_at
		FROM messages
		WHERE created_at >= $1 and created_at <= $2;
		`,
		input.StartDate,
		input.EndDate,
	)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	messages := make([]*models.Message, 0)

	for rows.Next() {
		var message models.Message
		var t time.Time

		err = rows.Scan(
			&message.ID,
			&message.Text,
			&t,
		)
		if err != nil {
			r.logger.Error(err)
			return nil, err
		}

		message.CreatedAt = models.DateTime(t)

		messages = append(messages, &message)
	}

	return messages, nil
}
