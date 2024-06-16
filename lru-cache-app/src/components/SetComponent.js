import React, { useState } from 'react';
import InputField from './InputField';
import Message from './Message';
import { setKeyValue } from '../api/api';

const SetComponent = () => {
  const [key, setKey] = useState('');
  const [value, setValue] = useState('');
  const [duration, setDuration] = useState('');
  const [message, setMessage] = useState({ error: '', success: '' });

  const handleSet = async () => {
    try {
        const data = await setKeyValue(key, value, duration);
        setMessage({ success: 'Value set successfully', error: '' });
        setTimeout(() => {
          setMessage({ success: '', error: '' });
        }, 3000);
      } catch (error) {
        console.error('Error setting data:', error);
        setMessage({ error: 'Failed to set data', success: '' });
        setTimeout(() => {
          setMessage({ success: '', error: '' });
        }, 3000);
      }
  };

  return (
    <div className="mb-3">
      <InputField placeholder="Key" value={key} onChange={(e) => setKey(e.target.value)} />
      <InputField placeholder="Value" value={value} onChange={(e) => setValue(e.target.value)} />
      <InputField placeholder="Duration (seconds)" value={duration} onChange={(e) => setDuration(e.target.value)} />
      <button className="btn btn-primary mb-2" onClick={handleSet}>Set</button>
      <Message message={message.error} type="danger" />
      <Message message={message.success} type="success" />
    </div>
  );
};

export default SetComponent;
