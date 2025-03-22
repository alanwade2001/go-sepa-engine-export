package service

import (
	"log/slog"
	"strconv"

	"github.com/alanwade2001/go-sepa-engine-data/model"
	"github.com/alanwade2001/go-sepa-engine-data/repository"
	"github.com/alanwade2001/go-sepa-iso/pacs_008_001_02"
)

type Export struct {
	reposMgr *repository.Manager
}

func NewExport(reposMgr *repository.Manager) *Export {
	document := &Export{
		reposMgr: reposMgr,
	}

	return document
}

func (d *Export) Export(id string) (*pacs_008_001_02.Document, error) {

	slog.Info("export", "id", id)
	sgModel := &model.SettlementGroup{}

	if ID, err := strconv.ParseUint(id, 10, 64); err != nil {
		slog.Error("failed to parse id", "err", err)
		return nil, err
	} else if sgEnt, err := d.reposMgr.SettlementGroup.FindByID(id); err != nil {
		slog.Error("failed to call find id", "err", err)
		return nil, err
	} else if err := sgModel.FromEntity(sgEnt); err != nil {
		slog.Error("failed to get model from entity", "err", err)
		return nil, err
	} else if sEntities, err := d.reposMgr.Settlement.FindSettlementsBySettlementGroupID(uint(ID)); err != nil {
		slog.Error("failed to find settlements", "err", err)
		return nil, err
	} else if sModels, err := model.FromEntities(sEntities); err != nil {
		slog.Error("failed to transfrom entities", "err", err)
		return nil, err
	} else {

		cdtTrfTxIves := make([]*pacs_008_001_02.CreditTransferTransactionInformation11, len(sModels))

		for _, v := range sModels {
			cdtTrfTxIves = append(cdtTrfTxIves, v.CdtTrfTxInf)
		}

		doc := &pacs_008_001_02.Document{
			FIToFICstmrCdtTrf: &pacs_008_001_02.FIToFICustomerCreditTransferV02{
				GrpHdr:      sgModel.GrpHdr,
				CdtTrfTxInf: cdtTrfTxIves,
			},
		}

		return doc, nil
	}

}
