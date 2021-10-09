package main

import "fmt"

type user struct {
	Name     string
	email    string
	id       string
	password string
}

type post struct {
	id      string
	cap     string
	img_url string
	time    string
}

func algo() {

	/*
				   string user_name,user_emal,user_pa,user_id_input;
				   ..............from this i get the user input from the instagram account

				   input_user := user{user_name,user_emal,user_id_input,user_pa}  //thorugh this I made the structure of the "user" .


			       string post_id ,post_cap, post_url,post_time;
				   .................from this i get the post information and store it in post structure;
		           input_post := post{post_id ,post_cap,post_url,post_time}  // through this I made the structure of the "post"

	*/
	u1 := user{"Brijesh", "sainibrijesh01@gmail.com", "saini01", "bri"}
	fmt.Println("User1: ", u1)

	u2 := user{"Brijesh", "sainibrijesh01@gmail.com", "rana02", "brijesh"}
	fmt.Println("User2: ", u2)

	u3 := user{"Deepak", "sainideepak@gmail.com", "kumar01", "brijesh"}
	fmt.Println("User2: ", u3)

	p1 := post{"rana02", "I am rana", "https://miro.medium.com/max/1400/1*mk1-6aYaf_Bes1E3Imhc0A.jpeg", "9th oct,2021 8:10:20"}
	fmt.Println("Post1: ", p1)

	p2 := post{"saini01", "I am Brijesh", "https://miro.medium.com/max/1400/1*mk1-6aYaf_Bes1E3Imhc0A.jpeg", "10th oct,2021 8:10:20"}
	fmt.Println("Post2: ", p2)

	p3 := post{"disha01", "I am disha", "https://images.ctfassets.net/hrltx12pl8hq/7yQR5uJhwEkRfjwMFJ7bUK/dc52a0913e8ff8b5c276177890eb0129/offset_comp_772626-opt.jpg?fit=fill&w=800&h=300", "10th oct,2021 8:00:20"}
	fmt.Println("Post3: ", p3)

	users := []user{u1, u2, u3}

	posts := []post{p1, p2, p3}

	//this for loop that I use match the user id with the post id and then all the posts related to that particular user will
	// all listed for that user .
	for _, us := range users {
		fmt.Printf("%#v\n", us)
		for _, po := range posts {

			if po.id == us.id {
				fmt.Printf("%#v\n", po)
			}
		}
	}
}
