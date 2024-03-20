package app

type App struct {
	Env *Env
}

func New() *App {
	cfg := &App{}
	cfg.Env = NewEnv()
	return cfg
}
