package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1AnnouncementsItemWithAnnouncementGetResponse all possible properties of an announcement resource.
type V1beta1AnnouncementsItemWithAnnouncementGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The body of an announcement, may contain rich media.
    content *string
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // An identifier for the resource, usually a UUID.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // List of images associated with the announcement.
    images []V1beta1AnnouncementsItemWithAnnouncementGetResponse_imagesable
    // A system specified name for the resource.
    name *string
    // The self reference for this resource.
    resourceUri *string
    // List of product services associated with the announcement.
    services []V1beta1AnnouncementsItemWithAnnouncementGetResponse_servicesable
    // A sentence summarising an announcement.
    title *string
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1AnnouncementsItemWithAnnouncementGetResponse instantiates a new V1beta1AnnouncementsItemWithAnnouncementGetResponse and sets the default values.
func NewV1beta1AnnouncementsItemWithAnnouncementGetResponse()(*V1beta1AnnouncementsItemWithAnnouncementGetResponse) {
    m := &V1beta1AnnouncementsItemWithAnnouncementGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1AnnouncementsItemWithAnnouncementGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AnnouncementsItemWithAnnouncementGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AnnouncementsItemWithAnnouncementGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The body of an announcement, may contain rich media.
// returns a *string when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetContent()(*string) {
    return m.content
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
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
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["images"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1AnnouncementsItemWithAnnouncementGetResponse_imagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1AnnouncementsItemWithAnnouncementGetResponse_imagesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1AnnouncementsItemWithAnnouncementGetResponse_imagesable)
                }
            }
            m.SetImages(res)
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
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1AnnouncementsItemWithAnnouncementGetResponse_servicesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1AnnouncementsItemWithAnnouncementGetResponse_servicesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1AnnouncementsItemWithAnnouncementGetResponse_servicesable)
                }
            }
            m.SetServices(res)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
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
    res["updatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *UUID when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetImages gets the images property value. List of images associated with the announcement.
// returns a []V1beta1AnnouncementsItemWithAnnouncementGetResponse_imagesable when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetImages()([]V1beta1AnnouncementsItemWithAnnouncementGetResponse_imagesable) {
    return m.images
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of product services associated with the announcement.
// returns a []V1beta1AnnouncementsItemWithAnnouncementGetResponse_servicesable when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetServices()([]V1beta1AnnouncementsItemWithAnnouncementGetResponse_servicesable) {
    return m.services
}
// GetTitle gets the title property value. A sentence summarising an announcement.
// returns a *string when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetTitle()(*string) {
    return m.title
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    if m.GetImages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImages()))
        for i, v := range m.GetImages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("images", cast)
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
    if m.GetServices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedAt", m.GetUpdatedAt())
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
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The body of an announcement, may contain rich media.
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetContent(value *string)() {
    m.content = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetImages sets the images property value. List of images associated with the announcement.
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetImages(value []V1beta1AnnouncementsItemWithAnnouncementGetResponse_imagesable)() {
    m.images = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of product services associated with the announcement.
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetServices(value []V1beta1AnnouncementsItemWithAnnouncementGetResponse_servicesable)() {
    m.services = value
}
// SetTitle sets the title property value. A sentence summarising an announcement.
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetTitle(value *string)() {
    m.title = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1AnnouncementsItemWithAnnouncementGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1AnnouncementsItemWithAnnouncementGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetImages()([]V1beta1AnnouncementsItemWithAnnouncementGetResponse_imagesable)
    GetName()(*string)
    GetResourceUri()(*string)
    GetServices()([]V1beta1AnnouncementsItemWithAnnouncementGetResponse_servicesable)
    GetTitle()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetContent(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetImages(value []V1beta1AnnouncementsItemWithAnnouncementGetResponse_imagesable)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetServices(value []V1beta1AnnouncementsItemWithAnnouncementGetResponse_servicesable)()
    SetTitle(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
