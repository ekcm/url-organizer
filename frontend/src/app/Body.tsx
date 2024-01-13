'use client'
import React, {useState} from "react";
import Card from "./Card";

const onButtonClick = async() => {
  console.log("Button clicked");

  // make get request to Go backend
  try{
    var url = "http://localhost:9090/v1/urlfolder/get/65983643a0c4a028d62b0bdb";
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok){
      throw new Error('Network response was not ok');
    }

    const data = await response.json();
    console.log("response from backend")
    console.log(data);  // data is a JSON object

  } catch (error) {
    console.log(error);
  }
};

const Body = () => {

  return (
    <div className="container-body">
      <button id="populate-button" onClick={onButtonClick}>
        <h6>Click here to populate URL</h6>
      </button>
      {/* <Card />
      <Card /> */}
    </div>
  )

};

export default Body;