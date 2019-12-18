import Link from 'next/link'
import { Button, Icon, Label } from 'semantic-ui-react'

const linkStyle = {
  marginRight: 15,
  color: "#fff",
}

const linkRight = {
  marginLeft: 15,
  float: "right"
}

const labelStyle = {
  marginLeft: 5
}

const divStyle = {
  paddingBottom: 20,
  marginBottom: 20,
  verticalAlign: "middle",
  borderBottom: "3px solid #fff"
}

export default function Header() {
  return (
    <div style={divStyle}>
      <Link href="/">
        <a style={linkStyle}>Home</a>
      </Link>
      <Link href="/builder">
        <a style={linkStyle}>Builder
        <Label style={labelStyle} color="blue">
            Beta
        </Label>
        </a>
      </Link>
      <Link href="/soon">
        <a style={linkStyle}>Validator
        <Label style={labelStyle} color="red">
            Coming Soon
        </Label>
        </a>
      </Link>
      <Link href="/soon">
        <Button style={linkRight} inverted icon labelPosition='left'>
          <Icon name='github' />
          Sign In
        </Button>
      </Link>
    </div>
  )
}