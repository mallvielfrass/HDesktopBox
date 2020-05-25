package hdesktopbox

//StructAPI retun FXToken - auth hdbox(http://5.61.48.15/partner_api) token
type StructAPI struct {
	FXToken string
}

//API entry point
func API(FXToken string) *StructAPI {
	return &StructAPI{

		FXToken: FXToken,
	}
}
