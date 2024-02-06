package main

type user struct {
	name  string
	admin bool
}

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

// defer keyword used in a function
func logAndDelete(users map[string]user, name string) (log string) {
	user, ok := users[name]
	defer delete(users, name)
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}

	return logDeleted
}

func main() {

}
