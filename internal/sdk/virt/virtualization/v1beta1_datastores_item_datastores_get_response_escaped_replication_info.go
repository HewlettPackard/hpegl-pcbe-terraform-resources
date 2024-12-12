package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresItemDatastoresGetResponse_replicationInfo replication groups information containing details of all replication partners. Applicable only for Protection Group type 'STORAGE_REPLICATION_GROUP'
type V1beta1DatastoresItemDatastoresGetResponse_replicationInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Id of the replication group.
    id *string
    // Name of the replication group.
    name *string
    // List of Volumes associated with vm protection group.
    partnerDetails []V1beta1DatastoresItemDatastoresGetResponse_replicationInfo_partnerDetailsable
    // Uri representing replication group in Storage Fleet
    resourceUri *string
    // type representing volume-set in Storage Fleet
    typeEscaped *string
}
// NewV1beta1DatastoresItemDatastoresGetResponse_replicationInfo instantiates a new V1beta1DatastoresItemDatastoresGetResponse_replicationInfo and sets the default values.
func NewV1beta1DatastoresItemDatastoresGetResponse_replicationInfo()(*V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) {
    m := &V1beta1DatastoresItemDatastoresGetResponse_replicationInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresItemDatastoresGetResponse_replicationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresItemDatastoresGetResponse_replicationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresItemDatastoresGetResponse_replicationInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["partnerDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresItemDatastoresGetResponse_replicationInfo_partnerDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresItemDatastoresGetResponse_replicationInfo_partnerDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresItemDatastoresGetResponse_replicationInfo_partnerDetailsable)
                }
            }
            m.SetPartnerDetails(res)
        }
        return nil
    }
    res["resourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceUri(val)
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
    return res
}
// GetId gets the id property value. Id of the replication group.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the replication group.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) GetName()(*string) {
    return m.name
}
// GetPartnerDetails gets the partnerDetails property value. List of Volumes associated with vm protection group.
// returns a []V1beta1DatastoresItemDatastoresGetResponse_replicationInfo_partnerDetailsable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) GetPartnerDetails()([]V1beta1DatastoresItemDatastoresGetResponse_replicationInfo_partnerDetailsable) {
    return m.partnerDetails
}
// GetResourceUri gets the resourceUri property value. Uri representing replication group in Storage Fleet
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. type representing volume-set in Storage Fleet
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
    if m.GetPartnerDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPartnerDetails()))
        for i, v := range m.GetPartnerDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("partnerDetails", cast)
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
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Id of the replication group.
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the replication group.
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) SetName(value *string)() {
    m.name = value
}
// SetPartnerDetails sets the partnerDetails property value. List of Volumes associated with vm protection group.
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) SetPartnerDetails(value []V1beta1DatastoresItemDatastoresGetResponse_replicationInfo_partnerDetailsable)() {
    m.partnerDetails = value
}
// SetResourceUri sets the resourceUri property value. Uri representing replication group in Storage Fleet
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. type representing volume-set in Storage Fleet
func (m *V1beta1DatastoresItemDatastoresGetResponse_replicationInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1DatastoresItemDatastoresGetResponse_replicationInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    GetPartnerDetails()([]V1beta1DatastoresItemDatastoresGetResponse_replicationInfo_partnerDetailsable)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetId(value *string)()
    SetName(value *string)()
    SetPartnerDetails(value []V1beta1DatastoresItemDatastoresGetResponse_replicationInfo_partnerDetailsable)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
