import React from 'react'
import { Button, Icon, Input, Select, DropdownProps, InputOnChangeData } from 'semantic-ui-react'
import { Field, FieldType } from '../models/crds'
import Require from './Require'
import FieldInput from './Input'

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

interface FieldInputGroupProps {
    field: Field;
    handleChange: any;
    removeSelf: any;
}

interface FieldInputGroupState {
    hover: boolean;
    fields: Array<Field>;
}

export default class FieldInputGroup extends React.Component<FieldInputGroupProps, FieldInputGroupState> {
    constructor(props: FieldInputGroupProps) {
        super(props);
        this.state = {
            hover: false,
            fields: this.props.field.children,
        }

        this.setName = this.setName.bind(this)
        this.setOptional = this.setOptional.bind(this)
        this.setType = this.setType.bind(this)
        this.addChild = this.addChild.bind(this)
        this.removeChild = this.removeChild.bind(this)
        this.removeSelf = this.removeSelf.bind(this)
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
        this.props.field.setType(data.value as FieldType)
        this.props.handleChange()
    }

    addChild() {
        this.props.field.addChild(new Field(Math.random(), "", FieldType.Object, this.props.field.depth+1, false))
        this.props.handleChange()
    }

    removeChild(f: Field) {
        const index: number = this.props.field.children.indexOf(f);
        this.props.field.children.splice(index, 1)
        this.props.handleChange()
    }

    removeSelf() {
        this.props.removeSelf(this.props.field)
    }

    render() {
        return (
            <div id={this.props.field.id.toString()} style={indenter(this.props.field.depth)}>
                <Input placeholder="fieldName" onChange={this.setName} value={this.props.field.name}></Input>
                <Select key={this.props.field.type} value={this.props.field.type} options={typeOptions} onChange={this.setType}></Select>
                <Require true={this.props.field.optional} onClick={this.setOptional}/>
                    <span>
                        <Button onClick={this.addChild} positive icon><Icon name="plus"/></Button>
                        <Button onClick={this.removeSelf} negative icon><Icon name="trash" /></Button>
                    </span>
                 {this.props.field.children.map(f =>
                    f.type === FieldType.Object ? <FieldInputGroup key={f.id} field={f} handleChange={this.props.handleChange} removeSelf={this.removeChild} /> :
                    <FieldInput key={f.id} field={f} handleChange={this.props.handleChange} removeSelf={this.removeChild}/>
                )}
            </div>
        )
    }
}
