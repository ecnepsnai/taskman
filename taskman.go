/*
Package taskman is an advanced task scheduler for go applications. It's built to model the Windows Task Scheduler
service and provide much greater control than what is offered by something like a cron expression.
*/
package taskman

import "github.com/ecnepsnai/logtic"

var log *logtic.Source

func init() {
	log = logtic.Connect("taskman")
}
