package output

import "ca-tech-dojo-go/pkg/cago/service/io"

type GetCharas struct {
	CharaName map[int]string
}

type AddUserChara struct{}

type GetUserCharas struct {
	Charas []io.UserChara
}
