import React, { useState } from 'react';
import Form from './Components/form';
import Reminder from './Components/reminder';
import { GoalProvider } from './Context/GoalContext';

const App = () => {
  const [showForm, setShowForm] = useState(true);

  return (
    <>
      <div className="text-center my-8">
        <h1 className="text-4xl font-bold text-blue-600">Welcome to WaterGuy</h1>
        <p className="text-xl text-gray-600 mt-2">Your personal water intake reminder</p>
      </div>
      <GoalProvider>
        {showForm && <Form />}
        <Reminder setShowForm={setShowForm} />
      </GoalProvider>
    </>
  );
};

export default App;
