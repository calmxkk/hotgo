package k8s

import (
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/cluster/list" method:"get" tags:"集群" summary:"获取集群列表"`
	adminin.RoleListInp
}

type ListRes struct {
	*adminin.RoleListModel
	form.PageRes
}
