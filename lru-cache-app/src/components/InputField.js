import React from 'react';

const InputField = ({ placeholder, value, onChange }) => (
  <input
    className="form-control mb-2"
    placeholder={placeholder}
    value={value}
    onChange={onChange}
  />
);

export default InputField;
