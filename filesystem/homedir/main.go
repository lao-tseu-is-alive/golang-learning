package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The user login is		: %s\n", usr.Username)
	fmt.Printf("The user name is		: %s\n", usr.Name)
	fmt.Printf("The user ID is			: %s\n", usr.Uid)
	fmt.Printf("The user group ID is		: %s\n", usr.Gid)
	fmt.Printf("The user home directory : %s\n", usr.HomeDir)
	fmt.Printf("The user belongs to this groups : \n")
	gidList, _ := usr.GroupIds()
	for gid := range gidList {
		//fmt.Printf("GID [%v]\t%T\t%s\n",gid, gid, gidList[gid])
		group, err := user.LookupGroupId(gidList[gid])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("GID[%s]\t%s\n", gidList[gid], group.Name)
	}
}
