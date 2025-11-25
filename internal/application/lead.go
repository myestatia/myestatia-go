// internal/application/lead_service.go
package application

import (
    "context"
    "github.com/myestatia/myestatia-go/internal/domain/entity"

)

type LeadLogic struct {}


func (s LeadLogic) Create(ctx context.Context,name string, email string) (entity.Lead, error) {
    return entity.Lead{Name: name, Email: email}, nil
}