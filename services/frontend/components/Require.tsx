import React from 'react'
import { Label } from 'semantic-ui-react'

const Require = (props: any) => <Label onClick={props.onClick} color={(props.true ? "pink" : "blue")}>{(props.true ? 'Optional' : 'Required')}</Label>

export default Require