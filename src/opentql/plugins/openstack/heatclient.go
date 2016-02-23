package openstack

import (
		"github.com/rackspace/gophercloud/openstack/orchestration/v1/stacks"
		"fmt"
        "github.com/rackspace/gophercloud"
        "github.com/rackspace/gophercloud/openstack"
		"opentql/drivers/OpenStackDriver"
		"reflect"
		"encoding/json"
		"io/ioutil"
		"os"
		"opentql/common"
)

func CreateStack()([]uint8,string){
	//error handling func
	defer func() {
        if err := recover(); err != nil {
            fmt.Println("stack creation failed:", err)
            panic(err)
            //caller shoud handle the mechanism to handle this exception
        }
    }()
	
	provider := OpenStackDriver.OpenstackClient()
	//creating template obj
	var template_desc stacks.CreateOpts 
	
	//generating absolute path for template file
	 pwd, err := os.Getwd()
     if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
     
    fmt.Println("dir-",pwd+"/src/opentql/plugins/openstack/templates/template1.json")
	
	//reading template json file
	b, err := ioutil.ReadFile(pwd+"/src/opentql/plugins/openstack/templates/template1.json")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("***",string(b),err)
	template_desc.Template = string(b)
	
	// generating random template name
	template_desc.Name = common.GenerateRandomString(12)
	
	// Note- region should come from UI
	// not using region name in gophercloud.EndpointOpts{}
	client, _ := openstack.NewOrchestrationV1(provider, gophercloud.EndpointOpts{})
	res := stacks.Create(client, template_desc)
	
	//fmt.Println(res.PrettyPrintJSON())
	if (res.Err!=nil){
		//raise exeption		
		panic(res.Err)	
	}
	result,_ := json.Marshal(res)
	fmt.Println("--->",reflect.TypeOf(result))
	
	//return stack in json format and its stack-name
	return result,template_desc.Name
	
	}


//this function is to check whether stack is created successfuly or not
//also return error messages for faulty resource
func GetStack( stack_name string, stack_id string){
	
	defer func() {
        if err := recover(); err != nil {
            fmt.Println("Get stack details failed:", err)
        }
    }()
	
	fmt.Println(stack_id)
	provider := OpenStackDriver.OpenstackClient()
	client, _ := openstack.NewOrchestrationV1(provider, gophercloud.EndpointOpts{})
	res := stacks.Get(client, stack_name, stack_id)
	fmt.Println(res.PrettyPrintJSON())
	result,_ := json.Marshal(res)
	
	var dat map[string]interface{}

    if err := json.Unmarshal(result, &dat); err != nil {
        panic(err)
    }
    fmt.Println("$$$$$",dat["Body"].([]interface{}))

	fmt.Println("--->",string(result))
	
	}