package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspAccountsPostResponse_onboardingTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Latest available onboarding Template
    latestVersion *string
    // A short message describing the features available in the new onboarding Template
    message *string
    // True if the account needs to be updated with the latest version
    upgradeNeeded *bool
    // Version of the onboarding Template applied on the account.
    versionApplied *string
}
// NewV1beta1CspAccountsPostResponse_onboardingTemplate instantiates a new V1beta1CspAccountsPostResponse_onboardingTemplate and sets the default values.
func NewV1beta1CspAccountsPostResponse_onboardingTemplate()(*V1beta1CspAccountsPostResponse_onboardingTemplate) {
    m := &V1beta1CspAccountsPostResponse_onboardingTemplate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsPostResponse_onboardingTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsPostResponse_onboardingTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsPostResponse_onboardingTemplate(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["latestVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestVersion(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["upgradeNeeded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeNeeded(val)
        }
        return nil
    }
    res["versionApplied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionApplied(val)
        }
        return nil
    }
    return res
}
// GetLatestVersion gets the latestVersion property value. Latest available onboarding Template
// returns a *string when successful
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) GetLatestVersion()(*string) {
    return m.latestVersion
}
// GetMessage gets the message property value. A short message describing the features available in the new onboarding Template
// returns a *string when successful
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) GetMessage()(*string) {
    return m.message
}
// GetUpgradeNeeded gets the upgradeNeeded property value. True if the account needs to be updated with the latest version
// returns a *bool when successful
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) GetUpgradeNeeded()(*bool) {
    return m.upgradeNeeded
}
// GetVersionApplied gets the versionApplied property value. Version of the onboarding Template applied on the account.
// returns a *string when successful
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) GetVersionApplied()(*string) {
    return m.versionApplied
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("latestVersion", m.GetLatestVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("upgradeNeeded", m.GetUpgradeNeeded())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("versionApplied", m.GetVersionApplied())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLatestVersion sets the latestVersion property value. Latest available onboarding Template
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) SetLatestVersion(value *string)() {
    m.latestVersion = value
}
// SetMessage sets the message property value. A short message describing the features available in the new onboarding Template
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) SetMessage(value *string)() {
    m.message = value
}
// SetUpgradeNeeded sets the upgradeNeeded property value. True if the account needs to be updated with the latest version
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) SetUpgradeNeeded(value *bool)() {
    m.upgradeNeeded = value
}
// SetVersionApplied sets the versionApplied property value. Version of the onboarding Template applied on the account.
func (m *V1beta1CspAccountsPostResponse_onboardingTemplate) SetVersionApplied(value *string)() {
    m.versionApplied = value
}
type V1beta1CspAccountsPostResponse_onboardingTemplateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLatestVersion()(*string)
    GetMessage()(*string)
    GetUpgradeNeeded()(*bool)
    GetVersionApplied()(*string)
    SetLatestVersion(value *string)()
    SetMessage(value *string)()
    SetUpgradeNeeded(value *bool)()
    SetVersionApplied(value *string)()
}
