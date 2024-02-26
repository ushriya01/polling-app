import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

interface Option {
  id: number;
  text: string;
  votes: number;
}

interface Poll {
  id: number;
  question: string;
  options: Option[];
  created_by: string;
}

const ListPolls: React.FC = () => {
  const [polls, setPolls] = useState<Poll[]>([]);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = localStorage.getItem('token');

        if (!token) {
          setError('Unauthorized access');
          return;
        }

        const response = await fetch('http://localhost:8080/api/polls', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Failed to fetch polls');
        }

        const data = await response.json();
        setPolls(data);
      } catch (error) {
        setError(error.message);
        console.error(error);
      }
    };

    fetchData();
  }, []);

  const handleVote = async (pollId: number, optionId: number) => {
    try {
      const token = localStorage.getItem('token');
  
      const response = await fetch('http://localhost:8080/api/polls/vote', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          poll_id: pollId,
          option_id: optionId,
        }),
      });
  
      if (!response.ok) {
        if (response.status === 409) {
          throw new Error('You have already voted for this poll');
        } else {
          throw new Error('Failed to vote');
        }
      }
      const updatedPolls = polls.map((poll) => {
        if (poll.id === pollId) {
          const updatedOptions = poll.options.map((option) => {
            if (option.id === optionId) {
              return { ...option, votes: option.votes + 1 };
            }
            return option;
          });
          return { ...poll, options: updatedOptions };
        }
        return poll;
      });
      setPolls(updatedPolls);
    } catch (error) {
      setError(error.message);
      console.error(error);
      setTimeout(() => {
        setError('');
      }, 2000);
    }
  };

  const handleEdit = (pollId: number) => {
    window.location.href = `/edit-poll/${pollId}`;
  };

  const handleDelete = async (pollId: number) => {
    try {
      const token = localStorage.getItem('token');

      const response = await fetch(`http://localhost:8080/api/polls/${pollId}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        if (response.status === 401) {
          throw new Error('You do not have access to delete this poll');
        } else {
          throw new Error('Failed to delete poll');
        }
      }      

      const updatedPolls = polls.filter((poll) => poll.id !== pollId);
      setPolls(updatedPolls);
    } catch (error) {
      setError(error.message);
      console.error(error);
    }
  };

  const handleShowSelectedUsers = async (pollId: number, optionId: number) => {
    try {
      const token = localStorage.getItem('token');

      const response = await fetch(`http://localhost:8080/api/polls/${pollId}/options/${optionId}/votes`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error('Failed to fetch selected users');
      }

      const data = await response.json();
      const selectedUsers = data.map((vote: any) => vote.user_id).join(', ');

      if (selectedUsers) {
        alert(`Users who selected this option: ${selectedUsers}`);
      }
    } catch (error) {
      setError(error.message);
      console.error(error);
    }
  };

  return (
    <div>
      <h2>Polling Dashboard</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <ul style={{ textAlign: 'left' }}>
        {polls === null || polls.length === 0 ? (
          <li>No polls available.</li>
        ) : (
          polls.map((poll) => (
            <li key={poll.id}>
              <h3>{poll.question}</h3>
              <ul>
                {poll.options.map((option) => (
                  <li key={option.id}>
                    {option.text} - Votes:{' '}
                    {option.votes > 0 ? (
                      <span
                        style={{ cursor: 'pointer', textDecoration: 'underline' }}
                        onClick={() => handleShowSelectedUsers(poll.id, option.id)}
                      >
                        {option.votes}
                      </span>
                    ) : (
                      <span>{option.votes}</span>
                    )}
                    <button onClick={() => handleVote(poll.id, option.id)}>Vote</button>
                  </li>
                ))}
              </ul>
              <button onClick={() => handleEdit(poll.id)}>Edit</button>
              <button onClick={() => handleDelete(poll.id)}>Delete</button>
            </li>
          ))
        )}
      </ul>
      <div style={{ textAlign: 'left', marginTop: '20px' }}>
        <Link to="/create-poll">
          <button>Create New Poll</button>
        </Link>
      </div>
    </div>
  );
};

export default ListPolls;
