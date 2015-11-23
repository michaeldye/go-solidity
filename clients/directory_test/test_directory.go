package main

import (
    "fmt"
    "repo.hovitos.engineering/MTN/go-solidity/contract_api"
    "os"
    )

func main() {
    fmt.Println("Starting directory client")

    if len(os.Args) < 4 {
        fmt.Printf("...terminating, only %v parameters were passed.\n",len(os.Args))
        os.Exit(1)
    }

    dir_contract := os.Args[1]
    fmt.Printf("using directory %v\n",dir_contract)
    registry_owner := os.Args[2]
    fmt.Printf("using account %v\n",registry_owner)

    // Establish the directory contract
    dirc := contract_api.SolidityContractFactory("directory")
    if _,err := dirc.Load_contract(registry_owner, ""); err != nil {
        fmt.Printf("...terminating, could not load directory contract: %v\n",err)
        os.Exit(1)
    }
    dirc.Set_contract_address(dir_contract)

    // Test to make sure the directory contract is invokable
    fmt.Printf("Retrieve contract for name 'a', should be zeroes.\n")
    p := make([]interface{},0,10)
    p = append(p,"a")
    if caddr,err := dirc.Invoke_method("get_entry",p); err == nil {
        fmt.Printf("Contract Address is %v\n",caddr)
    } else {
        fmt.Printf("Error invoking get_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Retrieve a list of all registered names, should have only the MTN platform entries.\n")
    p = make([]interface{},0,10)
    p = append(p,0)
    p = append(p,10)
    if nl,err := dirc.Invoke_method("get_names",p); err == nil {
        fmt.Printf("Registered names %v\n",nl)
    } else {
        fmt.Printf("Error invoking get_names: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Register 'a' with address 0x0000000000000000000000000000000000000010.\n")
    p = make([]interface{},0,10)
    p = append(p,"a")
    p = append(p,"0x0000000000000000000000000000000000000010")
    p = append(p,0)
    if _,err := dirc.Invoke_method("add_entry",p); err == nil {
        fmt.Printf("Registered 'a'.\n")
    } else {
        fmt.Printf("Error invoking add_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Retrieve 'a', should have address 10.\n")
    p = make([]interface{},0,10)
    p = append(p,"a")
    if aa,err := dirc.Invoke_method("get_entry",p); err == nil {
        fmt.Printf("Retrieved 'a', is %v.\n",aa)
    } else {
        fmt.Printf("Error invoking add_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Retrieve owner of 'a', should be %v.\n",registry_owner)
    p = make([]interface{},0,10)
    p = append(p,"a")
    p = append(p,0)
    if aa,err := dirc.Invoke_method("get_entry_owner",p); err == nil {
        fmt.Printf("Retrieved owner of 'a' %v.\n",aa)
    } else {
        fmt.Printf("Error invoking add_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Retrieve a list of all registered names, should have 'a' in it.\n")
    p = make([]interface{},0,10)
    p = append(p,0)
    p = append(p,10)
    if nl,err := dirc.Invoke_method("get_names",p); err == nil {
        fmt.Printf("Registered names %v\n",nl)
    } else {
        fmt.Printf("Error invoking get_names: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Register 'b' with address 0x0000000000000000000000000000000000000011.\n")
    p = make([]interface{},0,11)
    p = append(p,"b")
    p = append(p,"0x0000000000000000000000000000000000000011")
    p = append(p,0)
    if _,err := dirc.Invoke_method("add_entry",p); err == nil {
        fmt.Printf("Registered 'b'.\n")
    } else {
        fmt.Printf("Error invoking add_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Register 'c' with address 0x0000000000000000000000000000000000000012.\n")
    p = make([]interface{},0,11)
    p = append(p,"c")
    p = append(p,"0x0000000000000000000000000000000000000012")
    p = append(p,0)
    if _,err := dirc.Invoke_method("add_entry",p); err == nil {
        fmt.Printf("Registered 'c'.\n")
    } else {
        fmt.Printf("Error invoking add_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Register 'c' with address 0x0000000000000000000000000000000000000013, version 1.\n")
    p = make([]interface{},0,11)
    p = append(p,"c")
    p = append(p,"0x0000000000000000000000000000000000000013")
    p = append(p,1)
    if _,err := dirc.Invoke_method("add_entry",p); err == nil {
        fmt.Printf("Registered 'c, version 1'.\n")
    } else {
        fmt.Printf("Error invoking add_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Retrieve a list of all registered names, should have 'a,b,c,c' in it.\n")
    p = make([]interface{},0,10)
    p = append(p,0)
    p = append(p,10)
    if nl,err := dirc.Invoke_method("get_names",p); err == nil {
        fmt.Printf("Registered names %v\n",nl)
    } else {
        fmt.Printf("Error invoking get_names: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Retrieve verison 1 of 'c',should be 0x00..0013.\n")
    p = make([]interface{},0,10)
    p = append(p,"c")
    p = append(p,1)
    if nl,err := dirc.Invoke_method("get_entry_by_version",p); err == nil {
        fmt.Printf("Registered c version 1 as %v\n",nl)
    } else {
        fmt.Printf("Error invoking get_entry_by_version: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Delete 'b'.\n")
    p = make([]interface{},0,10)
    p = append(p,"b")
    p = append(p,0)
    if _,err := dirc.Invoke_method("delete_entry",p); err == nil {
        fmt.Printf("Deleted 'b'\n")
    } else {
        fmt.Printf("Error invoking delete_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Retrieve a list of all registered names, should have MTN contracts plus 'a,c,c' in it.\n")
    p = make([]interface{},0,10)
    p = append(p,0)
    p = append(p,10)
    if nl,err := dirc.Invoke_method("get_names",p); err == nil {
        fmt.Printf("Registered names %v\n",nl)
    } else {
        fmt.Printf("Error invoking get_names: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Delete 'c' version 0.\n")
    p = make([]interface{},0,10)
    p = append(p,"c")
    p = append(p,0)
    if _,err := dirc.Invoke_method("delete_entry",p); err == nil {
        fmt.Printf("Deleted 'c' version 0.\n")
    } else {
        fmt.Printf("Error invoking delete_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Delete 'c' version 1.\n")
    p = make([]interface{},0,10)
    p = append(p,"c")
    p = append(p,1)
    if _,err := dirc.Invoke_method("delete_entry",p); err == nil {
        fmt.Printf("Deleted 'c' version 1.\n")
    } else {
        fmt.Printf("Error invoking delete_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Delete 'a'.\n")
    p = make([]interface{},0,10)
    p = append(p,"a")
    p = append(p,0)
    if _,err := dirc.Invoke_method("delete_entry",p); err == nil {
        fmt.Printf("Deleted 'a'\n")
    } else {
        fmt.Printf("Error invoking delete_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Retrieve a list of all registered names, should be just the MTN platform entries.\n")
    p = make([]interface{},0,10)
    p = append(p,0)
    p = append(p,10)
    if nl,err := dirc.Invoke_method("get_names",p); err == nil {
        fmt.Printf("Registered names %v\n",nl)
    } else {
        fmt.Printf("Error invoking get_names: %v\n",err)
        os.Exit(1)
    }

    // ================= whisper directory tests ===================================
    // Find the whisper directory
    fmt.Printf("Retrieve contract for whisper directory.\n")
    p = make([]interface{},0,10)
    p = append(p,"whisper_directory")
    var wdaddr string
    if wda,err := dirc.Invoke_method("get_entry",p); err == nil {
        fmt.Printf("Contract Address is %v\n",wda)
        wdaddr = wda.(string)
    } else {
        fmt.Printf("Error invoking get_entry: %v\n",err)
        os.Exit(1)
    }

    // Establish the whisper directory contract
    wd := contract_api.SolidityContractFactory("whisper_directory")
    if _,err := wd.Load_contract(registry_owner, ""); err != nil {
        fmt.Printf("...terminating, could not load whisper directory contract: %v\n",err)
        os.Exit(1)
    }
    wd.Set_contract_address(wdaddr)

    fmt.Printf("Get entry at address 0x00..001, it's not there.\n")
    p = make([]interface{},0,10)
    p = append(p,"0x0000000000000000000000000000000000000001")
    if wa,err := wd.Invoke_method("get_entry",p); err == nil {
        fmt.Printf("Received %v.\n",wa)
    } else {
        fmt.Printf("Error invoking whisper get_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Add an entry for 0x0000deadbeef.\n")
    p = make([]interface{},0,10)
    p = append(p,"0x0000deadbeef")
    if _,err := wd.Invoke_method("add_entry",p); err == nil {
        fmt.Printf("Added an entry.\n")
    } else {
        fmt.Printf("Error invoking whisper add_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Get your current entry, should be 0x0000deadbeef.\n")
    p = make([]interface{},0,10)
    p = append(p,registry_owner)
    if wa,err := wd.Invoke_method("get_entry",p); err == nil {
        fmt.Printf("Received %v.\n",wa)
    } else {
        fmt.Printf("Error invoking whisper get_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Update your entry with 0x000012345678.\n")
    p = make([]interface{},0,10)
    p = append(p,"0x000012345678")
    if _,err := wd.Invoke_method("add_entry",p); err == nil {
        fmt.Printf("Updated the entry.\n")
    } else {
        fmt.Printf("Error invoking whisper add_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Get your current entry, should be 0x000012345678.\n")
    p = make([]interface{},0,10)
    p = append(p,registry_owner)
    if wa,err := wd.Invoke_method("get_entry",p); err == nil {
        fmt.Printf("Received %v.\n",wa)
    } else {
        fmt.Printf("Error invoking whisper get_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Delete your entry.\n")
    if _,err := wd.Invoke_method("delete_entry",nil); err == nil {
        fmt.Printf("Deleted the entry.\n")
    } else {
        fmt.Printf("Error invoking whisper add_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Printf("Get your current entry, should be empty string.\n")
    p = make([]interface{},0,10)
    p = append(p,registry_owner)
    if wa,err := wd.Invoke_method("get_entry",p); err == nil {
        fmt.Printf("Received %v.\n",wa)
    } else {
        fmt.Printf("Error invoking whisper get_entry: %v\n",err)
        os.Exit(1)
    }

    fmt.Println("Terminating directory test client")
}