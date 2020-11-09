package interactor

import (
	"ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/interactor/output"
	"ca-tech-dojo-go/pkg/cago/service"
	sInput "ca-tech-dojo-go/pkg/cago/service/input"
	"context"
)

// CharaInteractor CharaInteractor
type CharaInteractor interface {
	GetCharaList(ctx context.Context, chara *input.GetCharaList) (output.GetCharaList, error)
}

type charaInteractor struct {
	cs service.CharaService
}

// NewCharaInteractor NewCharaInteractor
func NewCharaInteractor(cs service.CharaService) CharaInteractor {
	return &charaInteractor{cs}
}

func (ci *charaInteractor) GetCharaList(ctx context.Context, chara *input.GetCharaList) (output.GetCharaList, error) {
	var outputGetCharaList output.GetCharaList

	// ユーザーの所持キャラ一覧を取得
	var inputGetUserCharas sInput.GetUserCharas
	inputGetUserCharas.UserID = chara.UserID
	outputGetUserChara, err := ci.cs.GetUserCharas(ctx, &inputGetUserCharas)
	if err != nil {
		return outputGetCharaList, err
	}

	charaIDMap := make(map[int]struct{})
	var charaIDs []int
	for _, chara := range outputGetUserChara.Charas {
		if _, ok := charaIDMap[chara.CharaID]; !ok {
			charaIDMap[chara.CharaID] = struct{}{}
			charaIDs = append(charaIDs, chara.CharaID)
		}
	}

	// キャラ名を取得
	var getCharasInput sInput.GetCharas
	getCharasInput.IDs = charaIDs
	getCharaOutput, err := ci.cs.GetCharas(ctx, &getCharasInput)
	if err != nil {
		return outputGetCharaList, err
	}

	// 結合
	for _, chara := range outputGetUserChara.Charas {
		var userChara output.UserChara
		userChara.CharaID = chara.CharaID
		userChara.ID = chara.ID
		userChara.Name = getCharaOutput.CharaName[userChara.CharaID]
		outputGetCharaList.Charas = append(outputGetCharaList.Charas, userChara)
	}

	return outputGetCharaList, nil
}
