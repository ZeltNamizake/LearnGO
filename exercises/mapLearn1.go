package main

import "fmt"

func main() {
    
    person := map[string]interface{}{
        "name": "Driyas",
        "age":  18,
        "hobbies": []string{"coding", "gaming"},
        "address": map[string]interface{}{
            "city": "Bangkok",
            "country": "Thailand",
            "landmarks": []string{"What Pho", "Grand Palace"},
        },
        "friends": []map[string]interface{}{
           {"name": "Budi", "age": 21},
           {"name": "Siti", "age": 19},
        },
    }

    // 1. access simple property
    name := person["name"].(string)
    age := person["age"].(int)
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)

    // 2. access array
    hobbies := person["hobbies"].([]string)
    fmt.Println("First hobby:", hobbies[0])
    fmt.Println("All hobbies:")
    for i, h := range hobbies {
        fmt.Printf("%d: %s\n", i, h)
    }

    // 3. Access nested object
    address := person["address"].(map[string]interface{})
    city := address["city"].(string)
    country := address["country"].(string)
    landmarks := address["landmarks"].([]string)
    fmt.Println("City:", city)
    fmt.Println("Country:", country)
    fmt.Println("Landmarks:", landmarks[0])

    // 4. Add new property
    person["email"] = "driyas@example.com"
    fmt.Println("Email:", person["email"])
    
    // List Friends
    friends := person["friends"].([]map[string]interface{})
    firstFriend := friends[0]
    // name := firstFriend["name"].(string)
//     age := firstFriend["age"].(int)
    fmt.Printf("Friend 1: %s, Age: %d", firstFriend["name"].(string), firstFriend["age"].(int))
    
    
}