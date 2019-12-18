import 'semantic-ui-css/semantic.min.css'
import Layout from '../components/Layout.js'

const h1style = {
    color: "#fff",
    textShadow: "2px 2px #a9a9a9",
    textAlign: "center"
}

const h3style = {
    color: "white",
    textAlign: "center"
}

const Soon = () => (
    <Layout>
        <h1 style={h1style}>Coming soon!</h1>
        <h3 style={h3style}>Help us <a href="https://github.com/crdsdev/crds">build</a> it!</h3>
    </Layout>
)

export default Soon;