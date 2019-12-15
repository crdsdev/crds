import { Button, Icon, Input, Select } from 'semantic-ui-react'

function indenter(i) {
    return {
        marginLeft: 30 * i,
        marginBottom: 20
    }
}

const typeOptions = [
    { key: 'bool', value: 'bool', text: 'bool' },
    { key: 'int32', value: 'int32', text: 'int32' },
    { key: 'int64', value: 'int64', text: 'int64' },
    { key: 'object', value: 'object', text: 'object' },
    { key: 'string', value: 'string', text: 'string' },
]

export default class FieldInput extends React.Component {
    constructor(props) {
        super(props);
        this.nameChange = this.nameChange.bind(this);
        this.typeChange = this.typeChange.bind(this);
        this.indentsInc = this.indentsInc.bind(this);
        this.indentsDec = this.indentsDec.bind(this);
        this.addField = this.addField.bind(this);
        this.removeSelf = this.removeSelf.bind(this);
    }

    nameChange(e) {
        this.props.updater(e.target.value, this.props.field.id, this.props.field.type, this.props.field.indents);
    }

    typeChange(e, { value }) {
        this.props.updater(this.props.field.name, this.props.field.id, value, this.props.field.indents);
    }

    indentsInc() {
        this.props.updater(this.props.field.name, this.props.field.id, this.props.field.type, this.props.field.indents + 1);
    }

    indentsDec() {
        if (this.props.field.indents > 0) {
            this.props.updater(this.props.field.name, this.props.field.id, this.props.field.type, this.props.field.indents - 1);
        }
    }

    addField() {
        this.props.adder(this.props.field)
    }

    removeSelf() {
        this.props.deleter(this.props.field.id)
    }

    render() {
        return (
            <div id={this.props.field.id} style={indenter(this.props.field.indents)}>
                <Input onChange={this.nameChange} value={this.props.field.name}></Input>
                <Select value={this.props.field.type} onChange={this.typeChange} options={typeOptions}></Select>
                <Button onClick={this.indentsDec} icon>\</Button>
                <Button onClick={this.indentsInc} icon>/</Button>
                <Button onClick={this.addField} positive icon>+</Button>
                <Button onClick={this.removeSelf} negative icon>x</Button>
            </div>
        )
    }
}
