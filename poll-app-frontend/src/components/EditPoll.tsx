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
    setOptions(prevOptions => [...prevOptions, { id: prevOptions.length + 1, text: '' }]);
  };

  const handleRemoveOption = (index: number) => {
    if (options.length > 2) {
      setOptions(prevOptions => prevOptions.filter((_, i) => i !== index));
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
  
      const updatedOptions = options.map(option => ({ text: option.text }));
  
      const response = await fetch(`http://localhost:8080/api/polls/${pollId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ question, options: updatedOptions }),
      });
  
      if (!response.ok) {
        if (response.status === 401) {
          throw new Error('You do not have access to update this poll');
        } else {
          throw new Error('Failed to update poll');
        }
      }      
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
        {options.map((option, index) => (
          <div key={option.id}>
            <input
              type="text"
              value={option.text}
              onChange={(e) => {
                const updatedOptions = options.map((opt, idx) =>
                  idx === index ? { ...opt, text: e.target.value } : opt
                );
                setOptions(updatedOptions);
              }}
            />
            {index >= 2 && (
              <button onClick={() => handleRemoveOption(index)}>Remove Option</button>
            )}
          </div>
        ))}
        {options.length < 2 && <button onClick={handleAddOption}>Add Option</button>}
        {options.length >= 2 && (
          <div>
            <button onClick={handleAddOption}>Add Option</button>
          </div>
        )}
      </div>
      <button onClick={handleUpdatePoll}>Update Poll</button>
    </div>
  );
};

export default EditPoll;
