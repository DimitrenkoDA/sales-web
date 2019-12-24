import React, { useState, useEffect } from "react"
import "./App.css"

const BASE_DEALERS_URL = "http://localhost:8080/dealers"
const BASE_SALEMANS_URL = "http://localhost:8080/salemans"
const BASE_SALEMAPS_URL = "http://localhost:8080/salemaps"

function App() {
  let [dealers, setDealers] = useState([])
  let [salemans, setSalemans] = useState([])
  let [salemaps, setSalemaps] = useState([])

  useEffect(() => {
    fetch(BASE_DEALERS_URL).then(res => res.json()).then(json => {
      let { dealers } = json
      setDealers(dealers)
    })
  }, [])

  useEffect(() => {
    fetch(BASE_SALEMANS_URL).then(res => res.json()).then(json => {
      let { salemans } = json
      setSalemans(salemans)
    })
  }, [])

  useEffect(() => {
    fetch(BASE_SALEMAPS_URL).then(res => res.json()).then(json => {
      let { salemaps } = json
      setSalemans(salemaps)
    })
  }, [])
  

  return (
    <div className="container">
      <div className="inner">
        <table className="table">
          <tr>
            <th className="header">Name</th>
            <th className="header">Address</th>
          </tr>
          {
            dealers.map(dealer => {
              return (
                <tr className="values">
                  <td>{dealer.name}</td>
                  <td>{dealer.address}</td>
                </tr>
              )
            })
          }
        </table>
      </div>
      <div className="inner">
        <table className="table">
          <tr>
            <th className="header">Name</th>
            <th className="header">Condition</th>
          </tr>
          {
            salemans.map(saleman => {
              return (
                <tr className="values">
                  <td>{saleman.name}</td>
                  <td>{saleman.condition}</td>
                </tr>
              )
            })
          }
        </table>
      </div>
      <div className="inner">
        <table className="table">
          <tr>
            <th className="header">MapID</th>
            <th className="header">ProdID</th>
            <th className="header">DateOfPost</th>
            <th className="header">Custumer</th>
            <th className="header">SalemanID</th>
            <th className="header">Quantity</th>
            <th className="header">SaleDate</th>
          </tr>
          {
            salemaps.map(salemap => {
              return (
                <tr className="values">
                  <td>{salemap.id}</td>
                  <td>{salemap.prod_id}</td>
                  <td>{salemap.dat}</td>
                  <td>{salemap.sub_id}</td>
                  <td>{salemap.man_code}</td>
                  <td>{salemap.quantity}</td>
                  <td>{salemap.sale_dat}</td>
                </tr>
              )
            })
          }
        </table>
      </div>
    </div>
  )
}

export default App
