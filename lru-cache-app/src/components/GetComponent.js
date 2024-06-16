import React, { useState } from 'react';
import InputField from './InputField';
import Message from './Message';
import { getKeyValue } from '../api/api';

const GetComponent = () => {
  const [key, setKey] = useState('');
  const [fetchedValue, setFetchedValue] = useState(null);
  const [message, setMessage] = useState({ error: '' });

  const handleGet = async () => {
    try {
        const data = await getKeyValue(key);
        if (data && data.value !== undefined) {
          setFetchedValue(data.value);
          setMessage({ error: '', message:'' });
        } else {
          setMessage({ error: 'Value not found', message: ''});
          setFetchedValue(null);
        }
        setTimeout(() => {
          setMessage({ error: '', message: ''});
          setFetchedValue(null);
        }, 3000);
      } catch (error) {
        console.error('Error fetching data:', error);
        setMessage({ error: 'Error fetching data', message:'' });
        setFetchedValue(null);
        setTimeout(() => {
          setMessage({ error: '', message : '' });
          setFetchedValue(null);
        }, 3000);
      }
  };

  return (
    <div className="mb-3">
      <InputField placeholder="Key" value={key} onChange={(e) => setKey(e.target.value)} />
      <button className="btn btn-success mb-2" onClick={handleGet}>Get</button>
      <Message message={message.error} type="danger" />
      {fetchedValue !== null && <Message message={`Value: ${fetchedValue}`} type="success" />}
    </div>
  );
};

export default GetComponent;
