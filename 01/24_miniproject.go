package main

type Room struct {
	ID int
	Type string
	Status bool
	BedCount int
	Price int
}

func main(){

}




func generateRooms() []Room{
	var rooms []Room = []Room{}
	rooms = append(rooms, Room{ID:1, Type: "Single", Status: true, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID:2, Type: "Single", Status: true, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID:3, Type: "Single", Status: true, BedCount: 2, Price: 200})
	rooms = append(rooms, Room{ID:4, Type: "Single", Status: true, BedCount: 2, Price: 100})
	rooms = append(rooms, Room{ID:5, Type: "Double", Status: true, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID:6, Type: "Double", Status: true, BedCount: 3, Price: 300})
	rooms = append(rooms, Room{ID:7, Type: "Double", Status: true, BedCount: 3, Price: 100})
	rooms = append(rooms, Room{ID:8, Type: "Standard", Status: true, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID:9, Type: "Standard", Status: true, BedCount: 4, Price: 100})
	rooms = append(rooms, Room{ID:10, Type: "Standard", Status: true, BedCount: 4, Price: 100})

	return rooms

}



