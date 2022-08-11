package service

import (
	"fmt"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
	"rohmat.co.id/serverconfig"
)

func RestructureScope() (err model.ErrorModel) {
	db := serverconfig.ServerAttribute.DBConnection
	// select mapping nexseller
	listID, err := dao.GetListIDMappingNexsellerForScope(db)
	if err.Error != nil {
		return
	}
	// insert data scope nexseller
	for i := 0; i < len(listID); i++ {
		_, _ = fmt.Fprintln(serverconfig.ServerAttribute.Write, fmt.Sprintf("restructure scope nexseller ke- %d", i+1))
		fmt.Println("restructure scope nexseller ke- ", i+1)
		dataScope := model.ScopeModel{
			Scope:       fmt.Sprintf("nexsoft.nx_id:%d", listID[i]),
			CreatedBy:   1,
			Description: fmt.Sprintf("mapping_nexseller %d", listID[i]),
		}
		err = dao.InsertDataScope(db, dataScope)
		if err.Error != nil {
			return
		}
	}
	// select nexchief account
	listID, err = dao.GetListIDNexchiefAccountForScope(db)
	if err.Error != nil {
		return
	}
	// insert data scope nexchief account
	for i := 0; i < len(listID); i++ {
		_, _ = fmt.Fprintln(serverconfig.ServerAttribute.Write, fmt.Sprintf("restructure scope nexchief account ke- %d", i+1))
		fmt.Println("restructure scope nexchief account ke- ", i+1)
		dataScope := model.ScopeModel{
			Scope:       fmt.Sprintf("nexsoft.nc_id:%d", listID[i]),
			CreatedBy:   1,
			Description: fmt.Sprintf("nexchief_account %d", listID[i]),
		}
		err = dao.InsertDataScope(db, dataScope)
		if err.Error != nil {
			return
		}
	}

	// select nexchief account
	listID, err = dao.GetListIDGeoTreeForScope(db)
	if err.Error != nil {
		return
	}
	// insert data scope nexchief account
	for i := 0; i < len(listID); i++ {
		_, _ = fmt.Fprintln(serverconfig.ServerAttribute.Write, fmt.Sprintf("restructure scope geo tree ke- %d", i+1))
		fmt.Println("restructure scope geo tree ke- ", i+1)
		dataScope := model.ScopeModel{
			Scope:       fmt.Sprintf("nexsoft.geo_id:%d", listID[i]),
			CreatedBy:   1,
			Description: fmt.Sprintf("geo_tree %d", listID[i]),
		}
		err = dao.InsertDataScope(db, dataScope)
		if err.Error != nil {
			return
		}
	}
	return
}
