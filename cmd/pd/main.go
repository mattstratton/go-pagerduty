package main

import (
	"fmt"
	"os"

	"github.com/PagerDuty/go-pagerduty/internal/pkg/pagerduty"
	"github.com/mitchellh/cli"
)

const (
	version = "0.0.0"
)

func loadCommands() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"ability list":  pagerduty.AbilityListCommand,
		"ability test":  pagerduty.AbilityTestCommand,
		"addon list":    pagerduty.AddonListCommand,
		"addon install": pagerduty.AddonInstallCommand,
		"addon show":    pagerduty.AddonShowCommand,
		"addon delete":  pagerduty.AddonDeleteCommand,
		"addon update":  pagerduty.AddonUpdateCommand,

		"escalation-policy list":   pagerduty.EscalationPolicyListCommand,
		"escalation-policy create": pagerduty.EscalationPolicyCreateCommand,
		"escalation-policy delete": pagerduty.EscalationPolicyDeleteCommand,

		"escalation-policy show":   pagerduty.EscalationPolicyShowCommand,
		"escalation-policy update": pagerduty.EscalationPolicyUpdateCommand,

		"incident list":        pagerduty.IncidentListCommand,
		"incident manage":      pagerduty.IncidentManageCommand,
		"incident show":        pagerduty.IncidentShowCommand,
		"incident note list":   pagerduty.IncidentNoteListCommand,
		"incident note create": pagerduty.IncidentNoteCreateCommand,
		"incident snooze":      pagerduty.IncidentSnoozeCommand,

		"log-entry list": pagerduty.LogEntryListCommand,
		"log-entry show": pagerduty.LogEntryShowCommand,

		"maintenance-window list":   pagerduty.MaintenanceWindowListCommand,
		"maintenance-window create": pagerduty.MaintenanceWindowCreateCommand,
		"maintenance-window delete": pagerduty.MaintenanceWindowDeleteCommand,
		"maintenance-window show":   pagerduty.MaintenanceWindowShowCommand,
		"maintenance-window update": pagerduty.MaintenanceWindowUpdateCommand,

		"notification list": pagerduty.NotificationListCommand,

		"oncall list": pagerduty.OncallListCommand,

		"schedule list":    pagerduty.ScheduleListCommand,
		"schedule create":  pagerduty.ScheduleCreateCommand,
		"schedule preview": pagerduty.SchedulePreviewCommand,
		"schedule delete":  pagerduty.ScheduleDeleteCommand,
		"schedule show":    pagerduty.ScheduleShowCommand,
		"schedule update":  pagerduty.ScheduleUpdateCommand,

		"schedule override list":   pagerduty.ScheduleOverrideListCommand,
		"schedule override create": pagerduty.ScheduleOverrideCreateCommand,
		"schedule override delete": pagerduty.ScheduleOverrideDeleteCommand,

		"schedule oncall list": pagerduty.ScheduleOncallListCommand,

		"service list":               pagerduty.ServiceListCommand,
		"service create":             pagerduty.ServiceCreateCommand,
		"service delete":             pagerduty.ServiceDeleteCommand,
		"service show":               pagerduty.ServiceShowCommand,
		"service update":             pagerduty.ServiceUpdateCommand,
		"service integration create": pagerduty.ServiceIntegrationCreateCommand,
		"service integration show":   pagerduty.ServiceIntegrationShowCommand,
		"service integration update": pagerduty.ServiceIntegrationUpdateCommand,

		"team list":                     pagerduty.TeamListCommand,
		"team create":                   pagerduty.TeamShowCommand,
		"team delete":                   pagerduty.TeamDeleteCommand,
		"team show":                     pagerduty.TeamShowCommand,
		"team update":                   pagerduty.TeamUpdateCommand,
		"team remove escalation-policy": pagerduty.TeamRemoveEscalationPolicyCommand,
		"team add escalation-policy":    pagerduty.TeamAddEscalationPolicyCommand,
		"team add user":                 pagerduty.TeamAddUserCommand,

		"user list":                     pagerduty.UserListCommand,
		"user create":                   pagerduty.UserCreateCommand,
		"user delete":                   pagerduty.UserDeleteCommand,
		"user show":                     pagerduty.UserShowCommand,
		"user update":                   pagerduty.UserUpdateCommand,
		"user contact-method list":      pagerduty.UserContactMethodListCommand,
		"user contact-method create":    pagerduty.UserContactMethodCreateCommand,
		"user contact-method delete":    pagerduty.UserContactMethodDeleteCommand,
		"user contact-method show":      pagerduty.UserContactMethodShowCommand,
		"user contact-method update":    pagerduty.UserContactMethodUpdateCommand,
		"user notification-rule list":   pagerduty.UserNotificationRuleListCommand,
		"user notification-rule create": pagerduty.UserNotificationRuleCreateCommand,
		"user notification-rule delete": pagerduty.UserNotificationRuleDeleteCommand,
		"user notification-rule show":   pagerduty.UserNotificationRuleShowCommand,
		"user notification-rule update": pagerduty.UserNotificationRuleUpdateCommand,
	}
}

func main() {
	os.Exit(invokeCLI(os.Args[1:]))
}

func invokeCLI(args []string) int {
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			fmt.Println(version)
			return 0
		}
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: loadCommands(),
		HelpFunc: cli.BasicHelpFunc("pd"),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
