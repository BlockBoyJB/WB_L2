package v1

import "event-server/internal/service"

func NewRouter(services *service.Services) {
	newEventRouter(services.Event)
}
