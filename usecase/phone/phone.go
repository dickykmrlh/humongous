package phone

import (
	"fmt"
	"math/rand"

	"github.com/dickymrlh/humongous/domain/phone"
)

func PlayAroundWithPhone(pc *phone.PhoneCollection) {
	populatePhones(800, 5550000, 5650000, pc)
}

func populatePhones(area, start, stop int, pc *phone.PhoneCollection) {
	i := start
	for i < stop {
		country := rand.Intn(8)
		num := int64((country * 1e10) + (area * 1e7) + i)
		fullNumber := fmt.Sprintf("+%d %d-%d", country, area, i)
		err := pc.InsertOne(phone.Phone{
			ID: num,
			Components: phone.Component{
				Country: country,
				Area:    area,
				Prefix:  (i / 1e4),
				Number:  i,
			},
			Display: fullNumber,
		})
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			i++
			fmt.Printf("inserted number: %s\n", fullNumber)
		}
	}

	fmt.Println("DONE")
}
