package user // Тот же пакет user

func displayUserPassword(u User) string {
	return u.password // Доступ к закрытому полю внутри того же пакета
}
