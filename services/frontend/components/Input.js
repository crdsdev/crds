import { Button, Icon, Input, Select } from 'semantic-ui-react'
import Require from './Require'

function indenter(i) {
    return {
        paddingLeft: 30 * i,
        paddingBottom: 3,
        paddingTop: 3,
        // borderBottom: "1px solid #7e7e7e"
    }
}

const typeOptions = [
    { key: 'bool', value: 'bool', text: 'bool' },
    { key: 'int32', value: 'int32', text: 'int32' },
    { key: 'int64', value: 'int64', text: 'int64' },
    { key: 'resource.Quantity', value: 'resource.Quantity', text: 'resource.Quantity' },
    { key: 'object', value: 'object', text: 'object' },
    { key: 'string', value: 'string', text: 'string' },
]

export default class FieldInput extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            hover: false
        }
        this.nameChange = this.nameChange.bind(this);
        this.typeChange = this.typeChange.bind(this);
        this.indentsInc = this.indentsInc.bind(this);
        this.indentsDec = this.indentsDec.bind(this);
        this.optionalSwitch = this.optionalSwitch.bind(this);
        this.addField = this.addField.bind(this);
        this.removeSelf = this.removeSelf.bind(this);
        this.hoverOn = this.hoverOn.bind(this);
        this.hoverOff = this.hoverOff.bind(this);
    }

    nameChange(e) {
        this.props.updater(e.target.value, this.props.field.id, this.props.field.type, this.props.field.optional, this.props.field.indents);
    }

    typeChange(e, { value }) {
        this.props.updater(this.props.field.name, this.props.field.id, value, this.props.field.optional, this.props.field.indents);
    }

    indentsInc() {
        this.props.updater(this.props.field.name, this.props.field.id, this.props.field.type, this.props.field.optional, this.props.field.indents + 1);
    }

    indentsDec() {
        if (this.props.field.indents > 0) {
            this.props.updater(this.props.field.name, this.props.field.id, this.props.field.type, this.props.field.optional, this.props.field.indents - 1);
        }
    }

    optionalSwitch() {
        this.props.updater(this.props.field.name, this.props.field.id, this.props.field.type, !this.props.field.optional, this.props.field.indents);
    }

    addField() {
        this.props.adder(this.props.field)
    }

    removeSelf() {
        this.props.deleter(this.props.field.id)
    }

    hoverOff() {
        this.setState({ hover: false })
    }

    hoverOn() {
        this.setState({ hover: true })
    }

    render() {
        return (
            <div id={this.props.field.id} style={indenter(this.props.field.indents)} onMouseEnter={this.hoverOn} onMouseLeave={this.hoverOff}>
                <Input placeholder="fieldName" onChange={this.nameChange} value={this.props.field.name}></Input>
                <Select value={this.props.field.type} onChange={this.typeChange} options={typeOptions}></Select>
                <Require onClick={this.optionalSwitch} true={this.props.field.optional} />
                {this.state.hover &&
                    <span>
                        <Button onClick={this.indentsDec} icon><Icon name="chevron left" /></Button>
                        <Button onClick={this.indentsInc} icon><Icon name="chevron right" /></Button>
                        <Button onClick={this.addField} positive icon><Icon name="plus" /></Button>
                        <Button onClick={this.removeSelf} negative icon><Icon name="trash" /></Button>
                    </span>
                }
            </div>
        )
    }
}
