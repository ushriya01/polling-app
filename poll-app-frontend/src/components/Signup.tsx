import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import '../index.css';

const Signup: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [signupSuccess, setSignupSuccess] = useState(false);

  const handleSignup = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (!response.ok) {
        if (response.status === 400) {
          // Bad request (validation error, etc.)
          throw new Error('Invalid username or password. Please try again.');
        } else if (response.status === 409) {
          // Conflict (username already exists, etc.)
          throw new Error('Username already exists. Please choose a different username.');
        } else {
          // Other errors
          throw new Error('An unexpected error occurred. Please try again later.');
        }
      }

      setSignupSuccess(true);
      setUsername('');
      setPassword('');
      setError('');
    } catch (error) {
      setError(error.message);
    }
  };

  return (
    <div className="signup-container">
      <h2>Signup</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {signupSuccess ? (
        <>
          <p style={{ color: 'green' }}>Signup successful. You can now sign in!</p> {/* Success message */}
          <p>
            Already have an account? <Link to="/signin">Sign in</Link> {/* Link to signin page */}
          </p>
        </>
      ) : (
        <>
          <div className="input-container">
            <input type="text" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
          </div>
          <div className="input-container">
            <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
          </div>
          <button onClick={handleSignup}>Signup</button>
        </>
      )}
    </div>
  );
};

export default Signup;
