export namespace define {
	
	export class Connection {
	    identity: string;
	    name: string;
	    addr: string;
	    port: string;
	    userName: string;
	    passWord: string;
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.identity = source["identity"];
	        this.name = source["name"];
	        this.addr = source["addr"];
	        this.port = source["port"];
	        this.userName = source["userName"];
	        this.passWord = source["passWord"];
	    }
	}
	export class CreateKeyValueRequest {
	    conn_Identity: string;
	    db: number;
	    key: string;
	    Kvalue: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateKeyValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.conn_Identity = source["conn_Identity"];
	        this.db = source["db"];
	        this.key = source["key"];
	        this.Kvalue = source["Kvalue"];
	        this.type = source["type"];
	    }
	}
	export class KeyListRequest {
	    conn_Identity: string;
	    db: number;
	    keyWord: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyListRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.conn_Identity = source["conn_Identity"];
	        this.db = source["db"];
	        this.keyWord = source["keyWord"];
	    }
	}
	export class KeyValueRequest {
	    conn_Identity: string;
	    db: number;
	    key: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.conn_Identity = source["conn_Identity"];
	        this.db = source["db"];
	        this.key = source["key"];
	    }
	}
	export class UpdateKeyValueRequest {
	    conn_Identity: string;
	    db: number;
	    key: string;
	    ttl: number;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateKeyValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.conn_Identity = source["conn_Identity"];
	        this.db = source["db"];
	        this.key = source["key"];
	        this.ttl = source["ttl"];
	        this.value = source["value"];
	    }
	}

}

