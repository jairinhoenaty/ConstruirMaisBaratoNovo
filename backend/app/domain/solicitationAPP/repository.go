package solicitationapp

type SolicitationAppRepository interface {
	Save(solicitation SolicitationApp) (*SolicitationApp, error)
	UpdateFeedback(idFirebase string, rating int, feedback string) (*SolicitationApp, error)
}
