import React, { useState } from 'react';

const CreatePoll: React.FC = () => {
  const [question, setQuestion] = useState('');
  const [options, setOptions] = useState<string[]>(['']);
  const [error, setError] = useState('');

  const handleAddOption = () => {
    setOptions(prevOptions => [...prevOptions, '']);
  };

  const handleOptionChange = (index: number, value: string) => {
    const newOptions = [...options];
    newOptions[index] = value;
    setOptions(newOptions);
  };

  const handleCreatePoll = async () => {
    try {
      if (question.trim() === '') {
        throw new Error('Please enter a question');
      }
      if (options.some(option => option.trim() === '')) {
        throw new Error('Please enter all options');
      }

      const pollData = {
        question,
        options: options.map(text => ({ text })),
      };

      const token = localStorage.getItem('token');

      const response = await fetch('http://localhost:8080/api/polls', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(pollData),
      });

      if (!response.ok) {
        throw new Error('Failed to create poll');
      }

      window.location.href = '/polls';
    } catch (error) {
      setError(error.message);
      console.error(error);
    }
  };

  return (
    <div>
      <h2>Create New Poll</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <div>
        <label>Question:</label>
        <input type="text" value={question} onChange={(e) => setQuestion(e.target.value)} />
      </div>
      <div>
        <label>Options:</label>
        {options.map((option, index) => (
          <div key={index}>
            <input
              type="text"
              value={option}
              onChange={(e) => handleOptionChange(index, e.target.value)}
            />
          </div>
        ))}
        <button onClick={handleAddOption}>Add Option</button>
      </div>
      <button onClick={handleCreatePoll}>Create Poll</button>
    </div>
  );
};

export default CreatePoll;
