package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // IQN of the adapter
    iqn *string
    // Model of the adapter
    model *string
    // Name of the adapter
    name *string
    // Type of the adapter
    typeEscaped *string
    // WWN of the adapter
    wwn *string
}
// NewV1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo instantiates a new V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo and sets the default values.
func NewV1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo()(*V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) {
    m := &V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["iqn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIqn(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["wwn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWwn(val)
        }
        return nil
    }
    return res
}
// GetIqn gets the iqn property value. IQN of the adapter
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) GetIqn()(*string) {
    return m.iqn
}
// GetModel gets the model property value. Model of the adapter
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) GetModel()(*string) {
    return m.model
}
// GetName gets the name property value. Name of the adapter
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the type property value. Type of the adapter
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetWwn gets the wwn property value. WWN of the adapter
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) GetWwn()(*string) {
    return m.wwn
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("iqn", m.GetIqn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wwn", m.GetWwn())
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
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIqn sets the iqn property value. IQN of the adapter
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) SetIqn(value *string)() {
    m.iqn = value
}
// SetModel sets the model property value. Model of the adapter
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) SetModel(value *string)() {
    m.model = value
}
// SetName sets the name property value. Name of the adapter
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the type property value. Type of the adapter
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetWwn sets the wwn property value. WWN of the adapter
func (m *V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfo) SetWwn(value *string)() {
    m.wwn = value
}
type V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIqn()(*string)
    GetModel()(*string)
    GetName()(*string)
    GetTypeEscaped()(*string)
    GetWwn()(*string)
    SetIqn(value *string)()
    SetModel(value *string)()
    SetName(value *string)()
    SetTypeEscaped(value *string)()
    SetWwn(value *string)()
}
