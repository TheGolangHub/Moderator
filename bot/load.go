package bot

import "github.com/PaulSonOfLars/gotgbot/v2/ext"

func Load(d *ext.Dispatcher) {
	loadAdmins(d)
	loadRules(d)
	loadBlist(d)
}

const (
	RulesGroup     = 1
	NormTalkGroup  = 2
	ProfanityGroup = 3
)
