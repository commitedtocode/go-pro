package openstack
/* 
this package contains the all methdos/operations related to image service.
Assume that OpenstackDriver will give here Provider object
*/

import (
		"fmt"
        "github.com/rackspace/gophercloud"
        "github.com/rackspace/gophercloud/openstack"
		"opentql/drivers/OpenStackDriver"
		"github.com/rackspace/gophercloud/pagination"
		"github.com/rackspace/gophercloud/openstack/compute/v2/images"
		"reflect"
		"encoding/json"
		
		)



func GetOSImages() ([]uint8){
	//getting provider client from openstack driver
	provider := OpenStackDriver.OpenstackClient()
	//apply filter to get list
	opts := images.ListOpts{}
	// Note- region should come from UI
	//client, _ := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{Region: "RegionOne",})
	// not using region name
	client, _ := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	// Retrieve a pager (i.e. a paginated collection)
	pager := images.ListDetail(client, opts)
	
	//fmt.Println(pager,"end")
	//fmt.Println(pager.AllPages());
	//fmt.Println(reflect.TypeOf(pager))
	
	
	// creating array of image object
	var image_array []images.Image
	var k int
	k=0
	
	// Define an anonymous function to be executed on each page's iteration, forming image list in image_array
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
							imageList, _ := images.ExtractImages(page)
							for _, i := range imageList {
								// "i" will be a images.Image
								fmt.Println(i.Name,"....",i.Status)
								//here i is of images.Image type
								image_array = append(image_array, i)
								k++
								}
							return true, fmt.Errorf("Required parameter")
						})
	 //intB, _ := json.Marshal(1)
  
	fmt.Print(err)
	image_list,_ := json.Marshal(image_array)
	fmt.Println(string(image_list))
	fmt.Println(reflect.TypeOf(image_list))
	
	//return result in json format
	return image_list
	
	}

func UploadOSImage(){
	fmt.Println("Image upload logic impemented here")
	//write logic for uploading tql configured image to openstack
	}
