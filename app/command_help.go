// Copyright (c) 2019-present AboutOne, Inc. All Rights Reserved.


package app

import (
	goi18n "github.com/mattermost/go-i18n/i18n"
	"za-white-screen/model"
)

type HelpProvider struct {
}

const (
	CMD_HELP = "help"
)

func init() {
	RegisterCommandProvider(&HelpProvider{})
}

func (h *HelpProvider) GetTrigger() string {
	return CMD_HELP
}

func (h *HelpProvider) GetCommand(a *App, T goi18n.TranslateFunc) *model.Command {
	return &model.Command{
		Trigger:          CMD_HELP,
		AutoComplete:     true,
		AutoCompleteDesc: T("api.command_help.desc"),
		DisplayName:      T("api.command_help.name"),
	}
}

func (h *HelpProvider) DoCommand(a *App, args *model.CommandArgs, message string) *model.CommandResponse {
	helpLink := *a.Config().SupportSettings.HelpLink

	if helpLink == "" {
		helpLink = model.SUPPORT_SETTINGS_DEFAULT_HELP_LINK
	}

	return &model.CommandResponse{GotoLocation: helpLink}
}
