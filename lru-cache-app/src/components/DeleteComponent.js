import React, { useState } from 'react';
import InputField from './InputField';
import Message from './Message';
import { deleteKeyValue } from '../api/api';

const DeleteComponent = () => {
  const [key, setKey] = useState('');
  const [message, setMessage] = useState({ error: '', success: '' });

  const handleDelete = async () => {
    try {
        const data = await deleteKeyValue(key);
        setMessage({ success: 'Value deleted successfully', error: '' });
        setKey('');
        setTimeout(() => {
          setMessage({ success: '', error: '' });
        }, 3000);
      } catch (error) {
        console.error('Error deleting data:', error);
        setMessage({ error: 'Failed to delete data', success: '' });
        setTimeout(() => {
          setMessage({ success: '', error: '' });
        }, 3000);
      }
  };

  return (
    <div className="mb-3">
      <InputField placeholder="Key" value={key} onChange={(e) => setKey(e.target.value)} />
      <button className="btn btn-danger mb-2" onClick={handleDelete}>Delete</button>
      <Message message={message.error} type="danger" />
      <Message message={message.success} type="success" />
    </div>
  );
};

export default DeleteComponent;
