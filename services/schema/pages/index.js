import 'semantic-ui-css/semantic.min.css'
import Layout from '../components/MyLayout.js'
import FieldInput from '../components/Input.js'

class field {
    constructor(id, name, type, optional, indents) {
        this.id = id
        this.name = name
        this.type = type
        this.optional = optional
        this.indents = indents
    }
}

export default class Builder extends React.Component {
    constructor(props) {
        super(props)
        var f = new field(0, "apiVersion", "string", false, 0)
        this.state = {
            inputs: [f],
            count: 1
        }
        this.updateField = this.updateField.bind(this)
        this.addField = this.addField.bind(this)
        this.deleteField = this.deleteField.bind(this)
    }

    updateField(name, id, type, optional, indents) {
        var inp = this.state.inputs
        var index = inp.findIndex(x => x.id === id)
        inp[index].name = name
        inp[index].type = type
        inp[index].optional = optional
        inp[index].indents = indents
        console.log(inp)
        this.setState({ inputs: inp })
    }

    addField(prev) {
        var f = new field(this.state.count, "", prev.type, true, prev.indents)
        var inp = this.state.inputs
        var index = inp.findIndex(x => x.id === prev.id)
        inp.splice(index + 1, 0, f)
        // inp[this.state.count] = f
        this.setState({ inputs: inp, count: this.state.count + 1 })
    }

    deleteField(id) {
        var inp = this.state.inputs
        var index = inp.findIndex(x => x.id === id)
        inp.splice(index, 1)
        this.setState({ inputs: inp, count: this.state.count + 1 })
    }

    render() {
        console.log(this.state.inputs)
        return (
            <Layout>
                {this.state.inputs.map(f =>
                    <FieldInput key={f.id} field={f} updater={this.updateField} deleter={this.deleteField} adder={this.addField}></FieldInput>
                )}
            </Layout>
        )
    }
}