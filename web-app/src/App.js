import React, { useState, useEffect } from "react"
import "./App.css"
import Tabs from "./Tabs"

const BASE_DEALERS_URL = "http://localhost:8080/dealers"
const BASE_DEALS_URL = "http://localhost:8080/deals"
const BASE_PRICELIST_URL = "http://localhost:8080/pricelist"
const BASE_SALEMANS_URL = "http://localhost:8080/salemans"
const BASE_SALEMAPS_URL = "http://localhost:8080/salemaps"

function App() {
  let [activeTab, setActiveTab] = useState("dealers")
  let [dealers, setDealers] = useState([])
  let [salemans, setSalemans] = useState([])
  let [salemaps, setSalemaps] = useState([])
  let [deals, setDeals] = useState([])
  let [pricelists, setPricelist] = useState([])
  let [top5, setTop5] = useState([])
  let [unsold, setUnsold] = useState([])
  let [salemanName, setSalemanName] = useState("")
  let [leftDate, setLeftDate] = useState("")
  let [rightDate, setRightDate] = useState("")

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

  useEffect(() => {
    fetch(BASE_DEALS_URL).then(res => res.json()).then(json => {
      let { deals } = json
      setDeals(deals)
    })
  }, [])

  useEffect(() => {
    fetch(BASE_PRICELIST_URL).then(res => res.json()).then(json => {
      let { pricelist } = json
      setPricelist(pricelist)
    })
  }, [])

  useEffect(() => {
    fetch(BASE_SALEMANS_URL + "/top5").then(res => res.json()).then(json => {
      let { top5 } = json
      setTop5(top5)
    })
  }, [])

  useEffect(() => {
    fetch(`${BASE_SALEMANS_URL}/unsold?saleman_name=${salemanName}&left_date=${leftDate}&right_date=${rightDate}`).then(res => res.json()).then(json => {
      let { unsold } = json
      setUnsold(unsold)
    })
  }, [salemanName, leftDate, rightDate] )

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
            onClick={() => setActiveTab("deals")}
            active={activeTab === "deals"}
          >
            Deals
          </Tabs.Switcher.Item>
          <Tabs.Switcher.Item
            onClick={() => setActiveTab("pricelist")}
            active={activeTab === "pricelist"}
          >
            Pricelist
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
          <Tabs.Switcher.Item
            onClick={() => setActiveTab("top5")}
            active={activeTab === "top5"}
          >
            Top5
          </Tabs.Switcher.Item>
          <Tabs.Switcher.Item
            onClick={() => setActiveTab("unsold")}
            active={activeTab === "unsold"}
          >
            Unsold
          </Tabs.Switcher.Item>
        </Tabs.Switcher>
        <Tabs.Tab active={activeTab === "dealers"}>
          <table className="table">
            <thead>
              <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Address</th>
                <th>Phone</th>
                <th>StatusID</th>
              </tr>
            </thead>
            <tbody>
              {
                dealers.map((dealer, key) => {
                  return (
                    <tr key={key}>
                      <td>{dealer.id}</td>
                      <td>{dealer.name}</td>
                      <td>{dealer.address}</td>
                      <td>{dealer.phone}</td>
                      <td>{dealer.status_id}</td>
                    </tr>
                  )
                })
              }
            </tbody>
          </table>
        </Tabs.Tab>
        <Tabs.Tab active={activeTab === "deals"}>
          <table className="table">
            <thead>
              <tr>
                <th>ID</th>
                <th>Started At</th>
                <th>Finished At</th>
                <th>CustumerID</th>
              </tr>
            </thead>
            <tbody>
              {
                deals.map((deal, key) => {
                  return (
                    <tr key={key}>
                      <td>{deal.id}</td>
                      <td>{deal.started_at}</td>
                      <td>{deal.finished_at}</td>
                      <td>{deal.sub_id}</td>
                    </tr>
                  )
                })
              }
            </tbody>
          </table>
        </Tabs.Tab>
        <Tabs.Tab active={activeTab === "pricelist"}>
          <table className="table">
            <thead>
              <tr>
                <th>ID</th>
                <th>Since</th>
                <th>Price</th>
              </tr>
            </thead>
            <tbody>
              {
                pricelists.map((pricelist, key) => {
                  return (
                    <tr key={key}>
                      <td>{pricelist.id}</td>
                      <td>{pricelist.since}</td>
                      <td>{pricelist.price}</td>
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
                <th>ID</th>
                <th>Name</th>
                <th>Pincode</th>
                <th>DealerID</th>
                <th>StatusID</th>
                <th>Condition</th>
              </tr>
            </thead>
            <tbody>
              {
                salemans.map((saleman, key) => {
                  return (
                    <tr key={key}>
                      <td>{saleman.id}</td>
                      <td>{saleman.name}</td>
                      <td>{saleman.code}</td>
                      <td>{saleman.dealer_id}</td>
                      <td>{saleman.status_id}</td>
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
                      <td>{salemap.since}</td>
                      <td>{salemap.sub_id}</td>
                      <td>{salemap.saleman_id}</td>
                      <td>{salemap.quantity}</td>
                      <td>{salemap.sale_date}</td>
                    </tr>
                  )
                })
              }
            </tbody>
          </table>
        </Tabs.Tab>
        <Tabs.Tab active={activeTab === "top5"}>
          <table className="table">
            <thead>
              <tr>
                <th>Name</th>
                <th>Cash</th>
                <th>Rank</th>
              </tr>
            </thead>
            <tbody>
              {
                top5.map((top5, key) => {
                  return (
                    <tr key={key}>
                      <td>{top5.name}</td>
                      <td>{top5.cash}</td>
                      <td>{top5.rank}</td>
                    </tr>
                  )
                })
              }
            </tbody>
          </table>
        </Tabs.Tab>
        <Tabs.Tab active={activeTab === "unsold"}>
          <form>
            <div className="input">
              <label>Saleman Name</label>
              <input
                value={salemanName}
                onChange={e => setSalemanName(e.target.value)}
              />
            </div>
            <div className="input">
              <label>Left Date</label>
              <input
                value={leftDate}
                onChange={e => setLeftDate(e.target.value)}
              />
            </div>
            <div className="input">
              <label>Right Date</label>
              <input
                value={rightDate}
                onChange={e => setRightDate(e.target.value)}
              />
            </div>
          </form>
          <table className="table">
            <thead>
              <tr>
                <th>ID</th>
                <th>Name</th>
              </tr>
            </thead>
            <tbody>
              {
                unsold.map((unsold, key) => {
                  return (
                    <tr key={key}>
                      <td>{unsold.id}</td>
                      <td>{unsold.name}</td>
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
