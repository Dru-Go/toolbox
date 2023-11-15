package constant

const Toolbox_usage = `
toolbox is a CLI application used to manage the materials and transactions of noah's workshop.

This application is built to 
	- manage the continuous synchronization of excel records to the datastores
	- calculate balance, unitPrice and weighted average of related transactions 

Usage:
  toolbox [command]

Available Commands:
  material    	create and get material in the datastore
  transaction 	calculate transaction prices
  import      	import data from one source and dump to the datastore
  help        	Help about any command
`
