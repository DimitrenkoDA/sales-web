import React, { useState, useEffect } from "react"
import "./App.css"

const BASE_URL = "http://localhost:8080/dealers"

function App() {
  let [dealers, setDealers] = useState([])

  useEffect(() => {
    fetch(BASE_URL).then(res => res.json()).then(json => {
      let { dealers } = json
      setDealers(dealers)
    })
  }, [])

  return (
    <div className="container">
      <table className="table">
        <tr>
          <th>Name</th>
          <th>Address</th>
        </tr>
        {
          dealers.map(dealer => {
            return (
              <tr>
                <td>{dealer.name}</td>
                <td>{dealer.address}</td>
              </tr>
            )
          })
        }
      </table>
    </div>
  )
}

export default App
