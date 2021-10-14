package bot

import "github.com/PaulSonOfLars/gotgbot/v2/ext"

func Load(d *ext.Dispatcher) {
	loadRules(d)
}
