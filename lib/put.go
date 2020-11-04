
package lib

import (
	Photo "../models/photo"
	PutService "../services"
)

func Put() {

	photo := Photo.Photo{
		1,
		1,
		"accusamus beatae ad facilis cum similique qui sunt",
		"https://via.placeholder.com/600/92c952",
		"https://via.placeholder.com/150/92c952",
	}

	PutService.PutPhoto(photo)

}
