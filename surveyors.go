package mineorea

import (

)

type surveyor struct {
	name string
}

func Surveyor() *surveyor {
	faker := newFaker()

	return &surveyor{name: faker.Name()}
}
