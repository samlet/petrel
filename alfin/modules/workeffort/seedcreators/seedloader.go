package seedcreators

import (
	"context"
	"github.com/fatih/color"
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

	if err := CreateTemporalExpression(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusItem(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateTermType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateUserPreference(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateTemporalExpressionAssoc(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyRelationshipType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyContentType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSecurityPermission(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateFixedAsset(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContactMechType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSkillType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContactMechTypePurpose(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSecurityGroupPermission(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortFixedAssetAssign(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateCommunicationEventPrpTyp(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyIdentificationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateCommunicationEventType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyQualType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyStatus(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffort(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortSkillStandard(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateUserLogin(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSecurityGroup(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusValidChange(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortAssoc(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyRole(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContactMechPurposeType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateEnumerationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateWorkEffortPartyAssignment(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyClassificationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateEnumeration(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateRoleType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateUserLoginSecurityGroup(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateParty(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePerson(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreatePartyContactMech(ctx); err != nil {
		log.Fatal(err)
	}
	if execUpdaters {
		prompt := color.New(color.FgGreen).PrintfFunc()
		warn := color.New(color.FgRed).PrintfFunc()
		if err := UpdateWorkEffortAssoc(ctx); err != nil {
			log.Println(err)
			warn("Update WorkEffortAssoc fail: %v\n", err)
		} else {
			prompt("Update WorkEffortAssoc ok.\n")
		}
		if err := UpdatePartyRole(ctx); err != nil {
			log.Println(err)
			warn("Update PartyRole fail: %v\n", err)
		} else {
			prompt("Update PartyRole ok.\n")
		}
		if err := UpdateContactMechPurposeType(ctx); err != nil {
			log.Println(err)
			warn("Update ContactMechPurposeType fail: %v\n", err)
		} else {
			prompt("Update ContactMechPurposeType ok.\n")
		}
		if err := UpdateEnumerationType(ctx); err != nil {
			log.Println(err)
			warn("Update EnumerationType fail: %v\n", err)
		} else {
			prompt("Update EnumerationType ok.\n")
		}
		if err := UpdateWorkEffortPartyAssignment(ctx); err != nil {
			log.Println(err)
			warn("Update WorkEffortPartyAssignment fail: %v\n", err)
		} else {
			prompt("Update WorkEffortPartyAssignment ok.\n")
		}
		if err := UpdatePartyClassificationType(ctx); err != nil {
			log.Println(err)
			warn("Update PartyClassificationType fail: %v\n", err)
		} else {
			prompt("Update PartyClassificationType ok.\n")
		}
		if err := UpdateEnumeration(ctx); err != nil {
			log.Println(err)
			warn("Update Enumeration fail: %v\n", err)
		} else {
			prompt("Update Enumeration ok.\n")
		}
		if err := UpdateRoleType(ctx); err != nil {
			log.Println(err)
			warn("Update RoleType fail: %v\n", err)
		} else {
			prompt("Update RoleType ok.\n")
		}
		if err := UpdatePartyType(ctx); err != nil {
			log.Println(err)
			warn("Update PartyType fail: %v\n", err)
		} else {
			prompt("Update PartyType ok.\n")
		}
		if err := UpdateParty(ctx); err != nil {
			log.Println(err)
			warn("Update Party fail: %v\n", err)
		} else {
			prompt("Update Party ok.\n")
		}
		if err := UpdatePerson(ctx); err != nil {
			log.Println(err)
			warn("Update Person fail: %v\n", err)
		} else {
			prompt("Update Person ok.\n")
		}
		if err := UpdatePartyContactMech(ctx); err != nil {
			log.Println(err)
			warn("Update PartyContactMech fail: %v\n", err)
		} else {
			prompt("Update PartyContactMech ok.\n")
		}
		if err := UpdateUserLoginSecurityGroup(ctx); err != nil {
			log.Println(err)
			warn("Update UserLoginSecurityGroup fail: %v\n", err)
		} else {
			prompt("Update UserLoginSecurityGroup ok.\n")
		}
		if err := UpdateStatusItem(ctx); err != nil {
			log.Println(err)
			warn("Update StatusItem fail: %v\n", err)
		} else {
			prompt("Update StatusItem ok.\n")
		}
		if err := UpdateTermType(ctx); err != nil {
			log.Println(err)
			warn("Update TermType fail: %v\n", err)
		} else {
			prompt("Update TermType ok.\n")
		}
		if err := UpdateUserPreference(ctx); err != nil {
			log.Println(err)
			warn("Update UserPreference fail: %v\n", err)
		} else {
			prompt("Update UserPreference ok.\n")
		}
		if err := UpdateTemporalExpressionAssoc(ctx); err != nil {
			log.Println(err)
			warn("Update TemporalExpressionAssoc fail: %v\n", err)
		} else {
			prompt("Update TemporalExpressionAssoc ok.\n")
		}
		if err := UpdateTemporalExpression(ctx); err != nil {
			log.Println(err)
			warn("Update TemporalExpression fail: %v\n", err)
		} else {
			prompt("Update TemporalExpression ok.\n")
		}
		if err := UpdatePartyContentType(ctx); err != nil {
			log.Println(err)
			warn("Update PartyContentType fail: %v\n", err)
		} else {
			prompt("Update PartyContentType ok.\n")
		}
		if err := UpdateSecurityPermission(ctx); err != nil {
			log.Println(err)
			warn("Update SecurityPermission fail: %v\n", err)
		} else {
			prompt("Update SecurityPermission ok.\n")
		}
		if err := UpdateFixedAsset(ctx); err != nil {
			log.Println(err)
			warn("Update FixedAsset fail: %v\n", err)
		} else {
			prompt("Update FixedAsset ok.\n")
		}
		if err := UpdateContactMechType(ctx); err != nil {
			log.Println(err)
			warn("Update ContactMechType fail: %v\n", err)
		} else {
			prompt("Update ContactMechType ok.\n")
		}
		if err := UpdateSkillType(ctx); err != nil {
			log.Println(err)
			warn("Update SkillType fail: %v\n", err)
		} else {
			prompt("Update SkillType ok.\n")
		}
		if err := UpdateContactMechTypePurpose(ctx); err != nil {
			log.Println(err)
			warn("Update ContactMechTypePurpose fail: %v\n", err)
		} else {
			prompt("Update ContactMechTypePurpose ok.\n")
		}
		if err := UpdateStatusType(ctx); err != nil {
			log.Println(err)
			warn("Update StatusType fail: %v\n", err)
		} else {
			prompt("Update StatusType ok.\n")
		}
		if err := UpdatePartyRelationshipType(ctx); err != nil {
			log.Println(err)
			warn("Update PartyRelationshipType fail: %v\n", err)
		} else {
			prompt("Update PartyRelationshipType ok.\n")
		}
		if err := UpdateCommunicationEventPrpTyp(ctx); err != nil {
			log.Println(err)
			warn("Update CommunicationEventPrpTyp fail: %v\n", err)
		} else {
			prompt("Update CommunicationEventPrpTyp ok.\n")
		}
		if err := UpdatePartyIdentificationType(ctx); err != nil {
			log.Println(err)
			warn("Update PartyIdentificationType fail: %v\n", err)
		} else {
			prompt("Update PartyIdentificationType ok.\n")
		}
		if err := UpdateCommunicationEventType(ctx); err != nil {
			log.Println(err)
			warn("Update CommunicationEventType fail: %v\n", err)
		} else {
			prompt("Update CommunicationEventType ok.\n")
		}
		if err := UpdateWorkEffortType(ctx); err != nil {
			log.Println(err)
			warn("Update WorkEffortType fail: %v\n", err)
		} else {
			prompt("Update WorkEffortType ok.\n")
		}
		if err := UpdatePartyQualType(ctx); err != nil {
			log.Println(err)
			warn("Update PartyQualType fail: %v\n", err)
		} else {
			prompt("Update PartyQualType ok.\n")
		}
		if err := UpdateSecurityGroupPermission(ctx); err != nil {
			log.Println(err)
			warn("Update SecurityGroupPermission fail: %v\n", err)
		} else {
			prompt("Update SecurityGroupPermission ok.\n")
		}
		if err := UpdateWorkEffortFixedAssetAssign(ctx); err != nil {
			log.Println(err)
			warn("Update WorkEffortFixedAssetAssign fail: %v\n", err)
		} else {
			prompt("Update WorkEffortFixedAssetAssign ok.\n")
		}
		if err := UpdateWorkEffort(ctx); err != nil {
			log.Println(err)
			warn("Update WorkEffort fail: %v\n", err)
		} else {
			prompt("Update WorkEffort ok.\n")
		}
		if err := UpdateWorkEffortSkillStandard(ctx); err != nil {
			log.Println(err)
			warn("Update WorkEffortSkillStandard fail: %v\n", err)
		} else {
			prompt("Update WorkEffortSkillStandard ok.\n")
		}
		if err := UpdateUserLogin(ctx); err != nil {
			log.Println(err)
			warn("Update UserLogin fail: %v\n", err)
		} else {
			prompt("Update UserLogin ok.\n")
		}
		if err := UpdateSecurityGroup(ctx); err != nil {
			log.Println(err)
			warn("Update SecurityGroup fail: %v\n", err)
		} else {
			prompt("Update SecurityGroup ok.\n")
		}
		if err := UpdatePartyStatus(ctx); err != nil {
			log.Println(err)
			warn("Update PartyStatus fail: %v\n", err)
		} else {
			prompt("Update PartyStatus ok.\n")
		}
		if err := UpdateStatusValidChange(ctx); err != nil {
			log.Println(err)
			warn("Update StatusValidChange fail: %v\n", err)
		} else {
			prompt("Update StatusValidChange ok.\n")
		}
	}
}
