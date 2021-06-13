
package helper

import (
    "context"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent"
    "log"

    "github.com/samlet/petrel/alfin/modules/workeffort/ent/communicationeventprptyp"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/communicationeventtype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechpurposetype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechtype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechtypepurpose"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/enumeration"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/enumerationtype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/fixedasset"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/party"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/partyclassificationtype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/partycontactmech"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/partycontenttype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/partyidentificationtype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/partyqualtype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/partyrelationshiptype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/partyrole"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/partystatus"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/partytype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/person"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/roletype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygroup"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygrouppermission"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/securitypermission"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/skilltype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/statusitem"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/statustype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/statusvalidchange"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/temporalexpression"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/temporalexpressionassoc"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/termtype"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/userlogin"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/userloginsecuritygroup"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/userpreference"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffort"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortassoc"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortfixedassetassign"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortpartyassignment"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortskillstandard"
    "github.com/samlet/petrel/alfin/modules/workeffort/ent/workefforttype"
)


func WorkEffortSkillStandardRef(ctx context.Context, stringId string) *ent.WorkEffortSkillStandard {
    client:=ent.FromContext(ctx)
    rec, err := client.WorkEffortSkillStandard.
        Query().
        Where(workeffortskillstandard.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find WorkEffortSkillStandard %s: %v", stringId, err)
    }
    return rec
}

func StatusValidChangeRef(ctx context.Context, stringId string) *ent.StatusValidChange {
    client:=ent.FromContext(ctx)
    rec, err := client.StatusValidChange.
        Query().
        Where(statusvalidchange.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find StatusValidChange %s: %v", stringId, err)
    }
    return rec
}

func PartyRelationshipTypeRef(ctx context.Context, stringId string) *ent.PartyRelationshipType {
    client:=ent.FromContext(ctx)
    rec, err := client.PartyRelationshipType.
        Query().
        Where(partyrelationshiptype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find PartyRelationshipType %s: %v", stringId, err)
    }
    return rec
}

func ContactMechPurposeTypeRef(ctx context.Context, stringId string) *ent.ContactMechPurposeType {
    client:=ent.FromContext(ctx)
    rec, err := client.ContactMechPurposeType.
        Query().
        Where(contactmechpurposetype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find ContactMechPurposeType %s: %v", stringId, err)
    }
    return rec
}

func PartyStatusRef(ctx context.Context, stringId string) *ent.PartyStatus {
    client:=ent.FromContext(ctx)
    rec, err := client.PartyStatus.
        Query().
        Where(partystatus.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find PartyStatus %s: %v", stringId, err)
    }
    return rec
}

func UserLoginRef(ctx context.Context, stringId string) *ent.UserLogin {
    client:=ent.FromContext(ctx)
    rec, err := client.UserLogin.
        Query().
        Where(userlogin.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find UserLogin %s: %v", stringId, err)
    }
    return rec
}

func SecurityGroupRef(ctx context.Context, stringId string) *ent.SecurityGroup {
    client:=ent.FromContext(ctx)
    rec, err := client.SecurityGroup.
        Query().
        Where(securitygroup.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find SecurityGroup %s: %v", stringId, err)
    }
    return rec
}

func SecurityPermissionRef(ctx context.Context, stringId string) *ent.SecurityPermission {
    client:=ent.FromContext(ctx)
    rec, err := client.SecurityPermission.
        Query().
        Where(securitypermission.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find SecurityPermission %s: %v", stringId, err)
    }
    return rec
}

func CommunicationEventTypeRef(ctx context.Context, stringId string) *ent.CommunicationEventType {
    client:=ent.FromContext(ctx)
    rec, err := client.CommunicationEventType.
        Query().
        Where(communicationeventtype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find CommunicationEventType %s: %v", stringId, err)
    }
    return rec
}

func PartyIdentificationTypeRef(ctx context.Context, stringId string) *ent.PartyIdentificationType {
    client:=ent.FromContext(ctx)
    rec, err := client.PartyIdentificationType.
        Query().
        Where(partyidentificationtype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find PartyIdentificationType %s: %v", stringId, err)
    }
    return rec
}

func PartyClassificationTypeRef(ctx context.Context, stringId string) *ent.PartyClassificationType {
    client:=ent.FromContext(ctx)
    rec, err := client.PartyClassificationType.
        Query().
        Where(partyclassificationtype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find PartyClassificationType %s: %v", stringId, err)
    }
    return rec
}

func WorkEffortAssocRef(ctx context.Context, stringId string) *ent.WorkEffortAssoc {
    client:=ent.FromContext(ctx)
    rec, err := client.WorkEffortAssoc.
        Query().
        Where(workeffortassoc.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find WorkEffortAssoc %s: %v", stringId, err)
    }
    return rec
}

func TemporalExpressionRef(ctx context.Context, stringId string) *ent.TemporalExpression {
    client:=ent.FromContext(ctx)
    rec, err := client.TemporalExpression.
        Query().
        Where(temporalexpression.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find TemporalExpression %s: %v", stringId, err)
    }
    return rec
}

func PartyContentTypeRef(ctx context.Context, stringId string) *ent.PartyContentType {
    client:=ent.FromContext(ctx)
    rec, err := client.PartyContentType.
        Query().
        Where(partycontenttype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find PartyContentType %s: %v", stringId, err)
    }
    return rec
}

func UserPreferenceRef(ctx context.Context, stringId string) *ent.UserPreference {
    client:=ent.FromContext(ctx)
    rec, err := client.UserPreference.
        Query().
        Where(userpreference.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find UserPreference %s: %v", stringId, err)
    }
    return rec
}

func ContactMechTypePurposeRef(ctx context.Context, stringId string) *ent.ContactMechTypePurpose {
    client:=ent.FromContext(ctx)
    rec, err := client.ContactMechTypePurpose.
        Query().
        Where(contactmechtypepurpose.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find ContactMechTypePurpose %s: %v", stringId, err)
    }
    return rec
}

func SecurityGroupPermissionRef(ctx context.Context, stringId string) *ent.SecurityGroupPermission {
    client:=ent.FromContext(ctx)
    rec, err := client.SecurityGroupPermission.
        Query().
        Where(securitygrouppermission.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find SecurityGroupPermission %s: %v", stringId, err)
    }
    return rec
}

func StatusItemRef(ctx context.Context, stringId string) *ent.StatusItem {
    client:=ent.FromContext(ctx)
    rec, err := client.StatusItem.
        Query().
        Where(statusitem.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find StatusItem %s: %v", stringId, err)
    }
    return rec
}

func EnumerationTypeRef(ctx context.Context, stringId string) *ent.EnumerationType {
    client:=ent.FromContext(ctx)
    rec, err := client.EnumerationType.
        Query().
        Where(enumerationtype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find EnumerationType %s: %v", stringId, err)
    }
    return rec
}

func TemporalExpressionAssocRef(ctx context.Context, stringId string) *ent.TemporalExpressionAssoc {
    client:=ent.FromContext(ctx)
    rec, err := client.TemporalExpressionAssoc.
        Query().
        Where(temporalexpressionassoc.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find TemporalExpressionAssoc %s: %v", stringId, err)
    }
    return rec
}

func PartyContactMechRef(ctx context.Context, stringId string) *ent.PartyContactMech {
    client:=ent.FromContext(ctx)
    rec, err := client.PartyContactMech.
        Query().
        Where(partycontactmech.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find PartyContactMech %s: %v", stringId, err)
    }
    return rec
}

func WorkEffortTypeRef(ctx context.Context, stringId string) *ent.WorkEffortType {
    client:=ent.FromContext(ctx)
    rec, err := client.WorkEffortType.
        Query().
        Where(workefforttype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find WorkEffortType %s: %v", stringId, err)
    }
    return rec
}

func PartyTypeRef(ctx context.Context, stringId string) *ent.PartyType {
    client:=ent.FromContext(ctx)
    rec, err := client.PartyType.
        Query().
        Where(partytype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find PartyType %s: %v", stringId, err)
    }
    return rec
}

func TermTypeRef(ctx context.Context, stringId string) *ent.TermType {
    client:=ent.FromContext(ctx)
    rec, err := client.TermType.
        Query().
        Where(termtype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find TermType %s: %v", stringId, err)
    }
    return rec
}

func CommunicationEventPrpTypRef(ctx context.Context, stringId string) *ent.CommunicationEventPrpTyp {
    client:=ent.FromContext(ctx)
    rec, err := client.CommunicationEventPrpTyp.
        Query().
        Where(communicationeventprptyp.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find CommunicationEventPrpTyp %s: %v", stringId, err)
    }
    return rec
}

func PartyRoleRef(ctx context.Context, stringId string) *ent.PartyRole {
    client:=ent.FromContext(ctx)
    rec, err := client.PartyRole.
        Query().
        Where(partyrole.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find PartyRole %s: %v", stringId, err)
    }
    return rec
}

func WorkEffortFixedAssetAssignRef(ctx context.Context, stringId string) *ent.WorkEffortFixedAssetAssign {
    client:=ent.FromContext(ctx)
    rec, err := client.WorkEffortFixedAssetAssign.
        Query().
        Where(workeffortfixedassetassign.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find WorkEffortFixedAssetAssign %s: %v", stringId, err)
    }
    return rec
}

func RoleTypeRef(ctx context.Context, stringId string) *ent.RoleType {
    client:=ent.FromContext(ctx)
    rec, err := client.RoleType.
        Query().
        Where(roletype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find RoleType %s: %v", stringId, err)
    }
    return rec
}

func StatusTypeRef(ctx context.Context, stringId string) *ent.StatusType {
    client:=ent.FromContext(ctx)
    rec, err := client.StatusType.
        Query().
        Where(statustype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find StatusType %s: %v", stringId, err)
    }
    return rec
}

func PartyRef(ctx context.Context, stringId string) *ent.Party {
    client:=ent.FromContext(ctx)
    rec, err := client.Party.
        Query().
        Where(party.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find Party %s: %v", stringId, err)
    }
    return rec
}

func PersonRef(ctx context.Context, stringId string) *ent.Person {
    client:=ent.FromContext(ctx)
    rec, err := client.Person.
        Query().
        Where(person.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find Person %s: %v", stringId, err)
    }
    return rec
}

func PartyQualTypeRef(ctx context.Context, stringId string) *ent.PartyQualType {
    client:=ent.FromContext(ctx)
    rec, err := client.PartyQualType.
        Query().
        Where(partyqualtype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find PartyQualType %s: %v", stringId, err)
    }
    return rec
}

func FixedAssetRef(ctx context.Context, stringId string) *ent.FixedAsset {
    client:=ent.FromContext(ctx)
    rec, err := client.FixedAsset.
        Query().
        Where(fixedasset.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find FixedAsset %s: %v", stringId, err)
    }
    return rec
}

func SkillTypeRef(ctx context.Context, stringId string) *ent.SkillType {
    client:=ent.FromContext(ctx)
    rec, err := client.SkillType.
        Query().
        Where(skilltype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find SkillType %s: %v", stringId, err)
    }
    return rec
}

func WorkEffortRef(ctx context.Context, stringId string) *ent.WorkEffort {
    client:=ent.FromContext(ctx)
    rec, err := client.WorkEffort.
        Query().
        Where(workeffort.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find WorkEffort %s: %v", stringId, err)
    }
    return rec
}

func UserLoginSecurityGroupRef(ctx context.Context, stringId string) *ent.UserLoginSecurityGroup {
    client:=ent.FromContext(ctx)
    rec, err := client.UserLoginSecurityGroup.
        Query().
        Where(userloginsecuritygroup.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find UserLoginSecurityGroup %s: %v", stringId, err)
    }
    return rec
}

func EnumerationRef(ctx context.Context, stringId string) *ent.Enumeration {
    client:=ent.FromContext(ctx)
    rec, err := client.Enumeration.
        Query().
        Where(enumeration.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find Enumeration %s: %v", stringId, err)
    }
    return rec
}

func WorkEffortPartyAssignmentRef(ctx context.Context, stringId string) *ent.WorkEffortPartyAssignment {
    client:=ent.FromContext(ctx)
    rec, err := client.WorkEffortPartyAssignment.
        Query().
        Where(workeffortpartyassignment.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find WorkEffortPartyAssignment %s: %v", stringId, err)
    }
    return rec
}

func ContactMechTypeRef(ctx context.Context, stringId string) *ent.ContactMechType {
    client:=ent.FromContext(ctx)
    rec, err := client.ContactMechType.
        Query().
        Where(contactmechtype.StringRefEQ(stringId)).
        // `Only` fails if no record found,
        // or more than 1 record returned.
        Only(ctx)
    if err != nil {
        log.Fatalf("Cannot find ContactMechType %s: %v", stringId, err)
    }
    return rec
}


