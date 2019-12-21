import FieldInput from "../components/Input";

export enum FieldType {
    Bool = 'bool' ,
    Int32 = 'int32',
    Int64 = 'int64',
    Quantity ='resource.Quantity',
    Object = 'object',
    String = 'string'
}

export enum CRDScope {
    Cluster = 'Cluster',
    Namespaced = 'Namespace'
}

export class Field {
    public id: number;
    public name: string;
    public type: FieldType;
    public optional = false;
    public depth: number;
    public children: Array<Field>;

    public constructor(id: number, name: string, type: FieldType, depth: number, optional = true) {
        this.id = id;
        this.name = name;
        this.type = type;
        this.depth = depth;
        this.optional = optional;
        this.children = [];
    }

    public addChild(f: Field) {
        if (this.type == FieldType.Object) {
            this.children.push(f)
        }
    }

    public setName(n: string) {
        this.name = n;
    }

    public setOptional(o: boolean) {
        this.optional = o;
    }

    public setType(t: FieldType) {
        this.type = t;
    }
}


export class CustomResource {
    public apiVersion: string;
    public kind: string;
    public scope: CRDScope;
    public schema: Array<Field>;

    public constructor(apiVersion: string, kind: string, scope: CRDScope) {
        this.apiVersion = apiVersion;
        this.kind = kind;
        this.scope = scope;
        this.schema = [];
    }

    public addField(f: Field) {
        this.schema.push(f);
    }

    public setSchema(f: Field[]) {
        this.schema = f
    }
}