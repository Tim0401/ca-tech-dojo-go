package service

import (
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/io"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
	"time"

	"golang.org/x/xerrors"
)

// CharaService CharaService
type CharaService interface {
	GetCharas(ctx context.Context, chara *input.GetCharas) (output.GetCharas, error)
	AddUserChara(ctx context.Context, chara *input.AddUserChara) (output.AddUserChara, error)
	GetUserCharas(ctx context.Context, chara *input.GetUserCharas) (output.GetUserCharas, error)
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
		return outputGetCharas, xerrors.Errorf("Call NewConnection: %w", err)
	}

	charaModels, err := con.Chara().FindByIDs(chara.IDs)
	if err != nil {
		return outputGetCharas, xerrors.Errorf("Call FindByIDs: %w", err)
	}

	// 格納
	outputGetCharas.CharaName = make(map[int]string, len(charaModels))
	for _, charaModel := range charaModels {
		outputGetCharas.CharaName[charaModel.ID] = charaModel.Name
	}

	return outputGetCharas, nil
}

func (cs *charaService) AddUserChara(ctx context.Context, chara *input.AddUserChara) (output.AddUserChara, error) {
	var outputAddUserChara output.AddUserChara

	con, err := cs.r.NewConnection()
	if err != nil {
		return outputAddUserChara, xerrors.Errorf("Call NewConnection: %w", err)
	}

	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.Chara().AddUserChara(chara.CharaIDs, time.Now(), chara.UserID)
		if err != nil {
			return xerrors.Errorf("Call AddUserChara: %w", err)
		}

		return nil
	})

	if err != nil {
		return outputAddUserChara, xerrors.Errorf("Call RunTransaction: %w", err)
	}
	return outputAddUserChara, nil
}

func (cs *charaService) GetUserCharas(ctx context.Context, chara *input.GetUserCharas) (output.GetUserCharas, error) {
	var outputGetUserCharas output.GetUserCharas

	con, err := cs.r.NewConnection()
	if err != nil {
		return outputGetUserCharas, xerrors.Errorf("Call NewConnection: %w", err)
	}

	userCharaModels, err := con.Chara().FindUserCharaByUserID(chara.UserID)
	if err != nil {
		return outputGetUserCharas, xerrors.Errorf("Call FindUserCharaByUserID: %w", err)
	}

	// 格納
	for _, userCharaModel := range userCharaModels {
		var userChara io.UserChara
		userChara.ID = userCharaModel.ID
		userChara.CharaID = userCharaModel.CharaID
		outputGetUserCharas.Charas = append(outputGetUserCharas.Charas, userChara)
	}

	return outputGetUserCharas, nil
}
