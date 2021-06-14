package seedtypes
import "github.com/samlet/petrel/services"


type UserLogin struct {
    UserLoginId *string `json:"userLoginId,omitempty" xml:"userLoginId,attr"`
    CurrentPassword *string `json:"currentPassword,omitempty" xml:"currentPassword,attr"`
    PasswordHint *string `json:"passwordHint,omitempty" xml:"passwordHint,attr"`
    IsSystem *byte `json:"isSystem,omitempty" xml:"isSystem,attr"`
    Enabled *byte `json:"enabled,omitempty" xml:"enabled,attr"`
    HasLoggedOut *byte `json:"hasLoggedOut,omitempty" xml:"hasLoggedOut,attr"`
    RequirePasswordChange *byte `json:"requirePasswordChange,omitempty" xml:"requirePasswordChange,attr"`
    LastCurrencyUom *string `json:"lastCurrencyUom,omitempty" xml:"lastCurrencyUom,attr"`
    LastLocale *string `json:"lastLocale,omitempty" xml:"lastLocale,attr"`
    LastTimeZone *string `json:"lastTimeZone,omitempty" xml:"lastTimeZone,attr"`
    DisabledDateTime *services.DateTime `json:"disabledDateTime,omitempty" xml:"disabledDateTime,attr"`
    SuccessiveFailedLogins *int64 `json:"successiveFailedLogins,omitempty" xml:"successiveFailedLogins,attr"`
    ExternalAuthId *string `json:"externalAuthId,omitempty" xml:"externalAuthId,attr"`
    UserLdapDn *string `json:"userLdapDn,omitempty" xml:"userLdapDn,attr"`
    DisabledBy *string `json:"disabledBy,omitempty" xml:"disabledBy,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
    PartyId *string `json:"partyId,omitempty" xml:"partyId,attr"`
}

