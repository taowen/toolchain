package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Team struct {
	Group []*Member `json:"group"`
}

type Member struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	teamA := &Team{
		Group: []*Member{
			&Member{
				Name: "Jordan",
				Age:  20,
			},
			&Member{
				Name: "James",
				Age:  24,
			},
		},
	}

	// shallowCopy(teamA)
	serializationGob(teamA)
	serializationJSON(teamA)
	reflection(teamA)
}

// "浅"拷贝
func shallowCopy(teamA *Team) {
	showTeam(teamA)
	teamB := &Team{}
	*teamB = *teamA
	teamB.Group[0].Name = "curry"
	showTeam(teamB)
	showTeam(teamA)
}

// 序列化-gob
func serializationGob(teamA *Team) {
	// showTeam(teamA)
	teamB := &Team{}
	buffer := &bytes.Buffer{}

	startTime := time.Now().Unix()
	for i := 0; i < 1000000; i++ {
		gob.NewEncoder(buffer).Encode(*teamA)
		gob.NewDecoder(bytes.NewBuffer(buffer.Bytes())).Decode(teamB)
	}
	fmt.Printf("gob time cost:%v\n", time.Now().Unix()-startTime)

	// gob.NewEncoder(buffer).Encode(*teamA)
	// gob.NewDecoder(bytes.NewBuffer(buffer.Bytes())).Decode(teamB)

	teamB.Group[0].Name = "curry"
	// showTeam(teamB)
	// showTeam(teamA)
}

// 序列化-json
func serializationJSON(teamA *Team) {
	// showTeam(teamA)
	teamB := &Team{}

	startTime := time.Now().Unix()
	for i := 0; i < 1000000; i++ {
		buffer, _ := json.Marshal(&teamA)
		json.Unmarshal([]byte(buffer), teamB)
	}
	fmt.Printf("json time cost:%v\n", time.Now().Unix()-startTime)

	// buffer, _ := json.Marshal(&teamA)
	// json.Unmarshal([]byte(buffer), teamB)

	teamB.Group[0].Name = "curry"
	// showTeam(teamB)
	// showTeam(teamA)
}

// 使用反射
func reflection(teamA *Team) {
	// showTeam(teamA)
	teamB := &Team{}

	startTime := time.Now().Unix()
	for i := 0; i < 1000000; i++ {
		original := reflect.ValueOf(teamA)
		cpy := reflect.New(original.Type()).Elem()
		copyRecursive(original, cpy)
		teamB = cpy.Interface().(*Team)
	}
	fmt.Printf("reflection time cost:%v\n", time.Now().Unix()-startTime)

	// original := reflect.ValueOf(teamA)
	// cpy := reflect.New(original.Type()).Elem()
	// // 将original拷贝至cpy中
	// copyRecursive(original, cpy)

	// // 需要进行类型断言，因此需要先转换为interface{}的格式、
	// teamB = cpy.Interface().(*Team)
	teamB.Group[0].Name = "curry"
	// showTeam(teamB)
	// showTeam(teamA)
}

// 迭代地进行拷贝
func copyRecursive(original, cpy reflect.Value) {
	// 根据反射获得的不同结果（类型）做不同的操作
	switch original.Kind() {
	case reflect.Ptr:
		// 指针类型
		// 获取指针所指的实际值，并拷贝给cpy
		originalValue := original.Elem()
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem())

	case reflect.Struct:
		// 结构体类型
		// 遍历结构体中的每一个成员
		for i := 0; i < original.NumField(); i++ {
			copyRecursive(original.Field(i), cpy.Field(i))
		}

	case reflect.Slice:
		// 切片类型
		// 在cpy中创建一个新的切片，并使用一个for循环拷贝
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}

	default:
		// 真正的拷贝操作
		cpy.Set(original)
	}
}

func showTeam(team *Team) {
	fmt.Printf("team address:%p\n", team)
	for index, itm := range team.Group {
		fmt.Printf("member %d\n", index)
		fmt.Printf("Name %v:, address: %p\n", itm.Name, &itm.Name)
	}
	fmt.Println("-------------------------")
}
