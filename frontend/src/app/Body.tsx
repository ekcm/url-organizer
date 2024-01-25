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
    console.log(data.message.urlLinks);  // data is a JSON object
    var urlLinks = data.message.urlLinks;

    for (var urlLink of urlLinks){
      console.log(urlLink)
      try{
        var urlLinksUrl = "http://localhost:9090/v1/url/get/" + urlLink;
        const response = await fetch(urlLinksUrl, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
        });
      
        if (!response.ok){
          throw new Error('Network response was not ok');
        }
        
        const data = await response.json();
        console.log("response from urlLink")
        console.log(data);
      
      }catch (error) {
        console.log(error);
      }
    }

  } catch (error) {
    console.log(error);
  }
};

const Body = () => {

  return (
    <div className="body">
      <div>
        <button id="populate-button" onClick={onButtonClick}>
          <h6>Click here to populate URL</h6>
        </button>
      </div>
      <div className="container-body">
        {/* <Card /> */}
      </div>
    </div>
  )

};

export default Body;