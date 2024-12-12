package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody attrbutes of Data Services Collector which are available for update.
type V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // User-defined name given to the Data Services Connector by the customer.
    displayName *string
    // List of configured DNS servers configured on the Data Services Connector.
    dns []V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_dnsable
    // NTP servers against which the Data Services Connector should synchronize.
    ntp []V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpable
    // The proxy property
    proxy V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_proxyable
    // List of search domains.
    searchDomains []string
}
// NewV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody instantiates a new V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody and sets the default values.
func NewV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody()(*V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) {
    m := &V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. User-defined name given to the Data Services Connector by the customer.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetDns gets the dns property value. List of configured DNS servers configured on the Data Services Connector.
// returns a []V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_dnsable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) GetDns()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_dnsable) {
    return m.dns
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["dns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_dnsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_dnsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_dnsable)
                }
            }
            m.SetDns(res)
        }
        return nil
    }
    res["ntp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpable)
                }
            }
            m.SetNtp(res)
        }
        return nil
    }
    res["proxy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_proxyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxy(val.(V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_proxyable))
        }
        return nil
    }
    res["searchDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSearchDomains(res)
        }
        return nil
    }
    return res
}
// GetNtp gets the ntp property value. NTP servers against which the Data Services Connector should synchronize.
// returns a []V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) GetNtp()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpable) {
    return m.ntp
}
// GetProxy gets the proxy property value. The proxy property
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_proxyable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) GetProxy()(V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_proxyable) {
    return m.proxy
}
// GetSearchDomains gets the searchDomains property value. List of search domains.
// returns a []string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) GetSearchDomains()([]string) {
    return m.searchDomains
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDns()))
        for i, v := range m.GetDns() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("dns", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNtp() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNtp()))
        for i, v := range m.GetNtp() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("ntp", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("proxy", m.GetProxy())
        if err != nil {
            return err
        }
    }
    if m.GetSearchDomains() != nil {
        err := writer.WriteCollectionOfStringValues("searchDomains", m.GetSearchDomains())
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
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. User-defined name given to the Data Services Connector by the customer.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDns sets the dns property value. List of configured DNS servers configured on the Data Services Connector.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) SetDns(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_dnsable)() {
    m.dns = value
}
// SetNtp sets the ntp property value. NTP servers against which the Data Services Connector should synchronize.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) SetNtp(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpable)() {
    m.ntp = value
}
// SetProxy sets the proxy property value. The proxy property
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) SetProxy(value V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_proxyable)() {
    m.proxy = value
}
// SetSearchDomains sets the searchDomains property value. List of search domains.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody) SetSearchDomains(value []string)() {
    m.searchDomains = value
}
type V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetDns()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_dnsable)
    GetNtp()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpable)
    GetProxy()(V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_proxyable)
    GetSearchDomains()([]string)
    SetDisplayName(value *string)()
    SetDns(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_dnsable)()
    SetNtp(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_ntpable)()
    SetProxy(value V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBody_proxyable)()
    SetSearchDomains(value []string)()
}
