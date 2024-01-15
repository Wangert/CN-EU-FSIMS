package service

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
)

func QueryAllFoodchains() ([]models.Foodchain, int64, error) {
	fcs := make([]models.Foodchain, 0)

	allMallGoods, _, err := GetAllMallGoods()
	if err != nil {
		return nil, 0, err
	}

	for _, good := range allMallGoods {
		p, err := query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(*good.FinalPID)).First()
		if err != nil {
			return nil, 0, err
		}
		packagePID := p.PrePID

		p, err = query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(packagePID)).First()
		if err != nil {
			return nil, 0, err
		}
		slaughterPID := p.PrePID

		p, err = query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(slaughterPID)).First()
		if err != nil {
			return nil, 0, err
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

	allPackageProducts, err := query.Q.PackageProduct.WithContext(context.Background()).
		Where(query.PackageProduct.State.Neq(TRANSPORT_END_PACPRO)).Find()
	if err != nil {
		return nil, 0, err
	}

	for _, pro := range allPackageProducts {
		packageBatch, err := query.Q.PackageBatch.WithContext(context.Background()).Where(query.PackageBatch.BatchNumber.Eq(pro.BatchNumber)).First()
		if err != nil {
			return nil, 0, err
		}
		packagePID := packageBatch.PID

		p, err := query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(packagePID)).First()
		if err != nil {
			return nil, 0, err
		}

		checkcode := p.CheckCode
		slaughterPID := p.PrePID

		p, err = query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(slaughterPID)).First()
		if err != nil {
			return nil, 0, err
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
		return nil, 0, err
	}

	for _, pro := range allSlaughterProducts {
		slaBatch, err := query.Q.SlaughterBatch.WithContext(context.Background()).Where(query.SlaughterBatch.BatchNumber.Eq(pro.BatchNumber)).First()
		if err != nil {
			return nil, 0, err
		}
		slaughterPID := slaBatch.PID

		p, err := query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(slaughterPID)).First()
		if err != nil {
			return nil, 0, err
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
		return nil, 0, err
	}

	for _, cow := range allCows {
		feedingBatch, err := query.Q.FeedingBatch.WithContext(context.Background()).Where(query.FeedingBatch.BatchNumber.Eq(*cow.BatchNumber)).First()
		if err != nil {
			return nil, 0, err
		}
		pasturePID := feedingBatch.PID

		p, err := query.Q.Procedure.WithContext(context.Background()).Where(query.Procedure.PID.Eq(pasturePID)).First()
		if err != nil {
			return nil, 0, err
		}
		checkcode := p.CheckCode

		fc := models.Foodchain{
			Checkcode:            checkcode,
			CurrentProductNumber: cow.Number,
			CurrentLastProcedure: "slaughter",
			PasturePID:           pasturePID,
			SlaughterPID:         "",
			PackagePID:           "",
			ColdchainPID:         "",
		}

		fcs = append(fcs, fc)
	}

	return fcs, int64(len(fcs)), nil
}
