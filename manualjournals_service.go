package reeleezee

func NewManualJournalsService(api *API) *ManualJournalsService {
	return &ManualJournalsService{api: api}
}

type ManualJournalsService struct {
	api *API
}
