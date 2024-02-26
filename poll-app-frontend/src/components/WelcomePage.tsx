import React from 'react';
import { Link } from 'react-router-dom';
import '../index.css';

const HomePage: React.FC = () => {
  return (
    <div className="home-container">
      <h1>Welcome to Poll App</h1>
      <div className="button-container">
        <Link to="/signup"><button className="button">Sign Up</button></Link>
        <span className="spacer"></span>
        <Link to="/signin"><button className="button">Sign In</button></Link>
      </div>
      <div className="decorations"></div>
    </div>
  );
};

export default HomePage;
