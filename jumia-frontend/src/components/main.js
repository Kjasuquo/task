import React, { useEffect, useState } from "react";
import instance from "../axios";
import "./main.css";

export default function Main() {
  const [data, setData] = useState([]);
  const [country, setCountry] = useState("")
  const [state, setState] = useState("")

  async function search() {
    try {
      let url = `/?country=${country}&state=${state}`;
      const response = await instance.get(url);
      if (response.status === 200) {
        setData(response.data);
      }
    } catch (error) {
      console.log(error);
    }
  }

  const handleCountry = (e)=>{
      setCountry(e.target.value)
  }
  const handleState = (e)=>{
      setState(e.target.value)
  }

  useEffect(() => {
    search();
  }, [state, country]);

  return (
    <>
      <main>
        <h1>PHONE NUMBERS</h1>

        <form>
          <div className="search">
            <select className="box" placeholder="Select Country" name="country" onChange = {handleCountry}>
              <option value="">Select Country</option>
              <option value="Morocco">Morocco</option>
              <option value="Mozambique">Mozambique</option>
              <option value="Uganda">Uganda</option>
              <option value="Ethiopia">Ethiopia</option>
              <option value="Cameroon">Cameroon</option>
            </select>
            
            <select className="box" placeholder="Phone Numbers" name="state" onChange = {handleState}>
              <option value="">Phone Numbers</option>
              <option value="OK">Valid</option>
              <option value="NOK">Not Valid</option>
            </select>
          </div>
        </form>

        <div class="">
          <table className="content-table">
            <thead>
              <tr>
                <th>Country</th>
                <th>State</th>
                <th>Country Code</th>
                <th>Phone Num.</th>
              </tr>
            </thead>
            <tbody>
              {data.length > 0 ? (
                data.map((number) => {
                  return (
                    <>
                      <tr>
                        <td>{number.country}</td>
                        <td>{number.valid}</td>
                        <td>{number.country_code}</td>
                        <td>{number.num}</td>
                        <td></td>
                      </tr>
                    </>
                  );
                })
              ) : (
                <p> Loading...</p>
              )}
            </tbody>
          </table>
        </div>
      </main>
    </>
  );
}
