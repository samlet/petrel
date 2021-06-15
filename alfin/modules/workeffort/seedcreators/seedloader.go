package seedcreators

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func LoadSeeds() {
	client, err := ent.Open("sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %!v(MISSING)", err)
	}
	defer client.Close()

	ctx := context.Background()
	ctx = ent.NewContext(ctx, client)
	ctx = cachecomp.NewContext(ctx)

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %!v(MISSING)", err)
	}

	if err := CreateUserPreference(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSecurityPermission(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContactMechTypePurpose(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSecurityGroup(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyStatus(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateCommunicationEventType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateCommunicationEventPrpTyp(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateParty(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateUserLogin(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateTemporalExpression(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyRelationshipType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusValidChange(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortAssoc(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyQualType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateRoleType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePerson(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyContactMech(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortPartyAssignment(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateEnumerationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyContentType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusItem(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContactMechPurposeType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateTermType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSecurityGroupPermission(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortSkillStandard(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffort(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateFixedAsset(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyIdentificationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSkillType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateUserLoginSecurityGroup(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContactMechType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortFixedAssetAssign(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyRole(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyClassificationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateTemporalExpressionAssoc(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateEnumeration(ctx); err != nil {
		log.Fatal(err)
	}
}
