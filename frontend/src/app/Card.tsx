'use client'
import React, {useState} from "react";

const Card = () => {
  return (
    <div>
      <div className="card-container">
        <h2 id="card-container-url-name">urlname</h2>
        <div id="url-link-container">
          <h6>url link</h6>
        </div>
        <button id="url-link-button">
          populate url
        </button>
      </div>
    </div>
  )
};

export default Card;