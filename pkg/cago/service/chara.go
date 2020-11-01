package service

import (
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/io"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
	"time"
)

// CharaService CharaService
type CharaService interface {
	GetCharas(ctx context.Context, user *input.GetCharas) (output.GetCharas, error)
	AddUserChara(ctx context.Context, user *input.AddUserChara) (output.AddUserChara, error)
}

type charaService struct {
	r repository.Repository
}

// NewCharaService NewCharaService
func NewCharaService(r repository.Repository) CharaService {
	return &charaService{r}
}

func (cs *charaService) GetCharas(ctx context.Context, chara *input.GetCharas) (output.GetCharas, error) {
	var outputGetCharas output.GetCharas

	con, err := cs.r.NewConnection()
	if err != nil {
		return outputGetCharas, err
	}
	defer con.Close()

	charaModels, err := con.Chara().FindByIDs(chara.IDs)
	if err != nil {
		return outputGetCharas, err
	}

	// 格納
	for _, charaModel := range charaModels {
		var chara io.Chara
		chara.ID = charaModel.ID
		chara.Name = charaModel.Name
		outputGetCharas.Charas = append(outputGetCharas.Charas, chara)
	}

	return outputGetCharas, err
}

func (cs *charaService) AddUserChara(ctx context.Context, chara *input.AddUserChara) (output.AddUserChara, error) {
	var outputAddUserChara output.AddUserChara

	con, err := cs.r.NewConnection()
	if err != nil {
		return outputAddUserChara, err
	}
	defer con.Close()

	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.Chara().AddUserChara(chara.CharaIDs, time.Now(), chara.UserID)
		if err != nil {
			return err
		}

		return nil
	})

	return outputAddUserChara, err
}
