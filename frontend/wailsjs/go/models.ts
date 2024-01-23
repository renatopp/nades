export namespace main {
	
	export class Nade {
	    id: string;
	    map: string;
	    type: string;
	    x: number;
	    y: number;
	
	    static createFrom(source: any = {}) {
	        return new Nade(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.map = source["map"];
	        this.type = source["type"];
	        this.x = source["x"];
	        this.y = source["y"];
	    }
	}
	export class Variant {
	    id: string;
	    nadeId: string;
	    name: string;
	    map: string;
	    description: string;
	    x: number;
	    y: number;
	    throwImage: string;
	    lineupImage: string;
	    positionImage: string;
	
	    static createFrom(source: any = {}) {
	        return new Variant(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nadeId = source["nadeId"];
	        this.name = source["name"];
	        this.map = source["map"];
	        this.description = source["description"];
	        this.x = source["x"];
	        this.y = source["y"];
	        this.throwImage = source["throwImage"];
	        this.lineupImage = source["lineupImage"];
	        this.positionImage = source["positionImage"];
	    }
	}

}

