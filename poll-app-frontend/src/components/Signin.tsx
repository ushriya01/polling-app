import React, { useState } from 'react';
import '../index.css';

const Signin: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSignin = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/signin', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (!response.ok) {
        if (response.status === 401) {
          // Unauthorized (invalid username or password)
          throw new Error('Invalid username or password. Please try again.');
        } else {
          // Other errors
          throw new Error('An unexpected error occurred. Please try again later.');
        }
      }

      const data = await response.json();
      const token = data.token;

      localStorage.setItem('token', token);

      window.location.href = '/polls';

      const pollsResponse = await fetch('http://localhost:8080/api/polls', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (!pollsResponse.ok) {
        throw new Error('Failed to fetch polls');
      }

    } catch (error) {
      setError(error.message);
    }
  };

  return (
    <div className="signin-container">
      <h2>Signin</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <div className="input-container">
        <input type="text" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
      </div>
      <div className="input-container">
        <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      </div>
      <button onClick={handleSignin}>Signin</button>
    </div>
  );
};

export default Signin;
