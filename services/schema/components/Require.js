import { Label } from 'semantic-ui-react'

const Require = props => <Label onClick={props.onClick} color={(props.true ? 'gray' : 'blue')}>{(props.true ? 'Optional' : 'Required')}</Label>

export default Require