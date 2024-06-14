import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';

const App = () => {
  const [setKey, setSetKey] = useState('');
  const [deleteKey, setDeleteKey] = useState('');
  const [setValueKey, setSetValueKey] = useState('');
  const [value, setValue] = useState('');
  const [duration, setDuration] = useState('');
  const [fetchedValue, setFetchedValue] = useState(null);
  const [errorSetMessage, setErrorSetMessage] = useState('');
  const [errorGetMessage, setErrorGetMessage] = useState('');
  const [errorDeleteMessage, setErrorDeleteMessage] = useState('');
  const [successMessage, setSuccessMessage] = useState('');
  const [successDeleteMessage, setSuccessDeleteMessage] = useState('');

  const handleSet = async () => {
    try {
      const response = await fetch(`http://localhost:8005/set?key=${setKey}&value=${value}&duration=${duration}`);
      if (!response.ok) {
        throw new Error('Failed to set data');
      }
      const data = await response.json();
      if (data && data.value !== undefined) {
       
        setTimeout(() => {
          setSuccessMessage(data.value);
        }, 3000); // Update to set only the value from data
        setErrorGetMessage('');
      } else {
        setErrorGetMessage('Value not found'); // Handle case where data.value is undefined
      }

      // Show success message
      // Clear success message after 3 seconds

    } catch (error) {
      console.error('Error setting data:', error);
      setErrorSetMessage('Failed to set data');
    }
  };

  const handleGet = async () => {
    try {
      const response = await fetch(`http://localhost:8005/get?key=${setValueKey}`);
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      if (data && data.value !== undefined) {
        setFetchedValue(data.value); // Update to set only the value from data
        setErrorGetMessage('');
      } else {
        setErrorGetMessage('Value not found'); // Handle case where data.value is undefined
      }
    } catch (error) {
      console.error('Error fetching data:', error);
      setErrorGetMessage('Error fetching data');
      setFetchedValue(null); // Reset fetchedValue on error
    }
  };

  const handleDelete = async () => {
    try {
      const response = await fetch(`http://localhost:8005/delete?key=${deleteKey}`, {
        mode: 'cors',
        method: 'DELETE',
      });
      console.log(response);
      if (!response.ok) {
        throw new Error('Failed to delete data');
      }
      // Optionally handle response data if needed

      const data = await response.json();
      console.log(data);
      if (data && data.value !== undefined) {
        setSuccessDeleteMessage(data.value);
        setDeleteKey('');
        setErrorDeleteMessage('');
      } else {
        setErrorDeleteMessage('Value not found'); // Handle case where data.value is undefined
      }

    } catch (error) {
      console.error('Error deleting data:', error);
      setErrorDeleteMessage('Failed to delete data');
    }
  };

  return (
    <div className="container mt-5">
      <h1 className="mb-4">LRU Cache</h1>


      <div className="mb-3">
        <input className="form-control mb-2" placeholder="Key" value={setKey} onChange={(e) => setSetKey(e.target.value)} />
        <input className="form-control mb-2" placeholder="Value" value={value} onChange={(e) => setValue(e.target.value)} />
        <input className="form-control mb-2" placeholder="Duration (seconds)" value={duration} onChange={(e) => setDuration(e.target.value)} />
        <button className="btn btn-primary mb-2" onClick={handleSet}>Set</button>
        {errorSetMessage && <div className="alert alert-danger">{errorSetMessage}</div>}
        {successMessage && <div className="alert alert-success">{successMessage}</div>}
      </div>
      <div className="mb-3">
        <input className="form-control mb-2" placeholder="Key" value={setValueKey} onChange={(e) => setSetValueKey(e.target.value)} />
        <button className="btn btn-success mb-2" onClick={handleGet}>Get</button>
        {errorGetMessage && <div className="alert alert-danger">{errorGetMessage}</div>}
        {fetchedValue !== null && <div className="alert alert-success">Value: {fetchedValue}</div>}
        {/* Ensure fetchedValue is not null before rendering */}
      </div>
      <div className="mb-3">
        <input className="form-control mb-2" placeholder="Key" value={deleteKey} onChange={(e) => setDeleteKey(e.target.value)} />
        <button className="btn btn-danger mb-2" onClick={handleDelete}>Delete</button>
        {errorDeleteMessage && <div className="alert alert-danger">{errorDeleteMessage}</div>}
        {successDeleteMessage && <div className="alert alert-success">{successDeleteMessage}</div>}
      </div>
    </div>
  );
};

export default App;
