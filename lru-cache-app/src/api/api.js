export const setKeyValue = async (key, value, duration) => {
    const response = await fetch(`http://localhost:8005/set?key=${key}&value=${value}&duration=${duration}`);
    if (!response.ok) {
      throw new Error('Failed to set data');
    }
    return response.json();
  };
  
  export const getKeyValue = async (key) => {
    const response = await fetch(`http://localhost:8005/get?key=${key}`);
    if (!response.ok) {
      throw new Error('Failed to fetch data');
    }
    return response.json();
  };
  
  export const deleteKeyValue = async (key) => {
    const response = await fetch(`http://localhost:8005/delete?key=${key}`, { method: 'DELETE' });
    if (!response.ok) {
      throw new Error('Failed to delete data');
    }
    return response.json();
  };
  