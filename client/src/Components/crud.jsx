import React, { useState } from 'react';
import { fetchUserValue, updateUserValue, deleteUser, createUserEntry } from './ApiService';

function App() {
  const [userID, setUserID] = useState('');
  const [value, setValue] = useState('');

  const handleCreate = async () => {
    try {
      const response = await createUserEntry(userID, parseInt(value, 10));
      console.log(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  const handleFetch = async () => {
    try {
      const response = await fetchUserValue(userID);
      console.log(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  const handleUpdate = async () => {
    try {

      // Parse the value to an integer
      const intValue = parseInt(value, 10);
      
      if (isNaN(intValue)) {
        console.error('Value must be a number');
        return;
      }
      const response = await updateUserValue(userID, intValue);
      console.log(response.data);

    } catch (error) {
      console.error(error);
    }
  };
  

  const handleDelete = async () => {
    try {
      const response = await deleteUser(userID);
      console.log(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="App">
    <div className='text-3xl font-bold underline text-center'>WELCOME TO WATER GUY </div>
      <input
        type="text"
        value={userID}
        onChange={(e) => setUserID(e.target.value)}
        placeholder="Enter userID"
      />
      <input
        type="text"
        value={value}
        onChange={(e) => setValue(e.target.value)}
        placeholder="Enter value"
      />
      <button onClick={handleCreate}>Create User Entry</button>
      <button onClick={handleFetch}>Fetch User Value</button>
      <button onClick={handleUpdate}>Update User Value</button>
      <button onClick={handleDelete}>Delete User</button>
    </div>
  );
}

export default App;
