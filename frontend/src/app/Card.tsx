'use client'
import React, {useState} from "react";

const populateUrl = () => {
  const h6Element = document.querySelector('.card-container h6');
  if (h6Element){
    const urlLink: string = h6Element.textContent || '';
    console.log(urlLink);

    window.open(urlLink, '_blank')
  }
}

const Card = () => {
  return (
    <div>
      <div className="card-container">
        <h2 id="card-container-url-name">urlname</h2>
        <div id="url-link-container" className="url-container">
          <h6>url link</h6>
        </div>
        <button id="url-link-button" onClick={populateUrl}>
          populate url
        </button>
      </div>
    </div>
  )
};

export default Card;