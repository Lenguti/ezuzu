package managerdb

import (
	"github.com/lenguti/ezuzu/business/core/manager"
)

type dbManager struct {
	ID        string `db:"id"`
	Entity    string `db:"entity"`
	CreatedAt int64  `db:"created_at"`
	UpdateAt  int64  `db:"updated_at"`
}

func toDBManager(m manager.Manager) dbManager {
	return dbManager{
		ID:        m.ID.String(),
		Entity:    m.Entity,
		CreatedAt: m.CreatedAt.Unix(),
		UpdateAt:  m.UpdatedAt.Unix(),
	}
}
