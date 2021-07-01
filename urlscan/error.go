package urlscan

func isNotFoundError(err error) bool {
	return err.Error() == "status: 404"
}
