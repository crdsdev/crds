import React from 'react'
import { Button, Icon, Input, Select, DropdownProps, InputOnChangeData } from 'semantic-ui-react'
import { Field, FieldType } from '../models/crds'
import Require from './Require'

function indenter(i: any) {
    return {
        paddingLeft: 30 * i,
        paddingBottom: 3,
        paddingTop: 3,
        // borderBottom: "1px solid #7e7e7e"
    }
}

const typeOptions = [
    { key: 'bool', value: FieldType.Bool, text: 'bool'},
    { key: 'int32', value: FieldType.Int32, text: 'int32' },
    { key: 'int64', value: FieldType.Int64, text: 'int64' },
    { key: 'resource.Quantity', value: FieldType.Quantity, text: 'resource.Quantity' },
    { key: 'object', value: FieldType.Object, text: 'object' },
    { key: 'string', value: FieldType.String, text: 'string' },
]

interface FieldInputProps {
    field: Field;
    handleChange: any;
    removeSelf: any;
}

interface FieldInputState {
    hover: boolean;
    fields: Array<Field>;
}

export default class FieldInput extends React.Component<FieldInputProps, FieldInputState> {
    constructor(props: FieldInputProps) {
        super(props);
        this.state = {
            hover: false,
            fields: this.props.field.children,
        }

        this.setName = this.setName.bind(this)
        this.setType = this.setType.bind(this)
        this.addChild = this.addChild.bind(this)
    }

    setName(event: React.ChangeEvent<HTMLInputElement>, data: InputOnChangeData) {
        this.props.field.setName(data.value)
        this.props.handleChange()
    }

    setOptional() {
        this.props.field.setOptional(!this.props.field.optional)
        this.props.handleChange()
    }

    setType(e: React.SyntheticEvent<HTMLElement, Event>, data: DropdownProps) {
        console.log(data.value as FieldType)
        this.props.field.setType(data.value as FieldType)
        this.props.handleChange()
    }

    addChild() {
        this.props.field.addChild(new Field(Math.random(), Math.random().toString(36).substring(2, 15), FieldType.Object, this.props.field.depth+1, false))
        this.props.handleChange()
    }

    render() {
        return (
            <div id={this.props.field.id.toString()} style={indenter(this.props.field.depth)}>
                <Input placeholder="fieldName" onChange={this.setName} value={this.props.field.name}></Input>
                <Select key={this.props.field.type} value={this.props.field.type} options={typeOptions} onChange={this.setType}></Select>
                <Require true={this.props.field.optional} onClick={this.setOptional}/>
                    <span>
                        <Button negative icon><Icon name="trash" /></Button>
                    </span>
            </div>
        )
    }
}
