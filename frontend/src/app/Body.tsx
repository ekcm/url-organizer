'use client'
import React, {useState} from "react";
import Card from "./Card";
import AddCard from "./Add-Card";
import ReactDOM from 'react-dom';

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
      // console.log(urlLink)
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
        // console.log(data);

        createCard(data.urlName, data.url);
      
      }catch (error) {
        console.log(error);
      }
    }

  } catch (error) {
    console.log(error);
  }
};

const addNewCard = () => {
  const addCardDiv = document.createElement('div');
  ReactDOM.render(<AddCard />, addCardDiv);
  const container = document.getElementById('add-url-container');

  if (container) {
    // Append the Card component to the container
    container.appendChild(addCardDiv);
  }
}

const createCard: (urlName:string, url:string) => void = (urlName, url) => {

  const cardDiv = document.createElement('div');
  ReactDOM.render(<Card />, cardDiv);
  const container = document.getElementsByClassName('container-body')[0] as HTMLElement | null;

  if (container) {
    // Append the Card component to the container
    container.appendChild(cardDiv);
    const h2Element = cardDiv.querySelector('h2');
    if (h2Element) {
      h2Element.textContent = urlName;
    }

    const h6Element = cardDiv.querySelector('h6');
    if (h6Element) {
      h6Element.textContent = url;
    }

  }
}

const Body = () => {

  return (
    <div className="body">
      <div>
        <button id="populate-button" onClick={onButtonClick}>
          <h6>Retrieve saved urls</h6>
        </button>
        <button id="add-new-url-button" onClick={addNewCard}>
          <h6>Add new url</h6>
        </button>
      </div>
      <div id="add-url-container"></div>
      <div className="container-body">
        {/* <Card/> */}
      </div>
    </div>
  )

};

export default Body;