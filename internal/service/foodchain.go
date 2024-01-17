package service

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
	"errors"
)

func QueryProductsByPid(pid string, ptype int) (*response.ResProductsInfo, error) {

	switch ptype {
	case PASTURE_TYPE:
		q := query.Q.FeedingBatch
		batch, err := q.WithContext(context.Background()).Where(q.PID.Eq(pid)).Preload(q.Cows).First()
		if err != nil {
			return nil, err
		}

		cowsInfo := make([]product.CowInfo, len(batch.Cows))
		for i, cow := range batch.Cows {
			cowsInfo[i] = product.ToCowInfo(&cow)
		}

		return &response.ResProductsInfo{CowsInfo: cowsInfo}, nil

	case SLAUGHTER_TYPE:
		q := query.SlaughterBatch
		batch, err := q.WithContext(context.Background()).Where(q.PID.Eq(pid)).Preload(q.Products).First()
		if err != nil {
			return nil, err
		}

		psInfo := make([]product.SlaughterProductInfo, len(batch.Products))
		for i, slaughterProduct := range batch.Products {
			psInfo[i] = product.ToSlaughterProductInfo(&slaughterProduct)
		}

		return &response.ResProductsInfo{SlaughterProductsInfo: psInfo}, nil
	case PACKAGE_TYPE:
		q1 := query.PackageBatch
		batch, err := q1.WithContext(context.Background()).Where(q1.PID.Eq(pid)).First()
		if err != nil {
			return nil, err
		}

		q2 := query.PackageProduct
		pp, err := q2.WithContext(context.Background()).Where(q2.BatchNumber.Eq(batch.BatchNumber)).Find()

		psInfo := make([]product.PackageProductInfo, len(pp))

		for i, p := range pp {
			psInfo[i] = product.ToPackageProductInfo(p)
		}

		return &response.ResProductsInfo{PackageProductsInfo: psInfo}, nil
	case COLDCHAIN_TRANSPORT_TYPE:
		q1 := query.PackageProductAndTransportPIDMap
		m, err := q1.WithContext(context.Background()).Where(q1.TransportPID.Eq(pid)).First()
		if err != nil {
			return nil, err
		}

		q2 := query.PackageProduct
		packProduct, err := q2.WithContext(context.Background()).Where(q2.Number.Eq(m.PackageProductNumber)).First()
		if err != nil {
			return nil, err
		}

		q3 := query.TransportBatch
		batch, err := q3.WithContext(context.Background()).Where(q3.BatchNumber.Eq(*packProduct.TransportBatchNumber)).First()
		if err != nil {
			return nil, err
		}

		q4 := query.Mall
		mall, err := q4.WithContext(context.Background()).Where(q4.Number.Eq(batch.MallNumber)).First()
		if err != nil {
			return nil, err
		}

		ccInfo := response.ColdchainInfo{
			TVNumber:   batch.TVNumber,
			MallNumber: batch.MallNumber,
			MallName:   mall.Name,
		}

		return &response.ResProductsInfo{ColdchainInfo: &ccInfo}, nil
	default:
		return nil, errors.New("type is error")
	}
}

func QueryAllFoodchains() ([]models.Foodchain, int64, int64, error) {
	fcs := make([]models.Foodchain, 0)

	allMallGoods, _, err := GetAllMallGoods()
	if err != nil {
		return nil, 0, 0, err
	}

	for _, good := range allMallGoods {
		p, err := query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(*good.FinalPID)).First()
		if err != nil {
			return nil, 0, 0, err
		}
		packagePID := p.PrePID

		p, err = query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(packagePID)).First()
		if err != nil {
			return nil, 0, 0, err
		}
		slaughterPID := p.PrePID

		p, err = query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(slaughterPID)).First()
		if err != nil {
			return nil, 0, 0, err
		}
		pasturePID := p.PrePID

		fc := models.Foodchain{
			Checkcode:            good.Checkcode,
			CurrentProductNumber: good.Number,
			CurrentLastProcedure: "end",
			PasturePID:           pasturePID,
			SlaughterPID:         slaughterPID,
			PackagePID:           packagePID,
			ColdchainPID:         *good.FinalPID,
		}

		fcs = append(fcs, fc)
	}

	cfcsCount := len(fcs)

	allPackageProducts, err := query.Q.PackageProduct.WithContext(context.Background()).
		Where(query.PackageProduct.State.Neq(TRANSPORT_END_PACPRO)).Find()
	if err != nil {
		return nil, 0, 0, err
	}

	for _, pro := range allPackageProducts {
		packageBatch, err := query.Q.PackageBatch.WithContext(context.Background()).Where(query.PackageBatch.BatchNumber.Eq(pro.BatchNumber)).First()
		if err != nil {
			return nil, 0, 0, err
		}
		packagePID := packageBatch.PID

		p, err := query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(packagePID)).First()
		if err != nil {
			return nil, 0, 0, err
		}

		checkcode := p.CheckCode
		slaughterPID := p.PrePID

		p, err = query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(slaughterPID)).First()
		if err != nil {
			return nil, 0, 0, err
		}

		pasturePID := p.PrePID

		fc := models.Foodchain{
			Checkcode:            checkcode,
			CurrentProductNumber: pro.Number,
			CurrentLastProcedure: "package",
			PasturePID:           pasturePID,
			SlaughterPID:         slaughterPID,
			PackagePID:           packagePID,
			ColdchainPID:         "",
		}

		fcs = append(fcs, fc)
	}

	allSlaughterProducts, err := query.Q.SlaughterProduct.WithContext(context.Background()).
		Where(query.SlaughterProduct.State.Neq(PACKAGE_END_SLAPRO)).Find()
	if err != nil {
		return nil, 0, 0, err
	}

	for _, pro := range allSlaughterProducts {
		slaBatch, err := query.Q.SlaughterBatch.WithContext(context.Background()).Where(query.SlaughterBatch.BatchNumber.Eq(pro.BatchNumber)).First()
		if err != nil {
			return nil, 0, 0, err
		}
		slaughterPID := slaBatch.PID

		p, err := query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(slaughterPID)).First()
		if err != nil {
			return nil, 0, 0, err
		}
		checkcode := p.CheckCode
		pasturePID := p.PrePID

		fc := models.Foodchain{
			Checkcode:            checkcode,
			CurrentProductNumber: pro.Number,
			CurrentLastProcedure: "slaughter",
			PasturePID:           pasturePID,
			SlaughterPID:         slaughterPID,
			PackagePID:           "",
			ColdchainPID:         "",
		}

		fcs = append(fcs, fc)
	}

	allCows, err := query.Q.Cow.WithContext(context.Background()).Where(query.Cow.State.Neq(SLAUGHTER_END_STATE_COW)).
		Where(query.Cow.State.Neq(INIT_STATE_COW)).Where(query.Cow.State.Neq(FEEDING_STATE_COW)).Find()
	if err != nil {
		return nil, 0, 0, err
	}

	for _, cow := range allCows {
		feedingBatch, err := query.Q.FeedingBatch.WithContext(context.Background()).Where(query.FeedingBatch.BatchNumber.Eq(*cow.BatchNumber)).First()
		if err != nil {
			return nil, 0, 0, err
		}
		pasturePID := feedingBatch.PID

		p, err := query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(pasturePID)).First()
		if err != nil {
			return nil, 0, 0, err
		}
		checkcode := p.CheckCode

		fc := models.Foodchain{
			Checkcode:            checkcode,
			CurrentProductNumber: cow.Number,
			CurrentLastProcedure: "pasture",
			PasturePID:           pasturePID,
			SlaughterPID:         "",
			PackagePID:           "",
			ColdchainPID:         "",
		}

		fcs = append(fcs, fc)
	}

	return fcs, int64(len(fcs)), int64(cfcsCount), nil
}

func QueryPidInfo(pid string) (string, string, string, string, error) {
	q := query.Procedure
	pro, err := q.WithContext(context.Background()).Where(q.PID.Eq(pid)).First()
	if err != nil {
		return "", "", "", "", err
	}
	starttimestamp := pro.StartTimestamp
	endtimestamp := pro.CompletedTimestamp
	op := pro.Operator
	p := query.FSIMSUser
	user, err := p.WithContext(context.Background()).Where(p.Account.Eq(op)).First()
	if err != nil {
		return "", "", "", "", err
	}
	starttime := starttimestamp.Format("2006-01-02 15:04:05")
	endtime := endtimestamp.Format("2006-01-02 15:04:05")
	address := user.Company
	housenumber := user.HouseNumber
	return starttime, endtime, address, housenumber, nil
}
