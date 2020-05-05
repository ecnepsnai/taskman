package taskman

// ActionFunc describes the function signature for an action function
type ActionFunc func(userData interface{}) error

// UserDataFunc describes the function signature for a method that returns the user data for this action
type UserDataFunc func(taskName string) (interface{}, error)

// Action describes an action to take when a trigger fires
type Action struct {
	UserDataProvider UserDataFunc
	Func             ActionFunc
	ErrorIsFatal     bool
}

func (a Action) runNow(taskName string) error {
	var userData interface{}
	if a.UserDataProvider != nil {
		d, err := a.UserDataProvider(taskName)
		if err != nil {
			log.Error("Error calling user data provider for task '%s': %s", taskName, err.Error())
			if a.ErrorIsFatal {
				return err
			}
			return nil
		}
		userData = d
	}

	if err := a.Func(userData); err != nil {
		log.Error("Error running action on task '%s': %s", taskName, err.Error())
		if a.ErrorIsFatal {
			return err
		}
	}

	return nil
}
