## TCP Socket.io-like library

Example usage

``` go

    //create new server
    s, err := NewServer(":6500")
    	if err != nil {
    		panic(err)
    	}
    
        //when new client connecting, server will call "connection" event.
    	err = s.On("connection", func(c Client) {
    		log.Printf("connected %s", c.ID())
    
    		err = c.On("test", func(data []byte) {
    			var st string
    			err := json.Unmarshal(data, &st)
    			if err != nil {
    				panic(err)
    			}
    			log.Printf("%s", st)
    		})
    
    		if err != nil {
    			panic(err)
    		}
    	})
    
    	if err != nil {
    		panic(err)
    	}
    
        //this method will block next code and wait when program finish or will called Stop() method, that it run in goroutine
    	go s.Start()
    
    	d, err := NewDial("localhost:6500")
    	if err != nil {
    		panicl(err)
    	}
    	err = d.On("connection", func(c Client) {
    		b, err := json.Marshal("test")
    		if err != nil {
    			panic(err)
    		}
    		if err := d.Emit("test", b); err != nil {
    			panic(err)
    		}
    	})
    
    	if err != nil {
    		panic(err)
    	}
    
        //this code wrote for will make sure what dial code finished
    	time.Sleep(5 * time.Second)
    	
    	//stop the server wait & clode tcp connect
    	s.Stop()
```