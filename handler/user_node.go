package handler

//func (uh *UserHandler) ListNodes(c *gin.Context) {
//	nodes, err := api.NewAPI().GetAllNodes()
//	if err != nil {
//		uh.logger.WithError(err).Error("list nodes error")
//		c.JSON(http.StatusInternalServerError, models.NewErrorResp("获取节点列表失败！"))
//		return
//	}
//	var nodeResps []*models.NodeResp
//	for _, node := range nodes {
//		nodeResp := &models.NodeResp{}
//		copier.Copy(nodeResp, node)
//		nodeResp.Available = state.GetLoader().GetNodeAvailable(node.Id)
//		nodeResps = append(nodeResps, nodeResp)
//	}
//	c.JSON(http.StatusOK, models.NewSuccessResp(nodeResps))
//}
