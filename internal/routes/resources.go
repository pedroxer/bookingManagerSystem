package routes

import (
	"encoding/json"
	proto_gen "github.com/pedroxer/BookingManagerSystem/internal/proto_gen/protos"
	"github.com/pedroxer/BookingManagerSystem/internal/routes/common_errors"
	"github.com/valyala/fasthttp"
	"strconv"
)

type resourceImpl struct {
	r *Router
}

func registerResource(r *Router) {
	resource := &resourceImpl{
		r: r,
	}
	// Workplace endpoints
	r.rtr.GET("/api/v1/resources/workplaces", resource.getWorkplaces)
	r.rtr.GET("/api/v1/resources/workplaces/{id}", resource.getWorkplaceById)
	r.rtr.GET("/api/v1/resources/workplaces/tag/{tag}", resource.getWorkplaceByUniqueTag)
	r.rtr.POST("/api/v1/resources/workplaces", resource.createWorkplace)
	r.rtr.PUT("/api/v1/resources/workplaces/{id}", resource.updateWorkplace)
	r.rtr.DELETE("/api/v1/resources/workplaces/{id}", resource.deleteWorkplace)

	// Parking Space endpoints
	r.rtr.GET("/api/v1/resources/parking", resource.getParkingSpaces)
	r.rtr.GET("/api/v1/resources/parking/{id}", resource.getParkingSpaceById)
	r.rtr.POST("/api/v1/resources/parking", resource.createParkingSpace)
	r.rtr.PUT("/api/v1/resources/parking/{id}", resource.updateParkingSpace)
	r.rtr.DELETE("/api/v1/resources/parking/{id}", resource.deleteParkingSpace)

	// Item endpoints
	r.rtr.GET("/api/v1/resources/items", resource.getItems)
	r.rtr.GET("/api/v1/resources/items/{id}", resource.getItemById)
	r.rtr.POST("/api/v1/resources/items", resource.createItem)
	r.rtr.PUT("/api/v1/resources/items/{id}", resource.updateItem)
	r.rtr.DELETE("/api/v1/resources/items/{id}", resource.deleteItem)
	//	r.rtr.POST("/api/v1/resources/items/{id}/attach", resource.attachItemToWorkplace)

}

func (r *resourceImpl) getWorkplaces(ctx *fasthttp.RequestCtx) {
	var req GetWorkplacesRequest

	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	workplaces, err := r.r.resourceClient.GetWorkplaces(ctx, &proto_gen.GetWorkplacesRequest{
		Zone:        req.Zone,
		Floor:       req.Floor,
		Type:        req.Type,
		Capacity:    req.Capacity,
		IsAvailable: req.IsAvailable,
		WithItems:   req.WithItems,
		Page:        req.Page,
	})
	if err != nil {
		r.r.logger.Warn(err.Error())
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}

	var resp GetWorkplacesResponse
	resp.PageSize = workplaces.PageSize
	resp.Page = workplaces.Page
	resp.TotalCount = workplaces.TotalCount
	for _, workplace := range workplaces.Workplaces {
		resp.Workplaces = append(resp.Workplaces, transformWorkplace(workplace))
	}
	res, err := json.Marshal(resp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) getWorkplaceById(ctx *fasthttp.RequestCtx) {
	var idRaw = ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	if id <= 0 {
		r.r.logger.Warn("invalid id")
		common_errors.FormError(ctx, "id is required", fasthttp.StatusBadRequest)
		return
	}
	workplace, err := r.r.resourceClient.GetWorkplaceById(ctx, &proto_gen.GetWorkplaceByIdRequest{
		Id: id,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}

	res, err := json.Marshal(transformWorkplace(workplace))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) getWorkplaceByUniqueTag(ctx *fasthttp.RequestCtx) {
	var tag = ctx.UserValue("tag").(string)
	workplace, err := r.r.resourceClient.GetWorkplaceByUniqueTag(ctx, &proto_gen.GetWorkplaceByUniqueTagRequest{
		UniqueTag: tag,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}

	res, err := json.Marshal(transformWorkplace(workplace))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) createWorkplace(ctx *fasthttp.RequestCtx) {
	var req CreateWorkplaceRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	workplace, err := r.r.resourceClient.CreateWorkplace(ctx, &proto_gen.CreateWorkplaceRequest{
		Address:           req.Address,
		Zone:              req.Zone,
		Floor:             req.Floor,
		Number:            req.Number,
		Type:              req.Type,
		Capacity:          req.Capacity,
		Description:       req.Description,
		IsAvailable:       req.IsAvailable,
		MaintenanceStatus: req.MaintenanceStatus,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	resp, err := json.Marshal(transformWorkplace(workplace))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(resp)
}

func (r *resourceImpl) updateWorkplace(ctx *fasthttp.RequestCtx) {
	var req UpdateWorkplaceRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	var idRaw = ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	if id <= 0 {
		r.r.logger.Warn("invalid id")
		common_errors.FormError(ctx, "id is required", fasthttp.StatusBadRequest)
		return
	}
	workplace, err := r.r.resourceClient.UpdateWorkplace(ctx, &proto_gen.UpdateWorkplaceRequest{
		Id:                id,
		Address:           req.Address,
		Zone:              req.Zone,
		Floor:             req.Floor,
		Number:            req.Number,
		Type:              req.Type,
		Capacity:          req.Capacity,
		Description:       req.Description,
		IsAvailable:       req.IsAvailable,
		MaintenanceStatus: req.MaintenanceStatus,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}

	res, err := json.Marshal(transformWorkplace(workplace))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) deleteWorkplace(ctx *fasthttp.RequestCtx) {
	var idRaw = ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	if id <= 0 {
		r.r.logger.Warn("invalid id")
		common_errors.FormError(ctx, "id is required", fasthttp.StatusBadRequest)
		return
	}
	resp, err := r.r.resourceClient.DeleteWorkplace(ctx, &proto_gen.DeleteWorkplaceRequest{
		Id: id,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	res, err := json.Marshal(resp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) getParkingSpaces(ctx *fasthttp.RequestCtx) {
	var req GetParkingSpacesRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}

	parkingSpaces, err := r.r.resourceClient.GetParkingSpaces(ctx, &proto_gen.GetParkingSpacesRequest{
		Address:     req.Address,
		Number:      req.Number,
		IsAvailable: req.IsAvailable,
		Zone:        req.Zone,
		Type:        req.Type,
		Page:        req.Page,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	parkingSpaceResp := GetParkingSpacesResponse{
		ParkingSpaces: make([]ParkingSpace, 0),
		TotalCount:    parkingSpaces.TotalCount,
		Page:          parkingSpaces.Page,
		PageSize:      parkingSpaces.PageSize,
	}
	for _, parkingSpace := range parkingSpaces.ParkingSpaces {
		parkingSpaceResp.ParkingSpaces = append(parkingSpaceResp.ParkingSpaces, transformParkingSpace(parkingSpace))
	}
	res, err := json.Marshal(parkingSpaceResp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) getParkingSpaceById(ctx *fasthttp.RequestCtx) {
	idRaw := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	if id <= 0 {
		r.r.logger.Warn("invalid id")
		common_errors.FormError(ctx, "id is required", fasthttp.StatusBadRequest)
		return
	}
	parkingSpace, err := r.r.resourceClient.GetParkingSpaceById(ctx, &proto_gen.GetParkingSpaceByIdRequest{
		Id: id,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	res, err := json.Marshal(transformParkingSpace(parkingSpace))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) createParkingSpace(ctx *fasthttp.RequestCtx) {
	var req CreateParkingSpaceRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	parkingSpace, err := r.r.resourceClient.CreateParkingSpace(ctx, &proto_gen.CreateParkingSpaceRequest{
		Number:      req.Number,
		Address:     req.Address,
		Type:        req.Type,
		Zone:        req.Zone,
		IsAvailable: req.IsAvailable,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	res, err := json.Marshal(transformParkingSpace(parkingSpace))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) updateParkingSpace(ctx *fasthttp.RequestCtx) {
	var req UpdateParkingSpaceRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	idRaw := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	if id <= 0 {
		r.r.logger.Warn("invalid id")
		common_errors.FormError(ctx, "id is required", fasthttp.StatusBadRequest)
		return
	}
	parkingSpace, err := r.r.resourceClient.UpdateParkingSpace(ctx, &proto_gen.UpdateParkingSpaceRequest{
		Id:          id,
		Number:      req.Number,
		Address:     req.Address,
		Type:        req.Type,
		Zone:        req.Zone,
		IsAvailable: req.IsAvailable,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	res, err := json.Marshal(transformParkingSpace(parkingSpace))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) deleteParkingSpace(ctx *fasthttp.RequestCtx) {
	idRaw := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	if id <= 0 {
		r.r.logger.Warn("invalid id")
		common_errors.FormError(ctx, "id is required", fasthttp.StatusBadRequest)
		return
	}
	_, err = r.r.resourceClient.DeleteParkingSpace(ctx, &proto_gen.DeleteParkingSpaceRequest{
		Id: id,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (r *resourceImpl) getItems(ctx *fasthttp.RequestCtx) {
	var req GetItemsRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	items, err := r.r.resourceClient.GetItems(ctx, &proto_gen.GetItemsRequest{
		Page:        req.Page,
		Type:        req.Type,
		Name:        req.Name,
		ConditionId: req.ConditionID,
		WorkplaceId: req.WorkplaceID,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	itemsResp := GetItemsResponse{
		Items:      make([]Item, 0),
		TotalCount: items.TotalCount,
		Page:       items.Page,
		PageSize:   items.PageSize,
	}
	for _, item := range items.Items {
		itemsResp.Items = append(itemsResp.Items, transformItem(item))
	}
	res, err := json.Marshal(itemsResp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) getItemById(ctx *fasthttp.RequestCtx) {
	idRaw := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	if id <= 0 {
		r.r.logger.Warn("invalid id")
		common_errors.FormError(ctx, "id is required", fasthttp.StatusBadRequest)
		return
	}
	item, err := r.r.resourceClient.GetItemById(ctx, &proto_gen.GetItemByIdRequest{
		Id: id,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	res, err := json.Marshal(transformItem(item))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) createItem(ctx *fasthttp.RequestCtx) {
	var req CreateItemRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	item, err := r.r.resourceClient.CreateItem(ctx, &proto_gen.CreateItemRequest{
		Type:        req.Type,
		Name:        req.Name,
		ConditionId: req.ConditionID,
		WorkplaceId: req.WorkplaceID,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	res, err := json.Marshal(transformItem(item))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) updateItem(ctx *fasthttp.RequestCtx) {
	var req UpdateItemRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	idRaw := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	item, err := r.r.resourceClient.UpdateItem(ctx, &proto_gen.UpdateItemRequest{
		Id:          id,
		Type:        req.Type,
		Name:        req.Name,
		ConditionId: req.ConditionID,
		WorkplaceId: req.WorkplaceID,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	res, err := json.Marshal(transformItem(item))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(res)
}

func (r *resourceImpl) deleteItem(ctx *fasthttp.RequestCtx) {
	idRaw := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	if id <= 0 {
		r.r.logger.Warn("invalid id")
		common_errors.FormError(ctx, "id is required", fasthttp.StatusBadRequest)
		return
	}
	_, err = r.r.resourceClient.DeleteItem(ctx, &proto_gen.DeleteItemRequest{
		Id: id,
	})
	if err != nil {
		r.r.logger.Warn(err)
		common_errors.FormError(ctx, common_errors.CastGrpcMessage(err), common_errors.CastGrpcErrors(err))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
}
func transformWorkplace(protoWorkplace *proto_gen.Workplace) Workplace {
	var workplace = Workplace{
		ID:                protoWorkplace.Id,
		Zone:              protoWorkplace.Zone,
		Floor:             protoWorkplace.Floor,
		Number:            protoWorkplace.Number,
		Type:              protoWorkplace.Type,
		Capacity:          protoWorkplace.Capacity,
		Description:       protoWorkplace.Description,
		IsAvailable:       protoWorkplace.IsAvailable,
		MaintenanceStatus: protoWorkplace.MaintenanceStatus,
		CreatedAt:         protoWorkplace.CreatedAt.AsTime(),
		UpdatedAt:         protoWorkplace.UpdatedAt.AsTime(),
		Items:             make([]Item, 0),
		UniqueTag:         protoWorkplace.UniqueTag,
	}
	for _, item := range protoWorkplace.Items {
		workplace.Items = append(workplace.Items, Item{
			ID:          item.Id,
			Type:        item.Type,
			Name:        item.Name,
			Condition:   item.Condition,
			CreatedAt:   item.CreatedAt.AsTime(),
			UpdatedAt:   item.UpdatedAt.AsTime(),
			WorkplaceID: item.WorkplaceId,
		})
	}
	return workplace
}

func transformParkingSpace(protoParkingSpace *proto_gen.ParkingSpace) ParkingSpace {
	return ParkingSpace{
		ID:          protoParkingSpace.Id,
		Number:      protoParkingSpace.Number,
		Address:     protoParkingSpace.Address,
		Zone:        protoParkingSpace.Zone,
		Type:        protoParkingSpace.Type,
		IsAvailable: protoParkingSpace.IsAvailable,
		CreatedAt:   protoParkingSpace.CreatedAt.AsTime(),
		UpdatedAt:   protoParkingSpace.UpdatedAt.AsTime(),
	}
}

func transformItem(protoItem *proto_gen.Item) Item {
	return Item{
		ID:          protoItem.Id,
		Type:        protoItem.Type,
		Name:        protoItem.Name,
		Condition:   protoItem.Condition,
		CreatedAt:   protoItem.CreatedAt.AsTime(),
		UpdatedAt:   protoItem.UpdatedAt.AsTime(),
		WorkplaceID: protoItem.WorkplaceId,
	}
}
