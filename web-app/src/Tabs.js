import React from "react"
import classnames from "classnames"
import "./Tabs.css"

const Tabs = ({ children, ...rest }) => {
  return (
    <div className="tabs" {...rest}>
      {children}
    </div>
  )
}

Tabs.Tab = ({ active, children, ...rest }) => {
  const className = classnames({ active }, "tab")

  return (
    <div className={className} {...rest}>
      {children}
    </div>
  )
}

Tabs.Switcher = ({ children, ...rest }) => {
  return (
    <div className="tabs-switcher" {...rest}>
      {children}
    </div>
  )
}

Tabs.Switcher.Item = ({ children, active, ...rest }) => {
  const className = classnames({ active }, "tabs-switcher-item")

  return (
    <div className={className} {...rest}>
      {children}
    </div>
  )
}


export default Tabs
