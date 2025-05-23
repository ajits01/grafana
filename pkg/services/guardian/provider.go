package guardian

import (
	"context"

	"github.com/grafana/grafana/pkg/apimachinery/identity"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/services/dashboards"
	"github.com/grafana/grafana/pkg/services/folder"
	"github.com/grafana/grafana/pkg/services/team"
	"github.com/grafana/grafana/pkg/setting"
)

type Provider struct{}

func ProvideService(
	cfg *setting.Cfg, ac accesscontrol.AccessControl,
	dashboardService dashboards.DashboardService, teamService team.Service,
	folderService folder.Service,
) *Provider {
	// TODO: Fix this hack, see https://github.com/grafana/grafana-enterprise/issues/2935
	InitAccessControlGuardian(cfg, ac, dashboardService, folderService, log.New("guardian"))
	return &Provider{}
}

func InitAccessControlGuardian(
	cfg *setting.Cfg, ac accesscontrol.AccessControl, dashboardService dashboards.DashboardService, folderService folder.Service, logger log.Logger,
) {
	New = func(ctx context.Context, dashId int64, orgId int64, user identity.Requester) (DashboardGuardian, error) {
		return NewAccessControlDashboardGuardian(ctx, cfg, dashId, user, ac, dashboardService, folderService, logger)
	}

	NewByFolder = func(ctx context.Context, f *folder.Folder, orgId int64, user identity.Requester) (DashboardGuardian, error) {
		return NewAccessControlFolderGuardian(ctx, cfg, f, user, ac, orgId, dashboardService, folderService)
	}
}
