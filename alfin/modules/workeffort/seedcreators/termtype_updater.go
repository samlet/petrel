package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateTermType(ctx context.Context) error {
	log.Println("TermType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.TermType
	failures := 0

	c = cache.Get("financial_term__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Financial").
		AddChildren(cache.Get("fin_payment_term__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_payment_disc__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_payment_fixday__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_late_fee_term__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_collect_term__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_nortn_item_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update financial_term__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("financial_term__termtype", c)
	}

	c = cache.Get("fin_payment_term__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Payment (net days)").
		SetParent(cache.Get("financial_term__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_pay_netdays_1__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_pay_netdays_2__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_pay_netdays_3__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_payment_term__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_payment_term__termtype", c)
	}

	c = cache.Get("fin_pay_netdays_1__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Payment net days, part 1").
		SetParent(cache.Get("fin_payment_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_pay_netdays_1__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_pay_netdays_1__termtype", c)
	}

	c = cache.Get("fin_pay_netdays_2__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Payment net days, part 2").
		SetParent(cache.Get("fin_payment_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_pay_netdays_2__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_pay_netdays_2__termtype", c)
	}

	c = cache.Get("fin_pay_netdays_3__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Payment net days, part 3").
		SetParent(cache.Get("fin_payment_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_pay_netdays_3__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_pay_netdays_3__termtype", c)
	}

	c = cache.Get("fin_payment_disc__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Payment (discounted if paid within specified days)").
		SetParent(cache.Get("financial_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_payment_disc__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_payment_disc__termtype", c)
	}

	c = cache.Get("fin_payment_fixday__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Payment (due on specified day of month)").
		SetParent(cache.Get("financial_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_payment_fixday__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_payment_fixday__termtype", c)
	}

	c = cache.Get("fin_late_fee_term__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Late Fee (percent)").
		SetParent(cache.Get("financial_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_late_fee_term__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_late_fee_term__termtype", c)
	}

	c = cache.Get("fin_collect_term__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Penalty For Collection Agency").
		SetParent(cache.Get("financial_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_collect_term__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_collect_term__termtype", c)
	}

	c = cache.Get("fin_nortn_item_term__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Non-Returnable Sales Item").
		SetParent(cache.Get("financial_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_nortn_item_term__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_nortn_item_term__termtype", c)
	}

	c = cache.Get("incentive__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Incentive").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update incentive__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("incentive__termtype", c)
	}

	c = cache.Get("legal_term__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Legal").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update legal_term__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("legal_term__termtype", c)
	}

	c = cache.Get("threshold__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Threshold").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update threshold__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("threshold__termtype", c)
	}

	c = cache.Get("clause_for_renewal__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Clause For Renewal").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update clause_for_renewal__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("clause_for_renewal__termtype", c)
	}

	c = cache.Get("agreement_terminatio__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Agreement Termination").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update agreement_terminatio__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("agreement_terminatio__termtype", c)
	}

	c = cache.Get("indemnification__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Indemnification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update indemnification__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("indemnification__termtype", c)
	}

	c = cache.Get("non_compete__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Non-Compete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update non_compete__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("non_compete__termtype", c)
	}

	c = cache.Get("exclusive_relationsh__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Exclusive Relationship").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update exclusive_relationsh__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("exclusive_relationsh__termtype", c)
	}

	c = cache.Get("commission_term__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Commission").
		AddChildren(cache.Get("fin_comm_fixed__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_comm_variable__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_comm_min__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fin_comm_max__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update commission_term__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("commission_term__termtype", c)
	}

	c = cache.Get("fin_comm_fixed__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Commission Term Fixed Per Unit").
		SetParent(cache.Get("commission_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_comm_fixed__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_comm_fixed__termtype", c)
	}

	c = cache.Get("fin_comm_variable__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Commission Term Variable").
		SetParent(cache.Get("commission_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_comm_variable__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_comm_variable__termtype", c)
	}

	c = cache.Get("fin_comm_min__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Commission Term Minimum Per Unit").
		SetParent(cache.Get("commission_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_comm_min__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_comm_min__termtype", c)
	}

	c = cache.Get("fin_comm_max__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Commission Term Maximum Per Unit").
		SetParent(cache.Get("commission_term__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fin_comm_max__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fin_comm_max__termtype", c)
	}

	c = cache.Get("inco_term__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Incoterm").
		AddChildren(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2000__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term__termtype", c)
	}

	c = cache.Get("inco_term_2020__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020").
		SetParent(cache.Get("inco_term__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_exw__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_fca__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_fas__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_fob__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_cpt__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_cfr__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_cif__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_cip__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_dpu__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_dap__termtype").(*ent.TermType)).
		AddChildren(cache.Get("inco_term_2020_ddp__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020__termtype", c)
	}

	c = cache.Get("inco_term_2020_exw__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Ex Works").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_exw__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_exw__termtype", c)
	}

	c = cache.Get("inco_term_2020_fca__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Free Carrier").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_fca__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_fca__termtype", c)
	}

	c = cache.Get("inco_term_2020_fas__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Free Alongside Ship").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_fas__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_fas__termtype", c)
	}

	c = cache.Get("inco_term_2020_fob__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Free On Board").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_fob__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_fob__termtype", c)
	}

	c = cache.Get("inco_term_2020_cpt__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Carriage Paid To").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_cpt__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_cpt__termtype", c)
	}

	c = cache.Get("inco_term_2020_cfr__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Cost and Freight").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_cfr__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_cfr__termtype", c)
	}

	c = cache.Get("inco_term_2020_cif__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Cost, Insurance and Freight").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_cif__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_cif__termtype", c)
	}

	c = cache.Get("inco_term_2020_cip__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Carriage and Insurance Paid to").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_cip__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_cip__termtype", c)
	}

	c = cache.Get("inco_term_2020_dpu__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Delivered at Place Unloaded").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_dpu__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_dpu__termtype", c)
	}

	c = cache.Get("inco_term_2020_dap__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Delivered at Place").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_dap__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_dap__termtype", c)
	}

	c = cache.Get("inco_term_2020_ddp__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2020 Delivered Duty Paid").
		SetParent(cache.Get("inco_term_2020__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2020_ddp__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2020_ddp__termtype", c)
	}

	c = cache.Get("inco_term_2010__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010").
		SetParent(cache.Get("inco_term__termtype").(*ent.TermType)).
		AddChildren(cache.Get("exw__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fca__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fas__termtype").(*ent.TermType)).
		AddChildren(cache.Get("fob__termtype").(*ent.TermType)).
		AddChildren(cache.Get("cfr__termtype").(*ent.TermType)).
		AddChildren(cache.Get("cif__termtype").(*ent.TermType)).
		AddChildren(cache.Get("cpt__termtype").(*ent.TermType)).
		AddChildren(cache.Get("cip__termtype").(*ent.TermType)).
		AddChildren(cache.Get("ddp__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2010__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2010__termtype", c)
	}

	c = cache.Get("exw__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010 Ex Works").
		SetParent(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update exw__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("exw__termtype", c)
	}

	c = cache.Get("fca__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010 Free Carrier").
		SetParent(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fca__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fca__termtype", c)
	}

	c = cache.Get("fas__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010 Free Alongside Ship").
		SetParent(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fas__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fas__termtype", c)
	}

	c = cache.Get("fob__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010 Free On Board").
		SetParent(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fob__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fob__termtype", c)
	}

	c = cache.Get("cfr__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010 Cost and Freight").
		SetParent(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cfr__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cfr__termtype", c)
	}

	c = cache.Get("cif__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010 Cost, Insurance and Freight").
		SetParent(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cif__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cif__termtype", c)
	}

	c = cache.Get("cpt__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010 Carriage Paid To").
		SetParent(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cpt__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cpt__termtype", c)
	}

	c = cache.Get("cip__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010 Carriage and Insurance Paid to").
		SetParent(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cip__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cip__termtype", c)
	}

	c = cache.Get("ddp__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2010 Delivered Duty Paid").
		SetParent(cache.Get("inco_term_2010__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ddp__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ddp__termtype", c)
	}

	c = cache.Get("inco_term_2000__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2000").
		SetParent(cache.Get("inco_term__termtype").(*ent.TermType)).
		AddChildren(cache.Get("daf__termtype").(*ent.TermType)).
		AddChildren(cache.Get("des__termtype").(*ent.TermType)).
		AddChildren(cache.Get("deq__termtype").(*ent.TermType)).
		AddChildren(cache.Get("ddu__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inco_term_2000__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inco_term_2000__termtype", c)
	}

	c = cache.Get("daf__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2000 Delivered At Frontier").
		SetParent(cache.Get("inco_term_2000__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update daf__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("daf__termtype", c)
	}

	c = cache.Get("des__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2000 Delivered Ex Ship").
		SetParent(cache.Get("inco_term_2000__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update des__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("des__termtype", c)
	}

	c = cache.Get("deq__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2000 Delivered Ex Quay").
		SetParent(cache.Get("inco_term_2000__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update deq__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("deq__termtype", c)
	}

	c = cache.Get("ddu__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Incoterm 2000 Delivered Duty Unpaid").
		SetParent(cache.Get("inco_term_2000__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ddu__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ddu__termtype", c)
	}

	c = cache.Get("purchasing__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Purchasing").
		AddChildren(cache.Get("purch_vendor_id__termtype").(*ent.TermType)).
		AddChildren(cache.Get("purch_freight__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purchasing__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purchasing__termtype", c)
	}

	c = cache.Get("purch_vendor_id__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Vendor Customer ID").
		SetParent(cache.Get("purchasing__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_vendor_id__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_vendor_id__termtype", c)
	}

	c = cache.Get("purch_freight__termtype").(*ent.TermType)
	c, err = c.Update().
		SetDescription("Preferred Freight").
		SetParent(cache.Get("purchasing__termtype").(*ent.TermType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_freight__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_freight__termtype", c)
	}

	c = cache.Get("other_term__termtype").(*ent.TermType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Other").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update other_term__termtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("other_term__termtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
