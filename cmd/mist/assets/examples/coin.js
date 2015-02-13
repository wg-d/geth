var contract = web3.eth.contractFromAbi([
    {
	"constant":false,
	"inputs":[
	    {"name":"_h","type":"hash256"}
	],
	"name":"confirm",
	"outputs":[],
	"type":"function"
    },{
	"constant":false,
	"inputs":[
	    {"name":_to,"type":"address"},
	    {"name":"_value","type":"uint256"},
	    {"name":"_data","type":"bytes"}
	],
	"name":"execute",
	"outputs":[
	    {"name":"_r","type":"hash256"}
	],
	"type":"function"
    },{
	"constant":false,
	"inputs":[
	    {"name":"_to","type":"address"}
	],"name":"kill",
	"outputs":[],
	"type":"function"
    },{
	"constant":false,
	"inputs":[
	    {"name":"_from","type":"address"},
	    {"name":"_to","type":"address"}
	],
	"name":"changeOwner",
	"outputs":[],
	"type":"function"
    },{
	"inputs":[
	    {"indexed":false,"name":"value","type":"uint256"}
	],
	"name":"CashIn",
	"type":"event"
    },{
	"inputs":[
	    {"indexed":true,"name":"out","type":"string32"},
	    {"indexed":false,"name":"owner","type":"address"},
	    {"indexed":false,"name":"value","type":"uint256"},
	    {"indexed":false,"name":"to","type":"address"}
	],
	"name":"SingleTransact",
	"type":"event"
    },{
	"inputs":[
	    {"indexed":true,"name":"out","type":"string32"},
	    {"indexed":false,"name":"owner","type":"address"},
	    {"indexed":false,"name":"operation","type":"hash256"},
	    {"indexed":false,"name":"value","type":"uint256"},
	    {"indexed":false,"name":"to","type":"address"}
	],
	"name":"MultiTransact",
	"type":"event"
    }
]);
