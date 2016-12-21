package devices




import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*WeaveDevicesHandleInvitationOK Successful response

swagger:response weaveDevicesHandleInvitationOK
*/
type WeaveDevicesHandleInvitationOK struct {
}

// NewWeaveDevicesHandleInvitationOK creates WeaveDevicesHandleInvitationOK with default headers values
func NewWeaveDevicesHandleInvitationOK() *WeaveDevicesHandleInvitationOK {
	return &WeaveDevicesHandleInvitationOK{}
}

// WriteResponse to the client
func (o *WeaveDevicesHandleInvitationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}