this is an experimental project for implement the adapter pattern in go for log libs.
The interface is simplified to use only common


Fixing the type as the interface type, as the inteface New return the Log type it has the implemented methods
````
import "github.com/edson-dev/adapters-log/zap"

var log Log
log = adapters.Log{}.New()

log.Debug("erro")
	
````
Instantiation- The structure are of the type adapter and the new method return a self interface type, the force the struct to implement all the interface methods
We can use 2 different types of log doing a naming convention over the import 
````
import "github.com/edson-dev/adapters-log/zap"
import log2 "github.com/edson-dev/adapters-log/log"

log := adapters.Log{}.New()
log2 := log2.Log{}.New()

````

Passing as parameters - the implementation import the Log type, so when import the concrete type the abstract type from the interface can be used. 
````
func print(i Log) {
	println(i.GetLevel())
}
````