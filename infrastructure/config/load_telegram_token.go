package iconfig

func (c *Config) LoadTelegramToken() {
	tg := &c.Watcher.Telegram

	if tg.Token == "" {
		tg.Token = secretFileRead(tg.TokenFile)
	}
}
