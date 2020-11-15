package service

import (
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/io"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"

	"math/rand"
)

// GachaService GachaService
type GachaService interface {
	GetGachaRate(ctx context.Context, user *input.GetGachaRate) (output.GetGachaRate, error)
	DrawGacha(ctx context.Context, user *input.DrawGacha) (output.DrawGacha, error)
}

type gachaService struct {
	r repository.Repository
}

// NewGachaService Gachaサービス作成
func NewGachaService(r repository.Repository) GachaService {
	return &gachaService{r}
}

func (gs *gachaService) GetGachaRate(ctx context.Context, gacha *input.GetGachaRate) (output.GetGachaRate, error) {
	var outputGachaRate output.GetGachaRate

	con, err := gs.r.NewConnection()
	if err != nil {
		return outputGachaRate, err
	}

	// 確率タイプ
	groupModels, err := con.GachaProbabilityGroup().FindByGachaType(gacha.GachaType)
	if err != nil {
		return outputGachaRate, err
	}

	groupIDs := make([]string, 0, len(groupModels))
	// 格納
	outputGachaRate.GroupProbability = new(io.GroupProbability)
	for _, groupModel := range groupModels {
		var gachaProbabilityGroup io.GachaProbabilityGroup
		groupIDs = append(groupIDs, groupModel.GachaProbabilityGroupID)
		gachaProbabilityGroup.ID = groupModel.GachaProbabilityGroupID
		gachaProbabilityGroup.Rate = groupModel.Rate
		outputGachaRate.GroupProbability.SumRate += groupModel.Rate
		outputGachaRate.GroupProbability.GachaProbabilityGroups = append(outputGachaRate.GroupProbability.GachaProbabilityGroups, &gachaProbabilityGroup)
	}

	// キャラごとの確率
	gachaModels, err := con.Gacha().FindByGroupIDs(groupIDs)
	if err != nil {
		return outputGachaRate, err
	}

	//格納
	charaRates := make(map[string]*io.CharaProbability, len(groupModels))
	for _, gachaModel := range gachaModels {
		var charaRate io.CharaRate
		charaRate.CharaID = gachaModel.CharaID
		charaRate.GachaProbabilityGroupID = gachaModel.GroupID
		charaRate.Rate = gachaModel.Rate
		if _, ok := charaRates[charaRate.GachaProbabilityGroupID]; !ok {
			charaRates[charaRate.GachaProbabilityGroupID] = new(io.CharaProbability)
		}
		charaRates[charaRate.GachaProbabilityGroupID].SumRate += charaRate.Rate
		charaRates[charaRate.GachaProbabilityGroupID].CharaRates = append(charaRates[charaRate.GachaProbabilityGroupID].CharaRates, &charaRate)
	}
	outputGachaRate.CharaProbability = charaRates

	return outputGachaRate, err
}

func (gs *gachaService) DrawGacha(ctx context.Context, gacha *input.DrawGacha) (output.DrawGacha, error) {
	var outputDrawGacha output.DrawGacha

	// どのレートタイプに当たるか抽選
	sum := 0
	targetNum := rand.Intn(gacha.GroupProbability.SumRate)
	var winningGroup string
	for _, group := range gacha.GroupProbability.GachaProbabilityGroups {
		sum += group.Rate
		if sum > targetNum {
			winningGroup = group.ID
			break
		}
	}

	// レートの中のどのキャラに当たるか抽選
	sum = 0
	targetNum = rand.Intn(gacha.CharaProbability[winningGroup].SumRate)
	for _, charaRate := range gacha.CharaProbability[winningGroup].CharaRates {
		sum += charaRate.Rate
		if sum > targetNum {
			outputDrawGacha.CharaID = charaRate.CharaID
			break
		}
	}

	return outputDrawGacha, nil
}
