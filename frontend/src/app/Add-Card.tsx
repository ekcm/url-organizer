'use client'
import React, {useState} from "react";

// const createCard = (urlName: string, url: string) => {
const createCard = async() => {
  console.log("createCard called");

  const urlNameInput = document.getElementById("url-name-input") as HTMLInputElement | null;
  const urlLinkInput = document.getElementById("url-link-input") as HTMLInputElement | null;

  var requestData = {};
  
  if (urlNameInput && urlLinkInput) {
    requestData = {
      urlName: urlNameInput.value,
      url: urlLinkInput.value,
    };
  
    // Now you can use requestData for your fetch or any other logic
  } else {
    console.error('One or both input elements not found.');
  }

  try{
    var url = "http://localhost:9090/v1/urlfolder/createAddUrlLink/65983643a0c4a028d62b0bdb"
    const response = await fetch(url, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(requestData),
    });

    if (response.ok) {
      const responseData = await response.json();
      console.log('Response:', responseData);
      urlNameInput!.value = '';
      urlLinkInput!.value = '';
    } else {
      // Handle errors
      console.error('Error:', response.statusText);
    }

  }catch(error){
    console.log(error)
  }
}

const AddCard = () => {
  return (
    <div>
      <div className="add-card-container">

        <div className="add-url-container">
          <label className="url-label">Url Name</label>
          <input type="text" id="url-name-input" name="url-name-input" className="url-input"/>
        </div>

        <div className="add-url-container" id="url-link-container">
          <label className="url-label">Url Link</label>
          <input type="text" id="url-link-input" name="url-link-input" className="url-input"/>
        </div>
        
        <button id="url-link-button" onClick={createCard}>
          Add new url
        </button>
      </div>
    </div>
  )
};

export default AddCard;