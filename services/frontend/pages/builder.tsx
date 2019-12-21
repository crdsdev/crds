import React from 'react'
import 'semantic-ui-css/semantic.min.css'
import { Button, Input } from 'semantic-ui-react'
import {CustomResource, Field, FieldType, CRDScope} from '../models/crds'
import Layout from '../components/Layout'
import FieldInputGroup from '../components/InputGroup'
import FieldInput from '../components/Input'

const floatStyle = {
    position: "fixed",
    right: "5%",
    bottom: "5%"
}

interface BuilderState {
    resource: CustomResource;
    count: number;
}

export default class Builder extends React.Component<{},BuilderState> {
    constructor(props: {}) {
        super(props)
        const c = new CustomResource("v1alpha1", "", CRDScope.Namespaced)
        this.state = {
            resource: c,
            count: 0
        }

        this.addField = this.addField.bind(this)
        this.handleChange = this.handleChange.bind(this)
        this.fieldRemover = this.fieldRemover.bind(this)
    }

    addField() {
        const c = this.state.resource
        c.schema.push(new Field(Math.random(), "", FieldType.Object, 0, true))
        this.setState({ resource: c, count: this.state.count +1})
    }

    handleChange() {
        this.setState({ count: this.state.count +1})
    }

    fieldRemover(f: Field) {
        const index: number = this.state.resource.schema.indexOf(f);
        if (index !== -1) {
            const fields: Field[] = this.state.resource.schema
            fields.splice(index, 1)
            const tempResource: CustomResource = this.state.resource
            tempResource.setSchema(fields)
            this.setState({ resource: tempResource })
        }
    }

    render() {
        return (
            <Layout>
                <h1 style={{ color: "white" }}>CRD Schema Builder</h1>
                <p style={{ color: "white" }}>Design your schema then generate your CustomResourceDefintion.</p>
                <Button color="blue" onClick={this.addField} inverted>Add Field</Button>
                <Input placeholder="kind" ></Input>
                <Input placeholder="apiVersion" ></Input>
                <hr></hr>
                {this.state.resource.schema.map(f =>
                    f.type === FieldType.Object ? <FieldInputGroup key={f.id} field={f} handleChange={this.handleChange} removeSelf={this.fieldRemover} /> :
                    <FieldInput key={f.id} field={f} handleChange={this.handleChange} removeSelf={this.fieldRemover}/>
                )}
                <Button onClick={() => console.log(this.state.resource)} style={floatStyle} color="green" inverted>Generate</Button>
            </Layout>
        )
    }
}