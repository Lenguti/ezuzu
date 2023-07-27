package v1

import "github.com/lenguti/ezuzu/business/core/manager"

// ClientManager - represents a client property manager entity.
type ClientManager struct {
	ID        string `json:"id"`
	Entity    string `json:"entity"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func toCoreNewManager(input CreateManagerRequest) manager.NewManager {
	newManager := manager.NewManager{
		Entity: input.Entity,
	}
	return newManager
}

func toClientManager(input manager.Manager) ClientManager {
	return ClientManager{
		ID:        input.ID.String(),
		Entity:    input.Entity,
		CreatedAt: input.CreatedAt.Unix(),
		UpdatedAt: input.UpdatedAt.Unix(),
	}
}
