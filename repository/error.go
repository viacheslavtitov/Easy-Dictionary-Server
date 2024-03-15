package repository

type encodingError string

func (e encodingError) Error() string {
	return string(e)
}
