package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse a subnet whose values are defined by, and synchronized with, a cloud service provider.
type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse struct {
    // Information about the cloud provider account where the resource is located.
    accountInfo V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_accountInfoable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Resource identifier assigned by the cloud provider.
    cspId *string
    // The cspInfo property
    cspInfo V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoable
    // Resource name assigned by the cloud provider, if any.
    cspName *string
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // An identifier for the resource, usually a UUID.
    id *string
    // A system specified name for the resource.
    name *string
    // The self reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
    // Information about the VPC to which the subnet is associated.
    vpcInfo V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_vpcInfoable
}
// V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo composed type wrapper for classes V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able, V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able
type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo struct {
    // Composed type representation for type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able
    v1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1 V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able
    // Composed type representation for type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able
    v1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2 V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able
}
// NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo instantiates a new V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo and sets the default values.
func NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo()(*V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo) {
    m := &V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo{
    }
    return m
}
// CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    if m.GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1() != nil {
        return m.GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1().GetFieldDeserializers()
    } else if m.GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2() != nil {
        return m.GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2().GetFieldDeserializers()
    }
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo) GetIsComposedType()(bool) {
    return true
}
// GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1 gets the V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1 property value. Composed type representation for type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo) GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able) {
    return m.v1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1
}
// GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2 gets the V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2 property value. Composed type representation for type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo) GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able) {
    return m.v1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1 sets the V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1 property value. Composed type representation for type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo) SetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able)() {
    m.v1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1 = value
}
// SetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2 sets the V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2 property value. Composed type representation for type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfo) SetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able)() {
    m.v1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2 = value
}
type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able)
    GetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able)
    SetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember1(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember1able)()
    SetV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseCspInfoMember2(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_cspInfoMember2able)()
}
// NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse instantiates a new V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse and sets the default values.
func NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse()(*V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) {
    m := &V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse(), nil
}
// GetAccountInfo gets the accountInfo property value. Information about the cloud provider account where the resource is located.
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_accountInfoable when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetAccountInfo()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_accountInfoable) {
    return m.accountInfo
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCspId gets the cspId property value. Resource identifier assigned by the cloud provider.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetCspId()(*string) {
    return m.cspId
}
// GetCspInfo gets the cspInfo property value. The cspInfo property
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoable when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetCspInfo()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoable) {
    return m.cspInfo
}
// GetCspName gets the cspName property value. Resource name assigned by the cloud provider, if any.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetCspName()(*string) {
    return m.cspName
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_accountInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountInfo(val.(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_accountInfoable))
        }
        return nil
    }
    res["cspId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspId(val)
        }
        return nil
    }
    res["cspInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspInfo(val.(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoable))
        }
        return nil
    }
    res["cspName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspName(val)
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["generation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
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
    res["vpcInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_vpcInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpcInfo(val.(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_vpcInfoable))
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetVpcInfo gets the vpcInfo property value. Information about the VPC to which the subnet is associated.
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_vpcInfoable when successful
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) GetVpcInfo()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_vpcInfoable) {
    return m.vpcInfo
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cspInfo", m.GetCspInfo())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountInfo sets the accountInfo property value. Information about the cloud provider account where the resource is located.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetAccountInfo(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_accountInfoable)() {
    m.accountInfo = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCspId sets the cspId property value. Resource identifier assigned by the cloud provider.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetCspId(value *string)() {
    m.cspId = value
}
// SetCspInfo sets the cspInfo property value. The cspInfo property
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetCspInfo(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoable)() {
    m.cspInfo = value
}
// SetCspName sets the cspName property value. Resource name assigned by the cloud provider, if any.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetCspName(value *string)() {
    m.cspName = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetVpcInfo sets the vpcInfo property value. Information about the VPC to which the subnet is associated.
func (m *V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse) SetVpcInfo(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_vpcInfoable)() {
    m.vpcInfo = value
}
type V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountInfo()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_accountInfoable)
    GetCspId()(*string)
    GetCspInfo()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoable)
    GetCspName()(*string)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    GetVpcInfo()(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_vpcInfoable)
    SetAccountInfo(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_accountInfoable)()
    SetCspId(value *string)()
    SetCspInfo(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_WithSubnetGetResponse_cspInfoable)()
    SetCspName(value *string)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
    SetVpcInfo(value V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse_vpcInfoable)()
}
