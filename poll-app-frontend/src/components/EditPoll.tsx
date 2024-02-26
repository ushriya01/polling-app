import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

interface Option {
  id: number;
  text: string;
}

const EditPoll: React.FC = () => {
  const { pollId } = useParams<{ pollId: string }>();
  const [question, setQuestion] = useState('');
  const [options, setOptions] = useState<Option[]>([]);
  const [newOption, setNewOption] = useState('');
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchPollData = async () => {
      try {
        const token = localStorage.getItem('token');

        const response = await fetch(`http://localhost:8080/api/polls/${pollId}`, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Failed to fetch poll data');
        }

        const pollData = await response.json();
        setQuestion(pollData.question);
        setOptions(pollData.options);
      } catch (error) {
        setError(error.message);
        console.error(error);
      }
    };

    fetchPollData();
  }, [pollId]);

  const handleAddOption = () => {
    if (newOption.trim() !== '') {
      setOptions(prevOptions => [...prevOptions, { id: options.length + 1, text: newOption }]);
      setNewOption('');
    }
  };

  const handleUpdatePoll = async () => {
    try {
      if (question.trim() === '') {
        throw new Error('Please enter a question');
      }
      if (options.some(option => option.text.trim() === '')) {
        throw new Error('Please enter all options');
      }

      const token = localStorage.getItem('token');

      const response = await fetch(`http://localhost:8080/api/polls/${pollId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ question, options }),
      });

      if (!response.ok) {
        throw new Error('Failed to update poll');
      }

      // Redirect user to the list of polls with updated data
      window.location.href = '/polls';
    } catch (error) {
      setError(error.message);
      console.error(error);
    }
  };

  return (
    <div>
      <h2>Edit Poll</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <div>
        <label>Question:</label>
        <input type="text" value={question} onChange={(e) => setQuestion(e.target.value)} />
      </div>
      <div>
        <label>Options:</label>
        {options.map((option) => (
          <div key={option.id}>
            <input
              type="text"
              value={option.text}
              onChange={(e) => {
                const newOptions = [...options];
                newOptions.find(o => o.id === option.id)!.text = e.target.value;
                setOptions(newOptions);
              }}
            />
          </div>
        ))}
        <div>
          <input
            type="text"
            value={newOption}
            onChange={(e) => setNewOption(e.target.value)}
            placeholder="New Option"
          />
          <button onClick={handleAddOption}>Add Option</button>
        </div>
      </div>
      <button onClick={handleUpdatePoll}>Update Poll</button>
    </div>
  );
};

export default EditPoll;
