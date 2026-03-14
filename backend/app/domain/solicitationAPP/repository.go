package solicitationapp

type SolicitationAppRepository interface {
	Save(solicitation SolicitationApp) (*SolicitationApp, error)
}
