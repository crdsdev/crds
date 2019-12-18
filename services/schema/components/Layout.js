import Header from './Header'

const layoutStyle = {
  padding: "2% 2% 2% 2%",
  border: '4px solid #DDD',
  backgroundColor: "black",
  height: "100vh",
  overflow: "auto"
}

export default function Layout(props) {
  return (
    <div style={layoutStyle}>
      <Header />
      {props.children}
    </div>
  )
}