import React, { useState, useEffect } from "react"
import "./App.css"
import Tabs from "./Tabs"

const BASE_DEALERS_URL = "http://localhost:8080/dealers"
const BASE_SALEMANS_URL = "http://localhost:8080/salemans"
const BASE_SALEMAPS_URL = "http://localhost:8080/salemaps"

function App() {
  let [activeTab, setActiveTab] = useState("dealers")
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
      setSalemaps(salemaps)
    })
  }, [])


  return (
    <div className="container">
      <Tabs>
        <Tabs.Switcher>
          <Tabs.Switcher.Item
            onClick={() => setActiveTab("dealers")}
            active={activeTab === "dealers"}
          >
            Dealers
          </Tabs.Switcher.Item>
          <Tabs.Switcher.Item
            onClick={() => setActiveTab("salemans")}
            active={activeTab === "salemans"}
          >
            Salemans
          </Tabs.Switcher.Item>
          <Tabs.Switcher.Item
            onClick={() => setActiveTab("salemaps")}
            active={activeTab === "salemaps"}
          >
            Salemaps
          </Tabs.Switcher.Item>
        </Tabs.Switcher>
        <Tabs.Tab active={activeTab === "dealers"}>
          <table className="table">
            <thead>
              <tr>
                <th>Name</th>
                <th>Address</th>
              </tr>
            </thead>
            <tbody>
              {
                dealers.map((dealer, key) => {
                  return (
                    <tr key={key}>
                      <td>{dealer.name}</td>
                      <td>{dealer.address}</td>
                    </tr>
                  )
                })
              }
            </tbody>
          </table>
        </Tabs.Tab>
        <Tabs.Tab active={activeTab === "salemans"}>
          <table className="table">
            <thead>
              <tr>
                <th>Name</th>
                <th>Condition</th>
              </tr>
            </thead>
            <tbody>
              {
                salemans.map((saleman, key) => {
                  return (
                    <tr key={key}>
                      <td>{saleman.name}</td>
                      <td>{saleman.condition}</td>
                    </tr>
                  )
                })
              }
            </tbody>
          </table>
        </Tabs.Tab>
        <Tabs.Tab active={activeTab === "salemaps"}>
          <table className="table">
            <thead>
              <tr>
                <th>MapID</th>
                <th>ProdID</th>
                <th>DateOfPost</th>
                <th>Custumer</th>
                <th>SalemanID</th>
                <th>Quantity</th>
                <th>SaleDate</th>
              </tr>
            </thead>
            <tbody>
              {
                salemaps.map((salemap, key) => {
                  return (
                    <tr key={key}>
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
            </tbody>
          </table>
        </Tabs.Tab>
      </Tabs>
    </div>
  )
}

export default App
