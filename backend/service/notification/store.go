package notification

import (
	"database/sql"
	"log"
	"time"

	"github.com/Jeno7u/studybud/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}


func (s *Store)BroadcastNotifications (){
    for {
        time.Sleep(5 * time.Second) 
        for userID, client := range clients {
            notification, err := s.getNotifications(userID)
			if err != nil {
				log.Println("Write error:", err)
                client.Conn.Close()
                delete(clients, userID)
			}

            err = client.Conn.WriteJSON(notification)
            if err != nil {
                log.Println("Write error:", err)
                client.Conn.Close()
                delete(clients, userID)
            }
        }
    }
}

func (s *Store) getNotifications(userID int) ([]types.Notification, error) {	
	rows, err := s.db.Query(`
		SELECT id, type, date, "dayOfWeek", "title"
		FROM notifications
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	
	var result []types.Notification

	for rows.Next() {
		notification := new(types.Notification)

		err := rows.Scan(
			&notification.ID,
			&notification.Type,
			&notification.Date,
			&notification.DayOfWeek,
			&notification.Title,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *notification)
	}

	return result, nil
}