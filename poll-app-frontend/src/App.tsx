import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import WelcomePage from './components/WelcomePage.tsx';
import Signup from './components/Signup.tsx';
import Signin from './components/Signin.tsx';
import ListPolls from './components/ListPolls.tsx';
import CreatePoll from './components/CreatePoll.tsx';
import EditPoll from './components/EditPoll.tsx';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
      <Route path="/" element={<WelcomePage />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/signin" element={<Signin />} />
        <Route path="/polls" element={<ListPolls />} />
        <Route path="/create-poll" element={<CreatePoll />} />
        <Route path="/edit-poll/:pollId" element={<EditPoll />} />
      </Routes>
    </Router>
  );
};

export default App;
