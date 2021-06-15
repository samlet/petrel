package seedcreators

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func LoadSeeds(execUpdaters bool) {
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
	if err := CreateParty(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusValidChange(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSecurityGroup(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateFixedAsset(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateTemporalExpression(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateUserLoginSecurityGroup(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyContentType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateCommunicationEventType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyStatus(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSecurityPermission(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyClassificationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortAssoc(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffort(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyContactMech(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateCommunicationEventPrpTyp(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePerson(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContactMechType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateEnumerationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateEnumeration(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortFixedAssetAssign(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyIdentificationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSecurityGroupPermission(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortPartyAssignment(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateRoleType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyRole(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContactMechTypePurpose(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortSkillStandard(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusItem(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateUserLogin(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateTermType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateTemporalExpressionAssoc(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyRelationshipType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyQualType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSkillType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContactMechPurposeType(ctx); err != nil {
		log.Fatal(err)
	}
	if execUpdaters {
		if err := UpdateRoleType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePartyRole(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateEnumerationType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateEnumeration(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateWorkEffortFixedAssetAssign(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePartyIdentificationType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateSecurityGroupPermission(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateWorkEffortPartyAssignment(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateUserLogin(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateContactMechTypePurpose(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateWorkEffortSkillStandard(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateStatusItem(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateContactMechPurposeType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateTermType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateTemporalExpressionAssoc(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePartyRelationshipType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePartyQualType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateSkillType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateUserLoginSecurityGroup(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateWorkEffortType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateUserPreference(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateParty(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateStatusValidChange(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateSecurityGroup(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateFixedAsset(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateTemporalExpression(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateCommunicationEventType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePartyContentType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateSecurityPermission(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePartyStatus(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePartyContactMech(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateStatusType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePartyClassificationType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateWorkEffortAssoc(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateWorkEffort(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateContactMechType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdateCommunicationEventPrpTyp(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePartyType(ctx); err != nil {
			log.Fatal(err)
		}
		if err := UpdatePerson(ctx); err != nil {
			log.Fatal(err)
		}
	}
}
