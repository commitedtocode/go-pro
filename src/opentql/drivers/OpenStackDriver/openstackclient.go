package OpenStackDriver 
/*
Openstack driver will return very basic common object of type Provider, 
which is used to create client for Nova, network  in plugin folder etc.

*/

import ("fmt"
        "github.com/rackspace/gophercloud"
        "github.com/rackspace/gophercloud/openstack"
        "github.com/rackspace/gophercloud/openstack/compute/v2/flavors"
)

// Option 1: Pass in the values yourself
func OpenstackClient() (r* gophercloud.ProviderClient) {
    opts := gophercloud.AuthOptions{
        IdentityEndpoint: "http://10.77.206.112:5000/v2.0",
        Username: "admin",
        Password: "openstack",
        TenantID: "72b71e23abf74d8498a2d0b5bb5bb10f",
    }
 
    
    // Option 2: Use a utility function to retrieve all your environment variables

    opts, err := openstack.AuthOptionsFromEnv()
    provider, err := openstack.AuthenticatedClient(opts)
    
    fmt.Println(provider);
    

    client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
                    Region: "RegionOne",
                    })
    fmt.Println(provider, err)
    fmt.Println(err)
    fmt.Println(client)
    flvs := flavors.ListOpts{ChangesSince: "2014-01-01T01:02:03Z", MinRAM: 4}
    fmt.Println(flvs)
    lists := flavors.ListDetail(client, flvs)
    
    fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$");
    fmt.Println(lists.Headers);
    
	fmt.Println("------------------------------------------------------");
		fmt.Print(lists);
	fmt.Println("------------------------------------------------------");	
    
    fmt.Println(lists)
    flavorID, err := flavors.Get(client, "1").Extract()
    fmt.Println(flavorID)
    */
    fmt.Println(provider,err)
    //fmt.Printf("%t",provider.EndpointLocator)
    return provider
    
    }
