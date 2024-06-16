import React from 'react';

const Message = ({ message, type }) => (
  message ? <div className={`alert alert-${type}`}>{message}</div> : null
);

export default Message;
