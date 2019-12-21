import React from 'react'
import 'semantic-ui-css/semantic.min.css'
import Layout from '../components/Layout'
import { CSSProperties } from 'react'

const h1style: CSSProperties = {
    color: "#fff",
    textShadow: "2px 2px #a9a9a9",
    textAlign: "center"
}

const pstyle: CSSProperties = {
    color: "#fff",
    textAlign: "center"
}

const Home = () => (
    <Layout>
        <h1 style={h1style}>CRDS: The CustomResourceDefinition Toolbox</h1>
        <p style={pstyle}>Your home for building, validating, and distributing <a href="https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/">CustomResourceDefinitions</a> and controllers.</p>
    </Layout>
)

export default Home;