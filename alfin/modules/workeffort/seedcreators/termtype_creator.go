package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateTermType(ctx context.Context) error {
	log.Println("TermType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.TermType

	c, err = client.TermType.Create().SetStringRef("financial_term__termtype").
		SetHasTable("No").
		SetDescription("Financial").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create financial_term__termtype: %v", err)
		return err
	}
	cache.Put("financial_term__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_payment_term__termtype").
		SetHasTable("No").
		SetDescription("Payment (net days)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_payment_term__termtype: %v", err)
		return err
	}
	cache.Put("fin_payment_term__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_pay_netdays_1__termtype").
		SetHasTable("No").
		SetDescription("Payment net days, part 1").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_pay_netdays_1__termtype: %v", err)
		return err
	}
	cache.Put("fin_pay_netdays_1__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_pay_netdays_2__termtype").
		SetHasTable("No").
		SetDescription("Payment net days, part 2").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_pay_netdays_2__termtype: %v", err)
		return err
	}
	cache.Put("fin_pay_netdays_2__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_pay_netdays_3__termtype").
		SetHasTable("No").
		SetDescription("Payment net days, part 3").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_pay_netdays_3__termtype: %v", err)
		return err
	}
	cache.Put("fin_pay_netdays_3__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_payment_disc__termtype").
		SetHasTable("No").
		SetDescription("Payment (discounted if paid within specified days)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_payment_disc__termtype: %v", err)
		return err
	}
	cache.Put("fin_payment_disc__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_payment_fixday__termtype").
		SetHasTable("No").
		SetDescription("Payment (due on specified day of month)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_payment_fixday__termtype: %v", err)
		return err
	}
	cache.Put("fin_payment_fixday__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_late_fee_term__termtype").
		SetHasTable("No").
		SetDescription("Late Fee (percent)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_late_fee_term__termtype: %v", err)
		return err
	}
	cache.Put("fin_late_fee_term__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_collect_term__termtype").
		SetHasTable("No").
		SetDescription("Penalty For Collection Agency").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_collect_term__termtype: %v", err)
		return err
	}
	cache.Put("fin_collect_term__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_nortn_item_term__termtype").
		SetHasTable("No").
		SetDescription("Non-Returnable Sales Item").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_nortn_item_term__termtype: %v", err)
		return err
	}
	cache.Put("fin_nortn_item_term__termtype", c)

	c, err = client.TermType.Create().SetStringRef("incentive__termtype").
		SetHasTable("No").
		SetDescription("Incentive").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create incentive__termtype: %v", err)
		return err
	}
	cache.Put("incentive__termtype", c)

	c, err = client.TermType.Create().SetStringRef("legal_term__termtype").
		SetHasTable("No").
		SetDescription("Legal").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create legal_term__termtype: %v", err)
		return err
	}
	cache.Put("legal_term__termtype", c)

	c, err = client.TermType.Create().SetStringRef("threshold__termtype").
		SetHasTable("No").
		SetDescription("Threshold").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create threshold__termtype: %v", err)
		return err
	}
	cache.Put("threshold__termtype", c)

	c, err = client.TermType.Create().SetStringRef("clause_for_renewal__termtype").
		SetHasTable("No").
		SetDescription("Clause For Renewal").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create clause_for_renewal__termtype: %v", err)
		return err
	}
	cache.Put("clause_for_renewal__termtype", c)

	c, err = client.TermType.Create().SetStringRef("agreement_terminatio__termtype").
		SetHasTable("No").
		SetDescription("Agreement Termination").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create agreement_terminatio__termtype: %v", err)
		return err
	}
	cache.Put("agreement_terminatio__termtype", c)

	c, err = client.TermType.Create().SetStringRef("indemnification__termtype").
		SetHasTable("No").
		SetDescription("Indemnification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create indemnification__termtype: %v", err)
		return err
	}
	cache.Put("indemnification__termtype", c)

	c, err = client.TermType.Create().SetStringRef("non_compete__termtype").
		SetHasTable("No").
		SetDescription("Non-Compete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create non_compete__termtype: %v", err)
		return err
	}
	cache.Put("non_compete__termtype", c)

	c, err = client.TermType.Create().SetStringRef("exclusive_relationsh__termtype").
		SetHasTable("No").
		SetDescription("Exclusive Relationship").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create exclusive_relationsh__termtype: %v", err)
		return err
	}
	cache.Put("exclusive_relationsh__termtype", c)

	c, err = client.TermType.Create().SetStringRef("commission_term__termtype").
		SetHasTable("No").
		SetDescription("Commission").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create commission_term__termtype: %v", err)
		return err
	}
	cache.Put("commission_term__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_comm_fixed__termtype").
		SetHasTable("No").
		SetDescription("Commission Term Fixed Per Unit").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_comm_fixed__termtype: %v", err)
		return err
	}
	cache.Put("fin_comm_fixed__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_comm_variable__termtype").
		SetHasTable("No").
		SetDescription("Commission Term Variable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_comm_variable__termtype: %v", err)
		return err
	}
	cache.Put("fin_comm_variable__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_comm_min__termtype").
		SetHasTable("No").
		SetDescription("Commission Term Minimum Per Unit").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_comm_min__termtype: %v", err)
		return err
	}
	cache.Put("fin_comm_min__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fin_comm_max__termtype").
		SetHasTable("No").
		SetDescription("Commission Term Maximum Per Unit").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fin_comm_max__termtype: %v", err)
		return err
	}
	cache.Put("fin_comm_max__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term__termtype").
		SetHasTable("No").
		SetDescription("Incoterm").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term__termtype: %v", err)
		return err
	}
	cache.Put("inco_term__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020__termtype").
		SetDescription("Incoterm 2020").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_exw__termtype").
		SetDescription("Incoterm 2020 Ex Works").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_exw__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_exw__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_fca__termtype").
		SetDescription("Incoterm 2020 Free Carrier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_fca__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_fca__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_fas__termtype").
		SetDescription("Incoterm 2020 Free Alongside Ship").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_fas__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_fas__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_fob__termtype").
		SetDescription("Incoterm 2020 Free On Board").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_fob__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_fob__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_cpt__termtype").
		SetDescription("Incoterm 2020 Carriage Paid To").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_cpt__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_cpt__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_cfr__termtype").
		SetDescription("Incoterm 2020 Cost and Freight").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_cfr__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_cfr__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_cif__termtype").
		SetDescription("Incoterm 2020 Cost, Insurance and Freight").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_cif__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_cif__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_cip__termtype").
		SetDescription("Incoterm 2020 Carriage and Insurance Paid to").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_cip__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_cip__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_dpu__termtype").
		SetDescription("Incoterm 2020 Delivered at Place Unloaded").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_dpu__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_dpu__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_dap__termtype").
		SetDescription("Incoterm 2020 Delivered at Place").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_dap__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_dap__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2020_ddp__termtype").
		SetDescription("Incoterm 2020 Delivered Duty Paid").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2020_ddp__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2020_ddp__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2010__termtype").
		SetDescription("Incoterm 2010").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2010__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2010__termtype", c)

	c, err = client.TermType.Create().SetStringRef("exw__termtype").
		SetDescription("Incoterm 2010 Ex Works").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create exw__termtype: %v", err)
		return err
	}
	cache.Put("exw__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fca__termtype").
		SetDescription("Incoterm 2010 Free Carrier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fca__termtype: %v", err)
		return err
	}
	cache.Put("fca__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fas__termtype").
		SetDescription("Incoterm 2010 Free Alongside Ship").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fas__termtype: %v", err)
		return err
	}
	cache.Put("fas__termtype", c)

	c, err = client.TermType.Create().SetStringRef("fob__termtype").
		SetDescription("Incoterm 2010 Free On Board").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fob__termtype: %v", err)
		return err
	}
	cache.Put("fob__termtype", c)

	c, err = client.TermType.Create().SetStringRef("cfr__termtype").
		SetDescription("Incoterm 2010 Cost and Freight").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cfr__termtype: %v", err)
		return err
	}
	cache.Put("cfr__termtype", c)

	c, err = client.TermType.Create().SetStringRef("cif__termtype").
		SetDescription("Incoterm 2010 Cost, Insurance and Freight").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cif__termtype: %v", err)
		return err
	}
	cache.Put("cif__termtype", c)

	c, err = client.TermType.Create().SetStringRef("cpt__termtype").
		SetDescription("Incoterm 2010 Carriage Paid To").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cpt__termtype: %v", err)
		return err
	}
	cache.Put("cpt__termtype", c)

	c, err = client.TermType.Create().SetStringRef("cip__termtype").
		SetDescription("Incoterm 2010 Carriage and Insurance Paid to").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cip__termtype: %v", err)
		return err
	}
	cache.Put("cip__termtype", c)

	c, err = client.TermType.Create().SetStringRef("ddp__termtype").
		SetDescription("Incoterm 2010 Delivered Duty Paid").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ddp__termtype: %v", err)
		return err
	}
	cache.Put("ddp__termtype", c)

	c, err = client.TermType.Create().SetStringRef("inco_term_2000__termtype").
		SetDescription("Incoterm 2000").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inco_term_2000__termtype: %v", err)
		return err
	}
	cache.Put("inco_term_2000__termtype", c)

	c, err = client.TermType.Create().SetStringRef("daf__termtype").
		SetDescription("Incoterm 2000 Delivered At Frontier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create daf__termtype: %v", err)
		return err
	}
	cache.Put("daf__termtype", c)

	c, err = client.TermType.Create().SetStringRef("des__termtype").
		SetDescription("Incoterm 2000 Delivered Ex Ship").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create des__termtype: %v", err)
		return err
	}
	cache.Put("des__termtype", c)

	c, err = client.TermType.Create().SetStringRef("deq__termtype").
		SetDescription("Incoterm 2000 Delivered Ex Quay").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create deq__termtype: %v", err)
		return err
	}
	cache.Put("deq__termtype", c)

	c, err = client.TermType.Create().SetStringRef("ddu__termtype").
		SetDescription("Incoterm 2000 Delivered Duty Unpaid").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ddu__termtype: %v", err)
		return err
	}
	cache.Put("ddu__termtype", c)

	c, err = client.TermType.Create().SetStringRef("purchasing__termtype").
		SetDescription("Purchasing").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purchasing__termtype: %v", err)
		return err
	}
	cache.Put("purchasing__termtype", c)

	c, err = client.TermType.Create().SetStringRef("purch_vendor_id__termtype").
		SetDescription("Vendor Customer ID").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_vendor_id__termtype: %v", err)
		return err
	}
	cache.Put("purch_vendor_id__termtype", c)

	c, err = client.TermType.Create().SetStringRef("purch_freight__termtype").
		SetDescription("Preferred Freight").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_freight__termtype: %v", err)
		return err
	}
	cache.Put("purch_freight__termtype", c)

	c, err = client.TermType.Create().SetStringRef("other_term__termtype").
		SetHasTable("No").
		SetDescription("Other").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_term__termtype: %v", err)
		return err
	}
	cache.Put("other_term__termtype", c)

	return nil
}
